
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-03 21:55:29.022895026 +0200 CEST
//code for table weburl

package gener


import (
	"restfest/db"
)

type Weburl struct {
    Id int64  `json:"id"`
    Url db.JSONNullString  `json:"url"`
    Zusatz db.JSONNullInt64  `json:"zusatz"`
    Created db.JSONNullString  `json:"created"`
    W_cr_date db.JSONNullString  `json:"w_cr_date"`
    W_upd_date db.JSONNullString  `json:"w_upd_date"`
    W_upd_uid db.JSONNullString  `json:"w_upd_uid"`
    W_cr_uid db.JSONNullString  `json:"w_cr_uid"`
}



func SQLWeburl(tab string,flag db.SQLOper)[]interface{}{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
      tab,
      "id",
    }
    case db.GenInsert:
    return    []interface{}{ tab,
       "url,zusatz,created,w_cr_date,w_cr_uid"  ,
       "$1,$2,$3,current_timestamp,'webSrv'" ,
        x,
      }
    default:
      return    []interface{}{ tab,
        "url=$1,zusatz=$2,created=$3,w_upd_date=current_timestamp,w_upd_uid='webSrv'"  ,
        "id=$4",
      x,
        }

  }

}


func EmptyWeburl() interface{}{
    return new(Weburl)
}


func ScannerWeburl()( []interface{}, interface{}){
	struT := new(Weburl)

return  []interface{}{
				&struT.Id,
				&struT.Url,
				&struT.Zusatz,
				&struT.Created,
				&struT.W_cr_date,
				&struT.W_upd_date,
				&struT.W_upd_uid,
				&struT.W_cr_uid,
			}, struT

}


func ScannerTWeburl()( []interface{},*Weburl){
	struT := new(Weburl)

return  []interface{}{
				&struT.Id,
				&struT.Url,
				&struT.Zusatz,
				&struT.Created,
				&struT.W_cr_date,
				&struT.W_upd_date,
				&struT.W_upd_uid,
				&struT.W_cr_uid,
			}, struT

}


func ROWInsertWeburl(inter interface{})[]interface{}{
  rt := inter.( *Weburl )
  return []interface{}{
	&rt.Url,
	&rt.Zusatz,
	&rt.Created,  }

}
