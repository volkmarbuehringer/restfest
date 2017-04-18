// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-18 21:50:02.865372246 +0200 CEST

package main


import (
"restfest/generrester"
"restfest/db"
"fmt"
)


func init(){
	
		
db.SQLFunMap["guges"]= generrester.SQLGuges
db.ROWInsertFunMap["guges"]= generrester.ROWInsertGuges
db.ROWQueryFunMap["guges"]= generrester.ROWQueryGuges
db.EmptyFunMap["guges"]=generrester.EmptyGuges
db.ParamFunMap["guges"]=generrester.EmptyParamGuges
db.ScannerFunMap["guges"]=generrester.ScannerGuges



		
db.SQLFunMap["guges1"]= generrester.SQLGuges1
db.ROWInsertFunMap["guges1"]= generrester.ROWInsertGuges1
db.ROWQueryFunMap["guges1"]= generrester.ROWQueryGuges1
db.EmptyFunMap["guges1"]=generrester.EmptyGuges1
db.ParamFunMap["guges1"]=generrester.EmptyParamGuges1
db.ScannerFunMap["guges1"]=generrester.ScannerGuges1



		
db.SQLFunMap["get_guges1_174200"]= generrester.SQLGet_guges1_174200
db.ROWInsertFunMap["get_guges1_174200"]= generrester.ROWInsertGet_guges1_174200
db.ROWQueryFunMap["get_guges1_174200"]= generrester.ROWQueryGet_guges1_174200
db.EmptyFunMap["get_guges1_174200"]=generrester.EmptyGet_guges1_174200
db.ParamFunMap["get_guges1_174200"]=generrester.EmptyParamGet_guges1_174200
db.ScannerFunMap["get_guges1_174200"]=generrester.ScannerGet_guges1_174200



		
db.SQLFunMap["lala"]= generrester.SQLLala
db.ROWInsertFunMap["lala"]= generrester.ROWInsertLala
db.ROWQueryFunMap["lala"]= generrester.ROWQueryLala
db.EmptyFunMap["lala"]=generrester.EmptyLala
db.ParamFunMap["lala"]=generrester.EmptyParamLala
db.ScannerFunMap["lala"]=generrester.ScannerLala

db.ConnectorFunMap["lala"]=generrester.ConnectorLala



		
db.SQLFunMap["lala1"]= generrester.SQLLala1
db.ROWInsertFunMap["lala1"]= generrester.ROWInsertLala1
db.ROWQueryFunMap["lala1"]= generrester.ROWQueryLala1
db.EmptyFunMap["lala1"]=generrester.EmptyLala1
db.ParamFunMap["lala1"]=generrester.EmptyParamLala1
db.ScannerFunMap["lala1"]=generrester.ScannerLala1



		
db.SQLFunMap["logger"]= generrester.SQLLogger
db.ROWInsertFunMap["logger"]= generrester.ROWInsertLogger
db.ROWQueryFunMap["logger"]= generrester.ROWQueryLogger
db.EmptyFunMap["logger"]=generrester.EmptyLogger
db.ParamFunMap["logger"]=generrester.EmptyParamLogger
db.ScannerFunMap["logger"]=generrester.ScannerLogger



		
db.SQLFunMap["pk_select"]= generrester.SQLPk_select
db.ROWInsertFunMap["pk_select"]= generrester.ROWInsertPk_select
db.ROWQueryFunMap["pk_select"]= generrester.ROWQueryPk_select
db.EmptyFunMap["pk_select"]=generrester.EmptyPk_select
db.ParamFunMap["pk_select"]=generrester.EmptyParamPk_select
db.ScannerFunMap["pk_select"]=generrester.ScannerPk_select



		
db.SQLFunMap["tedas"]= generrester.SQLTedas
db.ROWInsertFunMap["tedas"]= generrester.ROWInsertTedas
db.ROWQueryFunMap["tedas"]= generrester.ROWQueryTedas
db.EmptyFunMap["tedas"]=generrester.EmptyTedas
db.ParamFunMap["tedas"]=generrester.EmptyParamTedas
db.ScannerFunMap["tedas"]=generrester.ScannerTedas



		
db.SQLFunMap["testa1"]= generrester.SQLTesta1
db.ROWInsertFunMap["testa1"]= generrester.ROWInsertTesta1
db.ROWQueryFunMap["testa1"]= generrester.ROWQueryTesta1
db.EmptyFunMap["testa1"]=generrester.EmptyTesta1
db.ParamFunMap["testa1"]=generrester.EmptyParamTesta1
db.ScannerFunMap["testa1"]=generrester.ScannerTesta1



		
db.SQLFunMap["tester"]= generrester.SQLTester
db.ROWInsertFunMap["tester"]= generrester.ROWInsertTester
db.ROWQueryFunMap["tester"]= generrester.ROWQueryTester
db.EmptyFunMap["tester"]=generrester.EmptyTester
db.ParamFunMap["tester"]=generrester.EmptyParamTester
db.ScannerFunMap["tester"]=generrester.ScannerTester



		
db.SQLFunMap["t_master"]= generrester.SQLT_master
db.ROWInsertFunMap["t_master"]= generrester.ROWInsertT_master
db.ROWQueryFunMap["t_master"]= generrester.ROWQueryT_master
db.EmptyFunMap["t_master"]=generrester.EmptyT_master
db.ParamFunMap["t_master"]=generrester.EmptyParamT_master
db.ScannerFunMap["t_master"]=generrester.ScannerT_master



		
db.SQLFunMap["t_random"]= generrester.SQLT_random
db.ROWInsertFunMap["t_random"]= generrester.ROWInsertT_random
db.ROWQueryFunMap["t_random"]= generrester.ROWQueryT_random
db.EmptyFunMap["t_random"]=generrester.EmptyT_random
db.ParamFunMap["t_random"]=generrester.EmptyParamT_random
db.ScannerFunMap["t_random"]=generrester.ScannerT_random



		
db.SQLFunMap["vtester"]= generrester.SQLVtester
db.ROWInsertFunMap["vtester"]= generrester.ROWInsertVtester
db.ROWQueryFunMap["vtester"]= generrester.ROWQueryVtester
db.EmptyFunMap["vtester"]=generrester.EmptyVtester
db.ParamFunMap["vtester"]=generrester.EmptyParamVtester
db.ScannerFunMap["vtester"]=generrester.ScannerVtester



		
db.SQLFunMap["vweburl"]= generrester.SQLVweburl
db.ROWInsertFunMap["vweburl"]= generrester.ROWInsertVweburl
db.ROWQueryFunMap["vweburl"]= generrester.ROWQueryVweburl
db.EmptyFunMap["vweburl"]=generrester.EmptyVweburl
db.ParamFunMap["vweburl"]=generrester.EmptyParamVweburl
db.ScannerFunMap["vweburl"]=generrester.ScannerVweburl



		
db.SQLFunMap["get_weburl_173825"]= generrester.SQLGet_weburl_173825
db.ROWInsertFunMap["get_weburl_173825"]= generrester.ROWInsertGet_weburl_173825
db.ROWQueryFunMap["get_weburl_173825"]= generrester.ROWQueryGet_weburl_173825
db.EmptyFunMap["get_weburl_173825"]=generrester.EmptyGet_weburl_173825
db.ParamFunMap["get_weburl_173825"]=generrester.EmptyParamGet_weburl_173825
db.ScannerFunMap["get_weburl_173825"]=generrester.ScannerGet_weburl_173825



		
db.SQLFunMap["weburl"]= generrester.SQLWeburl
db.ROWInsertFunMap["weburl"]= generrester.ROWInsertWeburl
db.ROWQueryFunMap["weburl"]= generrester.ROWQueryWeburl
db.EmptyFunMap["weburl"]=generrester.EmptyWeburl
db.ParamFunMap["weburl"]=generrester.EmptyParamWeburl
db.ScannerFunMap["weburl"]=generrester.ScannerWeburl

db.ConnectorFunMap["weburl"]=generrester.ConnectorWeburl



		
db.SQLFunMap["get_weburl_173826"]= generrester.SQLGet_weburl_173826
db.ROWInsertFunMap["get_weburl_173826"]= generrester.ROWInsertGet_weburl_173826
db.ROWQueryFunMap["get_weburl_173826"]= generrester.ROWQueryGet_weburl_173826
db.EmptyFunMap["get_weburl_173826"]=generrester.EmptyGet_weburl_173826
db.ParamFunMap["get_weburl_173826"]=generrester.EmptyParamGet_weburl_173826
db.ScannerFunMap["get_weburl_173826"]=generrester.ScannerGet_weburl_173826



		
db.SQLFunMap["get_all_boo_173824"]= generrester.SQLGet_all_boo_173824
db.ROWInsertFunMap["get_all_boo_173824"]= generrester.ROWInsertGet_all_boo_173824
db.ROWQueryFunMap["get_all_boo_173824"]= generrester.ROWQueryGet_all_boo_173824
db.EmptyFunMap["get_all_boo_173824"]=generrester.EmptyGet_all_boo_173824
db.ParamFunMap["get_all_boo_173824"]=generrester.EmptyParamGet_all_boo_173824
db.ScannerFunMap["get_all_boo_173824"]=generrester.ScannerGet_all_boo_173824



db.FlagMap=map[string]int{	"guges":1,	"guges1":1,	"get_guges1_174200":3,	"lala":2,	"lala1":2,	"logger":1,	"pk_select":2,	"tedas":1,	"testa1":2,	"tester":1,	"t_master":1,	"t_random":1,	"vtester":2,	"vweburl":2,	"get_weburl_173825":3,	"weburl":4,	"get_weburl_173826":3,	"get_all_boo_173824":3, }

db.InitDB()
	fmt.Printf("geladene Tabellen: %d\n", len(db.SQLFunMap))
}
