
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-02 23:01:52.439870357 +0200 CEST
//code for table tedas

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

type dummyTedas time.Time

type TedasParams struct {
    Length int
    Offset int
}

type Tedas struct {
    Id int32  `json:"id"`
    Agg []int32  `json:"agg"`
}

type ArTedas []Tedas



var  TedasSQL  =db.ColumnLists{
	 "id,agg",
	  "agg"  ,
		 "$1" ,
	   "id" ,
		 "id=$2",
		"agg"	 ,
						"agg=$1"  ,

}


func (x *TedasParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Tedas) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x TedasParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Tedas) Columns() []string {
	return strings.Split(TedasSQL.All,",")
}


func (x *Tedas) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Agg,
			}

}

func (rt *Tedas)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Agg,
    
  }
}


func (rt *TedasParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Tedas) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst TedasParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  TedasSQL.All,
	     "tedas"	,
      TedasSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	TedasSQL.All,
   "tedas"	,
  		)


			
  }

	return ""

}




func ( dst Tedas)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"tedas"	,
			 TedasSQL.Inserts  ,
			 TedasSQL.BindsInsert ,
				TedasSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tedas"	,
					TedasSQL.PK,
						TedasSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tedas"	,
					TedasSQL.BindsUpdate  ,
								TedasSQL.PKUpdate,
					TedasSQL.All,)

			
  }

	return ""

}



type IterTedas struct {
	db.BaseCopy
	Tedas Tedas
}

type ArIterTedas struct {
	IterTedas
	ArTedas ArTedas
}


				

type MapTedas map[int32]Tedas

type MapIterTedas struct {
		IterTedas
		MapTedas  MapTedas
}

				func (t *MapIterTedas) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Tedas)
				 t.MapTedas=make(MapTedas)
					for t.IterTedas.Next(){
						t.MapTedas[t.Tedas.Id] = t.Tedas
					}
				 return len(t.MapTedas)
				}


	



func (t *ArIterTedas) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Tedas)
	for t.IterTedas.Next(){
		t.ArTedas = append(t.ArTedas,t.Tedas)
	}
 return len(t.ArTedas)
}

func (t *IterTedas) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Tedas)
return
}

func (t *IterTedas) Value() db.PgxGener  {

 return &t.Tedas
}

func (t *IterTedas) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Tedas)

}




			
