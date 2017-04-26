
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-24 23:12:09.271946769 +0200 CEST
//code for table get_all_boo_173824

package generteststruct



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"github.com/jackc/pgx"
			
)

type dummyGet_all_boo_173824 time.Time

type Get_all_boo_173824Params struct {
    P_start *int32 `schema:"p_start"`
    P_end *int32 `schema:"p_end"`
    P_len *int32 `schema:"p_len"`
    P_lala *string `schema:"p_lala"`
}


type Get_all_boo_173824 struct {
    Id int32  `json:"id"`
    Url *string  `json:"url"`
    Zusatz *int64  `json:"zusatz"`
    Created *time.Time  `json:"created"`
    W_cr_date *time.Time  `json:"w_cr_date"`
    W_upd_date *time.Time  `json:"w_upd_date"`
    W_upd_uid *string  `json:"w_upd_uid"`
    W_cr_uid *string  `json:"w_cr_uid"`
    Addtime *time.Time  `json:"addtime"`
    Addtime1 *time.Time  `json:"addtime1"`
    Flag bool  `json:"flag"`
    Test *float64  `json:"test"`
}

type ArGet_all_boo_173824 []Get_all_boo_173824



func (t Get_all_boo_173824) Columns() []string {
	return []string{	 "id" ,	 "url" ,	 "zusatz" ,	 "created" ,	 "w_cr_date" ,	 "w_upd_date" ,	 "w_upd_uid" ,	 "w_cr_uid" ,	 "addtime" ,	 "addtime1" ,	 "flag" ,	 "test" ,
			}
}


func (x *Get_all_boo_173824) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Url,
				&x.Zusatz,
				&x.Created,
				&x.W_cr_date,
				&x.W_upd_date,
				&x.W_upd_uid,
				&x.W_cr_uid,
				&x.Addtime,
				&x.Addtime1,
				&x.Flag,
				&x.Test,
			}

}

func (rt *Get_all_boo_173824)ROWInsert() db.InterPgx {
	  return db.InterPgx{
		
  }
}


func (rt *Get_all_boo_173824Params)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.P_start,
		&rt.P_end,
		&rt.P_len,
		&rt.P_lala,
}

}
				

func (dst *ArGet_all_boo_173824) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Get_all_boo_173824)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type BaseCopyGet_all_boo_173824 struct {
	Errc      error
	Inter    db.InterPgx
	Rows     *pgx.Rows
	Get_all_boo_173824 Get_all_boo_173824
}

func (t *BaseCopyGet_all_boo_173824) Values() ([]interface{}, error) {
	return t.Inter, t.Errc
}


func (t *BaseCopyGet_all_boo_173824) ValuesString() (record []string,err error) {
record, err = t.Inter.ConvertItoS()
return
}


func (t BaseCopyGet_all_boo_173824) Err() error {
	if t.Errc != io.EOF {
		return t.Errc
	}
	return nil
}

func (t *BaseCopyGet_all_boo_173824) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) error {

	t.Inter = t.Get_all_boo_173824.Scanner()
t.Rows = rows
	fmt.Println("vor copy", tab)
	copyCount, err := con.CopyFrom(
		[]string{tab},
		t.Get_all_boo_173824.Columns(),
		tt)

	fmt.Println("fertig", copyCount, err)
	if err != nil {
		return err
	}
	return nil
}

func (t *BaseCopyGet_all_boo_173824) Next() bool {
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



			


func SQLGet_all_boo_173824(flag db.SQLOper)string{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test"

  switch flag{
			

		case db.GenFunction:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
 "get_all_boo" ,
 "$1,$2,$3,$4"	,
		)


  }

	return ""

}