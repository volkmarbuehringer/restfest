
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-20 21:11:00.872008798 +0200 CEST
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

			

func ScannerVtesterI(struT *Vtester)  db.InterPgx  {

	return []interface{}{
				&struT.Name,
				&struT.Vorname,
				&struT.Id,
				&struT.Code,
				&struT.Lala,
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


func EmptyVtester() interface{}{
    return new(Vtester)
}


func EmptyParamVtester() interface{}{

 rt := new(VtesterParams)

	
		rt.Length=100
	

    return rt
}

func ScannerVtester()(  db.InterPgx , interface{}){
	struT := new(Vtester)

return  ScannerVtesterI(struT), struT

}


func ROWInsertVtester(inter interface{}) db.InterPgx {
					
  rt := inter.( *Vtester )
	
  return []interface{}{
				
	&rt.Vorname,
	&rt.Id,
	&rt.Code,
	&rt.Lala,

  }

}


func ROWQueryVtester(inter interface{}) db.InterPgx {
  rt := inter.( *VtesterParams )

  return []interface{}{
		&rt.Length,
		&rt.Offset,
}

}
