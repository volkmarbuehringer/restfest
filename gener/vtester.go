
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-10 21:37:34.924978355 +0200 CEST
//code for table vtester

package gener


import (
	"restfest/db"
)

type VtesterParams struct {
    Length db.JSONNullInt64 `schema:"length"`
    Offset db.JSONNullInt64 `schema:"offset"`
}


type Vtester struct {
    Name db.JSONString  `json:"name"`
    Vorname db.JSONString  `json:"vorname"`
    Id db.JSONNullInt64  `json:"id"`
    Code db.JSONString  `json:"code"`
    Lala db.JSONNullInt64  `json:"lala"`
}



func SQLVtester(flag db.SQLOper)[]interface{}{
  x :=   "name,vorname,id,code,lala"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
		       "vtester"	,
    	   "name" ,
    }
			
    case db.GenInsert:
    return    []interface{}{   "vtester"	,
       "vorname,id,code,lala"  ,
       "$1,$2,$3,$4" ,
        x,
      }
    default:
      return    []interface{}{   "vtester"	,
        "vorname=$1,id=$2,code=$3,lala=$4"  ,
        "name=$5",
      x,
        }

  }
			
			
}


func EmptyVtester() interface{}{
    return new(Vtester)
}


func EmptyParamVtester() interface{}{
    return new(VtesterParams)
}

func ScannerVtester()( []interface{}, interface{}){
	struT := new(Vtester)

return  []interface{}{
				&struT.Name,
				&struT.Vorname,
				&struT.Id,
				&struT.Code,
				&struT.Lala,
			}, struT

}


func ScannerTVtester()( []interface{},*Vtester){
	struT := new(Vtester)

return  []interface{}{
				&struT.Name,
				&struT.Vorname,
				&struT.Id,
				&struT.Code,
				&struT.Lala,
			}, struT

}


func ROWInsertVtester(inter interface{})[]interface{}{
					
  rt := inter.( *Vtester )
	
  return []interface{}{
				
	&rt.Vorname,
	&rt.Id,
	&rt.Code,
	&rt.Lala,

  }

}


func ROWQueryVtester(inter interface{})[]interface{}{
  rt := inter.( *VtesterParams )

 if !rt.Length.Valid {
	rt.Length.Int64=100
	rt.Length.Valid=true
} 

  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
