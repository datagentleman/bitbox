package bitbox

import (
	"reflect"
	"unsafe"
)

// Decode objects
func Decode(buf *Buffer, objects ...any) error {
	for _, obj := range objects {
		// Fast path - type cast
		handled, err := decodeFixed(buf, obj)
		if err != nil {
			return err
		}

		if handled {
			continue
		}

		// Slow path - reflections
		val := reflect.ValueOf(obj)

		if !isPointer(val.Kind()) || val.IsNil() || !val.IsValid() {
			return invalidValue(val)
		}

		val = reflect.Indirect(val)

		isPOD := false
		err = decode(buf, val, isPOD)
		if err != nil {
			return err
		}
	}
	return nil
}

func decode(buf *Buffer, val reflect.Value, isPOD bool) error {
	var err error

	switch val.Kind() {
	case
		reflect.Bool, reflect.Uintptr, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:

		size := int(val.Type().Size())
		buf.Read(toBytes(val, size))

	case reflect.Slice:
		err = decodeSlice(buf, val, isPOD)
	case reflect.Array:
		err = decodeArray(buf, val, isPOD)
	case reflect.Struct:
		err = decodeStruct(buf, val, isPOD)
	case reflect.String:
		l := uint32(0)
		buf.Decode(&l)
		b, err := buf.Next(int(l))
		if err != nil {
			return err
		}
		val.SetString(string(b))
	default:
		return invalidValue(val)
	}
	return err
}

// Decode structs.
func decodeStruct(buf *Buffer, val reflect.Value, isPOD bool) error {
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

		err := decode(buf, field, isPOD)
		if err != nil {
			return err
		}
	}
	return nil
}

// Decode arrays.
func decodeArray(buf *Buffer, val reflect.Value, isPOD bool) error {
	elem := val.Type().Elem()
	total := val.Len() * int(elem.Size())

	if isFixedType(elem.Kind()) {
		buf.Read(toBytes(val, int(total)))
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		err := decode(buf, val.Index(i), isPOD)
		if err != nil {
			return err
		}
	}
	return nil
}

// Decode slices.
func decodeSlice(buf *Buffer, val reflect.Value, isPOD bool) error {
	num := uint32(0)
	buf.Decode(&num)

	ensureLen(val, int(num))
	elem := val.Type().Elem()

	if isFixedType(elem.Kind()) {
		tsize := uint32(elem.Size())
		total := int(num * tsize)

		buf.Read(toBytes(val, total))
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		err := decode(buf, val.Index(i), isPOD)
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeFixedSlice[T any](buf *Buffer, out *[]T) error {
	l := uint32(0)
	buf.Decode(&l)
	n := int(l)

	if cap(*out) < n {
		*out = make([]T, n)
	} else {
		*out = (*out)[:n]
	}

	total := n * int(unsafe.Sizeof(*new(T)))
	s := unsafe.SliceData(*out)
	b := unsafe.Slice((*byte)(unsafe.Pointer(s)), total)

	buf.Read(b)
	return nil
}

func decodeFixedSlice2D[T any](buf *Buffer, out *[][]T) {
	l := uint32(0)
	buf.Decode(&l)
	n := int(l)

	if cap(*out) < n {
		*out = make([][]T, n)
	} else {
		*out = (*out)[:n]
	}

	for i := 0; i < n; i++ {
		decodeFixedSlice(buf, &(*out)[i])
	}
}

func decodeFixed(buf *Buffer, obj any) (bool, error) {
	switch val := obj.(type) {

	// Basic Pointers
	case *int8:
		buf.Read(ToBytes(val))
	case *int16:
		buf.Read(ToBytes(val))
	case *int32:
		buf.Read(ToBytes(val))
	case *int64:
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

	case *[]byte:
		decodeFixedSlice(buf, val)
	case *[]int8:
		decodeFixedSlice(buf, val)
	case *[]int16:
		decodeFixedSlice(buf, val)
	case *[]int32:
		decodeFixedSlice(buf, val)
	case *[]int64:
		decodeFixedSlice(buf, val)
	case *[]uint16:
		decodeFixedSlice(buf, val)
	case *[]uint32:
		decodeFixedSlice(buf, val)
	case *[]uint64:
		decodeFixedSlice(buf, val)
	case *[]float32:
		decodeFixedSlice(buf, val)
	case *[]float64:
		decodeFixedSlice(buf, val)
	case *[]complex64:
		decodeFixedSlice(buf, val)
	case *[]complex128:
		decodeFixedSlice(buf, val)
	case *[]uintptr:
		decodeFixedSlice(buf, val)
	case *[]bool:
		decodeFixedSlice(buf, val)

	case *[][]byte:
		decodeFixedSlice2D(buf, val)
	case *[][]int8:
		decodeFixedSlice2D(buf, val)
	case *[][]int16:
		decodeFixedSlice2D(buf, val)
	case *[][]int32:
		decodeFixedSlice2D(buf, val)
	case *[][]int64:
		decodeFixedSlice2D(buf, val)
	case *[][]uint16:
		decodeFixedSlice2D(buf, val)
	case *[][]uint32:
		decodeFixedSlice2D(buf, val)
	case *[][]uint64:
		decodeFixedSlice2D(buf, val)
	case *[][]float32:
		decodeFixedSlice2D(buf, val)
	case *[][]float64:
		decodeFixedSlice2D(buf, val)
	case *[][]complex64:
		decodeFixedSlice2D(buf, val)
	case *[][]complex128:
		decodeFixedSlice2D(buf, val)
	case *[][]uintptr:
		decodeFixedSlice2D(buf, val)
	case *[][]bool:
		decodeFixedSlice2D(buf, val)

	// String
	case *string:
		l := uint32(0)
		buf.Decode(&l)
		b, err := buf.Next(int(l))
		if err != nil {
			return false, err
		}
		*val = string(b)

	default:
		return false, nil
	}
	return true, nil
}
