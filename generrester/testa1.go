
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 11:19:25.05982242 +0200 CEST
//code for table testa1

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyTesta1 time.Time

type Testa1Params struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Testa1 struct {
    Id *int32  `json:"id"`
    Lulu *string  `json:"lulu"`
    Lulu1 *string  `json:"lulu1"`
    Ider *int32  `json:"ider"`
    Md5 *string  `json:"md5"`
    Gaga *string  `json:"gaga"`
    Fk1 *int32  `json:"fk1"`
    Ider1 *int32  `json:"ider1"`
    Lalu *string  `json:"lalu"`
    Zacka *string  `json:"zacka"`
    Fk2 *int32  `json:"fk2"`
    Name *string  `json:"name"`
    Vorname *string  `json:"vorname"`
}

type ArTesta1 []Testa1

func (x *Testa1) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Lulu,
				&x.Lulu1,
				&x.Ider,
				&x.Md5,
				&x.Gaga,
				&x.Fk1,
				&x.Ider1,
				&x.Lalu,
				&x.Zacka,
				&x.Fk2,
				&x.Name,
				&x.Vorname,
			}

}

func (rt *Testa1)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Lulu,	&rt.Lulu1,	&rt.Ider,	&rt.Md5,	&rt.Gaga,	&rt.Fk1,	&rt.Ider1,	&rt.Lalu,	&rt.Zacka,	&rt.Fk2,	&rt.Name,	&rt.Vorname,
    
  }
}


func (rt *Testa1Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}


func (dst *ArTesta1) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Testa1)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


			


func SQLTesta1(flag db.SQLOper)string{
  x :=   "id,lulu,lulu1,ider,md5,gaga,fk1,ider1,lalu,zacka,fk2,name,vorname"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "testa1"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "testa1"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"testa1"	,
			 "lulu,lulu1,ider,md5,gaga,fk1,ider1,lalu,zacka,fk2,name,vorname"  ,
			 "$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"testa1"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"testa1"	,
						"lulu=$1,lulu1=$2,ider=$3,md5=$4,gaga=$5,fk1=$6,ider1=$7,lalu=$8,zacka=$9,fk2=$10,name=$11,vorname=$12"  ,
						"id=$13",
					x)

			
  }

	return ""

}
