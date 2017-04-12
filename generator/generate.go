package generator

import (
	"fmt"
	"io"
	"os"
	"restfest/db"
	"strconv"
	"strings"
	"text/template"
	"time"

	"gopkg.in/inconshreveable/log15.v2"
)

type TabFlag struct {
	Table string
	Flag  int
}

func generateMap(t *template.Template, f io.Writer, arr []TabFlag) {
	if len(arr) > 0 {
		err := t.ExecuteTemplate(f, "mapper.tmpl", struct {
			Table     []TabFlag
			Timestamp time.Time
		}{arr, time.Now()})
		if err != nil {
			log15.Crit("DBFehler", "map", err)
			return
		}

	}

}

func generateStru(t *template.Template, table string, pk string, parameter string) {

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
		profB = []Prof{Prof{"length", "int64", ""}, Prof{"offset", "int64", ""}}
		flagger = true
	}
	f, err := os.Create("gener/" + namer + ".go")
	if err != nil {
		log15.Crit("DBFehler", "gener", err)
		return
	}

	defer f.Close()

	if rows, err := db.DBx.Query(sqlallcols, table); err != nil {
		log15.Crit("DBFehler", "query", err)
		return
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
				return
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

			if rows, err := db.DBx.Query(sqlfunctionparams, parameter); err != nil {
				log15.Crit("DBFehler", "query", err)
				return
			} else {
				defer rows.Close()

				for rows.Next() {
					prof := Prof{}
					if err := rows.Scan(&prof.Column, &prof.ColumnTrans); err != nil {
						log15.Crit("DBFehler", "scann", err)
						return
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
				return
			}

		}

	}

}

func Generator() {
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
		return
	}

	rows, err := db.DBx.Query(sqlalltabs)
	if err != nil {
		log15.Crit("DBFehler", "query", err)
		return
	}

	defer rows.Close()

	arr := []TabFlag{}

	for rows.Next() {
		var table string
		var pk string
		var flag int
		var parameter string
		if err = rows.Scan(&flag, &table, &pk, &parameter); err != nil {
			log15.Crit("DBFehler", "scan", err)
			return
		}

		fmt.Println("tabelle", table, parameter)
		generateStru(t, table, pk, parameter)
		if flag == 3 {
			table = parameter
		}
		flags := TabFlag{table, flag}
		arr = append(arr, flags)

	}

	f1, err1 := os.Create("mapper.go")
	if err1 != nil {
		log15.Crit("DBFehler", "create", err1)
		return
	}

	defer f1.Close()

	generateMap(t, f1, arr)
}
