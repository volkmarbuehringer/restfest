
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-03 21:55:29.034354172 +0200 CEST
//code for table testa1

package gener


import (
	"restfest/db"
)

type Testa1 struct {
    Id db.JSONNullInt64  `json:"id"`
    Lulu db.JSONNullString  `json:"lulu"`
    Lulu1 db.JSONNullString  `json:"lulu1"`
    Ider db.JSONNullInt64  `json:"ider"`
    Md5 db.JSONNullString  `json:"md5"`
    Gaga db.JSONNullString  `json:"gaga"`
    Fk1 db.JSONNullInt64  `json:"fk1"`
    Ider1 db.JSONNullInt64  `json:"ider1"`
    Lalu db.JSONNullString  `json:"lalu"`
    Zacka db.JSONNullString  `json:"zacka"`
    Fk2 db.JSONNullInt64  `json:"fk2"`
    Name db.JSONNullString  `json:"name"`
    Vorname db.JSONNullString  `json:"vorname"`
}



func SQLTesta1(tab string,flag db.SQLOper)[]interface{}{
  x :=   "id,lulu,lulu1,ider,md5,gaga,fk1,ider1,lalu,zacka,fk2,name,vorname"

  switch flag{
    case db.GenSelect:
    return []interface{}{
    x,
      tab,
      "id",
    }
    case db.GenInsert:
    return    []interface{}{ tab,
       "lulu,lulu1,ider,md5,gaga,fk1,ider1,lalu,zacka,fk2,name,vorname"  ,
       "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12" ,
        x,
      }
    default:
      return    []interface{}{ tab,
        "lulu=$1,lulu1=$2,ider=$3,md5=$4,gaga=$5,fk1=$6,ider1=$7,lalu=$8,zacka=$9,fk2=$10,name=$11,vorname=$12"  ,
        "id=$13",
      x,
        }

  }

}


func EmptyTesta1() interface{}{
    return new(Testa1)
}


func ScannerTesta1()( []interface{}, interface{}){
	struT := new(Testa1)

return  []interface{}{
				&struT.Id,
				&struT.Lulu,
				&struT.Lulu1,
				&struT.Ider,
				&struT.Md5,
				&struT.Gaga,
				&struT.Fk1,
				&struT.Ider1,
				&struT.Lalu,
				&struT.Zacka,
				&struT.Fk2,
				&struT.Name,
				&struT.Vorname,
			}, struT

}


func ScannerTTesta1()( []interface{},*Testa1){
	struT := new(Testa1)

return  []interface{}{
				&struT.Id,
				&struT.Lulu,
				&struT.Lulu1,
				&struT.Ider,
				&struT.Md5,
				&struT.Gaga,
				&struT.Fk1,
				&struT.Ider1,
				&struT.Lalu,
				&struT.Zacka,
				&struT.Fk2,
				&struT.Name,
				&struT.Vorname,
			}, struT

}


func ROWInsertTesta1(inter interface{})[]interface{}{
  rt := inter.( *Testa1 )
  return []interface{}{
	&rt.Lulu,
	&rt.Lulu1,
	&rt.Ider,
	&rt.Md5,
	&rt.Gaga,
	&rt.Fk1,
	&rt.Ider1,
	&rt.Lalu,
	&rt.Zacka,
	&rt.Fk2,
	&rt.Name,
	&rt.Vorname,  }

}