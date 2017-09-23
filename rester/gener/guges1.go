
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-09-22 21:57:58.001439203 +0200 CEST m=+0.412199308
//code for table guges1

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

type dummyGuges1 time.Time

type Guges1Params struct {
    Length int
    Offset int
}

type Guges1 struct {
    Id int32  `json:"id"`
    Agger ArWeburl  `json:"agger"`
    Texter *string  `json:"texter"`
    Zahler *int32  `json:"zahler"`
    Zahler2 *int32  `json:"zahler2"`
    Texter2 *string  `json:"texter2"`
}

type ArGuges1 []Guges1



var  Guges1SQL  =db.ColumnLists{
	 "id,agger,texter,zahler,zahler2,texter2",
	  "agger,texter,zahler,zahler2,texter2"  ,
		 "$1,$2,$3,$4,$5" ,
	   "id" ,
		 "id=$6",
		"agger,texter,zahler,zahler2,texter2"	 ,
						"agger=$1,texter=$2,zahler=$3,zahler2=$4,texter2=$5"  ,

}


func (x *Guges1Params) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Guges1) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x Guges1Params) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Guges1) Columns() []string {
	return strings.Split(Guges1SQL.All,",")
}


func (x *Guges1) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Agger,
				&x.Texter,
				&x.Zahler,
				&x.Zahler2,
				&x.Texter2,
			}

}

func (rt *Guges1)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Agger,	&rt.Texter,	&rt.Zahler,	&rt.Zahler2,	&rt.Texter2,
    
  }
}


func (rt *Guges1Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Guges1) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst Guges1Params)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  Guges1SQL.All,
	     "guges1"	,
      Guges1SQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	Guges1SQL.All,
   "guges1"	,
  		)


			
  }

	return ""

}




func ( dst Guges1)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"guges1"	,
			 Guges1SQL.Inserts  ,
			 Guges1SQL.BindsInsert ,
				Guges1SQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"guges1"	,
					Guges1SQL.PK,
						Guges1SQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"guges1"	,
					Guges1SQL.BindsUpdate  ,
								Guges1SQL.PKUpdate,
					Guges1SQL.All,)

			
  }

	return ""

}



type IterGuges1 struct {
	db.BaseCopy
	Guges1 Guges1
}

type ArIterGuges1 struct {
	IterGuges1
	ArGuges1 ArGuges1
}


				

type MapGuges1 map[int32]Guges1

type MapIterGuges1 struct {
		IterGuges1
		MapGuges1  MapGuges1
}

				func (t *MapIterGuges1) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Guges1)
				 t.MapGuges1=make(MapGuges1)
					for t.IterGuges1.Next(){
						t.MapGuges1[t.Guges1.Id] = t.Guges1
					}
				 return len(t.MapGuges1)
				}


	



func (t *ArIterGuges1) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Guges1)
	for t.IterGuges1.Next(){
		t.ArGuges1 = append(t.ArGuges1,t.Guges1)
	}
 return len(t.ArGuges1)
}

func (t *IterGuges1) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Guges1)
return
}

func (t *IterGuges1) Value() db.PgxGener  {

 return &t.Guges1
}

func (t *IterGuges1) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Guges1)

}




			
