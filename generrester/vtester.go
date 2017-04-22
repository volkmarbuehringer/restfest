
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 11:19:25.149014588 +0200 CEST
//code for table vtester

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
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


func (dst *ArVtester) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Vtester)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
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
