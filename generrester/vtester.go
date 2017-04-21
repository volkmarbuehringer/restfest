
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-21 22:14:12.715010582 +0200 CEST
//code for table vtester

package generrester



import (
	"restfest/db"
	"fmt"
	"time"

			
)

type dummyVtester time.Time

type VtesterParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Vtester struct {
    Name *string  `json:"name"`
    Vorname *string  `json:"vorname"`
    Id *int64  `json:"id"`
    Code *string  `json:"code"`
    Lala *string  `json:"lala"`
}

type ArVtester []Vtester

func (x *Vtester) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Name,
				&x.Vorname,
				&x.Id,
				&x.Code,
				&x.Lala,
			}

}

func (rt *Vtester)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Vorname,	&rt.Id,	&rt.Code,	&rt.Lala,
    
  }
}


func (rt *VtesterParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}


			


func SQLVtester(flag db.SQLOper)string{
  x :=   "name,vorname,id,code,lala"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "vtester"	,
      "name" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "vtester"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"vtester"	,
			 "vorname,id,code,lala"  ,
			 "$1,$2,$3,$4" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"vtester"	,
						"name",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"vtester"	,
						"vorname=$1,id=$2,code=$3,lala=$4"  ,
						"name=$5",
					x)

			
  }

	return ""

}
