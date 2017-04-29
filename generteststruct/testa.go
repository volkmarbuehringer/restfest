
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 23:06:05.546927913 +0200 CEST
//code for table testa

package generteststruct



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyTesta time.Time

type TestaParams struct {
    Length int
    Offset int
}

func (x *TestaParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Testa) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x TestaParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}

type Testa struct {
    Id int32  `json:"id"`
    T *map[string]interface{}  `json:"t"`
}

type ArTesta []Testa



func (t Testa) Columns() []string {
	return []string{	 "id" ,	 "t" ,
			}
}


func (x *Testa) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.T,
			}

}

func (rt *Testa)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.T,
    
  }
}


func (rt *TestaParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapTesta map[int32]Testa
				type MapIterTesta struct {
					IterTesta
					MapTesta  MapTesta
				}

				func (t *MapIterTesta) Next() bool {
					ok := t.IterTesta.Next()

				 return ok
				}

				func (t *MapIterTesta) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Testa)
				 t.MapTesta=make(MapTesta)
					for t.IterTesta.Next(){
						t.MapTesta[t.Testa.Id] = t.Testa
					}
				 return len(t.MapTesta)
				}


	

func (dst *ArTesta) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Testa)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterTesta struct {
	db.BaseCopy
	Testa Testa
}

type ArIterTesta struct {
	IterTesta
	ArTesta ArTesta
}



func (t *ArIterTesta) Next() bool {
	ok := t.IterTesta.Next()

 return ok
}

func (t *ArIterTesta) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Testa)
	for t.IterTesta.Next(){
		t.ArTesta = append(t.ArTesta,t.Testa)
	}
 return len(t.ArTesta)
}

func (t *IterTesta) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Testa)
return
}


func (t *IterTesta) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Testa)

}




			


func ( dst TestaParams)SQL(flag db.SQLOper)string{
  x :=   "id,t"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "testa"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "testa"	,
  		)


			
  }

	return ""

}




func ( dst Testa)SQL(flag db.SQLOper)string{
  x :=   "id,t"

  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"testa"	,
			 "t"  ,
			 "$1" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"testa"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"testa"	,
						"t=$1"  ,
						"id=$2",
					x)

			
  }

	return ""

}
