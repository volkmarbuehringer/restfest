
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 18:34:21.819291292 +0200 CEST
//code for table lala1

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyLala1 time.Time

type Lala1Params struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Lala1 struct {
    Id *int32  `json:"id"`
    Maxer *int32  `json:"maxer"`
    Miner *int32  `json:"miner"`
    Wa *time.Time  `json:"wa"`
    Dudu ArLala  `json:"dudu"`
}

type ArLala1 []Lala1


func (x *Lala1) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Maxer,
				&x.Miner,
				&x.Wa,
				&x.Dudu,
			}

}

func (rt *Lala1)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Maxer,	&rt.Miner,	&rt.Wa,	&rt.Dudu,
    
  }
}


func (rt *Lala1Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapLala1 map[int32]Lala1
func (dst *MapLala1) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Lala1)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[*x.Id] = *x
	}

	return nil
}
	

func (dst *ArLala1) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Lala1)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


			


func SQLLala1(flag db.SQLOper)string{
  x :=   "id,maxer,miner,wa,dudu"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "lala1"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "lala1"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"lala1"	,
			 "maxer,miner,wa,dudu"  ,
			 "$1,$2,$3,$4" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"lala1"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"lala1"	,
						"maxer=$1,miner=$2,wa=$3,dudu=$4"  ,
						"id=$5",
					x)

			
  }

	return ""

}
