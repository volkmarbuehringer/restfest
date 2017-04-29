// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 22:56:40.340833593 +0200 CEST

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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArCsvtest,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArGuges,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArGet_guges1_174200,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArGuges1,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArLala,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArLala1,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArLogger,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArLos,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArPk_select,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArTedas,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArTesta,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArTesta1,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArTester,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArT_master,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArT_random,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArVtester,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArVweburl,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArWeburl,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArGet_weburl_173826,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArGet_all_boo_173824,0)
				return &x
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
		func () db.PgxGenerAr {
		    x:=make(generrester.ArGet_weburl_173825,0)
				return &x
		},
		3, },
}

}
