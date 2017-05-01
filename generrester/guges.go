
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-01 12:49:44.413960283 +0200 CEST
//code for table guges

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"strings"
		"encoding/json"
	"github.com/jackc/pgx"
			
)

type dummyGuges time.Time

type GugesParams struct {
    Length int
    Offset int
}

type Guges struct {
    Id int32  `json:"id"`
    A Weburl  `json:"a"`
}

type ArGuges []Guges



var  GugesSQL  =db.ColumnLists{
	 "id,a",
	  "a"  ,
		 "$1" ,
	   "id" ,
		 "id=$2",
		"a"	 ,
						"a=$1"  ,

}


func (x *GugesParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Guges) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x GugesParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Guges) Columns() []string {
	return strings.Split(GugesSQL.All,",")
}


func (x *Guges) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.A,
			}

}

func (rt *Guges)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.A,
    
  }
}


func (rt *GugesParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Guges) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst GugesParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  GugesSQL.All,
	     "guges"	,
      GugesSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	GugesSQL.All,
   "guges"	,
  		)


			
  }

	return ""

}




func ( dst Guges)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"guges"	,
			 GugesSQL.Inserts  ,
			 GugesSQL.BindsInsert ,
				GugesSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"guges"	,
					GugesSQL.PK,
						GugesSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"guges"	,
					GugesSQL.BindsUpdate  ,
								GugesSQL.PKUpdate,
					GugesSQL.All,)

			
  }

	return ""

}



type IterGuges struct {
	db.BaseCopy
	Guges Guges
}

type ArIterGuges struct {
	IterGuges
	ArGuges ArGuges
}


				

type MapGuges map[int32]Guges

type MapIterGuges struct {
		IterGuges
		MapGuges  MapGuges
}

				func (t *MapIterGuges) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Guges)
				 t.MapGuges=make(MapGuges)
					for t.IterGuges.Next(){
						t.MapGuges[t.Guges.Id] = t.Guges
					}
				 return len(t.MapGuges)
				}


	



func (t *ArIterGuges) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Guges)
	for t.IterGuges.Next(){
		t.ArGuges = append(t.ArGuges,t.Guges)
	}
 return len(t.ArGuges)
}

func (t *IterGuges) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Guges)
return
}

func (t *IterGuges) Value() db.PgxGener  {

 return &t.Guges
}

func (t *IterGuges) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Guges)

}




			
