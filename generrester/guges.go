
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 22:56:39.867034226 +0200 CEST
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

func (x *GugesParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Guges) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x GugesParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

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
				type MapIterGuges struct {
					IterGuges
					MapGuges  MapGuges
				}

				func (t *MapIterGuges) Next() bool {
					ok := t.IterGuges.Next()

				 return ok
				}

				func (t *MapIterGuges) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Guges)
					for t.IterGuges.Next(){
						t.MapGuges[t.Guges.Id] = t.Guges
					}
				 return len(t.MapGuges)
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

type ArIterGuges struct {
	IterGuges
	ArGuges ArGuges
}



func (t *ArIterGuges) Next() bool {
	ok := t.IterGuges.Next()

 return ok
}

func (t *ArIterGuges) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Guges)
	for t.IterGuges.Next(){
		t.ArGuges = append(t.ArGuges,t.Guges)
	}
 return len(t.ArGuges)
}

func (t *IterGuges) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Guges)
return
}


func (t *IterGuges) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Guges)

}




			


func ( dst GugesParams)SQL(flag db.SQLOper)string{
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


			
  }

	return ""

}




func ( dst Guges)SQL(flag db.SQLOper)string{
  x :=   "id,a"

  switch flag{
			


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
