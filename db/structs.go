package db

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/jackc/pgx"
)

func Int16p(def string) (int16, error) {
	x, err := strconv.Atoi(def)
	if err != nil {

		return 0, err
	}
	return (int16)(x), nil
}

func Int32p(def string) (int32, error) {
	x, err := strconv.Atoi(def)
	if err != nil {

		return 0, err
	}
	return (int32)(x), nil
}
func Int64p(def string) (int64, error) {
	x, err := strconv.Atoi(def)
	if err != nil {

		return 0, err
	}
	return (int64)(x), nil
}
func Int(def string, errp *error) *int {
	x, err := strconv.Atoi(def)
	if err != nil {
		*errp = err
		return nil
	}
	return &x
}

func Int16(def string, errp *error) *int16 {
	x, err := strconv.Atoi(def)
	if err != nil {
		*errp = err
		return nil
	}
	var xx = (int16)(x)
	return &xx
}

func Int32(def string, errp *error) *int32 {
	x, err := strconv.Atoi(def)
	if err != nil {
		*errp = err
		return nil
	}
	var xx = (int32)(x)
	return &xx
}

func Int64(def string, errp *error) *int64 {
	x, err := strconv.Atoi(def)
	if err != nil {
		*errp = err
		return nil
	}
	var xx = (int64)(x)
	return &xx
}

func String(defaulter string) *string {
	if len(defaulter) > 0 {
		t := new(string)
		*t = defaulter
		return t
	}
	return nil
}

type BaseCopy struct {
	Errc  error
	Inter InterPgx
	Rows  *pgx.Rows
}

type ColumnLists struct {
	All            string
	Inserts        string
	BindsInsert    string
	PK             string
	PKUpdate       string
	BindsVarInsert string
	BindsUpdate    string
}

func (t *BaseCopy) NewCopy(rows *pgx.Rows, stru PgxGener) {

	t.Inter = stru.Scanner()
	t.Rows = rows
	return
}

func (t *BaseCopy) StartCopy(tab string, con *pgx.Conn, tt pgx.CopyFromSource, rows *pgx.Rows, stru PgxGener) (int, error) {

	t.Inter = stru.Scanner()
	t.Rows = rows

	copyCount, err := con.CopyFrom(
		[]string{tab},
		stru.Columns(),
		tt)

	if err != nil {
		return 0, err
	}
	return copyCount, nil
}

func (t *BaseCopy) Values() ([]interface{}, error) {
	return t.Inter, t.Errc
}

func (t *BaseCopy) ValuesString() (record []string, err error) {
	record, err = t.Inter.ConvertItoS()
	return
}

func (t BaseCopy) Err() error {
	if t.Errc != io.EOF {
		return t.Errc
	}
	return nil
}

func (t *BaseCopy) Next() bool {
	var ok bool
	if t.Rows == nil {
		return false
	}
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

func (arr InterPgx) ConvertItoS() (record []string, err error) {
	record = make([]string, len(arr))

	for i, val := range arr {
		switch val.(type) {
		case *string:

			record[i] = *(val.(*string))

		case **string:
			t := *(val.(**string))
			if t != nil {
				record[i] = *t
			} else {
				record[i] = ""
			}
		case *int32:
			record[i] = strconv.Itoa((int)(*val.(*int32)))
		case *int64:
			record[i] = strconv.Itoa((int)(*val.(*int64)))
		case *int16:
			record[i] = strconv.Itoa((int)(*val.(*int16)))
		case **int32:
			t := *(val.(**int32))
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case **int64:
			t := *(val.(**int64))
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case **int16:
			t := *(val.(**int16))
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case *float64:
			record[i] = strconv.FormatFloat(*val.(*float64), 'E', -1, 64)
		case **float64:
			t := *(val.(**float64))
			if t != nil {
				record[i] = strconv.FormatFloat(*t, 'E', -1, 64)
			} else {
				record[i] = ""
			}
		case *bool:
			record[i] = strconv.FormatBool(*val.(*bool))
		case **bool:
			t := *(val.(**bool))
			if t != nil {
				record[i] = strconv.FormatBool(*t)
			} else {
				record[i] = ""
			}
		case **time.Time:
			t := *(val.(**time.Time))
			if t != nil {
				record[i] = (*t).Format("2006-01-01")
			} else {
				record[i] = ""
			}
		case *time.Time:
			record[i] = (*val.(*time.Time)).Format("2006-01-01")
		default:
			if val != nil {
				err = fmt.Errorf("unknown %d %T", i, val)

			} else {
				err = fmt.Errorf("unknown %d ", i)

			}

		}

	}
	return
}

func (arr InterPgx) ConvertStoIS(stringer string, i int, val interface{}) error {
	var err error
	switch val.(type) {
	case *string:
		t := val.(*string)
		*t = stringer
	case **string:
		t := val.(**string)

		*t = String(stringer)
	case *int16:
		if len(stringer) > 0 {
			t := val.(*int16)
			*t, err = Int16p(stringer)
		}

	case *int32:
		if len(stringer) > 0 {
			t := val.(*int32)
			*t, err = Int32p(stringer)
		}
	case *int64:
		if len(stringer) > 0 {
			t := val.(*int64)
			*t, err = Int64p(stringer)
		}
	case *int:
		if len(stringer) > 0 {
			t := val.(*int)
			*t, err = strconv.Atoi(stringer)
		}
	case **int32:
		t := val.(**int32)
		if len(stringer) == 0 {
			*t = nil
		} else {
			*t = Int32(stringer, &err)
		}
	case **int:
		t := val.(**int)
		if len(stringer) == 0 {
			*t = nil
		} else {
			*t = Int(stringer, &err)
		}
	case **int16:
		t := val.(**int16)
		if len(stringer) == 0 {
			*t = nil
		} else {
			*t = Int16(stringer, &err)
		}
	case **int64:
		t := val.(**int64)
		if len(stringer) == 0 {
			*t = nil
		} else {
			*t = Int64(stringer, &err)
		}
	case *float64:
		t := val.(*float64)
		*t, err = strconv.ParseFloat(stringer, 64)

	case **float64:
		t := val.(**float64)
		if len(stringer) == 0 {
			*t = nil
		} else {
			*t = new(float64)
			**t, err = strconv.ParseFloat(stringer, 64)

		}

	case *bool:
		t := val.(*bool)
		*t, err = strconv.ParseBool(stringer)
	case **bool:
		t := val.(**bool)
		if len(stringer) == 0 {
			*t = nil
		} else {
			*t = new(bool)
			**t, err = strconv.ParseBool(stringer)

		}

	case **time.Time:
		t := val.(**time.Time)
		if len(stringer) == 0 {
			*t = nil
		} else {

			*t = new(time.Time)

			**t, err = time.Parse("2006-01-02", stringer)

		}
	case *time.Time:
		t := val.(*time.Time)
		*t, err = time.Parse("2006-01-02", stringer)
	default:
		if arr[i] != nil {
			return fmt.Errorf("unknown %d %T", i, arr[i])

		} else {
			return fmt.Errorf("unknown %d ", i)

		}

	}
	return err
}

func (arr InterPgx) ConvertAtoI(record [][]string) error {
	if len(arr) != len(record) {
		return fmt.Errorf("längen stimmen nicht %d %d", len(record), len(arr))
	}
	for i, val := range arr {
		if len(record[i]) > 0 {
			stringer := record[i][0]
			err := arr.ConvertStoIS(stringer, i, val)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (arr InterPgx) ConvertStoI(record []string) error {

	if len(arr) != len(record) {
		return fmt.Errorf("längen stimmen nicht %d %d", len(record), len(arr))
	}
	for i, val := range arr {
		stringer := record[i]
		err := arr.ConvertStoIS(stringer, i, val)
		if err != nil {
			return err
		}

	}
	return nil
}
