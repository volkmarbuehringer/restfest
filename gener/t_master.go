
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-11 19:28:42.798703647 +0200 CEST
//code for table t_master

package gener



import (
	"restfest/db"
	"fmt"
)

type T_masterParams struct {
    Length db.JSONNullInt64 `schema:"length"`
    Offset db.JSONNullInt64 `schema:"offset"`
}


type T_master struct {
    Id int64  `json:"id"`
    Lulu db.JSONString  `json:"lulu"`
    Lulu1 db.JSONString  `json:"lulu1"`
}



func SQLT_master(flag db.SQLOper)string{
  x :=   "id,lulu,lulu1"

  switch flag{
			
    case db.GenSelectID:
    return fmt.Sprintf("select %s from " + db.DBschema + ".%s where %s=$1",
    x,
	     "t_master"	,
      "id" ,
    )

		case db.GenSelectAll:
		return fmt.Sprintf("select %s from " + db.DBschema + ".%s order by %s limit $1 offset $2",
		x,
   "t_master"	,
  "id",
		)

		case db.GenInsert:
		return  fmt.Sprintf(   `insert into ` + db.DBschema + `.%s(%s)values(%s) returning %s`,
			"t_master"	,
			 "lulu,lulu1"  ,
			 "$1,$2" ,
				x)

				default:
				return  fmt.Sprintf(  `update ` + db.DBschema + `.%s set %s where %s returning %s`,
					"t_master"	,
						"lulu=$1,lulu1=$2"  ,
						"id=$3",
					x)

			
  }
			
			
}


func EmptyT_master() interface{}{
    return new(T_master)
}


func EmptyParamT_master() interface{}{
    return new(T_masterParams)
}

func ScannerT_master()( []interface{}, interface{}){
	struT := new(T_master)

return  []interface{}{
				&struT.Id,
				&struT.Lulu,
				&struT.Lulu1,
			}, struT

}


func ScannerTT_master()( []interface{},*T_master){
	struT := new(T_master)

return  []interface{}{
				&struT.Id,
				&struT.Lulu,
				&struT.Lulu1,
			}, struT

}


func ROWInsertT_master(inter interface{})[]interface{}{
					
  rt := inter.( *T_master )
	
  return []interface{}{
				
	&rt.Lulu,
	&rt.Lulu1,

  }

}


func ROWQueryT_master(inter interface{})[]interface{}{
  rt := inter.( *T_masterParams )

 if !rt.Length.Valid {
	rt.Length.Int64=100
	rt.Length.Valid=true
} 

  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
