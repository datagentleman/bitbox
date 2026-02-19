package bitbox

import (
	"reflect"
	"unsafe"
)

func Encode(buf *Buffer, objects ...any) {
	for _, obj := range objects {
		// Fast path - type cast
		if encodeFixed(buf, obj) {
			continue
		}

		// Slow path - reflections
		val := reflect.ValueOf(obj)
		val = reflect.Indirect(val)

		if !val.IsValid() {
			continue
		}

		encode(buf, val, false)
	}
}

// Encode slices with fixed elements. If slice elements are structs,
// they will be encoded as POD struct (they must be memory aligned).
func encodeSlice(buf *Buffer, val reflect.Value, isPOD bool) {
	elem := val.Type().Elem()

	// write number of elements
	count := uint32(val.Len())
	buf.Write(ToBytes(&count))

	if isFixedType(elem.Kind()) {
		val = addressable(val)

		size := elem.Size()
		total := count * uint32(size)

		buf.Write(toBytes(val, int(total)))
		return
	}

	for i := 0; i < val.Len(); i++ {
		encode(buf, val.Index(i), isPOD)
	}
}

// Encode POD structs. Ensure that your objects
// are pure POD and they are memory aligned.
func EncodePOD(buf *Buffer, object any) {
	val := reflect.ValueOf(object)
	val = reflect.Indirect(val)

	if !val.IsValid() {
		return
	}

	val = addressable(val)

	switch val.Kind() {
	case reflect.Struct:
		size := int(val.Type().Size())
		buf.Write(toBytes(val, size))

	// Handle POD with slices, arrays, nested slices,
	// named types, ...
	default:
		isPOD := true
		encode(buf, val, isPOD)
	}
}

// Encode array.
func encodeArray(buf *Buffer, val reflect.Value, isPOD bool) {
	elem := val.Type().Elem()

	if isFixedType(elem.Kind()) {
		size := int(elem.Size())
		total := uint32(val.Len() * size)

		val = addressable(val)
		buf.Write(toBytes(val, int(total)))

		return
	}

	for i := 0; i < val.Len(); i++ {
		encode(buf, val.Index(i), isPOD)
	}
}

func encodeStruct(buf *Buffer, val reflect.Value, isPOD bool) {
	val = reflect.Indirect(val)
	if !val.IsValid() {
		return
	}

	val = addressable(val)

	if isPOD {
		size := int(val.Type().Size())
		buf.Write(toBytes(val, size))
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		kind := field.Kind()

		if kind == reflect.Pointer {
			ptrFlag := uint8(0)

			if field.IsNil() {
				buf.Write(ToBytes(&ptrFlag))
				continue
			}

			ptrFlag = 1
			buf.Write(ToBytes(&ptrFlag))

			field = reflect.Indirect(field)
		}

		encode(buf, field, isPOD)
	}
}

// This also handle named types.
func encode(buf *Buffer, val reflect.Value, isPOD bool) {
	kind := val.Kind()

	switch kind {
	case
		reflect.Bool, reflect.Uintptr, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:

		size := int(val.Type().Size())
		buf.Write(toBytes(val, size))

	case reflect.Slice:
		encodeSlice(buf, val, isPOD)
	case reflect.Array:
		encodeArray(buf, val, isPOD)
	case reflect.String:
		encodeFixed(buf, val.String())
	case reflect.Struct:
		encodeStruct(buf, val, isPOD)
	}
}

// Encode basic types.
func encodeFixed(buf *Buffer, obj any) bool {
	switch val := obj.(type) {
	// Values
	case int8:
		buf.Write(ToBytes(&val))
	case int16:
		buf.Write(ToBytes(&val))
	case int32:
		buf.Write(ToBytes(&val))
	case int64:
		buf.Write(ToBytes(&val))
	case uint8:
		buf.Write(ToBytes(&val))
	case uint16:
		buf.Write(ToBytes(&val))
	case uint32:
		buf.Write(ToBytes(&val))
	case uint64:
		buf.Write(ToBytes(&val))
	case float32:
		buf.Write(ToBytes(&val))
	case float64:
		buf.Write(ToBytes(&val))
	case complex64:
		buf.Write(ToBytes(&val))
	case complex128:
		buf.Write(ToBytes(&val))
	case uintptr:
		buf.Write(ToBytes(&val))
	case bool:
		buf.Write(ToBytes(&val))

	// Pointers
	case *int8:
		buf.Write(ToBytes(val))
	case *int16:
		buf.Write(ToBytes(val))
	case *int32:
		buf.Write(ToBytes(val))
	case *int64:
		buf.Write(ToBytes(val))
	case *uint8:
		buf.Write(ToBytes(val))
	case *uint16:
		buf.Write(ToBytes(val))
	case *uint32:
		buf.Write(ToBytes(val))
	case *uint64:
		buf.Write(ToBytes(val))
	case *float32:
		buf.Write(ToBytes(val))
	case *float64:
		buf.Write(ToBytes(val))
	case *complex64:
		buf.Write(ToBytes(val))
	case *complex128:
		buf.Write(ToBytes(val))
	case *uintptr:
		buf.Write(ToBytes(val))
	case *bool:
		buf.Write(ToBytes(val))

	// Bytes
	case []byte:
		l := uint32(len(val))
		buf.Write(ToBytes(&l))
		buf.Write(val)

	case *[]byte:
		l := uint32(len(*val))
		buf.Write(ToBytes(&l))
		buf.Write(*val)

	// Strings
	case string:
		l := uint32(len(val))
		b := unsafe.Slice(unsafe.StringData(val), len(val))

		buf.Write(ToBytes(&l))
		buf.Write(b)

	case *string:
		l := uint32(len(*val))
		b := unsafe.Slice(unsafe.StringData(*val), len(*val))

		buf.Write(ToBytes(&l))
		buf.Write(b)

	default:
		return false
	}
	return true
}
