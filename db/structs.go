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
