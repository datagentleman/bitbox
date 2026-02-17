package benchmark

import (
	"bytes"
	"encoding/gob"
	"testing"

	bitbox "github.com/datagentleman/bitbox"
	"github.com/datagentleman/bitbox/test"
	"github.com/vmihailenco/msgpack/v5"
)

type txNamedArray [32]uint8
type txNamedSlice []uint16

type tx struct {
	ChainID    *uint64
	Nonce      uint64
	GasPrice   *uint64
	Gas        uint64
	To         *txNamedArray
	Value      *uint64
	Data       []byte
	AccessList txNamedSlice
}

const txLogicalSizeBytes = 96

func makeTx() tx {
	chainID := uint64(11155111)
	gasPrice := uint64(20_000_000_000)
	value := uint64(12345)
	to := txNamedArray{}
	for i := range to {
		to[i] = uint8(i + 1)
	}

	return tx{
		ChainID:    &chainID,
		Nonce:      42,
		GasPrice:   &gasPrice,
		Gas:        21000,
		To:         &to,
		Value:      &value,
		Data:       []byte{9, 8, 7, 6},
		AccessList: txNamedSlice{1, 3, 5, 7},
	}
}

func benchmarkTxBitbox(b *testing.B, in tx) {
	var out tx
	b.SetBytes(txLogicalSizeBytes * 2)
	b.ReportAllocs()

	buf := bitbox.NewBuffer(nil)
	bitbox.Encode(buf, &in)
	bitbox.Decode(buf, &out)
	test.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Clear()
		bitbox.Encode(buf, &in)
		bitbox.Decode(buf, &out)
	}
}

func benchmarkTxGob(b *testing.B, in tx) {
	var out tx
	b.SetBytes(txLogicalSizeBytes * 2)
	b.ReportAllocs()

	var wire bytes.Buffer
	enc := gob.NewEncoder(&wire)
	if err := enc.Encode(&in); err != nil {
		b.Fatalf("%v", err)
	}
	dec := gob.NewDecoder(bytes.NewReader(wire.Bytes()))
	if err := dec.Decode(&out); err != nil {
		b.Fatalf("%v", err)
	}
	test.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire.Reset()
		enc := gob.NewEncoder(&wire)
		if err := enc.Encode(&in); err != nil {
			b.Fatalf("%v", err)
		}
		dec := gob.NewDecoder(bytes.NewReader(wire.Bytes()))
		if err := dec.Decode(&out); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func benchmarkTxMsgPack(b *testing.B, in tx) {
	var out tx
	b.SetBytes(txLogicalSizeBytes * 2)
	b.ReportAllocs()

	wire, err := msgpack.Marshal(&in)
	if err != nil {
		b.Fatalf("%v", err)
	}
	if err := msgpack.Unmarshal(wire, &out); err != nil {
		b.Fatalf("%v", err)
	}
	test.AssertEqual(b, in, out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wire, err := msgpack.Marshal(&in)
		if err != nil {
			b.Fatalf("%v", err)
		}
		if err := msgpack.Unmarshal(wire, &out); err != nil {
			b.Fatalf("%v", err)
		}
	}
}

func BenchmarkEncodeDecodeTx(b *testing.B) {
	in := makeTx()

	b.Run("Bitbox", func(b *testing.B) {
		benchmarkTxBitbox(b, in)
	})

	b.Run("Gob", func(b *testing.B) {
		benchmarkTxGob(b, in)
	})

	b.Run("MsgPack", func(b *testing.B) {
		benchmarkTxMsgPack(b, in)
	})
}
