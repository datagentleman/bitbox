package benchmark

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"testing"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
	"github.com/vmihailenco/msgpack/v5"
)

type NamedType01 bool
type NamedType02 int8
type NamedType03 int16
type NamedType04 int32
type NamedType05 int64
type NamedType06 uint8
type NamedType07 uint16
type NamedType08 uint32
type NamedType09 uint64
type NamedType10 float32
type NamedType11 float64
type NamedType12 complex64
type NamedType13 complex128
type NamedType14 string
type NamedType15 []byte
type NamedType16 []uint64
type NamedType17 [1000]byte
type NamedType18 [4]uint32

type namedBenchCase struct {
	name       string
	setBytes   int64
	runBitbox  func(*testing.B)
	runGob     func(*testing.B)
	runBinary  func(*testing.B)
	runMsgPack func(*testing.B)
}

func BenchmarkEncodeDecodeNamedAndArrayBitbox(b *testing.B) {
	runNamedBenchTable(b, "bitbox")
}

func BenchmarkEncodeDecodeNamedAndArrayGob(b *testing.B) {
	runNamedBenchTable(b, "gob")
}

func BenchmarkEncodeDecodeNamedAndArrayBinary(b *testing.B) {
	runNamedBenchTable(b, "binary")
}

func BenchmarkEncodeDecodeNamedAndArrayMsgPack(b *testing.B) {
	runNamedBenchTable(b, "msgpack")
}

func runNamedBenchTable(b *testing.B, codec string) {
	for _, tc := range namedBenchCases {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			b.SetBytes(tc.setBytes)
			b.ReportAllocs()

			switch codec {
			case "bitbox":
				tc.runBitbox(b)
			case "gob":
				tc.runGob(b)
			case "binary":
				tc.runBinary(b)
			case "msgpack":
				tc.runMsgPack(b)
			default:
				b.Fatalf("unknown codec: %s", codec)
			}
		})
	}
}

func makePointerScalarCase[TNamed any](name string, setBytes int64, in TNamed) namedBenchCase {
	return namedBenchCase{
		name:     name,
		setBytes: setBytes,
		runBitbox: func(b *testing.B) {
			var out TNamed
			runBitboxRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runGob: func(b *testing.B) {
			var out TNamed
			runGobRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runBinary: func(b *testing.B) {
			var out TNamed
			runBinaryRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runMsgPack: func(b *testing.B) {
			var out TNamed
			runMsgPackRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
	}
}

func makeStringCase[TNamed ~string](name string, in TNamed) namedBenchCase {
	setBytes := 2 * int64(len(in))
	return namedBenchCase{
		name:     name,
		setBytes: setBytes,
		runBitbox: func(b *testing.B) {
			var out TNamed
			runBitboxRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runGob: func(b *testing.B) {
			var out TNamed
			runGobRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runBinary: func(b *testing.B) {
			var out string
			runBinaryRawStringRoundTrip(b, string(in), &out, func() { test.AssertEqual(b, in, TNamed(out)) })
		},
		runMsgPack: func(b *testing.B) {
			var out TNamed
			runMsgPackRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
	}
}

func makeBytesCase[TNamed ~[]byte](name string, in TNamed) namedBenchCase {
	setBytes := 2 * int64(len(in))
	return namedBenchCase{
		name:     name,
		setBytes: setBytes,
		runBitbox: func(b *testing.B) {
			var out TNamed
			runBitboxRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runGob: func(b *testing.B) {
			var out TNamed
			runGobRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runBinary: func(b *testing.B) {
			var out []byte
			runBinaryRawBytesRoundTrip(b, []byte(in), &out, func() { test.AssertEqual(b, in, TNamed(out)) })
		},
		runMsgPack: func(b *testing.B) {
			var out TNamed
			runMsgPackRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
	}
}

func makeSliceLenPrefixCase[TNamed ~[]TElem, TElem any](name string, elemSize int, in TNamed) namedBenchCase {
	setBytes := 2 * int64(len(in)) * int64(elemSize)

	return namedBenchCase{
		name:     name,
		setBytes: setBytes,
		runBitbox: func(b *testing.B) {
			var out TNamed
			runBitboxRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runGob: func(b *testing.B) {
			var out TNamed
			runGobRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runBinary: func(b *testing.B) {
			out := make([]TElem, len(in))
			runBinaryRoundTrip(b, []TElem(in), &out, func() { test.AssertEqual(b, in, TNamed(out)) })
		},
		runMsgPack: func(b *testing.B) {
			var out TNamed
			runMsgPackRoundTrip(b, in, &out, func() { test.AssertEqual(b, in, out) })
		},
	}
}

func makeNamedType17Case(name string, in NamedType17) namedBenchCase {
	setBytes := int64(2 * len(in))

	return namedBenchCase{
		name:     name,
		setBytes: setBytes,
		runBitbox: func(b *testing.B) {
			var out NamedType17
			runBitboxRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runGob: func(b *testing.B) {
			var out NamedType17
			runGobRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runBinary: func(b *testing.B) {
			var out NamedType17
			runBinaryRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runMsgPack: func(b *testing.B) {
			var out NamedType17
			runMsgPackRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
	}
}

func makeNamedType18Case(name string, in NamedType18) namedBenchCase {
	const namedUint32ArraySizeBytes = 16
	setBytes := int64(namedUint32ArraySizeBytes * 2)

	return namedBenchCase{
		name:     name,
		setBytes: setBytes,
		runBitbox: func(b *testing.B) {
			var out NamedType18
			runBitboxRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runGob: func(b *testing.B) {
			var out NamedType18
			runGobRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runBinary: func(b *testing.B) {
			var out NamedType18
			runBinaryRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
		runMsgPack: func(b *testing.B) {
			var out NamedType18
			runMsgPackRoundTrip(b, &in, &out, func() { test.AssertEqual(b, in, out) })
		},
	}
}

var namedBenchCases = func() []namedBenchCase {
	inArray := NamedType17{}
	for i := range inArray {
		inArray[i] = byte(i)
	}

	return []namedBenchCase{
		makePointerScalarCase("named_bool_pointer", 2, NamedType01(true)),
		makePointerScalarCase("named_int8_pointer", 2, NamedType02(-42)),
		makePointerScalarCase("named_int16_pointer", 4, NamedType03(-4242)),
		makePointerScalarCase("named_int32_pointer", 8, NamedType04(-424242)),
		makePointerScalarCase("named_int64_pointer", 16, NamedType05(-42424242)),
		makePointerScalarCase("named_uint8_pointer", 2, NamedType06(42)),
		makePointerScalarCase("named_uint16_pointer", 4, NamedType07(4242)),
		makePointerScalarCase("named_uint32_pointer", 8, NamedType08(0xDEADBEEF)),
		makePointerScalarCase("named_uint64_pointer", 16, NamedType09(42424242)),
		makePointerScalarCase("named_float32_pointer", 8, NamedType10(3.14)),
		makePointerScalarCase("named_float64_pointer", 16, NamedType11(6.28)),
		makePointerScalarCase("named_complex64_pointer", 16, NamedType12(complex(1.5, -2.5))),
		makePointerScalarCase("named_complex128_pointer", 32, NamedType13(complex(2.5, -3.5))),
		makeStringCase("named_string_value", NamedType14("bitbox-named")),
		makeBytesCase("named_bytes_value", NamedType15{1, 2, 3, 4, 5, 6, 7, 8}),
		makeSliceLenPrefixCase[NamedType16, uint64]("named_uint64_slice_value", 8, NamedType16{1, 2, 3, 4, 5, 6, 7, 8}),
		makeNamedType17Case("named_byte_array_pointer", inArray),
		makeNamedType18Case("named_uint32_array_pointer", NamedType18{1, 2, 3, 4}),
	}
}()

func runBitboxRoundTrip(b *testing.B, in any, out any, verify func()) {
	buf := bitbox.NewBuffer(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Clear()
		bitbox.Encode(buf, in)
		bitbox.Decode(buf, out)

		verify()
	}
}

func runBitboxLenPrefixPayload(b *testing.B, in any, expectedLen uint32) {
	buf := bitbox.NewBuffer(nil)
	var encodedLen uint32

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Clear()
		bitbox.Encode(buf, in)
		bitbox.Decode(buf, &encodedLen)

		test.AssertEqual(b, expectedLen, encodedLen)
	}
}

func runGobRoundTrip(b *testing.B, in any, out any, verify func()) {
	var wire bytes.Buffer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		enc := gob.NewEncoder(&wire)
		if err := enc.Encode(in); err != nil {
			b.Fatalf("%v", err)
		}
		dec := gob.NewDecoder(bytes.NewReader(wire.Bytes()))
		if err := dec.Decode(out); err != nil {
			b.Fatalf("%v", err)
		}

		verify()
	}
}

func runBinaryRoundTrip(b *testing.B, in any, out any, verify func()) {
	var wire bytes.Buffer
	r := bytes.NewReader(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		if err := binary.Write(&wire, binary.BigEndian, in); err != nil {
			b.Fatalf("%v", err)
		}
		r.Reset(wire.Bytes())
		if err := binary.Read(r, binary.BigEndian, out); err != nil {
			b.Fatalf("%v", err)
		}

		verify()
	}
}

func runBinaryRawStringRoundTrip(b *testing.B, in string, out *string, verify func()) {
	var wire bytes.Buffer
	r := bytes.NewReader(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		if _, err := wire.WriteString(in); err != nil {
			b.Fatalf("%v", err)
		}
		r.Reset(wire.Bytes())
		payload, err := io.ReadAll(r)
		if err != nil {
			b.Fatalf("%v", err)
		}
		*out = string(payload)
		verify()
	}
}

func runBinaryRawBytesRoundTrip(b *testing.B, in []byte, out *[]byte, verify func()) {
	var wire bytes.Buffer
	r := bytes.NewReader(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		if _, err := wire.Write(in); err != nil {
			b.Fatalf("%v", err)
		}
		r.Reset(wire.Bytes())
		payload, err := io.ReadAll(r)
		if err != nil {
			b.Fatalf("%v", err)
		}
		*out = append((*out)[:0], payload...)

		verify()
	}
}

func runMsgPackRoundTrip(b *testing.B, in any, out any, verify func()) {
	defer func() {
		if r := recover(); r != nil {
			b.Skipf("msgpack unsupported type %T: %v", in, r)
		}
	}()

	wire := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(wire)
	dec := msgpack.NewDecoder(wire)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		if err := enc.Encode(in); err != nil {
			b.Skipf("msgpack unsupported type %T: %v", in, err)
			return
		}
		if err := dec.Decode(out); err != nil {
			b.Skipf("msgpack decode unsupported type %T: %v", in, err)
			return
		}

		verify()
	}
}
