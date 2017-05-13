
// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by robots at
// 2017-05-13 09:31:50.992199216 +0200 CEST
//code for table lala

package gener



import (
	"restfest/db"
	"fmt"
	"time"
	"io"
	"strings"
		"encoding/json"
	"github.com/jackc/pgx"
			
			"github.com/jackc/pgx/pgtype"
		
)

type dummyLala time.Time

type LalaParams struct {
    Length int
    Offset int
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



var  LalaSQL  =db.ColumnLists{
	 "id,maxer,miner,da,haha,mist,lump",
	  "maxer,miner,da,haha,mist,lump"  ,
		 "$1,$2,$3,$4,$5,$6" ,
	   "id" ,
		 "id=$7",
		"maxer,miner,da,haha,mist,lump"	 ,
						"maxer=$1,miner=$2,da=$3,haha=$4,mist=$5,lump=$6"  ,

}


func (x *LalaParams) Scanner() db.InterPgx {
	return db.InterPgx{
				&x.Length,
				&x.Offset,
			}

}

func (x Lala) Reader(mapper map[string][]string) [][]string {
	return [][]string{
				}

}


func (x LalaParams) Reader(mapper map[string][]string) [][]string {
	return [][]string{
		mapper["length"],
		mapper["offset"],
			}

}



func (t Lala) Columns() []string {
	return strings.Split(LalaSQL.All,",")
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



func (t Lala) Writer(w io.Writer,trenner string)  error {
		w.Write([]byte(trenner))

if err := json.NewEncoder(w).Encode(t); err != nil {
	return err
}
	return nil
}




func ( dst LalaParams)SQL(flag db.SQLOper)string{


  switch flag{
			
    case db.GenSelectID,  db.GenSelectAll:
    return fmt.Sprintf(db.SQLPattern[flag],
  LalaSQL.All,
	     "lala"	,
      LalaSQL.PK ,
    )

		case db.GenSelectAll1:
		return fmt.Sprintf(db.SQLPattern[flag],
	LalaSQL.All,
   "lala"	,
  		)


			
  }

	return ""

}




func ( dst Lala)SQL(flag db.SQLOper)string{


  switch flag{
			


		case db.GenInsert:
		return  fmt.Sprintf( db.SQLPattern[flag],
			"lala"	,
			 LalaSQL.Inserts  ,
			 LalaSQL.BindsInsert ,
				LalaSQL.All,)

				case db.GenDelete:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"lala"	,
					LalaSQL.PK,
						LalaSQL.All,)
				case db.GenUpdate:
				return  fmt.Sprintf(db.SQLPattern[flag]  ,
					"lala"	,
					LalaSQL.BindsUpdate  ,
								LalaSQL.PKUpdate,
					LalaSQL.All,)

			
  }

	return ""

}



type IterLala struct {
	db.BaseCopy
	Lala Lala
}

type ArIterLala struct {
	IterLala
	ArLala ArLala
}


				

type MapLala map[int32]Lala

type MapIterLala struct {
		IterLala
		MapLala  MapLala
}

				func (t *MapIterLala) NewCopy(rows  *pgx.Rows) int {
				 t.BaseCopy.NewCopy(rows,&t.Lala)
				 t.MapLala=make(MapLala)
					for t.IterLala.Next(){
						t.MapLala[*t.Lala.Id] = t.Lala
					}
				 return len(t.MapLala)
				}


	



func (t *ArIterLala) NewCopy(rows  *pgx.Rows) int {
 t.BaseCopy.NewCopy(rows,&t.Lala)
	for t.IterLala.Next(){
		t.ArLala = append(t.ArLala,t.Lala)
	}
 return len(t.ArLala)
}

func (t *IterLala) NewCopy(rows  *pgx.Rows)  {

 t.BaseCopy.NewCopy(rows,&t.Lala)
return
}

func (t *IterLala) Value() db.PgxGener  {

 return &t.Lala
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



