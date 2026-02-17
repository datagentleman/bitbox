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

// Get pointer from val and cast it to []byte.
func toBytes(val reflect.Value, size int) []byte {
	// Slices are treated differently, we cannot take UnsafeAddr from them.
	if val.Kind() == reflect.Slice {
		dataPtr := unsafe.Pointer(val.Pointer())
		return unsafe.Slice((*byte)(dataPtr), size)
	}

	if !val.CanAddr() {
		return nil
	}

	ptr := unsafe.Pointer(val.UnsafeAddr())
	return unsafe.Slice((*byte)(ptr), size)
}

// Detect if value is fixed type.
func isFixedType(kind reflect.Kind) bool {
	switch kind {
	case
		reflect.Bool, reflect.Uintptr, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:

		return true
	default:
		return false
	}
}

func isStruct(kind reflect.Kind) bool {
	return kind == reflect.Struct
}

// Make slice of given type and size.
func MakeSlice(typ reflect.Type, size int) reflect.Value {
	return reflect.MakeSlice(typ, 0, size)
}
