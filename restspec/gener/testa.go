
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-06-06 22:13:10.867494987 +0200 CEST
//code for table testa

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

type dummyTesta time.Time

type TestaParams struct {
    Length int
    Offset int
}

type Testa struct {
    Id int32  `json:"id"`
    T *map[string]interface{}  `json:"t"`
}

type ArTesta []Testa



var  TestaSQL  =db.ColumnLists{
	 "id,t",
	  "t"  ,
		 "$1" ,
	   "id" ,
		 "id=$2",
		"t"	 ,
						"t=$1"  ,

}


func (x *TestaParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Testa) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x TestaParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Testa) Columns() []string {
	return strings.Split(TestaSQL.All,",")
}


func (x *Testa) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.T,
			}

}

func (rt *Testa)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.T,
    
  }
}


func (rt *TestaParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Testa) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst TestaParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  TestaSQL.All,
	     "testa"	,
      TestaSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	TestaSQL.All,
   "testa"	,
  		)


			
  }

	return ""

}




func ( dst Testa)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"testa"	,
			 TestaSQL.Inserts  ,
			 TestaSQL.BindsInsert ,
				TestaSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"testa"	,
					TestaSQL.PK,
						TestaSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"testa"	,
					TestaSQL.BindsUpdate  ,
								TestaSQL.PKUpdate,
					TestaSQL.All,)

			
  }

	return ""

}



type IterTesta struct {
	db.BaseCopy
	Testa Testa
}

type ArIterTesta struct {
	IterTesta
	ArTesta ArTesta
}


				

type MapTesta map[int32]Testa

type MapIterTesta struct {
		IterTesta
		MapTesta  MapTesta
}

				func (t *MapIterTesta) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Testa)
				 t.MapTesta=make(MapTesta)
					for t.IterTesta.Next(){
						t.MapTesta[t.Testa.Id] = t.Testa
					}
				 return len(t.MapTesta)
				}


	



func (t *ArIterTesta) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Testa)
	for t.IterTesta.Next(){
		t.ArTesta = append(t.ArTesta,t.Testa)
	}
 return len(t.ArTesta)
}

func (t *IterTesta) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Testa)
return
}

func (t *IterTesta) Value() db.PgxGener  {

 return &t.Testa
}

func (t *IterTesta) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Testa)

}




			
