package bitbox

import (
	"reflect"
	"unsafe"
)

func Encode(buf *Buffer, objects ...any) {
	for _, obj := range objects {
		if EncodeBasic(buf, obj) {
			continue
		}

		val := reflect.ValueOf(obj)
		val = reflect.Indirect(val)

		if !val.IsValid() {
			continue
		}

		switch val.Kind() {
		case reflect.Slice:
			EncodeSlice(buf, val)
		case reflect.Array:
			EncodeArray(buf, val)
		case reflect.Struct:
			EncodePODStruct(buf, val)
		default:
			EncodeNamedType(buf, val)
		}
	}
}

// Encode POD stucts.
func EncodePODStruct(buf *Buffer, val reflect.Value) {
	size := int(val.Type().Size())
	buf.Write(toBytes(val, size))
}

// Encode named types.
func EncodeNamedType(buf *Buffer, val reflect.Value) {
	kind := val.Kind()

	if !val.IsValid() || val.Type().Name() == "" {
		return
	}

	switch kind {
	case reflect.Slice:
		EncodeSlice(buf, val)
	case reflect.Array:
		EncodeArray(buf, val)
	case reflect.String:
		EncodeBasic(buf, val.String())
	default:
		if isPOD(kind) {
			size := int(val.Type().Size())
			buf.Write(toBytes(val, size))
		}
	}
}

// Encode POD slice. If slice elements are structs,
// they will be encoded as POD struct (they must be memory aligned).
func EncodeSlice(buf *Buffer, val reflect.Value) {
	elem := val.Type().Elem()

	if !isPOD(elem.Kind()) && !isStruct(elem.Kind()) {
		return
	}

	size := int(elem.Size())
	total := uint32(val.Len() * size)

	buf.Write(ToBytes(&total))
	buf.Write(toBytes(val, int(total)))
}

// Encode slice. If slice elements are structs,
// they will be encoded as POD struct (they must be memory aligned).
func EncodeArray(buf *Buffer, val reflect.Value) {
	elem := val.Type().Elem()

	if !isPOD(elem.Kind()) && !isStruct(elem.Kind()) {
		return
	}

	size := int(elem.Size())
	total := uint32(val.Len() * size)

	buf.Write(ToBytes(&total))
	buf.Write(toBytes(val, int(total)))
}

// Encode basic types.
func EncodeBasic(buf *Buffer, obj any) bool {
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
