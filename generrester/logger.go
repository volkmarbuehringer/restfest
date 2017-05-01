
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-01 11:30:28.19099496 +0200 CEST
//code for table logger

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"strings"
		"encoding/json"
	"github.com/jackc/pgx"
			
)

type dummyLogger time.Time

type LoggerParams struct {
    Length int
    Offset int
}

type Logger struct {
    Seq int32  `json:"seq"`
    Id *int32  `json:"id"`
    Typ *string  `json:"typ"`
    Tablename *string  `json:"tablename"`
    Rower *map[string]interface{}  `json:"rower"`
}

type ArLogger []Logger



var  LoggerSQL  =db.ColumnLists{
	 "seq,id,typ,tablename,rower",
	  "id,typ,tablename,rower"  ,
		 "$1,$2,$3,$4" ,
	   "seq" ,
		 "seq=$5",
		"id,typ,tablename,rower"	 ,
						"id=$1,typ=$2,tablename=$3,rower=$4"  ,

}


func (x *LoggerParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Logger) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x LoggerParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Logger) Columns() []string {
	return strings.Split(LoggerSQL.All,",")
}


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



func (t Logger) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst LoggerParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  LoggerSQL.All,
	     "logger"	,
      LoggerSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	LoggerSQL.All,
   "logger"	,
  		)


			
  }

	return ""

}




func ( dst Logger)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"logger"	,
			 LoggerSQL.Inserts  ,
			 LoggerSQL.BindsInsert ,
				LoggerSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"logger"	,
					LoggerSQL.PK,
						LoggerSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"logger"	,
					LoggerSQL.BindsUpdate  ,
								LoggerSQL.PKUpdate,
					LoggerSQL.All,)

			
  }

	return ""

}



type IterLogger struct {
	db.BaseCopy
	Logger Logger
}

type ArIterLogger struct {
	IterLogger
	ArLogger ArLogger
}


				

type MapLogger map[int32]Logger

type MapIterLogger struct {
		IterLogger
		MapLogger  MapLogger
}

				func (t *MapIterLogger) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Logger)
				 t.MapLogger=make(MapLogger)
					for t.IterLogger.Next(){
						t.MapLogger[t.Logger.Seq] = t.Logger
					}
				 return len(t.MapLogger)
				}


	



func (t *ArIterLogger) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Logger)
	for t.IterLogger.Next(){
		t.ArLogger = append(t.ArLogger,t.Logger)
	}
 return len(t.ArLogger)
}

func (t *IterLogger) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Logger)
return
}

func (t *IterLogger) Value() db.PgxGener  {

 return &t.Logger
}

func (t *IterLogger) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Logger)

}




			
