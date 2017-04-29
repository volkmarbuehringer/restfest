// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 22:56:40.348002947 +0200 CEST

package main


import (
"restfest/generrester"
"restfest/db"
"github.com/jackc/pgx"
"github.com/jackc/pgx/pgtype"
"fmt"
)


func init(){
	db.ConnectorFunMap=map[string]db.MapperFun2{
	  	 
	 "lala": func (con *pgx.Conn, oida pgtype.Oid,oid pgtype.Oid) error {
	 	con.ConnInfo.RegisterDataType(pgtype.DataType{
	 		Value: &generrester.ArLala{},
	 		Name:  "generrester.ArLala",
	 		Oid:  oida,
	 	})
	 	con.ConnInfo.RegisterDataType(pgtype.DataType{
	 		Value: &generrester.Lala{},
	 		Name:  "generrester.Lala",
	 		Oid:  oid,
	 	})
	 	return nil
	 } ,
	 
	  	 
	 "weburl": func (con *pgx.Conn, oida pgtype.Oid,oid pgtype.Oid) error {
	 	con.ConnInfo.RegisterDataType(pgtype.DataType{
	 		Value: &generrester.ArWeburl{},
	 		Name:  "generrester.ArWeburl",
	 		Oid:  oida,
	 	})
	 	con.ConnInfo.RegisterDataType(pgtype.DataType{
	 		Value: &generrester.Weburl{},
	 		Name:  "generrester.Weburl",
	 		Oid:  oid,
	 	})
	 	return nil
	 } ,
	 
}


db.InitDB()
	fmt.Printf("geladene Tabellen: %d\n", len(db.FunMap))
}
