package option

import "strconv"

// ================================ Uint ================================

// Uint is the option type for unsigned integers (uint).
type Uint *uint

// ZeroUint is a Uint with 0 as value.
var ZeroUint = NewUint(0)

// NewUint creates a new Uint using the value of the passed uint.
func NewUint(u uint) Uint { return &u }

// ================================ Int ================================

// Int is the option type for integers (int).
type Int *int

// ZeroInt is an Int with 0 as value.
var ZeroInt = NewInt(0)

// NewInt creates a new Int using the value of the passed int.
func NewInt(i int) Int { return &i }

// ================================ Float64 ================================

// Float64 is the option type for float64s.
type Float64 *float64

// ZeroInt is an Int with 0 as value.
var ZeroFloat64 = NewFloat64(0)

// NewInt creates a new Int using the value of the passed float64.
func NewFloat64(f float64) Float64 { return &f }

// ================================ NullableUint ================================

// NullableUint is a nullable version of an unsigned integer (uint).
type NullableUint = *NullableUintData

type NullableUintData struct {
	Val  uint
	Init bool
}

// NullUint serializes to JSON null.
var NullUint = &NullableUintData{}

// NewUint creates a new non-null NullableUint using the value of the passed uint.
func NewNullableUint(v uint) NullableUint {
	return &NullableUintData{
		Val:  v,
		Init: true,
	}
}

func (u NullableUintData) MarshalJSON() ([]byte, error) {
	if !u.Init {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(u.Val), 10)), nil
}

func (u *NullableUintData) UnmarshalJSON(json []byte) error {
	s := string(json)

	if s == "null" {
		*u = *NullUint
		return nil
	}

	v, err := strconv.ParseUint(s, 10, 64)

	u.Val = uint(v)
	u.Init = true

	return err
}

// ================================ NullableInt ================================

// NullableInt is a nullable version of an integer (int).
type NullableInt = *NullableIntData

type NullableIntData struct {
	Val  int
	Init bool
}

// NullInt serializes to JSON null.
var NullInt = &NullableIntData{}

// NewInt creates a new non-null NullableInt using the value of the passed int.
func NewNullableInt(v int) NullableInt {
	return &NullableIntData{
		Val:  v,
		Init: true,
	}
}

func (i NullableIntData) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(i.Val), 10)), nil
}

func (i *NullableIntData) UnmarshalJSON(json []byte) error {
	s := string(json)

	if s == "null" {
		*i = *NullInt
		return nil
	}

	v, err := strconv.ParseUint(s, 10, 64)

	i.Val = int(v)
	i.Init = true

	return err
}

// ================================ NullableFloat64 ================================

// NullableFloat64 is a nullable version of a float64.
type NullableFloat64 = *NullableFloat64Data

type NullableFloat64Data struct {
	Val  float64
	Init bool
}

// NullFloat64 serializes to JSON null.
var NullFloat64 = &NullableFloat64Data{}

// NewFloat64 creates a new non-null NullableFloat64 using the value of the passed float64.
func NewNullableFloat64(v float64) NullableFloat64 {
	return &NullableFloat64Data{
		Val:  v,
		Init: true,
	}
}

func (i NullableFloat64Data) MarshalJSON() ([]byte, error) {
	if !i.Init {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatFloat(float64(i.Val), 'g', -1, 64)), nil
}

func (i *NullableFloat64Data) UnmarshalJSON(json []byte) error {
	s := string(json)

	if s == "null" {
		*i = *NullFloat64
		return nil
	}

	v, err := strconv.ParseFloat(s, 64)

	i.Val = float64(v)
	i.Init = true

	return err
}
