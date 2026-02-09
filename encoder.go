package bitbox

import (
	"reflect"
	"unsafe"
)

// Encode objects
func Encode(buf *Buffer, objects ...any) *Buffer {
	for _, obj := range objects {
		encoded := true

		switch val := obj.(type) {
		// Bytes
		case []byte:
			l := uint32(len(val))
			buf.data = append(buf.data, ToBytes(&l)...)
			buf.data = append(buf.data, val...)

		case *[]byte:
			l := uint32(len(*val))
			buf.data = append(buf.data, ToBytes(&l)...)
			buf.data = append(buf.data, *val...)

		// Basic Types
		case int:
			buf.Write(ToBytes(&val))
		case int8:
			buf.Write(ToBytes(&val))
		case int16:
			buf.Write(ToBytes(&val))
		case int32:
			buf.Write(ToBytes(&val))
		case int64:
			buf.Write(ToBytes(&val))
		case uint:
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

		case *int:
			buf.Write(ToBytes(val))
		case *int8:
			buf.Write(ToBytes(val))
		case *int16:
			buf.Write(ToBytes(val))
		case *int32:
			buf.Write(ToBytes(val))
		case *int64:
			buf.Write(ToBytes(val))
		case *uint:
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

		// String
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
			encoded = false
		}

		// Fast path didn't encode anything, try with reflect.
		if !encoded {
			val := reflect.ValueOf(obj)
			val = reflect.Indirect(val)

			if !val.IsValid() {
				continue
			}

			if EncodeAlignedStruct(buf, val) {
				continue
			}
		}
	}

	return buf
}

// Encode aligned struct.
func EncodeAlignedStruct(buf *Buffer, val reflect.Value) bool {
	if val.Kind() == reflect.Struct {
		buf.Write(toBytes(val))
		return true
	}

	return false
}
