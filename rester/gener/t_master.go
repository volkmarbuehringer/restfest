
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-13 09:33:54.660581894 +0200 CEST
//code for table t_master

package gener



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"strings"
		"encoding/json"
	"github.com/jackc/pgx"
			
)

type dummyT_master time.Time

type T_masterParams struct {
    Length int
    Offset int
}

type T_master struct {
    Id int32  `json:"id"`
    Lulu *string  `json:"lulu"`
    Lulu1 *string  `json:"lulu1"`
}

type ArT_master []T_master



var  T_masterSQL  =db.ColumnLists{
	 "id,lulu,lulu1",
	  "lulu,lulu1"  ,
		 "$1,$2" ,
	   "id" ,
		 "id=$3",
		"lulu,lulu1"	 ,
						"lulu=$1,lulu1=$2"  ,

}


func (x *T_masterParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x T_master) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x T_masterParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t T_master) Columns() []string {
	return strings.Split(T_masterSQL.All,",")
}


func (x *T_master) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Lulu,
				&x.Lulu1,
			}

}

func (rt *T_master)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Lulu,	&rt.Lulu1,
    
  }
}


func (rt *T_masterParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t T_master) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst T_masterParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  T_masterSQL.All,
	     "t_master"	,
      T_masterSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	T_masterSQL.All,
   "t_master"	,
  		)


			
  }

	return ""

}




func ( dst T_master)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"t_master"	,
			 T_masterSQL.Inserts  ,
			 T_masterSQL.BindsInsert ,
				T_masterSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"t_master"	,
					T_masterSQL.PK,
						T_masterSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"t_master"	,
					T_masterSQL.BindsUpdate  ,
								T_masterSQL.PKUpdate,
					T_masterSQL.All,)

			
  }

	return ""

}



type IterT_master struct {
	db.BaseCopy
	T_master T_master
}

type ArIterT_master struct {
	IterT_master
	ArT_master ArT_master
}


				

type MapT_master map[int32]T_master

type MapIterT_master struct {
		IterT_master
		MapT_master  MapT_master
}

				func (t *MapIterT_master) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.T_master)
				 t.MapT_master=make(MapT_master)
					for t.IterT_master.Next(){
						t.MapT_master[t.T_master.Id] = t.T_master
					}
				 return len(t.MapT_master)
				}


	



func (t *ArIterT_master) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.T_master)
	for t.IterT_master.Next(){
		t.ArT_master = append(t.ArT_master,t.T_master)
	}
 return len(t.ArT_master)
}

func (t *IterT_master) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.T_master)
return
}

func (t *IterT_master) Value() db.PgxGener  {

 return &t.T_master
}

func (t *IterT_master) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.T_master)

}




			
