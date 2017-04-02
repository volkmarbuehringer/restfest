package gener

import (
	"database/sql"
	"encoding/json"
)

type SQLOper int

const (
	GenSelect SQLOper = iota
	GenInsert
	GenUpdate
)

type mapperFun func(rows *sql.Rows, len int) (stru []interface{}, err error)
type mapperFun1 func(interface{}) []interface{}
type mapperFun2 func(*sql.Stmt, interface{}, int) (*sql.Rows, error)

type JSONNullInt64 struct {
	sql.NullInt64
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
