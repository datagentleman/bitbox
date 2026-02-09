package bitbox

import "reflect"

// Decode objects
func Decode(buf *Buffer, objects ...any) {
	for _, obj := range objects {
		decoded := true

		switch val := obj.(type) {
		// Bytes
		case *[]byte:
			l := uint32(0)
			buf.Decode(&l)
			*val = append(*val, buf.Next(int(l))...)

		// Basic Types
		case *int:
			buf.Read(ToBytes(val))
		case *int8:
			buf.Read(ToBytes(val))
		case *int16:
			buf.Read(ToBytes(val))
		case *int32:
			buf.Read(ToBytes(val))
		case *int64:
			buf.Read(ToBytes(val))
		case *uint:
			buf.Read(ToBytes(val))
		case *uint8:
			buf.Read(ToBytes(val))
		case *uint16:
			buf.Read(ToBytes(val))
		case *uint32:
			buf.Read(ToBytes(val))
		case *uint64:
			buf.Read(ToBytes(val))
		case *float32:
			buf.Read(ToBytes(val))
		case *float64:
			buf.Read(ToBytes(val))
		case *complex64:
			buf.Read(ToBytes(val))
		case *complex128:
			buf.Read(ToBytes(val))
		case *uintptr:
			buf.Read(ToBytes(val))
		case *bool:
			buf.Read(ToBytes(val))

		// String
		case *string:
			l := uint32(0)
			buf.Decode(&l)

			*val = string(buf.Next(int(l)))
		default:
			decoded = false
		}

		// Decode didn't decode anything, lets try with reflections.
		if !decoded {
			val := reflect.ValueOf(obj)
			val = reflect.Indirect(val)

			if !val.IsValid() {
				continue
			}

			if DecodeAlignedStruct(buf, val) {
				continue
			}
		}
	}
}

func DecodeAlignedStruct(buf *Buffer, val reflect.Value) bool {
	if val.Kind() == reflect.Struct {
		buf.Read(toBytes(val))
		return true
	}

	return false
}
