
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-03 23:24:48.363924061 +0200 CEST
//code for table tester

package gener


import (
	"restfest/db"
)

type Tester struct {
    Name db.JSONNullString  `json:"name"`
    Vorname db.JSONNullString  `json:"vorname"`
    Id int64  `json:"id"`
    Code db.JSONNullString  `json:"code"`
    Lala db.JSONNullInt64  `json:"lala"`
}



func SQLTester(tab string,flag db.SQLOper)[]interface{}{
  x :=   "name,vorname,id,code,lala"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
      tab,
      "id",
    }
    case db.GenInsert:
    return    []interface{}{ tab,
       "name,vorname,code,lala"  ,
       "$1,$2,$3,$4" ,
        x,
      }
    default:
      return    []interface{}{ tab,
        "name=$1,vorname=$2,code=$3,lala=$4"  ,
        "id=$5",
      x,
        }

  }

}


func EmptyTester() interface{}{
    return new(Tester)
}


func ScannerTester()( []interface{}, interface{}){
	struT := new(Tester)

return  []interface{}{
				&struT.Name,
				&struT.Vorname,
				&struT.Id,
				&struT.Code,
				&struT.Lala,
			}, struT

}


func ScannerTTester()( []interface{},*Tester){
	struT := new(Tester)

return  []interface{}{
				&struT.Name,
				&struT.Vorname,
				&struT.Id,
				&struT.Code,
				&struT.Lala,
			}, struT

}


func ROWInsertTester(inter interface{})[]interface{}{
  rt := inter.( *Tester )
  return []interface{}{
	&rt.Name,
	&rt.Vorname,
	&rt.Code,
	&rt.Lala,  }

}
