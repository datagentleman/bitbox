package benchmark

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"strconv"
	"testing"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
	"github.com/vmihailenco/msgpack/v5"
)

func binaryWrite[T any](b *testing.B, wire *bytes.Buffer, in T) {
	switch v := any(in).(type) {
	case []byte:
		if _, err := wire.Write(v); err != nil {
			b.Fatalf("%v", err)
		}

	case string:
		if _, err := wire.WriteString(v); err != nil {
			b.Fatalf("%v", err)
		}

	case int:
		x := int64(v)
		if err := binary.Write(wire, binary.BigEndian, x); err != nil {
			b.Fatalf("%v", err)
		}

	case uint:
		x := uint64(v)
		if err := binary.Write(wire, binary.BigEndian, x); err != nil {
			b.Fatalf("%v", err)
		}

	case uintptr:
		x := uint64(v)
		if err := binary.Write(wire, binary.BigEndian, x); err != nil {
			b.Fatalf("%v", err)
		}

	default:
		if err := binary.Write(wire, binary.BigEndian, in); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func binaryRead[T any](b *testing.B, r *bytes.Reader, out *T) {
	switch any(*out).(type) {
	case []byte:
		buf, err := io.ReadAll(r)
		if err != nil {
			b.Fatalf("%v", err)
		}
		*out = any(buf).(T)

	case string:
		buf, err := io.ReadAll(r)
		if err != nil {
			b.Fatalf("%v", err)
		}
		*out = any(string(buf)).(T)

	case int:
		var x int64
		if err := binary.Read(r, binary.BigEndian, &x); err != nil {
			b.Fatalf("%v", err)
		}
		*out = any(int(x)).(T)

	case uint:
		var x uint64
		if err := binary.Read(r, binary.BigEndian, &x); err != nil {
			b.Fatalf("%v", err)
		}
		*out = any(uint(x)).(T)

	case uintptr:
		var x uint64
		if err := binary.Read(r, binary.BigEndian, &x); err != nil {
			b.Fatalf("%v", err)
		}
		*out = any(uintptr(x)).(T)

	default:
		if err := binary.Read(r, binary.BigEndian, out); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func benchmarkTypesBitbox[T any](b *testing.B, in T, setBytes int64) {
	var out T
	var zero T
	b.SetBytes(setBytes)
	b.ReportAllocs()

	buf := bitbox.NewBuffer([]byte{})
	bitbox.Encode(buf, in)
	bitbox.Decode(buf, &out)
	test.AssertEqual(b, in, out)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Clear()
		out = zero
		bitbox.Encode(buf, in)
		bitbox.Decode(buf, &out)
	}
}

func benchmarkTypesGob[T any](b *testing.B, in T, setBytes int64) {
	var out T
	var zero T
	b.SetBytes(setBytes)
	b.ReportAllocs()

	var wire bytes.Buffer
	enc := gob.NewEncoder(&wire)
	if err := enc.Encode(in); err != nil {
		b.Fatalf("%v", err)
	}

	r := bytes.NewReader(wire.Bytes())
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&out); err != nil {
		b.Fatalf("%v", err)
	}

	test.AssertEqual(b, in, out)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wire.Reset()
		out = zero
		enc := gob.NewEncoder(&wire)
		if err := enc.Encode(in); err != nil {
			b.Fatalf("%v", err)
		}

		r.Reset(wire.Bytes())
		dec := gob.NewDecoder(r)
		if err := dec.Decode(&out); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func benchmarkBinary[T any](b *testing.B, in T, setBytes int64) {
	var out T
	var zero T
	b.SetBytes(setBytes)
	b.ReportAllocs()

	var wire bytes.Buffer
	binaryWrite(b, &wire, in)

	r := bytes.NewReader(wire.Bytes())
	binaryRead(b, r, &out)

	test.AssertEqual(b, in, out)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wire.Reset()
		out = zero
		binaryWrite(b, &wire, in)

		r.Reset(wire.Bytes())
		binaryRead(b, r, &out)
	}
}

func benchmarkTypesMsgPack[T any](b *testing.B, in T, setBytes int64) {
	defer func() {
		if r := recover(); r != nil {
			b.Skipf("msgpack unsupported type %T: %v", in, r)
		}
	}()

	var out T
	var zero T
	b.SetBytes(setBytes)
	b.ReportAllocs()

	wire := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(wire)
	dec := msgpack.NewDecoder(wire)

	if err := enc.Encode(in); err != nil {
		b.Skipf("msgpack unsupported type %T: %v", in, err)
		return
	}
	if err := dec.Decode(&out); err != nil {
		b.Skipf("msgpack decode unsupported type %T: %v", in, err)
		return
	}
	test.AssertEqual(b, in, out)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wire.Reset()
		out = zero
		if err := enc.Encode(in); err != nil {
			b.Skipf("msgpack unsupported type %T: %v", in, err)
			return
		}
		if err := dec.Decode(&out); err != nil {
			b.Skipf("msgpack decode unsupported type %T: %v", in, err)
			return
		}
	}
}

func runBasicTypes[T any](b *testing.B, in T, setBytes int64) {
	b.Run("Bitbox", func(b *testing.B) {
		benchmarkTypesBitbox(b, in, setBytes)
	})

	b.Run("Gob", func(b *testing.B) {
		benchmarkTypesGob(b, in, setBytes)
	})

	b.Run("Binary", func(b *testing.B) {
		benchmarkBinary(b, in, setBytes)
	})

	b.Run("MsgPack", func(b *testing.B) {
		benchmarkTypesMsgPack(b, in, setBytes)
	})
}

func BenchmarkEncodeDecodeBasicTypes(b *testing.B) {
	b.Run("slice_128B", func(b *testing.B) {
		v := make([]byte, 128)
		for i := range v {
			v[i] = byte(i * 31)
		}
		runBasicTypes(b, v, int64(len(v))*2)
	})

	b.Run("slice_4KB", func(b *testing.B) {
		v := make([]byte, 4*1024)
		for i := range v {
			v[i] = byte(i * 31)
		}
		runBasicTypes(b, v, int64(len(v))*2)
	})

	b.Run("slice_64KB", func(b *testing.B) {
		v := make([]byte, 64*1024)
		for i := range v {
			v[i] = byte(i * 31)
		}
		runBasicTypes(b, v, int64(len(v))*2)
	})

	b.Run("bool", func(b *testing.B) {
		v := true
		runBasicTypes(b, v, 2)
	})

	b.Run("string", func(b *testing.B) {
		v := "bitbox-benchmark"
		runBasicTypes(b, v, int64(len(v))*2)
	})

	b.Run("int8", func(b *testing.B) {
		v := int8(-66)
		runBasicTypes(b, v, 2)
	})

	b.Run("int16", func(b *testing.B) {
		v := int16(-666)
		runBasicTypes(b, v, 4)
	})

	b.Run("int32", func(b *testing.B) {
		v := int32(-666)
		runBasicTypes(b, v, 8)
	})

	b.Run("int64", func(b *testing.B) {
		v := int64(-666)
		runBasicTypes(b, v, 16)
	})

	b.Run("uint8", func(b *testing.B) {
		v := uint8(66)
		runBasicTypes(b, v, 2)
	})

	b.Run("uint16", func(b *testing.B) {
		v := uint16(666)
		runBasicTypes(b, v, 4)
	})

	b.Run("uint32", func(b *testing.B) {
		v := uint32(666)
		runBasicTypes(b, v, 8)
	})

	b.Run("uint64", func(b *testing.B) {
		v := uint64(666)
		runBasicTypes(b, v, 16)
	})

	b.Run("uintptr", func(b *testing.B) {
		v := uintptr(666)
		runBasicTypes(b, v, int64((strconv.IntSize/8)*2))
	})

	b.Run("float32", func(b *testing.B) {
		v := float32(3.14159)
		runBasicTypes(b, v, 8)
	})

	b.Run("float64", func(b *testing.B) {
		v := float64(3.141592653589793)
		runBasicTypes(b, v, 16)
	})

	b.Run("complex64", func(b *testing.B) {
		v := complex64(complex(3.14, -2.71))
		runBasicTypes(b, v, 16)
	})

	b.Run("complex128", func(b *testing.B) {
		v := complex128(complex(3.1415926535, -2.7182818284))
		runBasicTypes(b, v, 32)
	})
}
