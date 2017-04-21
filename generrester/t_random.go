
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-22 00:33:09.150473468 +0200 CEST
//code for table t_random

package generrester



import (
	"restfest/db"
	"fmt"
	"time"

			
)

type dummyT_random time.Time

type T_randomParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type T_random struct {
    Id int32  `json:"id"`
    Md5 *string  `json:"md5"`
    Gaga *string  `json:"gaga"`
    Fk *int32  `json:"fk"`
}

type ArT_random []T_random

func (x *T_random) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Md5,
				&x.Gaga,
				&x.Fk,
			}

}

func (rt *T_random)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Md5,	&rt.Gaga,	&rt.Fk,
    
  }
}


func (rt *T_randomParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}


			


func SQLT_random(flag db.SQLOper)string{
  x :=   "id,md5,gaga,fk"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "t_random"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "t_random"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"t_random"	,
			 "md5,gaga,fk"  ,
			 "$1,$2,$3" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"t_random"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"t_random"	,
						"md5=$1,gaga=$2,fk=$3"  ,
						"id=$4",
					x)

			
  }

	return ""

}
