// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-01 11:30:28.568350476 +0200 CEST

package main


import (
"restfest/generrester"
"restfest/db"

)


func init(){
	db.FunMap=map[string]db.TFunMap{
	
	"csvtest":{
		func () db.PgxGener{
		    return new(generrester.Csvtest)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Csvtest)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.CsvtestParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterCsvtest )
		},
		1, },
	
	"guges":{
		func () db.PgxGener{
		    return new(generrester.Guges)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Guges)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.GugesParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterGuges )
		},
		1, },
	
	"get_guges1_174200":{
		func () db.PgxGener{
		    return new(generrester.Get_guges1_174200)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Get_guges1_174200)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Get_guges1_174200Params)
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterGet_guges1_174200 )
		},
		3, },
	
	"guges1":{
		func () db.PgxGener{
		    return new(generrester.Guges1)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Guges1)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Guges1Params)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterGuges1 )
		},
		1, },
	
	"lala":{
		func () db.PgxGener{
		    return new(generrester.Lala)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Lala)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.LalaParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterLala )
		},
		2, },
	
	"lala1":{
		func () db.PgxGener{
		    return new(generrester.Lala1)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Lala1)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Lala1Params)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterLala1 )
		},
		2, },
	
	"logger":{
		func () db.PgxGener{
		    return new(generrester.Logger)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Logger)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.LoggerParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterLogger )
		},
		1, },
	
	"los":{
		func () db.PgxGener{
		    return new(generrester.Los)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Los)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.LosParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterLos )
		},
		1, },
	
	"pk_select":{
		func () db.PgxGener{
		    return new(generrester.Pk_select)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Pk_select)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Pk_selectParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterPk_select )
		},
		2, },
	
	"tedas":{
		func () db.PgxGener{
		    return new(generrester.Tedas)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Tedas)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.TedasParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterTedas )
		},
		1, },
	
	"testa":{
		func () db.PgxGener{
		    return new(generrester.Testa)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Testa)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.TestaParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterTesta )
		},
		1, },
	
	"testa1":{
		func () db.PgxGener{
		    return new(generrester.Testa1)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Testa1)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Testa1Params)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterTesta1 )
		},
		2, },
	
	"tester":{
		func () db.PgxGener{
		    return new(generrester.Tester)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Tester)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.TesterParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterTester )
		},
		1, },
	
	"t_master":{
		func () db.PgxGener{
		    return new(generrester.T_master)
		},
		func () db.PgxGenerIns{
		    return new(generrester.T_master)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.T_masterParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterT_master )
		},
		1, },
	
	"t_random":{
		func () db.PgxGener{
		    return new(generrester.T_random)
		},
		func () db.PgxGenerIns{
		    return new(generrester.T_random)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.T_randomParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterT_random )
		},
		1, },
	
	"vtester":{
		func () db.PgxGener{
		    return new(generrester.Vtester)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Vtester)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.VtesterParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterVtester )
		},
		2, },
	
	"vweburl":{
		func () db.PgxGener{
		    return new(generrester.Vweburl)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Vweburl)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.VweburlParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterVweburl )
		},
		2, },
	
	"weburl":{
		func () db.PgxGener{
		    return new(generrester.Weburl)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Weburl)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.WeburlParams)
			
				rt.Length=100
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterWeburl )
		},
		4, },
	
	"get_weburl_173826":{
		func () db.PgxGener{
		    return new(generrester.Get_weburl_173826)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Get_weburl_173826)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Get_weburl_173826Params)
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterGet_weburl_173826 )
		},
		3, },
	
	"get_all_boo_173824":{
		func () db.PgxGener{
		    return new(generrester.Get_all_boo_173824)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Get_all_boo_173824)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Get_all_boo_173824Params)
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterGet_all_boo_173824 )
		},
		3, },
	
	"get_weburl_173825":{
		func () db.PgxGener{
		    return new(generrester.Get_weburl_173825)
		},
		func () db.PgxGenerIns{
		    return new(generrester.Get_weburl_173825)
		},
		func () db.PgxGenerIns{
		 rt := new(generrester.Get_weburl_173825Params)
			
		    return rt
		},
		func () db.Iterator {
				return new( generrester.IterGet_weburl_173825 )
		},
		3, },
}

}
