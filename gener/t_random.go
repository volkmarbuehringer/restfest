
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-09 23:08:53.354391052 +0200 CEST
//code for table t_random

package gener


import (
	"restfest/db"
)

type T_randomParams struct {
    Length db.JSONNullInt64
    Offset db.JSONNullInt64
}


type T_random struct {
    Id int64  `json:"id"`
    Md5 db.JSONString  `json:"md5"`
    Gaga db.JSONString  `json:"gaga"`
    Fk db.JSONNullInt64  `json:"fk"`
}



func SQLT_random(tab string,flag db.SQLOper)[]interface{}{
  x :=   "id,md5,gaga,fk"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
      tab,
      "id",
    }
    case db.GenInsert:
    return    []interface{}{ tab,
       "md5,gaga,fk"  ,
       "$1,$2,$3" ,
        x,
      }
    default:
      return    []interface{}{ tab,
        "md5=$1,gaga=$2,fk=$3"  ,
        "id=$4",
      x,
        }

  }

}


func EmptyT_random() interface{}{
    return new(T_random)
}


func EmptyParamT_random() interface{}{
    return new(T_randomParams)
}

func ScannerT_random()( []interface{}, interface{}){
	struT := new(T_random)

return  []interface{}{
				&struT.Id,
				&struT.Md5,
				&struT.Gaga,
				&struT.Fk,
			}, struT

}


func ScannerTT_random()( []interface{},*T_random){
	struT := new(T_random)

return  []interface{}{
				&struT.Id,
				&struT.Md5,
				&struT.Gaga,
				&struT.Fk,
			}, struT

}


func ROWInsertT_random(inter interface{})[]interface{}{
  rt := inter.( *T_random )
  return []interface{}{
	&rt.Md5,
	&rt.Gaga,
	&rt.Fk,  }

}


func ROWQueryT_random(inter interface{})[]interface{}{
  rt := inter.( *T_randomParams )


  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
