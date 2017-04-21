
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-21 19:51:48.932060942 +0200 CEST
//code for table get_all_boo_173824

package generrester



import (
	"restfest/db"
	"fmt"
	"time"

			
)

type dummyGet_all_boo_173824 time.Time

type Get_all_boo_173824Params struct {
    P_start *int32 `schema:"p_start"`
    P_end *int32 `schema:"p_end"`
    P_len *int32 `schema:"p_len"`
    P_lala *string `schema:"p_lala"`
}


type Get_all_boo_173824 struct {
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

type ArGet_all_boo_173824 []Get_all_boo_173824

func (x *Get_all_boo_173824) Scanner() db.InterPgx {
	return []interface{}{
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

func (rt *Get_all_boo_173824)ROWInsert() db.InterPgx {
	  return []interface{}{
				
  }
}


func (rt *Get_all_boo_173824Params)ROWInsert() db.InterPgx {

  return []interface{}{
		&rt.P_start,
		&rt.P_end,
		&rt.P_len,
		&rt.P_lala,
}

}


			


func SQLGet_all_boo_173824(flag db.SQLOper)string{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test"

  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
 "get_all_boo" ,
 "$1,$2,$3,$4"	,
		)


  }

	return ""

}


func EmptyGet_all_boo_173824() db.PgxGener{
    return new(Get_all_boo_173824)
}

func EmptyInsGet_all_boo_173824() db.PgxGenerIns{
    return new(Get_all_boo_173824)
}

func EmptyParamGet_all_boo_173824() db.PgxGenerIns{

 rt := new(Get_all_boo_173824Params)

	

    return rt
}
