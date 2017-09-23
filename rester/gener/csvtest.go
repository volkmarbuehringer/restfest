
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-09-22 21:57:57.951352548 +0200 CEST m=+0.362112463
//code for table csvtest

package gener



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"strings"
		"encoding/json"
	"github.com/jackc/pgx"
			
)

type dummyCsvtest time.Time

type CsvtestParams struct {
    Length int
    Offset int
}

type Csvtest struct {
    Year int32  `json:"year"`
    Anzsic string  `json:"anzsic"`
    Area string  `json:"area"`
    Geo *int32  `json:"geo"`
    Ec *int32  `json:"ec"`
    Id int32  `json:"id"`
}

type ArCsvtest []Csvtest



var  CsvtestSQL  =db.ColumnLists{
	 "year,anzsic,area,geo,ec,id",
	  "year,anzsic,area,geo,ec"  ,
		 "$1,$2,$3,$4,$5" ,
	   "id" ,
		 "id=$6",
		"year,anzsic,area,geo,ec"	 ,
						"year=$1,anzsic=$2,area=$3,geo=$4,ec=$5"  ,

}


func (x *CsvtestParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Csvtest) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x CsvtestParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Csvtest) Columns() []string {
	return strings.Split(CsvtestSQL.All,",")
}


func (x *Csvtest) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Year,
				&x.Anzsic,
				&x.Area,
				&x.Geo,
				&x.Ec,
				&x.Id,
			}

}

func (rt *Csvtest)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Year,	&rt.Anzsic,	&rt.Area,	&rt.Geo,	&rt.Ec,
    
  }
}


func (rt *CsvtestParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Csvtest) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst CsvtestParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  CsvtestSQL.All,
	     "csvtest"	,
      CsvtestSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	CsvtestSQL.All,
   "csvtest"	,
  		)


			
  }

	return ""

}




func ( dst Csvtest)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"csvtest"	,
			 CsvtestSQL.Inserts  ,
			 CsvtestSQL.BindsInsert ,
				CsvtestSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"csvtest"	,
					CsvtestSQL.PK,
						CsvtestSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"csvtest"	,
					CsvtestSQL.BindsUpdate  ,
								CsvtestSQL.PKUpdate,
					CsvtestSQL.All,)

			
  }

	return ""

}



type IterCsvtest struct {
	db.BaseCopy
	Csvtest Csvtest
}

type ArIterCsvtest struct {
	IterCsvtest
	ArCsvtest ArCsvtest
}


				

type MapCsvtest map[int32]Csvtest

type MapIterCsvtest struct {
		IterCsvtest
		MapCsvtest  MapCsvtest
}

				func (t *MapIterCsvtest) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Csvtest)
				 t.MapCsvtest=make(MapCsvtest)
					for t.IterCsvtest.Next(){
						t.MapCsvtest[t.Csvtest.Id] = t.Csvtest
					}
				 return len(t.MapCsvtest)
				}


	



func (t *ArIterCsvtest) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Csvtest)
	for t.IterCsvtest.Next(){
		t.ArCsvtest = append(t.ArCsvtest,t.Csvtest)
	}
 return len(t.ArCsvtest)
}

func (t *IterCsvtest) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Csvtest)
return
}

func (t *IterCsvtest) Value() db.PgxGener  {

 return &t.Csvtest
}

func (t *IterCsvtest) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Csvtest)

}




			