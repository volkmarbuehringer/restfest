
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-16 21:19:11.813308543 +0200 CEST
//code for table get_weburl_173825

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

type dummyGet_weburl_173825 time.Time

type Get_weburl_173825Params struct {
    P_start *int32
    P_end *int32
    P_len *int32
    P_lala *string
}

type Get_weburl_173825 struct {
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

type ArGet_weburl_173825 []Get_weburl_173825



var  Get_weburl_173825SQL  =db.ColumnLists{
	 "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test",
	  "id,url,zusatz,created,w_cr_date,w_cr_uid,addtime,addtime1,flag,test"  ,
		 "$1,$2,$3,$4,current_timestamp,'webSrv',$5,$6,$7,$8" ,
	   "get_weburl" ,
		 "get_weburl=$9",
		"$1,$2,$3,$4"	 ,
						"id=$1,url=$2,zusatz=$3,created=$4,w_upd_date=current_timestamp,w_upd_uid='webSrv',addtime=$5,addtime1=$6,flag=$7,test=$8"  ,

}


func (x *Get_weburl_173825Params) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.P_start,
				&x.P_end,
				&x.P_len,
				&x.P_lala,
			}

}

func (x Get_weburl_173825) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x Get_weburl_173825Params) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["p_start"],
		mapper["p_end"],
		mapper["p_len"],
		mapper["p_lala"],
			}

}



func (t Get_weburl_173825) Columns() []string {
	return strings.Split(Get_weburl_173825SQL.All,",")
}


func (x *Get_weburl_173825) Scanner() db.InterPgx {
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

func (rt *Get_weburl_173825)ROWInsert() db.InterPgx {
	  return db.InterPgx{
		
  }
}


func (rt *Get_weburl_173825Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.P_start,
		&rt.P_end,
		&rt.P_len,
		&rt.P_lala,
}

}



func (t Get_weburl_173825) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst Get_weburl_173825Params)SQL(flag db.SQLOper)string{


  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
	Get_weburl_173825SQL.All,
Get_weburl_173825SQL.PK,
 Get_weburl_173825SQL.BindsVarInsert,
		)


  }

	return ""

}




func ( dst Get_weburl_173825)SQL(flag db.SQLOper)string{


  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		Get_weburl_173825SQL.All,
Get_weburl_173825SQL.PK,
Get_weburl_173825SQL.BindsVarInsert,
		)


  }

	return ""

}



type IterGet_weburl_173825 struct {
	db.BaseCopy
	Get_weburl_173825 Get_weburl_173825
}

type ArIterGet_weburl_173825 struct {
	IterGet_weburl_173825
	ArGet_weburl_173825 ArGet_weburl_173825
}


				



func (t *ArIterGet_weburl_173825) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Get_weburl_173825)
	for t.IterGet_weburl_173825.Next(){
		t.ArGet_weburl_173825 = append(t.ArGet_weburl_173825,t.Get_weburl_173825)
	}
 return len(t.ArGet_weburl_173825)
}

func (t *IterGet_weburl_173825) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Get_weburl_173825)
return
}

func (t *IterGet_weburl_173825) Value() db.PgxGener  {

 return &t.Get_weburl_173825
}

func (t *IterGet_weburl_173825) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Get_weburl_173825)

}




			
