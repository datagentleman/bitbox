package bitbox

import (
	"reflect"
)

// Decode objects
func Decode(buf *Buffer, objects ...any) {
	for _, obj := range objects {
		// Fast path - type cast
		if decodeFixed(buf, obj) {
			continue
		}

		// Slow path - reflections
		v := reflect.ValueOf(obj)
		v = reflect.Indirect(v)

		if !v.IsValid() {
			continue
		}

		decode(buf, v)
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
		decodeSlice(buf, val, false)
	case reflect.Array:
		decodeArray(buf, val)
	case reflect.Struct:
		decodeStruct(buf, val)
	case reflect.String:
		l := uint32(0)
		buf.Decode(&l)
		val.SetString(string(buf.Next(int(l))))
	}
}

// Decode structs.
func decodeStruct(buf *Buffer, val reflect.Value) {
	if !val.IsValid() {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.Pointer {
			ptrFlag := uint8(1)
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

// Decode arrays.
func decodeArray(buf *Buffer, val reflect.Value) {
	elem := val.Type().Elem()
	total := val.Len() * int(elem.Size())

	switch elem.Kind() {
	// nested arrays
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			decodeArray(buf, val.Index(i))
		}
	default:
		buf.Read(toBytes(val, int(total)))
	}
}

// Decode slices.
func decodeSlice(buf *Buffer, val reflect.Value, isPOD bool) {
	num := uint32(0)
	buf.Decode(&num)

	ensureLen(val, int(num))
	elem := val.Type().Elem()

	switch elem.Kind() {
	// nested slice
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			decodeSlice(buf, val.Index(i), isPOD)
		}
	case reflect.Struct:
		if isPOD {
			tsize := uint32(elem.Size())
			total := int(num * tsize)

			buf.Read(toBytes(val, int(total)))
			return
		}

		for i := 0; i < val.Len(); i++ {
			decodeStruct(buf, val.Index(i))
		}

	default:
		tsize := uint32(elem.Size())
		total := int(num * tsize)

		buf.Read(toBytes(val, int(total)))
	}
}

func decodeFixed(buf *Buffer, obj any) bool {
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
		return false
	}
	return true
}
