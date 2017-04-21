
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 00:33:08.990550423 +0200 CEST
//code for table logger

package generrester



import (
	"restfest/db"
	"fmt"
	"time"

			
)

type dummyLogger time.Time

type LoggerParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Logger struct {
    Seq int32  `json:"seq"`
    Id *int32  `json:"id"`
    Typ *string  `json:"typ"`
    Tablename *string  `json:"tablename"`
    Rower *map[string]interface{}  `json:"rower"`
}

type ArLogger []Logger

func (x *Logger) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Seq,
				&x.Id,
				&x.Typ,
				&x.Tablename,
				&x.Rower,
			}

}

func (rt *Logger)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Id,	&rt.Typ,	&rt.Tablename,	&rt.Rower,
    
  }
}


func (rt *LoggerParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}


			


func SQLLogger(flag db.SQLOper)string{
  x :=   "seq,id,typ,tablename,rower"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "logger"	,
      "seq" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "logger"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"logger"	,
			 "id,typ,tablename,rower"  ,
			 "$1,$2,$3,$4" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"logger"	,
						"seq",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"logger"	,
						"id=$1,typ=$2,tablename=$3,rower=$4"  ,
						"seq=$5",
					x)

			
  }

	return ""

}
