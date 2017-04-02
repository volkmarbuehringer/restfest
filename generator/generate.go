package generator

import (
	"fmt"
	"io"
	"log"
	"os"
	"restfest/db"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func generateMap(t *template.Template, f io.Writer, arr []string) {
	if len(arr) > 0 {
		err := t.ExecuteTemplate(f, "mapper.tmpl", struct {
			Table     []string
			Timestamp time.Time
		}{arr, time.Now()})
		die(err)
	}

}

func generateStru(t *template.Template, table string, pk string) {

	f, err := os.Create("gener/" + table + ".go")
	die(err)
	defer f.Close()

	if rows, err := db.DB.Query(sqlallcols, table); err != nil {
		log.Fatal(err)
	} else {
		defer rows.Close()
		type Prof struct {
			Column      string
			ColumnTrans string
		}
		profA := []Prof{}
		ColumnsInsert := []string{}
		ColumnsUpdate := []string{}
		BindsInsert := []string{}
		BindsVarInsert := []string{}
		BindsUpdate := []string{}
		Columns := []string{}
		var pkBind string
		for inser, upder := 0, 0; rows.Next(); {

			prof := Prof{}
			if err := rows.Scan(&prof.Column, &prof.ColumnTrans); err != nil {
				log.Fatal(err)
			}
			profA = append(profA, prof)
			Columns = append(Columns, prof.Column)
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
		if len(profA) > 0 {
			err := t.ExecuteTemplate(f, "stru.tmpl", struct {
				Table          string
				PK             string
				Timestamp      time.Time
				PKBind         string
				Cols           []Prof
				Columns        []string
				ColumnsInsert  []string
				BindsVarInsert []string
				ColumnsUpdate  []string
				BindsInsert    []string
				BindsUpdate    []string
			}{
				table,
				pk,
				time.Now(),
				pkBind,
				profA,
				Columns,
				ColumnsInsert,
				BindsVarInsert,
				ColumnsUpdate,
				BindsInsert,
				BindsUpdate,
			})
			die(err)
		}

	}

}

func Generator() {
	os.RemoveAll("gener")
	os.Mkdir("gener", 0777)

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title":  strings.Title,
		"joiner": func(x []string) string { return strings.Join(x, ",") },
	}
	// Create a new template and parse the letter into it.
	t, err1 := template.New("stru").Funcs(funcMap).ParseGlob("templates/*")
	die(err1)

	rows, err := db.DB.Query(sqlalltabs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tabelle")
	defer rows.Close()
	arr := []string{}
	for rows.Next() {
		var table string
		var pk string
		if err := rows.Scan(&table, &pk); err != nil {
			log.Fatal(err)
		}
		arr = append(arr, table)
		fmt.Println("tabelle", table)
		generateStru(t, table, pk)
		die(err1)
	}

	f1, err1 := os.Create("gener/mapper.go")
	die(err1)
	defer f1.Close()

	generateMap(t, f1, arr)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
