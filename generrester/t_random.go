
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 11:15:22.075617734 +0200 CEST
//code for table t_random

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
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



func (t T_random) Columns() []string {
	return []string{	 "id" ,	 "md5" ,	 "gaga" ,	 "fk" ,
			}
}


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
				
				type MapT_random map[int32]T_random
func (dst *MapT_random) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( T_random)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[x.Id] = *x
	}

	return nil
}
	

func (dst *ArT_random) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( T_random)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterT_random struct {
	db.BaseCopy
	T_random T_random
}



func (t *IterT_random) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.T_random)
return
}


func (t *IterT_random) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.T_random)

}




			


func ( dst T_randomParams)SQL(flag db.SQLOper)string{
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


			
  }

	return ""

}




func ( dst T_random)SQL(flag db.SQLOper)string{
  x :=   "id,md5,gaga,fk"

  switch flag{
			


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
