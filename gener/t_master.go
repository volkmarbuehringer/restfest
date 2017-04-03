
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-03 23:38:38.328761791 +0200 CEST
//code for table t_master

package gener


import (
	"restfest/db"
)

type T_master struct {
    Id int64  `json:"id"`
    Lulu db.JSONNullString  `json:"lulu"`
    Lulu1 db.JSONNullString  `json:"lulu1"`
}



func SQLT_master(tab string,flag db.SQLOper)[]interface{}{
  x :=   "id,lulu,lulu1"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
      tab,
      "id",
    }
    case db.GenInsert:
    return    []interface{}{ tab,
       "lulu,lulu1"  ,
       "$1,$2" ,
        x,
      }
    default:
      return    []interface{}{ tab,
        "lulu=$1,lulu1=$2"  ,
        "id=$3",
      x,
        }

  }

}


func EmptyT_master() interface{}{
    return new(T_master)
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
	&rt.Lulu1,  }

}
