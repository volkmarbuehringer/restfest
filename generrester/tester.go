
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 11:05:01.907445929 +0200 CEST
//code for table tester

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyTester time.Time

type TesterParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Tester struct {
    Name *string  `json:"name"`
    Vorname *string  `json:"vorname"`
    Id int64  `json:"id"`
    Code *string  `json:"code"`
    Lala *string  `json:"lala"`
}

type ArTester []Tester



func (t Tester) Columns() []string {
	return []string{	 "name" ,	 "vorname" ,	 "id" ,	 "code" ,	 "lala" ,
			}
}


func (x *Tester) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Name,
				&x.Vorname,
				&x.Id,
				&x.Code,
				&x.Lala,
			}

}

func (rt *Tester)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Name,	&rt.Vorname,	&rt.Code,	&rt.Lala,
    
  }
}


func (rt *TesterParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapTester map[int64]Tester
func (dst *MapTester) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Tester)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[x.Id] = *x
	}

	return nil
}
	

func (dst *ArTester) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Tester)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterTester struct {
	db.BaseCopy
	Tester Tester
}



func (t *IterTester) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Tester)
return
}


func (t *IterTester) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Tester)

}




			


func ( dst TesterParams)SQL(flag db.SQLOper)string{
  x :=   "name,vorname,id,code,lala"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "tester"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "tester"	,
  		)


			
  }

	return ""

}




func ( dst Tester)SQL(flag db.SQLOper)string{
  x :=   "name,vorname,id,code,lala"

  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"tester"	,
			 "name,vorname,code,lala"  ,
			 "$1,$2,$3,$4" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tester"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"tester"	,
						"name=$1,vorname=$2,code=$3,lala=$4"  ,
						"id=$5",
					x)

			
  }

	return ""

}
