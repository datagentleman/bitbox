package bitbox

import (
	"reflect"
	"unsafe"
)

// Encode objects
func Encode(objects ...any) *Buffer {
	buf := NewBuffer([]byte{})

	for _, obj := range objects {
		val := reflect.ValueOf(obj)
		val = reflect.Indirect(val) // indirect pointers

		if !val.IsValid() {
			continue // skip nil pointers
		}

		// Encode []byte
		if IsBytes(val) {
			l := uint32(len(val.Bytes()))

			buf.data = append(buf.data, ToBytes(&l)...) // length prefix
			buf.data = append(buf.data, val.Bytes()...) // bytes
			continue
		}

		// Encode string
		if val.Kind() == reflect.String {
			s := val.String()
			l := uint32(len(s))

			buf.data = append(buf.data, ToBytes(&l)...) // length prefix
			buf.data = append(buf.data, s...)           // string bytes
			continue
		}

		// Encode basic types
		buf.data = append(buf.data, toBytes(val)...)
		continue
	}

	return buf
}

// Decode objects
func Decode(buf *Buffer, objects ...any) {
	for _, obj := range objects {
		val := reflect.ValueOf(obj)
		if !val.IsValid() || val.Kind() != reflect.Pointer || val.IsNil() {
			continue
		}

		val = val.Elem() // use value instead of pointer

		switch {
		case IsBytes(val):
			l := uint32(0)
			buf.Decode(&l)
			val.SetBytes(buf.Take(int(l)))
		case val.Kind() == reflect.String:
			l := uint32(0)
			buf.Decode(&l)
			val.SetString(string(buf.Take(int(l))))
		default:
			// Decode fixed-size basic types (including aliases) and structs.
			// Structs must be memory aligned.
			buf.Copy(toBytes(val))
		}
	}
}

// Get pointer to fixed type (including structs) and cast it to []byte.
// When passing structs, make sure they are memory aligned otherwise
// you would have extra bytes.
func ToBytes[T any](obj *T) []byte {
	size := unsafe.Sizeof(*obj)
	return unsafe.Slice((*byte)(unsafe.Pointer(obj)), size)
}

// Get pointer from reflect.Value and cast it to []byte.
// Value must be addressable.
func toBytes(val reflect.Value) []byte {
	if !val.CanAddr() {
		panic("value is not addressable")
	}

	ptr := unsafe.Pointer(val.UnsafeAddr())
	size := val.Type().Size()

	return unsafe.Slice((*byte)(ptr), size)
}

// Check if we deal with byte slice/array
func IsBytes(val reflect.Value) bool {
	if IsSlice(val) || IsArray(val) {
		return val.Type().Elem().Kind() == reflect.Uint8
	}

	return false
}

// Check if we have slice
func IsSlice(val reflect.Value) bool {
	return val.Kind() == reflect.Slice
}

// Check if we have array
func IsArray(val reflect.Value) bool {
	return val.Kind() == reflect.Array
}
