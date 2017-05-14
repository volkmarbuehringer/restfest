// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-14 12:17:25.683079384 +0200 CEST

package main

import (
"restfest/rester/gener"
"restfest/db"

)



func init(){
	db.FunMap=map[string]db.TFunMap{
	
	"csvtest":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Csvtest)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Csvtest)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.CsvtestParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterCsvtest )
		},
		Flag:		1, },
	
	"guges":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Guges)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Guges)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.GugesParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterGuges )
		},
		Flag:		1, },
	
	"guges1":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Guges1)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Guges1)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Guges1Params)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterGuges1 )
		},
		Flag:		1, },
	
	"get_guges1":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Get_guges1)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Get_guges1)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Get_guges1Params)
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterGet_guges1 )
		},
		Flag:		3, },
	
	"lala":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Lala)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Lala)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.LalaParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterLala )
		},
		Flag:		2, },
	
	"lala1":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Lala1)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Lala1)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Lala1Params)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterLala1 )
		},
		Flag:		2, },
	
	"logger":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Logger)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Logger)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.LoggerParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterLogger )
		},
		Flag:		1, },
	
	"los":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Los)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Los)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.LosParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterLos )
		},
		Flag:		1, },
	
	"pk_select":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Pk_select)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Pk_select)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Pk_selectParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterPk_select )
		},
		Flag:		2, },
	
	"tedas":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Tedas)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Tedas)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.TedasParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterTedas )
		},
		Flag:		1, },
	
	"testa":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Testa)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Testa)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.TestaParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterTesta )
		},
		Flag:		1, },
	
	"testa1":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Testa1)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Testa1)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Testa1Params)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterTesta1 )
		},
		Flag:		2, },
	
	"tester":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Tester)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Tester)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.TesterParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterTester )
		},
		Flag:		1, },
	
	"t_master":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.T_master)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.T_master)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.T_masterParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterT_master )
		},
		Flag:		1, },
	
	"t_random":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.T_random)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.T_random)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.T_randomParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterT_random )
		},
		Flag:		1, },
	
	"vtester":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Vtester)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Vtester)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.VtesterParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterVtester )
		},
		Flag:		2, },
	
	"vweburl":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Vweburl)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Vweburl)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.VweburlParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterVweburl )
		},
		Flag:		2, },
	
	"get_weburl_173825":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Get_weburl_173825)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Get_weburl_173825)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Get_weburl_173825Params)
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterGet_weburl_173825 )
		},
		Flag:		3, },
	
	"weburl":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Weburl)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Weburl)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.WeburlParams)
			
				rt.Length=100
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterWeburl )
		},
		Flag:		4, },
	
	"get_all_boo":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Get_all_boo)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Get_all_boo)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Get_all_booParams)
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterGet_all_boo )
		},
		Flag:		3, },
	
	"get_weburl":{
		EmptyFun:		func () db.PgxGener{
		    return new(gener.Get_weburl)
		},
		EmptyInsFun:  func () db.PgxGenerIns{
		    return new(gener.Get_weburl)
		},
				ParamFun: func () db.PgxGenerIns{
		 rt := new(gener.Get_weburlParams)
			
		    return rt
		},
	 Iterator:	func () db.Iterator {
				return new(gener.IterGet_weburl )
		},
		Flag:		3, },
}

}
