package bitbox

import (
	"reflect"
)

// Decode objects
func Decode(buf *Buffer, objects ...any) {
	for _, obj := range objects {
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
			v := reflect.ValueOf(obj)
			v = reflect.Indirect(v)

			if !v.IsValid() {
				continue
			}

			decode(buf, v)
		}
	}
}

func decode(buf *Buffer, val reflect.Value) {
	switch val.Kind() {
	case
		reflect.Bool, reflect.Uintptr, reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:

		size := int(val.Type().Size())
		buf.Read(toBytes(val, size))

	case reflect.Slice:
		decodeSlice(buf, val)
	case reflect.Array:
		decodeArray(buf, val)

	case reflect.Struct:
		size := int(val.Type().Size())
		buf.Read(toBytes(val, size))

	case reflect.String:
		l := uint32(0)
		buf.Decode(&l)
		val.SetString(string(buf.Next(int(l))))
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

			decode(buf, field)
		}
	}
}

// Decode arrays.
func decodeArray(buf *Buffer, val reflect.Value) {
	elem := val.Type().Elem()
	total := val.Len() * int(elem.Size())

	buf.Read(toBytes(val, int(total)))
}

// Decode slices.
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
