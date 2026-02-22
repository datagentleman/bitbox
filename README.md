```
 _______  ___  _______    _______  _______  __   __ 
|  _    ||   ||       |  |  _    ||       ||  |_|  |
| |_|   ||   ||_     _|  | |_|   ||   _   ||       |
|       ||   |  |   |    |       ||  | |  ||       |
|  _   | |   |  |   |    |  _   | |  |_|  | |     | 
| |_|   ||   |  |   |    | |_|   ||       ||   _   |
|_______||___|  |___|    |_______||_______||__| |__|



Bitbox is a tiny, extremely fast, low-overhead binary encoding/decoding package for Go.
It provides a universal byte format where data can be easily moved across different languages and platforms,
as long as all systems use the same endianness.
```

# Currently we are faster and more memory efficient than

 | MsgPack | Binary | Gob |
 |:-------:|:------:|:---:|
 |   ✅    |   ✅   | ✅  |

# Endianness warning

Bitbox writes/read all data using the machine’s native endianness (typically little-endian).
To ensure correct decoding, both encoder and decoder must run on platforms with the same endianness.

# Quick Examples

```go
// Simple example

leet1 := int64(1337)
leet2 := int64(0)

bit := bitbox.NewBuffer(nil)
bit.Encode(leet1)
bit.Decode(&leet2)


// We can also call bitbox with multiple values

name1 := "bitbox"
name2 := ""

age1 := int8(1)
age2 := int8(0)

methods1 := []string{"Encode", "Decode"}
methods2 := []string{}

bit.Clear()
bit.Encode(name1, age1, methods1)
bit.Decode(&name2, &age2, &methods2)


// Bitbox is especially fast when using with PODs (plain old data)

type POD struct {
  A uint64
  B uint32
  C uint16
  D uint16
}

pod1 := POD{A: uint64(1337), B: uint32(1447), C: uint16(1557), D: uint16(1667)}
pod2 := POD{}

bit.Clear()
bit.EncodePOD(&pod1)
bit.DecodePOD(&pod2)


// We can also use it with normal structs

type Tx struct {
  Price    uint64
  Sender   *[20]byte
  Receiver []byte
}

tx1 := Tx{Price: uint64(100), Sender: &[20]byte{1, 3, 3, 7}, Receiver: []byte{1, 4, 4, 7}}
tx2 := Tx{}

bit.Clear()
bit.Encode(&tx1)
bit.Decode(&tx2)
```

# Things to do

* add missing types (Maps, int, uint, math.big, ...)
* clean POD code
* clean benchmarks
* clean tests


# Run Tests

Run all tests in the module (to run them without cache use --count=1):

```bash
go test ./...
```

# Run Benchmarks

```bash
go test -run '^$' -bench Benchmark -benchmem -count=1 ./benchmark
```

The benchmark output includes:

- `ns/op`: time per benchmark operation
- `MB/s`: throughput in megabytes per second
- `B/op` and `allocs/op`: allocation metrics


For all benchmark tables (including basic types and named scalar/pointer types), 
see `BENCHMARKS.md`.


# Benchmark Results (Slices, Arrays)

| Type | Codec | ns/op | MB/s | B/op | allocs/op |
|-----:|:------|------:|-----:|-----:|----------:|
| slice_128B | Bitbox | 76.93 | 3327.52 | 156 | 3 |
| slice_128B | Gob | 1195 | 214.15 | 2056 | 30 |
| slice_128B | Binary | 106.7 | 2398.70 | 512 | 1 |
| slice_128B | MsgPack | 76.93 | 3327.85 | 152 | 2 |
| slice_4KB | Bitbox | 649.2 | 12618.74 | 4124 | 3 |
| slice_4KB | Gob | 2868 | 2856.39 | 15464 | 30 |
| slice_4KB | Binary | 2326 | 3522.10 | 17408 | 7 |
| slice_4KB | MsgPack | 676.4 | 12111.88 | 4120 | 2 |
| slice_64KB | Bitbox | 7987 | 16411.23 | 65564 | 3 |
| slice_64KB | Gob | 26978 | 4858.51 | 214633 | 30 |
| slice_64KB | Binary | 42987 | 3049.10 | 284930 | 16 |
| slice_64KB | MsgPack | 8815 | 14868.98 | 65560 | 2 |

# Benchmark Results (Named Slices, Arrays)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bytes_value | Bitbox | 289.0 | 55.37 | 52 | 3 |
| named_bytes_value | Gob | 1496 | 10.69 | 1760 | 30 |
| named_bytes_value | Binary | 375.3 | 42.63 | 560 | 3 |
| named_bytes_value | MsgPack | 305.3 | 52.40 | 48 | 2 |
| named_uint64_slice_value | Bitbox | 379.6 | 337.17 | 52 | 3 |
| named_uint64_slice_value | Gob | 9976 | 12.83 | 7752 | 189 |
| named_uint64_slice_value | Binary | 483.1 | 264.93 | 176 | 4 |
| named_uint64_slice_value | MsgPack | 879.9 | 145.48 | 72 | 3 |
| named_byte_array_pointer | Bitbox | 10943 | 182.76 | 2048 | 2 |
| named_byte_array_pointer | Gob | 35475 | 56.38 | 16537 | 194 |
| named_byte_array_pointer | Binary | 20372 | 98.17 | 4096 | 4 |
| named_byte_array_pointer | MsgPack | 10851 | 184.32 | 2096 | 4 |
| named_uint32_array_pointer | Bitbox | 252.4 | 126.78 | 32 | 2 |
| named_uint32_array_pointer | Gob | 9767 | 3.28 | 7728 | 189 |
| named_uint32_array_pointer | Binary | 349.2 | 91.64 | 64 | 4 |
| named_uint32_array_pointer | MsgPack | 487.3 | 65.67 | 32 | 2 |

# Benchmark Results (Nested Slices and Arrays)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| array_100x32_byte | Bitbox | 6567 | 974.56 | 6400 | 101 |
| array_100x32_byte | Gob | 71102 | 90.01 | 40736 | 302 |
| array_100x32_byte | Binary | 32851 | 194.82 | 6400 | 2 |
| array_100x32_byte | MsgPack | 10328 | 619.67 | 5600 | 101 |
| slice_2d_bytes | Bitbox | 5517 | 2320.22 | 9516 | 203 |
| slice_2d_bytes | Gob | 25844 | 495.29 | 53448 | 496 |
| slice_2d_bytes | Binary | --- | --- | --- | --- |
| slice_2d_bytes | MsgPack | 9432 | 1357.04 | 11872 | 106 |
| slice_2d_uint64 | Bitbox | 12834 | 7978.63 | 54316 | 203 |
| slice_2d_uint64 | Gob | 107248 | 954.80 | 175393 | 609 |
| slice_2d_uint64 | Binary | --- | --- | --- | --- |
| slice_2d_uint64 | MsgPack | 361018 | 283.64 | 115072 | 506 |
| slice_2d_tx | Bitbox | 194678 | 394.50 | 94512 | 4124 |
| slice_2d_tx | Gob | 487012 | 157.70 | 326205 | 5620 |
| slice_2d_tx | Binary | --- | --- | --- | --- |
| slice_2d_tx | MsgPack | 497737 | 154.30 | 166944 | 5006 |
| array_100x100_tx | Bitbox | 5031835 | 381.57 | 3206659 | 110001 |
| array_100x100_tx | Gob | 11173568 | 171.83 | 8783965 | 120412 |
| array_100x100_tx | Binary | --- | --- | --- | --- |
| array_100x100_tx | MsgPack | 11894120 | 161.42 | 2966658 | 120001 |

# Benchmark Results (POD structs)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 5646 | 566.75 | 52 | 3 |
| EncodeDecodeStruct | Gob | 30771 | 103.99 | 19940 | 232 |
| EncodeDecodeStruct | BinaryWriteRead | 13166 | 243.05 | 3696 | 12 |
| EncodeDecodeStruct | MsgPack | 40368 | 79.27 | 72 | 3 |

# Benchmark Results (Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeTx | Bitbox | 1267 | 151.58 | 200 | 4 |
| EncodeDecodeTx | Gob | 16086 | 11.94 | 10160 | 262 |
| EncodeDecodeTx | MsgPack | 2002 | 95.90 | 264 | 5 |

