
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 00:00:09.021896811 +0200 CEST
//code for table weburl

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
			"github.com/jackc/pgx/pgtype"
		
)

type dummyWeburl time.Time

type WeburlParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Weburl struct {
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

type ArWeburl []Weburl



func (t Weburl) Columns() []string {
	return []string{	 "id" ,	 "url" ,	 "zusatz" ,	 "created" ,	 "w_cr_date" ,	 "w_upd_date" ,	 "w_upd_uid" ,	 "w_cr_uid" ,	 "addtime" ,	 "addtime1" ,	 "flag" ,	 "test" ,
			}
}


func (x *Weburl) Scanner() db.InterPgx {
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

func (rt *Weburl)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Url,	&rt.Zusatz,	&rt.Created,	&rt.Addtime,	&rt.Addtime1,	&rt.Flag,	&rt.Test,
    
  }
}


func (rt *WeburlParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapWeburl map[int32]Weburl
func (dst *MapWeburl) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Weburl)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[x.Id] = *x
	}

	return nil
}
	

func (dst *ArWeburl) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Weburl)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterWeburl struct {
	db.BaseCopy
	Weburl Weburl
}



func (t *IterWeburl) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Weburl)
return
}


func (t *IterWeburl) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Weburl)

}




			


func (src *ArWeburl) AssignTo(dst interface{}) error {

	if src != nil {
		ttt, ok := dst.(*ArWeburl)
		if !ok {
				return fmt.Errorf("cannot assign %v to %T", src, dst)
		}
		*ttt = *src
	}
	return nil
}

func (dst *ArWeburl) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to ArWeburl", src)
}

func (dst *ArWeburl) Get() interface{} {
	return dst
}





func (src *Weburl) AssignTo(dst interface{}) error {
	if src != nil {
		ttt, ok := dst.(*Weburl)
		if !ok {
				return fmt.Errorf("cannot assign %v to %T", src, dst)
		}
		*ttt = *src
	}
	return nil
}

func (dst *Weburl) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to Weburl", src)
}

func (dst *Weburl) Get() interface{} {
return dst
}

func (dst *Weburl) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
			return nil
	}

	struT := new(Weburl)

d:=  struT.Scanner()

	err := d.DecodeBinary(ci,src)
	if err != nil {
		return err
	}
	*dst = *struT
	return nil
}



func (dst *ArWeburl) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		return nil
	}
	elements := make(ArWeburl, 0)
	funcer := func (result *ArWeburl) func() db.InterPgx {
		return func() db.InterPgx {
			pos := len(*result)
			*result = append(*result,Weburl{})
			return (*result)[pos].Scanner()
		}
	}

	helperfun := funcer(&elements)
	err := db.Helper(ci, src, helperfun)
 if err != nil {
	 return err
 }
	*dst = elements
	return nil
}





func ( dst WeburlParams)SQL(flag db.SQLOper)string{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "weburl"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "weburl"	,
  		)


			
  }

	return ""

}




func ( dst Weburl)SQL(flag db.SQLOper)string{
  x :=   "id,url,zusatz,created,w_cr_date,w_upd_date,w_upd_uid,w_cr_uid,addtime,addtime1,flag,test"

  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"weburl"	,
			 "url,zusatz,created,w_cr_date,w_cr_uid,addtime,addtime1,flag,test"  ,
			 "$1,$2,$3,current_timestamp,'webSrv',$4,$5,$6,$7" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"weburl"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"weburl"	,
						"url=$1,zusatz=$2,created=$3,w_upd_date=current_timestamp,w_upd_uid='webSrv',addtime=$4,addtime1=$5,flag=$6,test=$7"  ,
						"id=$8",
					x)

			
  }

	return ""

}
