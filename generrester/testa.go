
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 18:34:21.910841548 +0200 CEST
//code for table testa

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyTesta time.Time

type TestaParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Testa struct {
    Id int32  `json:"id"`
    T *map[string]interface{}  `json:"t"`
}

type ArTesta []Testa


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
func (dst *MapTesta) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Testa)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[x.Id] = *x
	}

	return nil
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


			


func SQLTesta(flag db.SQLOper)string{
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
