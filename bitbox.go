package bitbox

import (
	"reflect"
	"unsafe"
)

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
