
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-27 00:17:36.860758716 +0200 CEST
//code for table guges

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyGuges time.Time

type GugesParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Guges struct {
    Id int32  `json:"id"`
    A Weburl  `json:"a"`
}

type ArGuges []Guges



func (t Guges) Columns() []string {
	return []string{	 "id" ,	 "a" ,
			}
}


func (x *Guges) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.A,
			}

}

func (rt *Guges)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.A,
    
  }
}


func (rt *GugesParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapGuges map[int32]Guges
func (dst *MapGuges) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Guges)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[x.Id] = *x
	}

	return nil
}
	

func (dst *ArGuges) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Guges)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterGuges struct {
	db.BaseCopy
	Guges Guges
}




func (t *IterGuges) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Guges)
return
}


func (t *IterGuges) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Guges)

}




			


func SQLGuges(flag db.SQLOper)string{
  x :=   "id,a"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "guges"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "guges"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"guges"	,
			 "a"  ,
			 "$1" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"guges"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"guges"	,
						"a=$1"  ,
						"id=$2",
					x)

			
  }

	return ""

}
