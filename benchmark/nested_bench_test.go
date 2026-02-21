package benchmark

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"testing"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
	"github.com/vmihailenco/msgpack/v5"
)

func benchmarkNestedBitbox[T any](b *testing.B, in T, setBytes int64) {
	var out T
	var zero T

	b.SetBytes(setBytes)
	b.ReportAllocs()

	buf := bitbox.NewBuffer(nil)
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

func benchmarkNestedGob[T any](b *testing.B, in T, setBytes int64) {
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

func benchmarkNestedMsgPack[T any](b *testing.B, in T, setBytes int64) {
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

func benchmarkNestedBinaryArray100x32Byte(b *testing.B, in [100][32]byte, setBytes int64) {
	var out [100][32]byte

	b.SetBytes(setBytes)
	b.ReportAllocs()

	var wire bytes.Buffer
	if err := binary.Write(&wire, binary.BigEndian, &in); err != nil {
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
		if err := binary.Write(&wire, binary.BigEndian, &in); err != nil {
			b.Fatalf("%v", err)
		}
		r.Reset(wire.Bytes())
		if err := binary.Read(r, binary.BigEndian, &out); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func encodeBinary2DBytes(b *testing.B, wire *bytes.Buffer, in [][]byte) {
	rows := uint32(len(in))
	if err := binary.Write(wire, binary.BigEndian, rows); err != nil {
		b.Fatalf("%v", err)
	}
	for _, row := range in {
		l := uint32(len(row))
		if err := binary.Write(wire, binary.BigEndian, l); err != nil {
			b.Fatalf("%v", err)
		}
		if _, err := wire.Write(row); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func decodeBinary2DBytes(b *testing.B, r *bytes.Reader, out *[][]byte) {
	var rows uint32
	if err := binary.Read(r, binary.BigEndian, &rows); err != nil {
		b.Fatalf("%v", err)
	}

	n := int(rows)
	if cap(*out) < n {
		*out = make([][]byte, n)
	} else {
		*out = (*out)[:n]
	}

	for i := 0; i < n; i++ {
		var l uint32
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			b.Fatalf("%v", err)
		}
		m := int(l)
		if cap((*out)[i]) < m {
			(*out)[i] = make([]byte, m)
		} else {
			(*out)[i] = (*out)[i][:m]
		}
		if _, err := r.Read((*out)[i]); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func benchmarkNestedBinary2DBytes(b *testing.B, in [][]byte, setBytes int64) {
	var out [][]byte

	b.SetBytes(setBytes)
	b.ReportAllocs()

	var wire bytes.Buffer
	encodeBinary2DBytes(b, &wire, in)
	r := bytes.NewReader(wire.Bytes())
	decodeBinary2DBytes(b, r, &out)
	test.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		encodeBinary2DBytes(b, &wire, in)
		r.Reset(wire.Bytes())
		decodeBinary2DBytes(b, r, &out)
	}
}

func encodeBinary2DUint64(b *testing.B, wire *bytes.Buffer, in [][]uint64) {
	rows := uint32(len(in))
	if err := binary.Write(wire, binary.BigEndian, rows); err != nil {
		b.Fatalf("%v", err)
	}
	for _, row := range in {
		l := uint32(len(row))
		if err := binary.Write(wire, binary.BigEndian, l); err != nil {
			b.Fatalf("%v", err)
		}
		if l == 0 {
			continue
		}
		if err := binary.Write(wire, binary.BigEndian, row); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func decodeBinary2DUint64(b *testing.B, r *bytes.Reader, out *[][]uint64) {
	var rows uint32
	if err := binary.Read(r, binary.BigEndian, &rows); err != nil {
		b.Fatalf("%v", err)
	}

	n := int(rows)
	if cap(*out) < n {
		*out = make([][]uint64, n)
	} else {
		*out = (*out)[:n]
	}

	for i := 0; i < n; i++ {
		var l uint32
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			b.Fatalf("%v", err)
		}
		m := int(l)
		if cap((*out)[i]) < m {
			(*out)[i] = make([]uint64, m)
		} else {
			(*out)[i] = (*out)[i][:m]
		}
		if m == 0 {
			continue
		}
		if err := binary.Read(r, binary.BigEndian, (*out)[i]); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func benchmarkNestedBinary2DUint64(b *testing.B, in [][]uint64, setBytes int64) {
	var out [][]uint64

	b.SetBytes(setBytes)
	b.ReportAllocs()

	var wire bytes.Buffer
	encodeBinary2DUint64(b, &wire, in)
	r := bytes.NewReader(wire.Bytes())
	decodeBinary2DUint64(b, r, &out)
	test.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		encodeBinary2DUint64(b, &wire, in)
		r.Reset(wire.Bytes())
		decodeBinary2DUint64(b, r, &out)
	}
}

func makeNestedArray100x32Byte() [100][32]byte {
	var out [100][32]byte
	for i := range out {
		for j := range out[i] {
			out[i][j] = byte((i*31 + j*17) & 0xff)
		}
	}
	return out
}

func makeNested2DBytes() [][]byte {
	rows := 100
	cols := 64
	out := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		row := make([]byte, cols)
		for j := 0; j < cols; j++ {
			row[j] = byte((i*29 + j*13) & 0xff)
		}
		out[i] = row
	}
	return out
}

func makeNested2DUint64() [][]uint64 {
	rows := 100
	cols := 64
	out := make([][]uint64, rows)
	for i := 0; i < rows; i++ {
		row := make([]uint64, cols)
		for j := 0; j < cols; j++ {
			row[j] = uint64(i*1000 + j)
		}
		out[i] = row
	}
	return out
}

func makeNested2DTx() [][]tx {
	rows := 40
	cols := 10
	out := make([][]tx, rows)
	for i := 0; i < rows; i++ {
		row := make([]tx, cols)
		for j := 0; j < cols; j++ {
			v := makeTx()
			v.Nonce = uint64(i*cols + j)
			row[j] = v
		}
		out[i] = row
	}
	return out
}

func makeNestedArray100x100Tx() [100][100]tx {
	var out [100][100]tx
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			v := makeTx()
			v.Nonce = uint64(i*100 + j)
			out[i][j] = v
		}
	}
	return out
}

func BenchmarkEncodeDecodeNested(b *testing.B) {
	array100x32 := makeNestedArray100x32Byte()
	slice2DBytes := makeNested2DBytes()
	slice2DU64 := makeNested2DUint64()
	slice2DTx := makeNested2DTx()
	array100x100Tx := makeNestedArray100x100Tx()

	b.Run("array_100x32_byte", func(b *testing.B) {
		setBytes := int64(100*32) * 2

		b.Run("Bitbox", func(b *testing.B) {
			benchmarkNestedBitbox(b, array100x32, setBytes)
		})
		b.Run("Gob", func(b *testing.B) {
			benchmarkNestedGob(b, array100x32, setBytes)
		})
		b.Run("Binary", func(b *testing.B) {
			benchmarkNestedBinaryArray100x32Byte(b, array100x32, setBytes)
		})
		b.Run("MsgPack", func(b *testing.B) {
			benchmarkNestedMsgPack(b, array100x32, setBytes)
		})
	})

	b.Run("slice_2d_bytes", func(b *testing.B) {
		setBytes := int64(100*64) * 2

		b.Run("Bitbox", func(b *testing.B) {
			benchmarkNestedBitbox(b, slice2DBytes, setBytes)
		})
		b.Run("Gob", func(b *testing.B) {
			benchmarkNestedGob(b, slice2DBytes, setBytes)
		})
		b.Run("Binary", func(b *testing.B) {
			b.Skip("binary unsupported for [][]byte")
		})
		b.Run("MsgPack", func(b *testing.B) {
			benchmarkNestedMsgPack(b, slice2DBytes, setBytes)
		})
	})

	b.Run("slice_2d_uint64", func(b *testing.B) {
		setBytes := int64(100*64*8) * 2

		b.Run("Bitbox", func(b *testing.B) {
			benchmarkNestedBitbox(b, slice2DU64, setBytes)
		})
		b.Run("Gob", func(b *testing.B) {
			benchmarkNestedGob(b, slice2DU64, setBytes)
		})
		b.Run("Binary", func(b *testing.B) {
			b.Skip("binary unsupported for [][]uint64")
		})
		b.Run("MsgPack", func(b *testing.B) {
			benchmarkNestedMsgPack(b, slice2DU64, setBytes)
		})
	})

	b.Run("slice_2d_tx", func(b *testing.B) {
		setBytes := int64(len(slice2DTx)*len(slice2DTx[0])*txLogicalSizeBytes) * 2

		b.Run("Bitbox", func(b *testing.B) {
			benchmarkNestedBitbox(b, slice2DTx, setBytes)
		})
		b.Run("Gob", func(b *testing.B) {
			benchmarkNestedGob(b, slice2DTx, setBytes)
		})
		b.Run("Binary", func(b *testing.B) {
			b.Skip("binary unsupported for [][]tx")
		})
		b.Run("MsgPack", func(b *testing.B) {
			benchmarkNestedMsgPack(b, slice2DTx, setBytes)
		})
	})

	b.Run("array_100x100_tx", func(b *testing.B) {
		setBytes := int64(100*100*txLogicalSizeBytes) * 2

		b.Run("Bitbox", func(b *testing.B) {
			benchmarkNestedBitbox(b, array100x100Tx, setBytes)
		})
		b.Run("Gob", func(b *testing.B) {
			benchmarkNestedGob(b, array100x100Tx, setBytes)
		})
		b.Run("Binary", func(b *testing.B) {
			b.Skip("binary unsupported for [100][100]tx")
		})
		b.Run("MsgPack", func(b *testing.B) {
			benchmarkNestedMsgPack(b, array100x100Tx, setBytes)
		})
	})
}
