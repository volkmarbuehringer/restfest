
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 00:00:08.673304442 +0200 CEST
//code for table get_guges1_174200

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
)

type dummyGet_guges1_174200 time.Time

type Get_guges1_174200Params struct {
    P_start *int32 `schema:"p_start"`
    P_end *int32 `schema:"p_end"`
    P_len *int32 `schema:"p_len"`
    P_lala *string `schema:"p_lala"`
}


type Get_guges1_174200 struct {
    Id int32  `json:"id"`
    Agger ArWeburl  `json:"agger"`
    Texter *string  `json:"texter"`
    Zahler *int32  `json:"zahler"`
    Zahler2 *int32  `json:"zahler2"`
    Texter2 *string  `json:"texter2"`
}

type ArGet_guges1_174200 []Get_guges1_174200



func (t Get_guges1_174200) Columns() []string {
	return []string{	 "id" ,	 "agger" ,	 "texter" ,	 "zahler" ,	 "zahler2" ,	 "texter2" ,
			}
}


func (x *Get_guges1_174200) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Agger,
				&x.Texter,
				&x.Zahler,
				&x.Zahler2,
				&x.Texter2,
			}

}

func (rt *Get_guges1_174200)ROWInsert() db.InterPgx {
	  return db.InterPgx{
		
  }
}


func (rt *Get_guges1_174200Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.P_start,
		&rt.P_end,
		&rt.P_len,
		&rt.P_lala,
}

}
				

func (dst *ArGet_guges1_174200) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Get_guges1_174200)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterGet_guges1_174200 struct {
	db.BaseCopy
	Get_guges1_174200 Get_guges1_174200
}



func (t *IterGet_guges1_174200) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Get_guges1_174200)
return
}


func (t *IterGet_guges1_174200) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Get_guges1_174200)

}




			


func ( dst Get_guges1_174200Params)SQL(flag db.SQLOper)string{
  x :=   "id,agger,texter,zahler,zahler2,texter2"

  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
 "get_guges1" ,
 "$1,$2,$3,$4"	,
		)


  }

	return ""

}




func ( dst Get_guges1_174200)SQL(flag db.SQLOper)string{
  x :=   "id,agger,texter,zahler,zahler2,texter2"

  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
 "get_guges1" ,
 "$1,$2,$3,$4"	,
		)


  }

	return ""

}
