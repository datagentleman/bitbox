package benchmark

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"testing"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/vmihailenco/msgpack/v5"
)

// 4 fixed-size fields, aligned to 8 bytes, total size = 16 bytes.
type aligned4Fields struct {
	A uint64
	B uint32
	C uint16
	D uint16
}

const aligned4FieldsSizeBytes = 16

func benchmarkStructBitbox(b *testing.B, in []aligned4Fields) {
	out := []aligned4Fields{}
	b.SetBytes(int64(aligned4FieldsSizeBytes * len(in) * 2))
	b.ReportAllocs()

	buf := bitbox.NewBuffer([]byte{})
	bitbox.EncodePOD(buf, &in)
	bitbox.DecodePOD(buf, &out)
	bitbox.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Clear()
		bitbox.EncodePOD(buf, &in)
		bitbox.DecodePOD(buf, &out)
	}
}

func benchmarkStructGob(b *testing.B, in []aligned4Fields) {
	var out []aligned4Fields
	b.SetBytes(int64(aligned4FieldsSizeBytes * len(in) * 2))
	b.ReportAllocs()

	var wire bytes.Buffer
	enc := gob.NewEncoder(&wire)
	if err := enc.Encode(&in); err != nil {
		b.Fatalf("%v", err)
	}

	r := bytes.NewReader(wire.Bytes())
	dec := gob.NewDecoder(r)
	if err := dec.Decode(&out); err != nil {
		b.Fatalf("%v", err)
	}
	bitbox.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		enc := gob.NewEncoder(&wire)
		enc.Encode(&in)

		r.Reset(wire.Bytes())
		dec := gob.NewDecoder(r)
		dec.Decode(&out)
	}
}

func benchmarkStructBinary(b *testing.B, in []aligned4Fields) {
	out := make([]aligned4Fields, len(in))
	b.SetBytes(int64(aligned4FieldsSizeBytes * len(in) * 2))
	b.ReportAllocs()

	var wire bytes.Buffer
	if err := binary.Write(&wire, binary.BigEndian, &in); err != nil {
		b.Fatalf("%v", err)
	}

	r := bytes.NewReader(wire.Bytes())
	if err := binary.Read(r, binary.BigEndian, &out); err != nil {
		b.Fatalf("%v", err)
	}

	bitbox.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		binary.Write(&wire, binary.BigEndian, &in)
		r.Reset(wire.Bytes())
		binary.Read(r, binary.BigEndian, &out)
	}
}

func benchmarkStructMsgPack(b *testing.B, in []aligned4Fields) {
	var out []aligned4Fields
	b.SetBytes(int64(aligned4FieldsSizeBytes * len(in) * 2))
	b.ReportAllocs()

	wire := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(wire)
	dec := msgpack.NewDecoder(wire)

	if err := enc.Encode(&in); err != nil {
		b.Fatalf("%v", err)
	}
	if err := dec.Decode(&out); err != nil {
		b.Fatalf("%v", err)
	}
	bitbox.AssertEqual(b, in, out)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wire.Reset()
		enc.Encode(&in)
		dec.Decode(&out)
	}
}

func BenchmarkEncodeDecodeStruct(b *testing.B) {
	pod := aligned4Fields{
		A: 0x1122334455667788,
		B: 0x99AABBCC,
		C: 0xDDEE,
		D: 0xFF00,
	}

	in := make([]aligned4Fields, 0, 100)
	for i := 0; i < 100; i++ {
		in = append(in, pod)
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

	b.Run("MsgPack", func(b *testing.B) {
		benchmarkStructMsgPack(b, in)
	})
}
