
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-20 21:11:00.896032513 +0200 CEST
//code for table vweburl

package generrester



import (
	"restfest/db"
	"fmt"
	"time"

			
)

type dummyVweburl time.Time

type VweburlParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Vweburl struct {
    Id *int32  `json:"id"`
    Url *string  `json:"url"`
    Zusatz *int64  `json:"zusatz"`
    Created *time.Time  `json:"created"`
    W_cr_date *time.Time  `json:"w_cr_date"`
    W_upd_date *time.Time  `json:"w_upd_date"`
    W_upd_uid *string  `json:"w_upd_uid"`
    W_cr_uid *string  `json:"w_cr_uid"`
}

type ArVweburl []Vweburl

			

func ScannerVweburlI(struT *Vweburl)  db.InterPgx  {

	return []interface{}{
				&struT.Id,
				&struT.Url,
				&struT.Zusatz,
				&struT.Created,
				&struT.W_cr_date,
				&struT.W_upd_date,
				&struT.W_upd_uid,
				&struT.W_cr_uid,
			}

}


func SQLVweburl(flag db.SQLOper)string{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "vweburl"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "vweburl"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"vweburl"	,
			 "url,zusatz,created,w_cr_date,w_cr_uid"  ,
			 "$1,$2,$3,current_timestamp,'webSrv'" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"vweburl"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"vweburl"	,
						"url=$1,zusatz=$2,created=$3,w_upd_date=current_timestamp,w_upd_uid='webSrv'"  ,
						"id=$4",
					x)

			
  }

	return ""

}


func EmptyVweburl() interface{}{
    return new(Vweburl)
}


func EmptyParamVweburl() interface{}{

 rt := new(VweburlParams)

	
		rt.Length=100
	

    return rt
}

func ScannerVweburl()(  db.InterPgx , interface{}){
	struT := new(Vweburl)

return  ScannerVweburlI(struT), struT

}


func ROWInsertVweburl(inter interface{}) db.InterPgx {
					
  rt := inter.( *Vweburl )
	
  return []interface{}{
				
	&rt.Url,
	&rt.Zusatz,
	&rt.Created,

  }

}


func ROWQueryVweburl(inter interface{}) db.InterPgx {
  rt := inter.( *VweburlParams )

  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
