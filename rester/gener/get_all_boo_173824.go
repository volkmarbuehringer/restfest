
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-11 23:01:22.497525641 +0200 CEST
//code for table get_all_boo

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

type dummyGet_all_boo time.Time

type Get_all_booParams struct {
    P_start *int32
    P_end *int32
    P_len *int32
    P_lala *string
}

type Get_all_boo struct {
    Id int32  `json:"id"`
    Url *string  `json:"url"`
    Zusatz *int64  `json:"zusatz"`
    Created *time.Time  `json:"created"`
    W_cr_date *time.Time  `json:"w_cr_date"`
    W_upd_date *time.Time  `json:"w_upd_date"`
    W_upd_uid *string  `json:"w_upd_uid"`
    W_cr_uid *string  `json:"w_cr_uid"`
    Addtime *time.Time  `json:"addtime"`
    Addtime1 *time.Time  `json:"addtime1"`
    Flag bool  `json:"flag"`
    Test *float64  `json:"test"`
}

type ArGet_all_boo []Get_all_boo



var  Get_all_booSQL  =db.ColumnLists{
	 "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test",
	  "id,url,zusatz,created,w_cr_date,w_cr_uid,addtime,addtime1,flag,test"  ,
		 "$1,$2,$3,$4,current_timestamp,'webSrv',$5,$6,$7,$8" ,
	   "get_all_boo" ,
		 "get_all_boo=$9",
		"$1,$2,$3,$4"	 ,
						"id=$1,url=$2,zusatz=$3,created=$4,w_upd_date=current_timestamp,w_upd_uid='webSrv',addtime=$5,addtime1=$6,flag=$7,test=$8"  ,

}


func (x *Get_all_booParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.P_start,
				&x.P_end,
				&x.P_len,
				&x.P_lala,
			}

}

func (x Get_all_boo) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x Get_all_booParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["p_start"],
		mapper["p_end"],
		mapper["p_len"],
		mapper["p_lala"],
			}

}



func (t Get_all_boo) Columns() []string {
	return strings.Split(Get_all_booSQL.All,",")
}


func (x *Get_all_boo) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Url,
				&x.Zusatz,
				&x.Created,
				&x.W_cr_date,
				&x.W_upd_date,
				&x.W_upd_uid,
				&x.W_cr_uid,
				&x.Addtime,
				&x.Addtime1,
				&x.Flag,
				&x.Test,
			}

}

func (rt *Get_all_boo)ROWInsert() db.InterPgx {
	  return db.InterPgx{
		
  }
}


func (rt *Get_all_booParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.P_start,
		&rt.P_end,
		&rt.P_len,
		&rt.P_lala,
}

}



func (t Get_all_boo) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst Get_all_booParams)SQL(flag db.SQLOper)string{


  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
	Get_all_booSQL.All,
Get_all_booSQL.PK,
 Get_all_booSQL.BindsVarInsert,
		)


  }

	return ""

}




func ( dst Get_all_boo)SQL(flag db.SQLOper)string{


  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		Get_all_booSQL.All,
Get_all_booSQL.PK,
Get_all_booSQL.BindsVarInsert,
		)


  }

	return ""

}



type IterGet_all_boo struct {
	db.BaseCopy
	Get_all_boo Get_all_boo
}

type ArIterGet_all_boo struct {
	IterGet_all_boo
	ArGet_all_boo ArGet_all_boo
}


				



func (t *ArIterGet_all_boo) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Get_all_boo)
	for t.IterGet_all_boo.Next(){
		t.ArGet_all_boo = append(t.ArGet_all_boo,t.Get_all_boo)
	}
 return len(t.ArGet_all_boo)
}

func (t *IterGet_all_boo) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Get_all_boo)
return
}

func (t *IterGet_all_boo) Value() db.PgxGener  {

 return &t.Get_all_boo
}

func (t *IterGet_all_boo) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Get_all_boo)

}




			
