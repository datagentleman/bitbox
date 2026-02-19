package bitbox_test

import (
	"testing"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
)

type alignedStruct struct {
	A uint64
	B uint64
	C uint32
	D float32
}

type NamedTypeInt1 int
type NamedTypeInt2 int8
type NamedTypeInt3 int16
type NamedTypeInt4 int32
type NamedTypeInt5 int64

type NamedTypeUint1 uint
type NamedTypeUint2 uint8
type NamedTypeUint3 uint16
type NamedTypeUint4 uint32
type NamedTypeUint5 uint64
type NamedTypeUint6 uintptr

type NamedTypeFloat1 float32
type NamedTypeFloat2 float64

type NamedTypeComplex1 complex64
type NamedTypeComplex2 complex128

type NamedTypeString1 string
type NamedTypeBool1 bool
type NamedTypeByte1 []byte

type NamedTypeSlice1 []uint16
type NamedTypeArray1 [4]uint16
type NamedTypeArray2 [32]uint8

type Tx struct {
	ChainID    *uint64
	Nonce      uint64
	GasPrice   *uint64
	Gas        uint64
	To         *NamedTypeArray2
	Value      *uint64
	Data       []byte
	AccessList NamedTypeSlice1
}

func runTest[T any](t *testing.T, name string, in T) {
	t.Helper()

	t.Run(name+"/values", func(t *testing.T) {
		var out T

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})

	t.Run(name+"/pointers", func(t *testing.T) {
		var out T

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})
}

func runEncoderDecoder[T any](t *testing.T, name string, in T) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		var out T

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})
}

func TestFixedTypes(t *testing.T) {
	cases := []struct {
		name string
		run  func(*testing.T)
	}{
		// int
		{name: "int8", run: func(t *testing.T) { runTest(t, "int8", int8(-8)) }},
		{name: "int16", run: func(t *testing.T) { runTest(t, "int16", int16(-16)) }},
		{name: "int32", run: func(t *testing.T) { runTest(t, "int32", int32(-32)) }},
		{name: "int64", run: func(t *testing.T) { runTest(t, "int64", int64(-64)) }},

		// uint
		{name: "uint8", run: func(t *testing.T) { runTest(t, "uint8", uint8(8)) }},
		{name: "uint16", run: func(t *testing.T) { runTest(t, "uint16", uint16(16)) }},
		{name: "uint32", run: func(t *testing.T) { runTest(t, "uint32", uint32(32)) }},
		{name: "uint64", run: func(t *testing.T) { runTest(t, "uint64", uint64(64)) }},
		{name: "uintptr", run: func(t *testing.T) { runTest(t, "uintptr", uintptr(128)) }},

		// float
		{name: "float32", run: func(t *testing.T) { runTest(t, "float32", float32(3.14)) }},
		{name: "float64", run: func(t *testing.T) { runTest(t, "float64", float64(6.28)) }},

		// complex
		{name: "complex64", run: func(t *testing.T) { runTest(t, "complex64", complex64(complex(1.5, -2.5))) }},
		{name: "complex128", run: func(t *testing.T) { runTest(t, "complex128", complex128(complex(2.5, -3.5))) }},

		// bool, string, bytes
		{name: "string", run: func(t *testing.T) { runTest(t, "string", "bitbox") }},
		{name: "bytes", run: func(t *testing.T) { runTest(t, "bytes", []byte{1, 2, 3, 4, 5}) }},
		{name: "bool", run: func(t *testing.T) { runTest(t, "bool", true) }},
	}

	for _, tc := range cases {
		t.Run(tc.name, tc.run)
	}
}

func TestSlices(t *testing.T) {
	runTest(t, "slice_uint16", []uint16{1, 2, 3, 4})
	runTest(t, "slice_slice_byte", [][]byte{{1, 2, 3}, nil, {4, 5, 6, 7}})
	runTest(t, "slice_slice_slice_uint64", [][][]uint64{{{1, 2}, nil, {3, 4, 5}}, nil, {{6}, {7, 8, 9}}})
}

func TestArrays(t *testing.T) {
	runTest(t, "array", [4]uint16{1, 2, 3, 4})

	inByte := [32][32]byte{}
	for i := range inByte {
		for j := range inByte[i] {
			inByte[i][j] = byte((i + j) % 256)
		}
	}
	runTest(t, "2_nested_arrays", inByte)

	inU32 := [20][20][20]uint32{}
	for i := range inU32 {
		for j := range inU32[i] {
			for k := range inU32[i][j] {
				inU32[i][j][k] = uint32(i*10000 + j*100 + k)
			}
		}
	}
	runTest(t, "3_nested_arrays", inU32)
}

func TestAlignedStruct(t *testing.T) {
	in := alignedStruct{A: 123, B: 456, C: 999, D: 7.25}
	out := alignedStruct{}

	buf := bitbox.NewBuffer(nil)
	bitbox.Encode(buf, &in)
	bitbox.Decode(buf, &out)

	test.AssertEqual(t, in, out)
}

func TestStruct(t *testing.T) {
	t.Run("nil pointers", func(t *testing.T) {
		in := Tx{
			Nonce:      7,
			Gas:        21000,
			Data:       []byte{1, 2, 3},
			AccessList: NamedTypeSlice1{4, 5, 6},
		}
		out := Tx{}

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})

	t.Run("with pointers", func(t *testing.T) {
		chainID := uint64(11155111)
		gasPrice := uint64(20_000_000_000)
		value := uint64(12345)
		to := NamedTypeArray2{}

		in := Tx{
			ChainID:    &chainID,
			Nonce:      42,
			GasPrice:   &gasPrice,
			Gas:        21000,
			To:         &to,
			Value:      &value,
			Data:       []byte{9, 8, 7, 6},
			AccessList: NamedTypeSlice1{1, 3, 5, 7},
		}
		out := Tx{}

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})
}

func TestIntAndUint(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		in := int(-7)
		out := int(0)

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})

	t.Run("uint", func(t *testing.T) {
		in := uint(7)
		out := uint(0)

		buf := bitbox.NewBuffer(nil)
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)

		test.AssertEqual(t, in, out)
	})
}

func TestNamedTypes(t *testing.T) {
	// int
	runEncoderDecoder(t, "named_int", NamedTypeInt1(-7))
	runEncoderDecoder(t, "named_int8", NamedTypeInt2(-8))
	runEncoderDecoder(t, "named_int16", NamedTypeInt3(-16))
	runEncoderDecoder(t, "named_int32", NamedTypeInt4(-32))
	runEncoderDecoder(t, "named_int64", NamedTypeInt5(-64))

	// uint
	runEncoderDecoder(t, "named_uint", NamedTypeUint1(7))
	runEncoderDecoder(t, "named_uint8", NamedTypeUint2(8))
	runEncoderDecoder(t, "named_uint16", NamedTypeUint3(16))
	runEncoderDecoder(t, "named_uint32", NamedTypeUint4(32))
	runEncoderDecoder(t, "named_uint64", NamedTypeUint5(64))
	runEncoderDecoder(t, "named_uintptr", NamedTypeUint6(128))

	// float
	runEncoderDecoder(t, "named_float32", NamedTypeFloat1(3.14))
	runEncoderDecoder(t, "named_float64", NamedTypeFloat2(6.28))

	// complex
	runEncoderDecoder(t, "named_complex64", NamedTypeComplex1(complex(1.5, -2.5)))
	runEncoderDecoder(t, "named_complex128", NamedTypeComplex2(complex(2.5, -3.5)))

	// slice, array
	runEncoderDecoder(t, "named_slice_uint16", NamedTypeSlice1{1, 2, 3, 4})
	runEncoderDecoder(t, "named_array_uint16", NamedTypeArray1{1, 2, 3, 4})
	runEncoderDecoder(t, "named_array_uint16", NamedTypeArray2{1, 1, 1, 1})

	// string, bytes, bool
	runEncoderDecoder(t, "named_string", NamedTypeString1("named-string"))
	runEncoderDecoder(t, "named_bool", NamedTypeBool1(true))
	runEncoderDecoder(t, "named_bytes", NamedTypeByte1{10, 20, 30, 40})
}
