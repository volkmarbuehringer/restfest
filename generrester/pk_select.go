
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-03 22:46:24.560089852 +0200 CEST
//code for table pk_select

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

type dummyPk_select time.Time

type Pk_selectParams struct {
    Length int
    Offset int
}

type Pk_select struct {
    Table_name *string  `json:"table_name"`
    Column_name *string  `json:"column_name"`
}

type ArPk_select []Pk_select



var  Pk_selectSQL  =db.ColumnLists{
	 "table_name,column_name",
	  "column_name"  ,
		 "$1" ,
	   "table_name" ,
		 "table_name=$2",
		"column_name"	 ,
						"column_name=$1"  ,

}


func (x *Pk_selectParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Pk_select) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x Pk_selectParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Pk_select) Columns() []string {
	return strings.Split(Pk_selectSQL.All,",")
}


func (x *Pk_select) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Table_name,
				&x.Column_name,
			}

}

func (rt *Pk_select)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Column_name,
    
  }
}


func (rt *Pk_selectParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}



func (t Pk_select) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst Pk_selectParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  Pk_selectSQL.All,
	     "pk_select"	,
      Pk_selectSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	Pk_selectSQL.All,
   "pk_select"	,
  		)


			
  }

	return ""

}




func ( dst Pk_select)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"pk_select"	,
			 Pk_selectSQL.Inserts  ,
			 Pk_selectSQL.BindsInsert ,
				Pk_selectSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"pk_select"	,
					Pk_selectSQL.PK,
						Pk_selectSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"pk_select"	,
					Pk_selectSQL.BindsUpdate  ,
								Pk_selectSQL.PKUpdate,
					Pk_selectSQL.All,)

			
  }

	return ""

}



type IterPk_select struct {
	db.BaseCopy
	Pk_select Pk_select
}

type ArIterPk_select struct {
	IterPk_select
	ArPk_select ArPk_select
}


				

type MapPk_select map[string]Pk_select

type MapIterPk_select struct {
		IterPk_select
		MapPk_select  MapPk_select
}

				func (t *MapIterPk_select) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Pk_select)
				 t.MapPk_select=make(MapPk_select)
					for t.IterPk_select.Next(){
						t.MapPk_select[*t.Pk_select.Table_name] = t.Pk_select
					}
				 return len(t.MapPk_select)
				}


	



func (t *ArIterPk_select) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Pk_select)
	for t.IterPk_select.Next(){
		t.ArPk_select = append(t.ArPk_select,t.Pk_select)
	}
 return len(t.ArPk_select)
}

func (t *IterPk_select) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Pk_select)
return
}

func (t *IterPk_select) Value() db.PgxGener  {

 return &t.Pk_select
}

func (t *IterPk_select) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Pk_select)

}




			
