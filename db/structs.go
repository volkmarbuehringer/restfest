package db

import (
	"fmt"
	"strconv"
	"time"
)

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

func (arr InterPgx) ConvertStoI(record []string) error {
	var err error

	if len(arr) != len(record) {
		return fmt.Errorf("längen stimmen nicht %d %d", len(record), len(arr))
	}
	for i, val := range arr {
		stringer := record[i]
		switch val.(type) {
		case *string:
			t := val.(*string)
			*t = stringer
		case **string:
			t := val.(**string)

			if len(stringer) == 0 {
				*t = nil
			} else {

				*t = new(string)
				**t = stringer
			}
		case *int16:
			t := val.(*int16)
			var g int
			g, err = strconv.Atoi(stringer)
			*t = (int16)(g)
		case *int32:
			t := val.(*int32)
			var g int
			g, err = strconv.Atoi(stringer)
			*t = (int32)(g)
		case *int64:
			t := val.(*int64)
			var g int
			g, err = strconv.Atoi(stringer)
			*t = (int64)(g)
		case **int32:
			t := val.(**int32)
			if len(stringer) == 0 {
				*t = nil
			} else {
				var tt int
				tt, err = strconv.Atoi(stringer)
				var g = (int32)(tt)
				*t = &g
			}
		case **int16:
			t := val.(**int16)
			if len(stringer) == 0 {
				*t = nil
			} else {
				var tt int
				tt, err = strconv.Atoi(stringer)
				var g = (int16)(tt)
				*t = &g
			}
		case **int64:
			t := val.(**int64)
			if len(stringer) == 0 {
				*t = nil
			} else {
				var tt int
				tt, err = strconv.Atoi(stringer)
				var g = (int64)(tt)
				*t = &g
			}
		case *float64:
			t := val.(*float64)
			*t, err = strconv.ParseFloat(stringer, 64)

		case **float64:
			t := val.(**float64)
			if len(record[i]) == 0 {
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
			if len(record[i]) == 0 {
				*t = nil
			} else {
				*t = new(bool)
				**t, err = strconv.ParseBool(stringer)

			}

		case **time.Time:
			t := val.(**time.Time)
			if len(record[i]) == 0 {
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

		if err != nil {

			return err
		}

	}
	return err
}
