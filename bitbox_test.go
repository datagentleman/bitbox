package bitbox

import (
	"testing"

	"github.com/datagentleman/bitbox/test"
)

type alignedStruct4Fields struct {
	A uint64
	B uint32
	C uint16
	D uint16
}

func TestEncodeDecodeUint64(t *testing.T) {
	u1 := uint64(666)
	u2 := uint64(0)

	buf := Encode(&u1)
	Decode(buf, &u2)

	test.Assert(t, u1, u2)
}

func TestEncodeDecodeBool(t *testing.T) {
	v1 := true
	v2 := false

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeString(t *testing.T) {
	v1 := "bitbox"
	v2 := ""

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeInt(t *testing.T) {
	v1 := int(-666)
	v2 := int(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeInt8(t *testing.T) {
	v1 := int8(-66)
	v2 := int8(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeInt16(t *testing.T) {
	v1 := int16(-666)
	v2 := int16(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeInt32(t *testing.T) {
	v1 := int32(-666)
	v2 := int32(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeInt64(t *testing.T) {
	v1 := int64(-666)
	v2 := int64(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeUint(t *testing.T) {
	v1 := uint(666)
	v2 := uint(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeUint8(t *testing.T) {
	v1 := uint8(66)
	v2 := uint8(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeUint16(t *testing.T) {
	v1 := uint16(666)
	v2 := uint16(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeUint32(t *testing.T) {
	v1 := uint32(666)
	v2 := uint32(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeByte(t *testing.T) {
	v1 := byte(66)
	v2 := byte(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeRune(t *testing.T) {
	v1 := rune('世')
	v2 := rune(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeUintptr(t *testing.T) {
	v1 := uintptr(666)
	v2 := uintptr(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeFloat32(t *testing.T) {
	v1 := float32(3.14159)
	v2 := float32(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeFloat64(t *testing.T) {
	v1 := float64(3.141592653589793)
	v2 := float64(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeComplex64(t *testing.T) {
	v1 := complex64(complex(3.14, -2.71))
	v2 := complex64(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeComplex128(t *testing.T) {
	v1 := complex128(complex(3.1415926535, -2.7182818284))
	v2 := complex128(0)

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.Assert(t, v1, v2)
}

func TestEncodeDecodeAlignedStruct4Fields(t *testing.T) {
	v1 := alignedStruct4Fields{
		A: 0x1122334455667788,
		B: 0x99AABBCC,
		C: 0xDDEE,
		D: 0xFF00,
	}
	v2 := alignedStruct4Fields{}

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.AssertEqual(t, v1, v2)
}

func TestEncodeDecodeByteSlice(t *testing.T) {
	v1 := []byte{1, 2, 3, 4, 5, 6}
	v2 := []byte{}

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.AssertEqual(t, v1, v2)
}

func TestEncodeDecodeUint32Array(t *testing.T) {
	v1 := [4]uint32{1, 2, 3, 4}
	v2 := [4]uint32{}

	buf := Encode(&v1)
	Decode(buf, &v2)

	test.AssertEqual(t, v1, v2)
}
