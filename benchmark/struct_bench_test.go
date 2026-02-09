package benchmark

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"testing"
	"unsafe"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
)

// 4 fixed-size fields, aligned to 8 bytes, total size = 16 bytes.
type aligned4Fields struct {
	A uint64
	B uint32
	C uint16
	D uint16
}

func benchmarkStructBitbox(b *testing.B, in aligned4Fields) {
	var out aligned4Fields
	b.SetBytes(int64(unsafe.Sizeof(in)) * 2)
	b.ReportAllocs()

	buf := bitbox.NewBuffer([]byte{})
	buf = bitbox.Encode(buf, &in)
	bitbox.Decode(buf, &out)
	test.AssertEqual(b, in, out)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Clear()
		buf = bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)
	}
}

func benchmarkStructGob(b *testing.B, in aligned4Fields) {
	var out aligned4Fields
	b.SetBytes(int64(unsafe.Sizeof(in)) * 2)
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

func benchmarkStructBinary(b *testing.B, in aligned4Fields) {
	var out aligned4Fields
	b.SetBytes(int64(unsafe.Sizeof(in)) * 2)
	b.ReportAllocs()

	var wire bytes.Buffer
	if err := binary.Write(&wire, binary.BigEndian, in); err != nil {
		b.Fatalf("%v", err)
	}

	r := bytes.NewReader(wire.Bytes())
	if err := binary.Read(r, binary.BigEndian, &out); err != nil {
		b.Fatalf("%v", err)
	}

	test.AssertEqual(b, in, out)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wire.Reset()
		if err := binary.Write(&wire, binary.BigEndian, in); err != nil {
			b.Fatalf("%v", err)
		}

		r.Reset(wire.Bytes())
		if err := binary.Read(r, binary.BigEndian, &out); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func BenchmarkEncodeDecodeStruct(b *testing.B) {
	in := aligned4Fields{
		A: 0x1122334455667788,
		B: 0x99AABBCC,
		C: 0xDDEE,
		D: 0xFF00,
	}

	b.Run("Bitbox", func(b *testing.B) {
		benchmarkStructBitbox(b, in)
	})

	b.Run("Gob", func(b *testing.B) {
		benchmarkStructGob(b, in)
	})

	b.Run("BinaryWriteRead", func(b *testing.B) {
		benchmarkStructBinary(b, in)
	})
}
