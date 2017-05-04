
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-04 22:18:25.474414868 +0200 CEST
//code for table los

package genrestspec



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"strings"
		"encoding/json"
	"github.com/jackc/pgx"
			
)

type dummyLos time.Time

type LosParams struct {
    Length int
    Offset int
}

type Los struct {
    Id int32  `json:"id"`
    Url *string  `json:"url"`
    Zusatz *int64  `json:"zusatz"`
    Created *time.Time  `json:"created"`
    W_cr_date *time.Time  `json:"w_cr_date"`
    W_upd_date *time.Time  `json:"w_upd_date"`
    W_upd_uid *string  `json:"w_upd_uid"`
    W_cr_uid *string  `json:"w_cr_uid"`
    Addtime *time.Time  `json:"addtime"`
    Addtime1 *time.Time  `json:"addtime1"`
    Flag *bool  `json:"flag"`
    Test *float64  `json:"test"`
    L_iban *string  `json:"l_iban"`
}

type ArLos []Los



var  LosSQL  =db.ColumnLists{
	 "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test,l_iban",
	  "url,zusatz,created,w_cr_date,w_cr_uid,addtime,addtime1,flag,test,l_iban"  ,
		 "$1,$2,$3,current_timestamp,'webSrv',$4,$5,$6,$7,$8" ,
	   "id" ,
		 "id=$9",
		"url,zusatz,created,addtime,addtime1,flag,test,l_iban"	 ,
						"url=$1,zusatz=$2,created=$3,w_upd_date=current_timestamp,w_upd_uid='webSrv',addtime=$4,addtime1=$5,flag=$6,test=$7,l_iban=$8"  ,

}


func (x *LosParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Los) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x LosParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Los) Columns() []string {
	return strings.Split(LosSQL.All,",")
}


func (x *Los) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Url,
				&x.Zusatz,
				&x.Created,
				&x.W_cr_date,
				&x.W_upd_date,
				&x.W_upd_uid,
				&x.W_cr_uid,
				&x.Addtime,
				&x.Addtime1,
				&x.Flag,
				&x.Test,
				&x.L_iban,
			}

}

func (rt *Los)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Url,	&rt.Zusatz,	&rt.Created,	&rt.Addtime,	&rt.Addtime1,	&rt.Flag,	&rt.Test,	&rt.L_iban,
    
  }
}


func (rt *LosParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Los) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst LosParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  LosSQL.All,
	     "los"	,
      LosSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	LosSQL.All,
   "los"	,
  		)


			
  }

	return ""

}




func ( dst Los)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"los"	,
			 LosSQL.Inserts  ,
			 LosSQL.BindsInsert ,
				LosSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"los"	,
					LosSQL.PK,
						LosSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"los"	,
					LosSQL.BindsUpdate  ,
								LosSQL.PKUpdate,
					LosSQL.All,)

			
  }

	return ""

}



type IterLos struct {
	db.BaseCopy
	Los Los
}

type ArIterLos struct {
	IterLos
	ArLos ArLos
}


				

type MapLos map[int32]Los

type MapIterLos struct {
		IterLos
		MapLos  MapLos
}

				func (t *MapIterLos) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Los)
				 t.MapLos=make(MapLos)
					for t.IterLos.Next(){
						t.MapLos[t.Los.Id] = t.Los
					}
				 return len(t.MapLos)
				}


	



func (t *ArIterLos) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Los)
	for t.IterLos.Next(){
		t.ArLos = append(t.ArLos,t.Los)
	}
 return len(t.ArLos)
}

func (t *IterLos) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Los)
return
}

func (t *IterLos) Value() db.PgxGener  {

 return &t.Los
}

func (t *IterLos) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Los)

}




			
