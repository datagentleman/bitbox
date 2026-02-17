package bitbox_test

import (
	"testing"
	"unsafe"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
)

type alignedStruct struct {
	A uint64
	B uint64
	C uint32
	D float32
}

type namedInt8 int8
type namedUint32 uint32
type namedString string
type namedBytes []byte
type namedUint16Slice []uint16
type namedUint16Array [4]uint16

func runTest[T any](t *testing.T, name string, in T) {
	t.Helper()

	t.Run(name+"/encode_value_decode_pointer", func(t *testing.T) {
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)

		var out T
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, in, out)
	})

	t.Run(name+"/encode_pointer_decode_pointer", func(t *testing.T) {
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)

		var out T
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, in, out)
	})
}

func runDecodeNonPointerNoop[T any](t *testing.T, name string, in T) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)

		remainingBefore := buf.Len()
		bitbox.Decode(buf, in)
		remainingAfter := buf.Len()

		test.AssertEqual(t, remainingBefore, remainingAfter)
	})
}

func TestEncodeDecodeBasicTypes(t *testing.T) {
	cases := []struct {
		name string
		run  func(*testing.T)
	}{
		{name: "bool", run: func(t *testing.T) { runTest(t, "bool", true) }},
		{name: "int8", run: func(t *testing.T) { runTest(t, "int8", int8(-8)) }},
		{name: "int16", run: func(t *testing.T) { runTest(t, "int16", int16(-16)) }},
		{name: "int32", run: func(t *testing.T) { runTest(t, "int32", int32(-32)) }},
		{name: "int64", run: func(t *testing.T) { runTest(t, "int64", int64(-64)) }},
		{name: "uint8", run: func(t *testing.T) { runTest(t, "uint8", uint8(8)) }},
		{name: "uint16", run: func(t *testing.T) { runTest(t, "uint16", uint16(16)) }},
		{name: "uint32", run: func(t *testing.T) { runTest(t, "uint32", uint32(32)) }},
		{name: "uint64", run: func(t *testing.T) { runTest(t, "uint64", uint64(64)) }},
		{name: "uintptr", run: func(t *testing.T) { runTest(t, "uintptr", uintptr(128)) }},
		{name: "float32", run: func(t *testing.T) { runTest(t, "float32", float32(3.14)) }},
		{name: "float64", run: func(t *testing.T) { runTest(t, "float64", float64(6.28)) }},
		{name: "complex64", run: func(t *testing.T) { runTest(t, "complex64", complex64(complex(1.5, -2.5))) }},
		{name: "complex128", run: func(t *testing.T) { runTest(t, "complex128", complex128(complex(2.5, -3.5))) }},
		{name: "string", run: func(t *testing.T) { runTest(t, "string", "bitbox") }},
		{name: "bytes", run: func(t *testing.T) { runTest(t, "bytes", []byte{1, 2, 3, 4, 5}) }},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.run)
	}
}

func TestEncodeDecodeAlignedStruct(t *testing.T) {
	in := alignedStruct{A: 123, B: 456, C: 999, D: 7.25}

	buf := bitbox.NewBuffer(nil)
	bitbox.Encode(buf, &in)

	var out alignedStruct
	bitbox.Decode(buf, &out)
	test.AssertEqual(t, in, out)
}

func TestDecodeNonPointerIsNoop(t *testing.T) {
	cases := []struct {
		name string
		run  func(*testing.T)
	}{
		{name: "int32", run: func(t *testing.T) { runDecodeNonPointerNoop(t, "int32", int32(123)) }},
		{name: "string", run: func(t *testing.T) { runDecodeNonPointerNoop(t, "string", "decode-me") }},
		{name: "bytes", run: func(t *testing.T) { runDecodeNonPointerNoop(t, "bytes", []byte{7, 8, 9}) }},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.run)
	}
}

func TestUnsupportedIntUintAreNoop(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		in := int(-7)
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)
		test.AssertEqual(t, 0, buf.Len())
	})

	t.Run("uint", func(t *testing.T) {
		in := uint(7)
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)
		test.AssertEqual(t, 0, buf.Len())
	})
}

func TestIntUintPointerRoundTrip(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		in := int(-7)
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)

		var out int
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, in, out)
	})

	t.Run("uint", func(t *testing.T) {
		in := uint(7)
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)

		var out uint
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, in, out)
	})
}

func TestNamedBasicTypesPointerRoundTrip(t *testing.T) {
	t.Run("named int8 pointer decodes into int8", func(t *testing.T) {
		in := namedInt8(-42)
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)

		var out int8
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, int8(in), out)
	})

	t.Run("named uint32 pointer decodes into uint32", func(t *testing.T) {
		in := namedUint32(0xDEADBEEF)
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)

		var out uint32
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, uint32(in), out)
	})
}

func TestNamedStringAndBytesRoundTrip(t *testing.T) {
	t.Run("named string value decodes into string", func(t *testing.T) {
		in := namedString("named-string")
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)

		var out string
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, string(in), out)
	})

	t.Run("named []byte value decodes into []byte", func(t *testing.T) {
		in := namedBytes{10, 20, 30, 40}
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)

		var out []byte
		bitbox.Decode(buf, &out)
		test.AssertEqual(t, []byte(in), out)
	})
}

func TestNamedSlicesAndArraysEncodeLayout(t *testing.T) {
	t.Run("named []uint16 encodes length prefix + payload", func(t *testing.T) {
		in := namedUint16Slice{1, 2, 3, 4}
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)

		total := uint32(len(in) * int(unsafe.Sizeof(uint16(0))))
		test.AssertEqual(t, int(unsafe.Sizeof(total))+len(in)*int(unsafe.Sizeof(uint16(0))), buf.Len())

		var encodedLen uint32
		bitbox.Decode(buf, &encodedLen)
		test.AssertEqual(t, total, encodedLen)
		test.AssertEqual(t, int(total), len(buf.Data()))
	})

	t.Run("named [4]uint16 pointer encodes length prefix + payload", func(t *testing.T) {
		in := namedUint16Array{1, 2, 3, 4}
		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)

		total := uint32(len(in) * int(unsafe.Sizeof(uint16(0))))
		test.AssertEqual(t, int(unsafe.Sizeof(total))+len(in)*int(unsafe.Sizeof(uint16(0))), buf.Len())

		var encodedLen uint32
		bitbox.Decode(buf, &encodedLen)
		test.AssertEqual(t, total, encodedLen)
		test.AssertEqual(t, int(total), len(buf.Data()))
	})
}
