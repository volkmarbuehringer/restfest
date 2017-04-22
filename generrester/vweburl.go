
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 11:19:25.172490468 +0200 CEST
//code for table vweburl

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
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

func (x *Vweburl) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Url,
				&x.Zusatz,
				&x.Created,
				&x.W_cr_date,
				&x.W_upd_date,
				&x.W_upd_uid,
				&x.W_cr_uid,
			}

}

func (rt *Vweburl)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Url,	&rt.Zusatz,	&rt.Created,
    
  }
}


func (rt *VweburlParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}


func (dst *ArVweburl) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Vweburl)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
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
