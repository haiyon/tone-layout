package utils

import (
	"net/mail"
	"reflect"
	"regexp"
)

// IsPhoneValid - verify phone number
func IsPhoneValid(phone string) bool {
	// reference: https://learnku.com/articles/31543
	phoneCN, _ := regexp.Compile(`^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`)
	if phoneCN.MatchString(phone) {
		return true
	}
	global, _ := regexp.Compile(`\+(9[976]\d|8[987530]\d|6[987]\d|5[90]\d|42\d|3[875]\d|
2[98654321]\d|9[8543210]|8[6421]|6[6543210]|5[87654321]|
4[987654310]|3[9643210]|2[70]|7|1)\d{1,14}$`)
	return global.MatchString(phone)
}

// IsEmailValid - verify email
func IsEmailValid(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}

// IsNil - verify nil
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// IsNotNil verify is not nil
func IsNotNil(i any) bool {
	return !IsNil(i)
}

// IsEmpty verify is empty
func IsEmpty(i any) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		vi = vi.Elem()
	}
	switch vi.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return vi.Len() == 0
	case reflect.Bool:
		return !vi.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vi.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vi.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return vi.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return vi.IsNil()
	case reflect.Struct:
		return vi.NumField() == 0
	}
	return false
}

// IsNotEmpty verify is not empty
func IsNotEmpty(i any) bool {
	return !IsEmpty(i)
}

// BoolPointer convert a bool pointer
func BoolPointer(b bool) *bool {
	return &b
}

// IsTrue verify is true
func IsTrue(i any) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		vi = vi.Elem()
	}
	switch vi.Kind() {
	case reflect.Bool:
		return vi.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return vi.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return vi.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return vi.Float() != 0
	}
	return false
}

// IsFalse verify is false
func IsFalse(i any) bool {
	return !IsTrue(i)
}

// IsEqual verify is equal
func IsEqual(i, j any) bool {
	return reflect.DeepEqual(i, j)
}

// IsNotEqual verify is not equal
func IsNotEqual(i, j any) bool {
	return !IsEqual(i, j)
}

// IsGreater verify is greater
func IsGreater(i, j any) bool {
	ii := reflect.ValueOf(i)
	jj := reflect.ValueOf(j)
	if ii.Kind() == reflect.Ptr {
		ii = ii.Elem()
	}
	if jj.Kind() == reflect.Ptr {
		jj = jj.Elem()
	}
	switch ii.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return ii.Int() > jj.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return ii.Uint() > jj.Uint()
	case reflect.Float32, reflect.Float64:
		return ii.Float() > jj.Float()
	}
	return false
}

// IsGreaterOrEqual verify is greater or equal
func IsGreaterOrEqual(i, j any) bool {
	return IsGreater(i, j) || IsEqual(i, j)
}

// IsLess verify is less
func IsLess(i, j any) bool {
	return !IsGreaterOrEqual(i, j)
}

// IsLessOrEqual verify is less or equal
func IsLessOrEqual(i, j any) bool {
	return !IsGreater(i, j)
}

// IsIn verify is in
func IsIn(i any, j []any) bool {
	for _, v := range j {
		if IsEqual(i, v) {
			return true
		}
	}
	return false
}

// IsNotIn verify is not in
func IsNotIn(i any, j []any) bool {
	return !IsIn(i, j)
}

// IsContains verify is contains
func IsContains(i any, j []any) bool {
	for _, v := range j {
		if IsEqual(i, v) {
			return true
		}
	}
	return false
}

// IsNotContains verify is not contains
func IsNotContains(i any, j []any) bool {
	return !IsContains(i, j)
}

// IsInString verify is in string
func IsInString(i string, j []string) bool {
	for _, v := range j {
		if i == v {
			return true
		}
	}
	return false
}

// IsNotInString verify is not in string
func IsNotInString(i string, j []string) bool {
	return !IsInString(i, j)
}

// IsInArray verify is in array
func IsInArray(i any, j []any) bool {
	for _, v := range j {
		if IsEqual(i, v) {
			return true
		}
	}
	return false
}
