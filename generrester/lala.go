
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-04-29 11:05:01.717410496 +0200 CEST
//code for table lala

package generrester



import (
	"restfest/db"
	"fmt"
	"time"
	"github.com/jackc/pgx"
			
			"github.com/jackc/pgx/pgtype"
		
)

type dummyLala time.Time

type LalaParams struct {
    Length int `schema:"length"`
    Offset int `schema:"offset"`
}


type Lala struct {
    Id *int32  `json:"id"`
    Maxer *int32  `json:"maxer"`
    Miner *int32  `json:"miner"`
    Da *time.Time  `json:"da"`
    Haha ArWeburl  `json:"haha"`
    Mist ArWeburl  `json:"mist"`
    Lump ArWeburl  `json:"lump"`
}

type ArLala []Lala



func (t Lala) Columns() []string {
	return []string{	 "id" ,	 "maxer" ,	 "miner" ,	 "da" ,	 "haha" ,	 "mist" ,	 "lump" ,
			}
}


func (x *Lala) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Id,
				&x.Maxer,
				&x.Miner,
				&x.Da,
				&x.Haha,
				&x.Mist,
				&x.Lump,
			}

}

func (rt *Lala)ROWInsert() db.InterPgx {
	  return db.InterPgx{
			&rt.Maxer,	&rt.Miner,	&rt.Da,	&rt.Haha,	&rt.Mist,	&rt.Lump,
    
  }
}


func (rt *LalaParams)ROWInsert() db.InterPgx {
  return db.InterPgx{
		&rt.Length,
		&rt.Offset,
}

}
				
				type MapLala map[int32]Lala
func (dst *MapLala) Scanner( rows *pgx.Rows) error {

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Lala)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		(*dst)[*x.Id] = *x
	}

	return nil
}
	

func (dst *ArLala) Scanner( rows *pgx.Rows) error {
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		x := new( Lala)
		if err := rows.Scan(x.Scanner()...); err != nil {
			return err
		}
		*dst = append(*dst, *x)
	}
	return nil
}


type IterLala struct {
	db.BaseCopy
	Lala Lala
}



func (t *IterLala) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Lala)
return
}


func (t *IterLala) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource,rows  *pgx.Rows) (int, error) {

	return t.BaseCopy.StartCopy(tab, con, tt, rows, &t.Lala)

}




			


func (src *ArLala) AssignTo(dst interface{}) error {

	if src != nil {
		ttt, ok := dst.(*ArLala)
		if !ok {
				return fmt.Errorf("cannot assign %v to %T", src, dst)
		}
		*ttt = *src
	}
	return nil
}

func (dst *ArLala) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to ArLala", src)
}

func (dst *ArLala) Get() interface{} {
	return dst
}





func (src *Lala) AssignTo(dst interface{}) error {
	if src != nil {
		ttt, ok := dst.(*Lala)
		if !ok {
				return fmt.Errorf("cannot assign %v to %T", src, dst)
		}
		*ttt = *src
	}
	return nil
}

func (dst *Lala) Set(src interface{}) error {
	return fmt.Errorf("cannot convert %v to Lala", src)
}

func (dst *Lala) Get() interface{} {
return dst
}

func (dst *Lala) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
			return nil
	}

	struT := new(Lala)

d:=  struT.Scanner()

	err := d.DecodeBinary(ci,src)
	if err != nil {
		return err
	}
	*dst = *struT
	return nil
}



func (dst *ArLala) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		return nil
	}
	elements := make(ArLala, 0)
	funcer := func (result *ArLala) func() db.InterPgx {
		return func() db.InterPgx {
			pos := len(*result)
			*result = append(*result,Lala{})
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





func ( dst LalaParams)SQL(flag db.SQLOper)string{
  x :=   "id,maxer,miner,da,haha,mist,lump"

  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
    x,
	     "lala"	,
      "id" ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
		x,
   "lala"	,
  		)


			
  }

	return ""

}




func ( dst Lala)SQL(flag db.SQLOper)string{
  x :=   "id,maxer,miner,da,haha,mist,lump"

  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"lala"	,
			 "maxer,miner,da,haha,mist,lump"  ,
			 "$1,$2,$3,$4,$5,$6" ,
				x)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"lala"	,
						"id",
						x)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"lala"	,
						"maxer=$1,miner=$2,da=$3,haha=$4,mist=$5,lump=$6"  ,
						"id=$7",
					x)

			
  }

	return ""

}
