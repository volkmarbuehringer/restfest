
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-11 20:26:40.042373018 +0200 CEST
//code for table pk_select

package gener



import (
	"restfest/db"
	"fmt"
)

type Pk_selectParams struct {
    Length db.JSONNullInt64 `schema:"length"`
    Offset db.JSONNullInt64 `schema:"offset"`
}


type Pk_select struct {
    Table_name db.JSONString  `json:"table_name"`
    Column_name db.JSONString  `json:"column_name"`
}



func SQLPk_select(flag db.SQLOper)string{
  x :=   "table_name,column_name"

  switch flag{
			
    case db.GenSelectID:
    return fmt.Sprintf(db.SQLPattern[db.GenSelectID],
    x,
	     "pk_select"	,
      "table_name" ,
    )

		case db.GenSelectAll:
		return fmt.Sprintf(db.SQLPattern[db.GenSelectAll],
		x,
   "pk_select"	,
  "table_name",
		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[db.GenInsert],
			"pk_select"	,
			 "column_name"  ,
			 "$1" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[db.GenDelete]  ,
					"pk_select"	,
						"table_name",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[db.GenUpdate]  ,
					"pk_select"	,
						"column_name=$1"  ,
						"table_name=$2",
					x)

			
  }

	return ""

}


func EmptyPk_select() interface{}{
    return new(Pk_select)
}


func EmptyParamPk_select() interface{}{

 rt := new(Pk_selectParams)

	 if !rt.Length.Valid {
		rt.Length.Int64=100
		rt.Length.Valid=true
	} 

    return rt
}

func ScannerPk_select()( []interface{}, interface{}){
	struT := new(Pk_select)

return  []interface{}{
				&struT.Table_name,
				&struT.Column_name,
			}, struT

}


func ScannerTPk_select()( []interface{},*Pk_select){
	struT := new(Pk_select)

return  []interface{}{
				&struT.Table_name,
				&struT.Column_name,
			}, struT

}


func ROWInsertPk_select(inter interface{})[]interface{}{
					
  rt := inter.( *Pk_select )
	
  return []interface{}{
				
	&rt.Column_name,

  }

}


func ROWQueryPk_select(inter interface{})[]interface{}{
  rt := inter.( *Pk_selectParams )

  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
