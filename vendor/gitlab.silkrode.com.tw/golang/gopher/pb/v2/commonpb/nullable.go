package commonpb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"time"
)

// UnmarshalJSON ...
func (a *NullDouble) UnmarshalJSON(b []byte) error {
	var val float64
	if err := json.Unmarshal(b, &val); err != nil {
		a.Valid = false
		return err
	}
	a.Float64 = val
	return nil
}

// MarshalJSON ...
func (a NullDouble) MarshalJSON() ([]byte, error) {
	if !a.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(a.Float64)
}

// Scan implements the Scanner interface.
func (a *NullDouble) Scan(value interface{}) error {
	var i sql.NullFloat64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*a = NullDouble{Valid: false, Float64: i.Float64}
	} else {
		*a = NullDouble{Valid: true, Float64: i.Float64}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (a NullDouble) Value() (driver.Value, error) {
	if !a.Valid {
		return nil, nil
	}
	return a.Float64, nil
}

// UnmarshalJSON ...
func (a *NullInt64) UnmarshalJSON(b []byte) error {
	var val int64
	if err := json.Unmarshal(b, &val); err != nil {
		a.Valid = false
		return err
	}
	a.Int64 = val
	return nil
}

// MarshalJSON ...
func (a NullInt64) MarshalJSON() ([]byte, error) {
	if !a.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(a.Int64)
}

// Scan implements the Scanner interface.
func (a *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*a = NullInt64{Valid: false, Int64: i.Int64}
	} else {
		*a = NullInt64{Valid: true, Int64: i.Int64}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (a NullInt64) Value() (driver.Value, error) {
	if !a.Valid {
		return nil, nil
	}
	return a.Int64, nil
}

// UnmarshalJSON ...
func (a *NullInt32) UnmarshalJSON(b []byte) error {
	var val int32
	if err := json.Unmarshal(b, &val); err != nil {
		a.Valid = false
		return err
	}
	a.Int32 = val
	return nil
}

// MarshalJSON ...
func (a NullInt32) MarshalJSON() ([]byte, error) {
	if !a.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(a.Int32)
}

// Scan implements the Scanner interface.
func (a *NullInt32) Scan(value interface{}) error {
	var i sql.NullInt32
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*a = NullInt32{Valid: false, Int32: i.Int32}
	} else {
		*a = NullInt32{Valid: true, Int32: i.Int32}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (a NullInt32) Value() (driver.Value, error) {
	if !a.Valid {
		return nil, nil
	}
	return a.Int32, nil
}

// UnmarshalJSON ...
func (a *NullString) UnmarshalJSON(b []byte) error {
	var val string
	if err := json.Unmarshal(b, &val); err != nil {
		a.Valid = false
		return err
	}
	a.String_ = val
	return nil
}

// MarshalJSON ...
func (a NullString) MarshalJSON() ([]byte, error) {
	if !a.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(a.String_)
}

// Scan implements the Scanner interface.
func (a *NullString) Scan(value interface{}) error {
	var i sql.NullString
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*a = NullString{Valid: false, String_: i.String}
	} else {
		*a = NullString{Valid: true, String_: i.String}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (a NullString) Value() (driver.Value, error) {
	if !a.Valid {
		return nil, nil
	}
	return a.String_, nil
}

// UnmarshalJSON ...
func (a *NullFloat) UnmarshalJSON(b []byte) error {
	var val float32
	if err := json.Unmarshal(b, &val); err != nil {
		a.Valid = false
		return err
	}
	a.Float32 = val
	return nil
}

// MarshalJSON ...
func (a NullFloat) MarshalJSON() ([]byte, error) {
	if !a.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(a.Float32)
}

// UnmarshalJSON ...
func (a *NullTime) UnmarshalJSON(b []byte) error {
	var val time.Time
	if err := json.Unmarshal(b, &val); err != nil {
		a.Valid = false
		return err
	}
	a.Time = val
	return nil
}

// MarshalJSON ...
func (a NullTime) MarshalJSON() ([]byte, error) {
	if !a.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(a.Time)
}

// Scan implements the Scanner interface.
func (a *NullTime) Scan(value interface{}) error {
	var i sql.NullTime
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*a = NullTime{Valid: false, Time: i.Time}
	} else {
		*a = NullTime{Valid: true, Time: i.Time}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (a NullTime) Value() (driver.Value, error) {
	if !a.Valid {
		return nil, nil
	}
	return a.Time, nil
}
