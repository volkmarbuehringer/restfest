package generator

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
)

const basePro string = "restfest"

type mapperStru struct {
	Package   string
	Directory string
	Table     []*TabFlag
	Timestamp time.Time
}

type TabFlag struct {
	Table     string
	Flag      int
	TFlag     bool
	PK        string
	PKType    string
	PKPtr     string
	Parameter string
	Worker    bool
	Flagger   bool
	Specific  string
}

var db *pgx.Conn
var gendir string = "../" + os.Args[1] + "/gener"
var tflag bool

func generateMap(t *template.Template, arr []*TabFlag) error {
	f1, err := os.Create("mapper.tmp")
	if err != nil {
		log15.Crit("map", "create", err)
		return err
	}

	defer f1.Close()
	if len(arr) > 0 {
		if err := t.ExecuteTemplate(f1, "mapper.tmpl", mapperStru{os.Args[1], basePro + "/" + os.Args[1] + "/gener", arr, time.Now()}); err != nil {
			log15.Crit("template map", "map", err)
			return err
		}

	}
	return nil
}

func generateMapTyp(t *template.Template, arr []*TabFlag) error {
	f, err := os.Create("mappertyp.tmp")
	if err != nil {
		log15.Crit("DBFehler", "create", err)
		return err
	}

	defer f.Close()
	if len(arr) > 0 {
		if err := t.ExecuteTemplate(f, "mappertyp.tmpl", mapperStru{os.Args[1], basePro + "/" + os.Args[1] + "/gener", arr, time.Now()}); err != nil {
			log15.Crit("DBFehler", "map", err)
			return err
		}

	}
	return nil
}

func generateStru(t *template.Template, row *TabFlag) error {

	namer := row.Table
	flagger := false
	type Prof struct {
		Column      string
		ColumnTrans string
		ColumnT     string
	}
	profA := []Prof{}
	profB := []Prof{}
	if len(row.Parameter) > 0 {
		namer = row.Parameter

	} else {
		profB = []Prof{Prof{"length", "int", ""}, Prof{"offset", "int", ""}}
		flagger = true
		row.Flagger = true
	}
	if rows, err := db.Query(sqlallcols, row.Table); err != nil {
		log15.Crit("DBFehler", "query", err)
		return err
	} else {
		defer rows.Close()

		ColumnsInsert := []string{}
		ColumnsUpdate := []string{}
		BindsInsert := []string{}
		BindsVarInsert := []string{}
		BindsUpdate := []string{}
		Columns := []string{}
		ColumnsT := []string{}

		var pkBind string
		for inser, upder := 0, 0; rows.Next(); {

			prof := Prof{}
			if err := rows.Scan(&prof.Column, &prof.ColumnTrans, &prof.ColumnT); err != nil {
				log15.Crit("DBFehler", "scan1", err)
				return err
			}
			profA = append(profA, prof)
			Columns = append(Columns, prof.Column)
			ColumnsT = append(ColumnsT, prof.ColumnT)
			switch {
			case prof.Column == row.PK:
				if g := dbSequenzer(row.Table); len(g) > 0 {
					ColumnsInsert = append(ColumnsInsert, prof.Column)
					BindsInsert = append(BindsInsert, g)
				}

				if prof.ColumnTrans[0:1] == "*" {
					row.PKType = prof.ColumnTrans[1:]
					row.PKPtr = "*"
				} else {
					row.PKType = prof.ColumnTrans
					row.PKPtr = ""
				}
			case strings.Contains(prof.Column, "_cr_date"):
				ColumnsInsert = append(ColumnsInsert, prof.Column)
				BindsInsert = append(BindsInsert, dbTimestamp)
			case strings.Contains(prof.Column, "_upd_date"):
				ColumnsUpdate = append(ColumnsUpdate, prof.Column)
				BindsUpdate = append(BindsUpdate, prof.Column+"="+dbTimestamp)
			case strings.Contains(prof.Column, "_upd_uid"):
				ColumnsUpdate = append(ColumnsUpdate, prof.Column)
				BindsUpdate = append(BindsUpdate, prof.Column+"='webSrv'")
			case strings.Contains(prof.Column, "_cr_uid"):
				ColumnsInsert = append(ColumnsInsert, prof.Column)
				BindsInsert = append(BindsInsert, "'webSrv'")
			default:
				inser++
				upder++
				ColumnsInsert = append(ColumnsInsert, prof.Column)
				BindsVarInsert = append(BindsVarInsert, prof.Column)
				ColumnsUpdate = append(ColumnsUpdate, prof.Column)
				BindsInsert = append(BindsInsert, BindVar+strconv.Itoa(inser))
				BindsUpdate = append(BindsUpdate, prof.Column+EqBindVar+strconv.Itoa(upder))
				pkBind = row.PK + EqBindVar + strconv.Itoa(upder+1)
			}
		}
		{

			if rows, err := db.Query(sqlfunctionparams, row.Specific); err != nil {
				log15.Crit("DBFehler", "query", err)
				return err
			} else {
				defer rows.Close()

				for rows.Next() {

					prof := Prof{}
					if err := rows.Scan(&prof.Column, &prof.ColumnTrans); err != nil {
						log15.Crit("DBFehler", "scann2", err)
						return err
					}
					profB = append(profB, prof)
				}
				if !flagger {
					BindsVarInsert = make([]string, len(profB))
					for i := 0; i < len(profB); i++ {
						BindsVarInsert[i] = "$" + strconv.Itoa(i+1)
					}
				}
			}

		}
		if (len(profA) > 0 && flagger) || (len(profA) > 0 && len(profB) > 0) {

			f, err := os.Create(gendir + "/" + row.Specific + ".go")
			if err != nil {
				log15.Crit("DBFehler", "gener", err)
				return err
			}

			defer f.Close()

			err = t.ExecuteTemplate(f, "stru.tmpl", struct {
				Package        string
				Flagger        bool
				FlaggerUdt     bool
				Table          string
				PK             string
				PKType         string
				PKPtr          string
				Timestamp      time.Time
				PKBind         string
				Cols           []Prof
				Columns        []string
				ColumnsT       []string
				ColumnsInsert  []string
				BindsVarInsert []string
				ColumnsUpdate  []string
				BindsInsert    []string
				BindsUpdate    []string
				Parameters     []Prof
			}{
				os.Args[1],
				flagger,
				row.TFlag,
				namer,
				row.PK,
				row.PKType,
				row.PKPtr,
				time.Now(),
				pkBind,
				profA,
				Columns,
				ColumnsT,
				ColumnsInsert,
				BindsVarInsert,
				ColumnsUpdate,
				BindsInsert,
				BindsUpdate,
				profB,
			})

			if err != nil {
				log15.Crit("DBFehler", "temp", err)
				return err
			}
			row.Worker = true
		} else {
			row.Worker = false
		}

	}
	return nil
}

func dbGen() (arr []*TabFlag, err error) {

	rows, err := db.Query(sqlalltabs)
	if err != nil {
		log15.Crit("DBFehler", "querymain", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var row TabFlag
		if err = rows.Scan(&row.Flag, &row.Table, &row.PK, &row.Parameter, &row.TFlag, &row.Specific); err != nil {
			log15.Crit("DBFehler", "scanstruct", err)
			return
		}
		if row.TFlag {
			tflag = true
		}
		arr = append(arr, &row)

	}
	return
}

func Generator() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(pwd)

	defer db.Close()
	arr, err := dbGen()
	if err != nil {
		return err
	}
	//	os.RemoveAll(gendir)
	os.Mkdir(gendir, 0777)
	os.Remove("mapper.go")
	os.Remove("mappertyp.go")
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title":  strings.Title,
		"joiner": func(x []string) string { return strings.Join(x, ",") },
	}
	// Create a new template and parse the letter into it.
	t, err1 := template.New("stru").Funcs(funcMap).ParseGlob("../templates/*")
	if err1 != nil {
		log15.Crit("DBFehler Temp.", "temp1", err1)
		return err1
	}

	for _, row := range arr {
		fmt.Println("tabelle", row.Table, row.Parameter, row.TFlag)
		if err = generateStru(t, row); err != nil {
			fmt.Println("fehler bei struct", err)
			return err
		}
		if row.Flag == 3 {
			row.Table = row.Parameter
		}

	}
	if os.Args[2] != "1" {

		if err := generateMap(t, arr); err != nil {
			fmt.Println("fehler bei map", err)
			return err
		}
		if err = os.Rename("mapper.tmp", "mapper.go"); err != nil {
			return err
		}
	}
	if tflag {

		if err := generateMapTyp(t, arr); err != nil {
			fmt.Println("fehler bei map", err)
		}
		if err = os.Rename("mappertyp.tmp", "mappertyp.go"); err != nil {
			return err
		}
	}
	return nil
}

func init() {
	connConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		log15.Crit("DB", "parse", err)
		os.Exit(1)
	}
	connConfig.LogLevel = pgx.LogLevelWarn
	if db, err = pgx.Connect(connConfig); err != nil {
		log15.Crit("DB", "connect", err)
		os.Exit(1)
	}
}
