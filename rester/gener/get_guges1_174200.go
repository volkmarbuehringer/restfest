
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-11 23:01:22.124752057 +0200 CEST
//code for table get_guges1

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

type dummyGet_guges1 time.Time

type Get_guges1Params struct {
    P_start *int32
    P_end *int32
    P_len *int32
    P_lala *string
}

type Get_guges1 struct {
    Id int32  `json:"id"`
    Agger ArWeburl  `json:"agger"`
    Texter *string  `json:"texter"`
    Zahler *int32  `json:"zahler"`
    Zahler2 *int32  `json:"zahler2"`
    Texter2 *string  `json:"texter2"`
}

type ArGet_guges1 []Get_guges1



var  Get_guges1SQL  =db.ColumnLists{
	 "id,agger,texter,zahler,zahler2,texter2",
	  "id,agger,texter,zahler,zahler2,texter2"  ,
		 "$1,$2,$3,$4,$5,$6" ,
	   "get_guges1" ,
		 "get_guges1=$7",
		"$1,$2,$3,$4"	 ,
						"id=$1,agger=$2,texter=$3,zahler=$4,zahler2=$5,texter2=$6"  ,

}


func (x *Get_guges1Params) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.P_start,
				&x.P_end,
				&x.P_len,
				&x.P_lala,
			}

}

func (x Get_guges1) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x Get_guges1Params) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["p_start"],
		mapper["p_end"],
		mapper["p_len"],
		mapper["p_lala"],
			}

}



func (t Get_guges1) Columns() []string {
	return strings.Split(Get_guges1SQL.All,",")
}


func (x *Get_guges1) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Agger,
				&x.Texter,
				&x.Zahler,
				&x.Zahler2,
				&x.Texter2,
			}

}

func (rt *Get_guges1)ROWInsert() db.InterPgx {
	  return db.InterPgx{
		
  }
}


func (rt *Get_guges1Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.P_start,
		&rt.P_end,
		&rt.P_len,
		&rt.P_lala,
}

}



func (t Get_guges1) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst Get_guges1Params)SQL(flag db.SQLOper)string{


  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
	Get_guges1SQL.All,
Get_guges1SQL.PK,
 Get_guges1SQL.BindsVarInsert,
		)


  }

	return ""

}




func ( dst Get_guges1)SQL(flag db.SQLOper)string{


  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		Get_guges1SQL.All,
Get_guges1SQL.PK,
Get_guges1SQL.BindsVarInsert,
		)


  }

	return ""

}



type IterGet_guges1 struct {
	db.BaseCopy
	Get_guges1 Get_guges1
}

type ArIterGet_guges1 struct {
	IterGet_guges1
	ArGet_guges1 ArGet_guges1
}


				



func (t *ArIterGet_guges1) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Get_guges1)
	for t.IterGet_guges1.Next(){
		t.ArGet_guges1 = append(t.ArGet_guges1,t.Get_guges1)
	}
 return len(t.ArGet_guges1)
}

func (t *IterGet_guges1) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Get_guges1)
return
}

func (t *IterGet_guges1) Value() db.PgxGener  {

 return &t.Get_guges1
}

func (t *IterGet_guges1) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Get_guges1)

}




			
