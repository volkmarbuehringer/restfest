
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-30 10:13:09.826442338 +0200 CEST
//code for table tester

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
		"encoding/json"
	"github.com/jackc/pgx"
			
)

type dummyTester time.Time

type TesterParams struct {
    Length int
    Offset int
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

type Tester struct {
    Name *string  `json:"name"`
    Vorname *string  `json:"vorname"`
    Id int64  `json:"id"`
    Code *string  `json:"code"`
    Lala *string  `json:"lala"`
}

type ArTester []Tester



func (t Tester) Columns() []string {
	return []string{	 "name" ,	 "vorname" ,	 "id" ,	 "code" ,	 "lala" ,
			}
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


	



type IterTester struct {
	db.BaseCopy
	Tester Tester
}

type ArIterTester struct {
	IterTester
	ArTester ArTester
}

func (t Tester) Writer(w io.Writer)  error {
if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
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




			


func ( dst TesterParams)SQL(flag db.SQLOper)string{
  x :=   "name,vorname,id,code,lala"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "tester"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "tester"	,
  		)


			
  }

	return ""

}




func ( dst Tester)SQL(flag db.SQLOper)string{
  x :=   "name,vorname,id,code,lala"

  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"tester"	,
			 "name,vorname,code,lala"  ,
			 "$1,$2,$3,$4" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tester"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tester"	,
						"name=$1,vorname=$2,code=$3,lala=$4"  ,
						"id=$5",
					x)

			
  }

	return ""

}
