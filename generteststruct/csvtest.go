
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-24 23:12:08.825540816 +0200 CEST
//code for table csvtest

package generteststruct



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"github.com/jackc/pgx"
			
)

type dummyCsvtest time.Time

type CsvtestParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Csvtest struct {
    Year int32  `json:"year"`
    Anzsic string  `json:"anzsic"`
    Area string  `json:"area"`
    Geo *int32  `json:"geo"`
    Ec *int32  `json:"ec"`
    Id int32  `json:"id"`
}

type ArCsvtest []Csvtest



func (t Csvtest) Columns() []string {
	return []string{	 "year" ,	 "anzsic" ,	 "area" ,	 "geo" ,	 "ec" ,	 "id" ,
			}
}


func (x *Csvtest) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Year,
				&x.Anzsic,
				&x.Area,
				&x.Geo,
				&x.Ec,
				&x.Id,
			}

}

func (rt *Csvtest)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Year,	&rt.Anzsic,	&rt.Area,	&rt.Geo,	&rt.Ec,
    
  }
}


func (rt *CsvtestParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapCsvtest map[int32]Csvtest
func (dst *MapCsvtest) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Csvtest)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[x.Id] = *x
	}

	return nil
}
	

func (dst *ArCsvtest) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Csvtest)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type BaseCopyCsvtest struct {
	Errc      error
	Inter    db.InterPgx
	Rows     *pgx.Rows
	Csvtest Csvtest
}

func (t *BaseCopyCsvtest) Values() ([]interface{}, error) {
	return t.Inter, t.Errc
}


func (t *BaseCopyCsvtest) ValuesString() (record []string,err error) {
record, err = t.Inter.ConvertItoS()
return
}


func (t BaseCopyCsvtest) Err() error {
	if t.Errc != io.EOF {
		return t.Errc
	}
	return nil
}

func (t *BaseCopyCsvtest) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) error {

	t.Inter = t.Csvtest.Scanner()
t.Rows = rows
	fmt.Println("vor copy", tab)
	copyCount, err := con.CopyFrom(
		[]string{tab},
		t.Csvtest.Columns(),
		tt)

	fmt.Println("fertig", copyCount, err)
	if err != nil {
		return err
	}
	return nil
}

func (t *BaseCopyCsvtest) Next() bool {
	var ok bool
	for {
		ok = t.Rows.Next()
		if !ok {
			break
		}
		t.Rows.Scan(t.Inter...)
		break
	}
	t.Errc = t.Rows.Err()
	if t.Errc != nil {
		return false
	}
	return ok

}



			


func SQLCsvtest(flag db.SQLOper)string{
  x :=   "year,anzsic,area,geo,ec,id"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "csvtest"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "csvtest"	,
  		)

		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"csvtest"	,
			 "year,anzsic,area,geo,ec"  ,
			 "$1,$2,$3,$4,$5" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"csvtest"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"csvtest"	,
						"year=$1,anzsic=$2,area=$3,geo=$4,ec=$5"  ,
						"id=$6",
					x)

			
  }

	return ""

}