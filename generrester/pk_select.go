
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-21 22:14:12.57314495 +0200 CEST
//code for table pk_select

package generrester



import (
	"restfest/db"
	"fmt"
	"time"

			
)

type dummyPk_select time.Time

type Pk_selectParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Pk_select struct {
    Table_name *string  `json:"table_name"`
    Column_name *string  `json:"column_name"`
}

type ArPk_select []Pk_select

func (x *Pk_select) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Table_name,
				&x.Column_name,
			}

}

func (rt *Pk_select)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Column_name,
    
  }
}


func (rt *Pk_selectParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}


			


func SQLPk_select(flag db.SQLOper)string{
  x :=   "table_name,column_name"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "pk_select"	,
      "table_name" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "pk_select"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"pk_select"	,
			 "column_name"  ,
			 "$1" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"pk_select"	,
						"table_name",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"pk_select"	,
						"column_name=$1"  ,
						"table_name=$2",
					x)

			
  }

	return ""

}
