package bitbox

import (
	"reflect"
)

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

		// Basic Pointers
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

			if val.Kind() != reflect.Pointer {
				continue
			}

			val = reflect.Indirect(val)

			if !val.IsValid() {
				continue
			}

			switch val.Kind() {
			case reflect.Slice:
				decodeSlice(buf, val)
			case reflect.Array:
				decodeArray(buf, val)
			case reflect.Struct:
				decodePOD(buf, val)
			case reflect.String:
				decodeString(buf, val)
			default:
				if isFixedType(val.Kind()) {
					decodeFixed(buf, val)
				}
			}
		}
	}
}

// Decode structs.
func DecodeStruct(buf *Buffer, objects ...any) {
	for _, obj := range objects {
		val := reflect.ValueOf(obj)

		if val.Kind() != reflect.Pointer || val.IsNil() {
			continue
		}

		val = val.Elem()
		if !val.IsValid() || val.Kind() != reflect.Struct {
			continue
		}

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)

			if field.Kind() == reflect.Pointer {
				ptrFlag := uint8(0)
				buf.Read(ToBytes(&ptrFlag))

				if ptrFlag == 0 {
					field.Set(reflect.Zero(field.Type()))
					continue
				}

				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}

				field = field.Elem()
			}

			switch field.Kind() {
			case reflect.Slice:
				decodeSlice(buf, field)
			case reflect.Array:
				decodeArray(buf, field)
			case reflect.String:
				decodeString(buf, field)
			default:
				if isFixedType(field.Kind()) {
					size := int(field.Type().Size())
					buf.Read(toBytes(field, size))
				}
			}
		}
	}
}

// Decode arrays - reflect style.
func decodeArray(buf *Buffer, val reflect.Value) {
	elem := val.Type().Elem()
	total := val.Len() * int(elem.Size())

	buf.Read(toBytes(val, int(total)))
}

// Decode slices - reflect style.
func decodeSlice(buf *Buffer, val reflect.Value) {
	elem := val.Type().Elem()
	tsize := uint32(elem.Size())

	total := uint32(0)
	buf.Decode(&total)

	n := int(total / tsize)
	if val.Cap() < n {
		val.Set(MakeSlice(val.Type(), n))
	}

	val.SetLen(n)
	buf.Read(toBytes(val, int(total)))
}

func decodePOD(buf *Buffer, val reflect.Value) {
	size := int(val.Type().Size())
	buf.Read(toBytes(val, size))
}

func decodeString(buf *Buffer, val reflect.Value) {
	l := uint32(0)
	buf.Decode(&l)
	val.SetString(string(buf.Next(int(l))))
}

func decodeFixed(buf *Buffer, val reflect.Value) {
	size := int(val.Type().Size())
	buf.Read(toBytes(val, size))
}
