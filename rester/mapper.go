// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-21 21:36:41.215182827 +0200 CEST

package main


import (
"restfest/generrester"
"restfest/db"

)


func init(){
	db.FunMap=map[string]db.TFunMap{
	
	"guges":{		generrester.SQLGuges,
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

		1, },
	
	"guges1":{		generrester.SQLGuges1,
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

		1, },
	
	"get_guges1_174200":{		generrester.SQLGet_guges1_174200,
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

		3, },
	
	"lala":{		generrester.SQLLala,
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

		2, },
	
	"lala1":{		generrester.SQLLala1,
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

		2, },
	
	"logger":{		generrester.SQLLogger,
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

		1, },
	
	"pk_select":{		generrester.SQLPk_select,
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

		2, },
	
	"tedas":{		generrester.SQLTedas,
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

		1, },
	
	"testa":{		generrester.SQLTesta,
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

		1, },
	
	"testa1":{		generrester.SQLTesta1,
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

		2, },
	
	"tester":{		generrester.SQLTester,
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

		1, },
	
	"t_master":{		generrester.SQLT_master,
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

		1, },
	
	"t_random":{		generrester.SQLT_random,
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

		1, },
	
	"vtester":{		generrester.SQLVtester,
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

		2, },
	
	"vweburl":{		generrester.SQLVweburl,
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

		2, },
	
	"get_weburl_173825":{		generrester.SQLGet_weburl_173825,
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

		3, },
	
	"weburl":{		generrester.SQLWeburl,
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

		4, },
	
	"get_weburl_173826":{		generrester.SQLGet_weburl_173826,
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

		3, },
	
	"get_all_boo_173824":{		generrester.SQLGet_all_boo_173824,
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

		3, },
}

}
