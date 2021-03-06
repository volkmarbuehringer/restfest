
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-09-22 21:57:58.236523101 +0200 CEST m=+0.647283267
//code for table tester

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

type dummyTester time.Time

type TesterParams struct {
    Length int
    Offset int
}

type Tester struct {
    Name *string  `json:"name"`
    Vorname *string  `json:"vorname"`
    Id int64  `json:"id"`
    Code *string  `json:"code"`
    Lala *string  `json:"lala"`
}

type ArTester []Tester



var  TesterSQL  =db.ColumnLists{
	 "name,vorname,id,code,lala",
	  "name,vorname,code,lala"  ,
		 "$1,$2,$3,$4" ,
	   "id" ,
		 "id=$5",
		"name,vorname,code,lala"	 ,
						"name=$1,vorname=$2,code=$3,lala=$4"  ,

}


func (x *TesterParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Tester) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x TesterParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Tester) Columns() []string {
	return strings.Split(TesterSQL.All,",")
}


func (x *Tester) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Name,
				&x.Vorname,
				&x.Id,
				&x.Code,
				&x.Lala,
			}

}

func (rt *Tester)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Name,	&rt.Vorname,	&rt.Code,	&rt.Lala,
    
  }
}


func (rt *TesterParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Tester) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst TesterParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  TesterSQL.All,
	     "tester"	,
      TesterSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	TesterSQL.All,
   "tester"	,
  		)


			
  }

	return ""

}




func ( dst Tester)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"tester"	,
			 TesterSQL.Inserts  ,
			 TesterSQL.BindsInsert ,
				TesterSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tester"	,
					TesterSQL.PK,
						TesterSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tester"	,
					TesterSQL.BindsUpdate  ,
								TesterSQL.PKUpdate,
					TesterSQL.All,)

			
  }

	return ""

}



type IterTester struct {
	db.BaseCopy
	Tester Tester
}

type ArIterTester struct {
	IterTester
	ArTester ArTester
}


				

type MapTester map[int64]Tester

type MapIterTester struct {
		IterTester
		MapTester  MapTester
}

				func (t *MapIterTester) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Tester)
				 t.MapTester=make(MapTester)
					for t.IterTester.Next(){
						t.MapTester[t.Tester.Id] = t.Tester
					}
				 return len(t.MapTester)
				}


	



func (t *ArIterTester) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Tester)
	for t.IterTester.Next(){
		t.ArTester = append(t.ArTester,t.Tester)
	}
 return len(t.ArTester)
}

func (t *IterTester) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Tester)
return
}

func (t *IterTester) Value() db.PgxGener  {

 return &t.Tester
}

func (t *IterTester) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Tester)

}




			
