package values

// Values works as GoCalculate's main list of values
type Values interface {
	// Returns the Raw Value slice. Used mainly for use with range
	values() []Value

	// Returns the index of val. If val is not in Values it returns -1.
	IndexOf(val Value) int

	// Set the Value val at index
	Set(index int, val Value)

	// Append Value val to Values
	Append(val Value)

	// Returns the Value at index
	Get(index int) Value

	// Returns a subset of Values from start to finish
	Subset(start, finish int) Values

	// Returns a copy of Values.
	Copy() Values

	// Return Length of Values
	Len() int

	// Returns the Core Type of the Values. i.e the highest ranking Type
	Type() Type
}

type values struct {
	vals     []Value
	length   int
	coreType Type
}

func (v *values) setValues(vals []Value) {
	v.vals = make([]Value, len(vals))
	v.length = len(vals)
	v.coreType = Real
	for index, val := range vals {
		if val != nil {
			v.Set(index, val)
		}
	}
}

func (v *values) values() []Value { return v.vals }

func (v *values) Type() Type { return v.coreType }

func (v *values) Len() int { return v.length }

func (v *values) Set(index int, val Value) {
	if !val.IsZero() {
		if v.Type() < val.Type() {
			v.coreType = val.Type()
		}
		v.vals[index] = val
	} else {
		v.vals[index] = nil
	}
}

func (v *values) Get(index int) Value {
	val := v.vals[index]
	if val == nil {
		return Zero()
	}
	return val
}

func (v *values) Append(val Value) {
	vals := append(v.values(), val)
	v.setValues(vals)
}

func (v *values) Copy() Values {
	vals := new(values)
	vElements := make([]Value, len(v.vals))
	for index, val := range v.values() {
		if val != nil {
			vElements[index] = val
		}
	}
	vals.length = v.Len()
	vals.coreType = v.Type()
	vals.vals = vElements
	return vals
}

func (v *values) Subset(start, finish int) Values {
	vals := new(values)
	subVals := make([]Value, len(v.vals[start:finish+1]))
	copy(subVals, v.vals[start:finish+1])
	vals.setValues(subVals)
	return vals
}

func (v *values) IndexOf(val Value) int {
	for index, value := range v.values() {
		if value != nil {
			if value.Type() == Complex && value.Complex() == val.Complex() {
				return index
			} else if value.Real() == val.Real() {
				return index
			}
		}
	}
	return -1
}

// NewValues will return a new Values
// Type is Real
func NewValues(length int) Values {
	newValues := new(values)
	newValues.vals = make([]Value, length)
	newValues.length = length
	newValues.coreType = Real
	return newValues
}

// MakeValuesAlt returns a Values type, but requires a framework []Value slice
func MakeValuesAlt(vals []Value) Values {
	newValues := new(values)
	if vals == nil {
		vals = make([]Value, 0)
	}
	newValues.setValues(vals)
	return newValues
}

// MakeValues returns a Values type, takes in a slice of interfaces.
// in an interface is not a supported type, that interface will be forced to the zero
// Value
func MakeValues(vals ...interface{}) Values {
	var values []Value
	for _, val := range vals {
		values = append(values, MakeValue(val))
	}
	return MakeValuesAlt(values)
}

// RetrieveValues returns a copy of the []Value of vals Values
func RetrieveValues(vals Values) []Value {
	return vals.Copy().values()
}
