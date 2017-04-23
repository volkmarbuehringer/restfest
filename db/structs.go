package db

import (
	"fmt"
	"strconv"
	"time"
)

func (arr InterPgx) ConvertItoS() (record []string, err error) {
	record = make([]string, len(arr))

	for i := range arr {
		switch arr[i].(type) {
		case string:
			record[i] = arr[i].(string)
		case *string:
			t := arr[i].(*string)
			if t != nil {
				record[i] = *t
			} else {
				record[i] = ""
			}
		case int32:
			record[i] = strconv.Itoa((int)(arr[i].(int32)))
		case int64:
			record[i] = strconv.Itoa((int)(arr[i].(int64)))
		case *int32:
			t := arr[i].(*int32)
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case *int64:
			t := arr[i].(*int64)
			if t != nil {
				record[i] = strconv.Itoa((int)(*t))
			} else {
				record[i] = ""
			}
		case float64:
			record[i] = strconv.FormatFloat(arr[i].(float64), 'E', -1, 64)
		case *float64:
			t := arr[i].(*float64)
			if t != nil {
				record[i] = strconv.FormatFloat(*t, 'E', -1, 64)
			} else {
				record[i] = ""
			}
		case bool:
			record[i] = strconv.FormatBool(arr[i].(bool))
		case *time.Time:
			t := arr[i].(*time.Time)
			if t != nil {
				record[i] = t.Format("2006-01-01")
			} else {
				record[i] = ""
			}
		case time.Time:
			record[i] = (arr[i].(time.Time)).Format("2006-01-01")
		default:
			if arr[i] != nil {
				err = fmt.Errorf("unknown %d %T", i, arr[i])

			} else {
				err = fmt.Errorf("unknown %d ", i)

			}

		}

	}
	return
}

func (arrr *InterPgx) ConvertStoI(record []string) error {
	var err error
	arr := *arrr

	if len(arr) != len(record) {
		err = fmt.Errorf("längen stimmen nicht %d %d", len(record), len(arr))
		return err
	}
	for i := range arr {

		switch arr[i].(type) {
		case string:
			arr[i] = record[i]
		case *string:
			if len(record[i]) == 0 {
				arr[i] = nil
			} else {

				arr[i] = new(string)
				arr[i] = record[i]
			}
		case int32, int64, int:
			arr[i], err = strconv.Atoi(record[i])

		case *int32:
			if len(record[i]) == 0 {
				arr[i] = nil
			} else {
				var tt int
				tt, err = strconv.Atoi(record[i])
				arr[i] = (int32)(tt)
			}

		case *int64:
			if len(record[i]) == 0 {
				arr[i] = nil
			} else {
				var tt int
				tt, err = strconv.Atoi(record[i])
				arr[i] = (int64)(tt)
			}
		case float64:
			arr[i], err = strconv.ParseFloat(record[i], 64)
		case *float64:
			if len(record[i]) == 0 {
				arr[i] = nil
			} else {
				t := new(float64)
				*t, err = strconv.ParseFloat(record[i], 64)
				arr[i] = t
			}
		case bool:
			arr[i], err = strconv.ParseBool(record[i])
		case *time.Time:
			if len(record[i]) == 0 {
				arr[i] = nil
			} else {

				t := new(time.Time)

				*t, err = time.Parse(record[i], "2006-01-01")
				arr[i] = *t
			}
		case time.Time:
			arr[i], err = time.Parse(record[i], "2006-01-01")
		default:
			if arr[i] != nil {
				err = fmt.Errorf("unknown %d %T", i, arr[i])

			} else {
				err = fmt.Errorf("unknown %d ", i)

			}

		}
		if err != nil {
			fmt.Println("errr", err)
			return err
		}

	}
	return err
}
