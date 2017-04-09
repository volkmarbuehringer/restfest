
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-09 23:08:53.280832997 +0200 CEST
//code for table weburl

package gener


import (
	"restfest/db"
)

type WeburlParams struct {
    Length db.JSONNullInt64
    Offset db.JSONNullInt64
}


type Weburl struct {
    Id int64  `json:"id"`
    Url db.JSONString  `json:"url"`
    Zusatz db.JSONNullInt64  `json:"zusatz"`
    Created db.NullTime  `json:"created"`
    W_cr_date db.NullTime  `json:"w_cr_date"`
    W_upd_date db.NullTime  `json:"w_upd_date"`
    W_upd_uid db.JSONString  `json:"w_upd_uid"`
    W_cr_uid db.JSONString  `json:"w_cr_uid"`
    Addtime db.NullTime  `json:"addtime"`
    Addtime1 db.NullTime  `json:"addtime1"`
    Flag bool  `json:"flag"`
}



func SQLWeburl(tab string,flag db.SQLOper)[]interface{}{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
      tab,
      "id",
    }
    case db.GenInsert:
    return    []interface{}{ tab,
       "url,zusatz,created,w_cr_date,w_cr_uid,addtime,addtime1,flag"  ,
       "$1,$2,$3,current_timestamp,'webSrv',$4,$5,$6" ,
        x,
      }
    default:
      return    []interface{}{ tab,
        "url=$1,zusatz=$2,created=$3,w_upd_date=current_timestamp,w_upd_uid='webSrv',addtime=$4,addtime1=$5,flag=$6"  ,
        "id=$7",
      x,
        }

  }

}


func EmptyWeburl() interface{}{
    return new(Weburl)
}


func EmptyParamWeburl() interface{}{
    return new(WeburlParams)
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
				&struT.Addtime,
				&struT.Addtime1,
				&struT.Flag,
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
				&struT.Addtime,
				&struT.Addtime1,
				&struT.Flag,
			}, struT

}


func ROWInsertWeburl(inter interface{})[]interface{}{
  rt := inter.( *Weburl )
  return []interface{}{
	&rt.Url,
	&rt.Zusatz,
	&rt.Created,
	&rt.Addtime,
	&rt.Addtime1,
	&rt.Flag,  }

}


func ROWQueryWeburl(inter interface{})[]interface{}{
  rt := inter.( *WeburlParams )


  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
