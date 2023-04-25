package utils

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

type NullString sql.NullString

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return json.Marshal("")
}

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// Value - Implementation of valuer for database/sql
func (ns NullString) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as string.
	return ns.String, nil
}

type NullInt32 sql.NullInt32

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int32)
	}
	return json.Marshal(0)
}

// Scan implements the Scanner interface for NullString
func (ni *NullInt32) Scan(value interface{}) error {
	var n sql.NullInt32
	if err := n.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt32{n.Int32, false}
	} else {
		*ni = NullInt32{n.Int32, true}
	}

	return nil
}

// Value - Implementation of valuer for database/sql
func (ni NullInt32) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as Int32.
	return ni.Int32, nil
}

type NullTime sql.NullTime

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	}
	return json.Marshal("")
}

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	var n sql.NullTime
	if err := n.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{n.Time, false}
	} else {
		*nt = NullTime{n.Time, true}
	}

	return nil
}

// Value - Implementation of valuer for database/sql
func (ni NullTime) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as Int32.
	return ni.Time, nil
}

type NullFloat64 sql.NullFloat64

func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Float64)
	}
	return json.Marshal(0.0)
}

// Scan implements the Scanner interface for NullString
func (ni *NullFloat64) Scan(value interface{}) error {
	var n sql.NullFloat64
	if err := n.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullFloat64{n.Float64, false}
	} else {
		*ni = NullFloat64{n.Float64, true}
	}

	return nil
}

// Value - Implementation of valuer for database/sql
func (ni NullFloat64) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as FLoat64.
	return ni.Float64, nil
}
