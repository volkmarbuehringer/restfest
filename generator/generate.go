package generator

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/jackc/pgx"
	"gopkg.in/inconshreveable/log15.v2"
)

type TabFlag struct {
	Table     string
	Flag      int
	PK        string
	Parameter string
}

var db *pgx.Conn

func generateMap(t *template.Template, f io.Writer, arr []*TabFlag) error {
	if len(arr) > 0 {
		if err := t.ExecuteTemplate(f, "mapper.tmpl", struct {
			Table     []*TabFlag
			Timestamp time.Time
		}{arr, time.Now()}); err != nil {
			log15.Crit("DBFehler", "map", err)
			return err
		}

	}
	return nil
}

func generateStru(t *template.Template, table string, pk string, parameter string) error {

	namer := table
	flagger := false
	type Prof struct {
		Column      string
		ColumnTrans string
		ColumnT     string
	}
	profA := []Prof{}
	profB := []Prof{}
	if len(parameter) > 0 {
		namer = parameter

	} else {
		profB = []Prof{Prof{"length", "int", ""}, Prof{"offset", "int", ""}}
		flagger = true
	}
	f, err := os.Create("gener/" + namer + ".go")
	if err != nil {
		log15.Crit("DBFehler", "gener", err)
		return err
	}

	defer f.Close()

	if rows, err := db.Query(sqlallcols, table); err != nil {
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
				log15.Crit("DBFehler", "scan", err)
				return err
			}
			profA = append(profA, prof)
			Columns = append(Columns, prof.Column)
			ColumnsT = append(ColumnsT, prof.ColumnT)
			switch {
			case prof.Column == pk:
				if g := dbSequenzer(table); len(g) > 0 {
					ColumnsInsert = append(ColumnsInsert, prof.Column)
					BindsInsert = append(BindsInsert, g)
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
				pkBind = pk + EqBindVar + strconv.Itoa(upder+1)
			}
		}
		{

			if rows, err := db.Query(sqlfunctionparams, parameter); err != nil {
				log15.Crit("DBFehler", "query", err)
				return err
			} else {
				defer rows.Close()

				for rows.Next() {
					prof := Prof{}
					if err := rows.Scan(&prof.Column, &prof.ColumnTrans); err != nil {
						log15.Crit("DBFehler", "scann", err)
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
		if len(profA) > 0 {
			err := t.ExecuteTemplate(f, "stru.tmpl", struct {
				Flagger        bool
				Table          string
				PK             string
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
				flagger,
				namer,
				pk,
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
		if err = rows.Scan(&row.Flag, &row.Table, &row.PK, &row.Parameter); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return
		}

		arr = append(arr, &row)

	}
	return
}

func Generator() error {
	defer db.Close()
	arr, err := dbGen()
	os.RemoveAll("gener")
	os.Mkdir("gener", 0777)
	os.Remove("mapper.go")
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title":  strings.Title,
		"joiner": func(x []string) string { return strings.Join(x, ",") },
	}
	// Create a new template and parse the letter into it.
	t, err1 := template.New("stru").Funcs(funcMap).ParseGlob("templates/*")
	if err1 != nil {
		log15.Crit("DBFehler", "temp1", err1)
		return err1
	}

	for _, row := range arr {
		fmt.Println("tabelle", row.Table, row.Parameter)
		if err = generateStru(t, row.Table, row.PK, row.Parameter); err != nil {
			return err
		}
		if row.Flag == 3 {
			row.Table = row.Parameter
		}

	}

	f1, err1 := os.Create("mapper.go")
	if err1 != nil {
		log15.Crit("DBFehler", "create", err1)
		return err
	}

	defer f1.Close()

	return generateMap(t, f1, arr)

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
