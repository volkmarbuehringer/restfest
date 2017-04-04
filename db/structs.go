package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type SQLOper int

const (
	GenSelect SQLOper = iota
	GenInsert
	GenUpdate
)

type MapperFun func(rows *sql.Rows, len int) (stru []interface{}, err error)
type MapperFun1 func(interface{}) []interface{}

type String string

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (v NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	}
	return json.Marshal(nil)

}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JSONNullInt64 struct {
	sql.NullInt64
}

type JSONTime struct {
	time.Time
}

type JSONString struct {
	String
}

func (n JSONString) Value() (driver.Value, error) {
	if len(n.String) == 0 {
		return nil, nil
	} else {
		return string(n.String), nil
	}
}

func (n *JSONString) Scan(value interface{}) error {
	if value == nil {
		n.String = ""
	} else {
		x, ok := value.(string)

		if !ok {
			return fmt.Errorf("string nicht ok")
		}
		n.String = String(x)
	}
	return nil

}

func (v JSONString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String)
}

func (v *JSONString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *String
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.String = *x
	} else {
		v.String = ""
	}

	return nil
}

func (v JSONNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)

}

func (v *JSONNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JSONNullFloat64 struct {
	sql.NullFloat64
}

func (v JSONNullFloat64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Float64)
	}
	return json.Marshal(nil)

}

func (v *JSONNullFloat64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *float64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Float64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JSONNullString struct {
	sql.NullString
}

func (v JSONNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)

}

func (v *JSONNullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

var SQLFunMap = map[string]func(string, SQLOper) []interface{}{}

var ROWInsertFunMap = map[string]MapperFun1{}

var EmptyFunMap = map[string]func() interface{}{}

var ScannerFunMap = map[string]func() ([]interface{}, interface{}){}
