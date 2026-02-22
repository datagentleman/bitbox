```
 _______  ___  _______    _______  _______  __   __ 
|  _    ||   ||       |  |  _    ||       ||  |_|  |
| |_|   ||   ||_     _|  | |_|   ||   _   ||       |
|       ||   |  |   |    |       ||  | |  ||       |
|  _   | |   |  |   |    |  _   | |  |_|  | |     | 
| |_|   ||   |  |   |    | |_|   ||       ||   _   |
|_______||___|  |___|    |_______||_______||__| |__|

```

# bitbox

Bitbox is a tiny, extremely fast, low-overhead binary encoding/decoding package for Go.
It provides a universal byte format where data can be easily moved across different languages and platforms,
as long as all systems use the same endianness.

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


# Benchmark Results (Basic Types, Slices)

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
| bool | Bitbox | 10.63 | 188.22 | 0 | 0 |
| bool | Gob | 1005 | 1.99 | 1552 | 25 |
| bool | Binary | 39.06 | 51.20 | 2 | 2 |
| bool | MsgPack | 10.22 | 195.71 | 0 | 0 |
| string | Bitbox | 56.75 | 563.87 | 36 | 3 |
| string | Gob | 1068 | 29.97 | 1640 | 28 |
| string | Binary | 129.0 | 248.12 | 528 | 2 |
| string | MsgPack | 51.78 | 618.04 | 32 | 2 |
| int8 | Bitbox | 10.00 | 199.94 | 0 | 0 |
| int8 | Gob | 1004 | 1.99 | 1552 | 25 |
| int8 | Binary | 34.41 | 58.11 | 2 | 2 |
| int8 | MsgPack | 27.36 | 73.11 | 0 | 0 |
| int16 | Bitbox | 16.38 | 244.14 | 2 | 1 |
| int16 | Gob | 992.2 | 4.03 | 1552 | 26 |
| int16 | Binary | 42.62 | 93.84 | 4 | 2 |
| int16 | MsgPack | 51.31 | 77.95 | 2 | 1 |
| int32 | Bitbox | 17.96 | 445.45 | 4 | 1 |
| int32 | Gob | 1063 | 7.53 | 1560 | 26 |
| int32 | Binary | 35.97 | 222.41 | 8 | 2 |
| int32 | MsgPack | 44.65 | 179.19 | 4 | 1 |
| int64 | Bitbox | 18.58 | 861.09 | 8 | 1 |
| int64 | Gob | 1020 | 15.68 | 1568 | 26 |
| int64 | Binary | 38.73 | 413.13 | 16 | 2 |
| int64 | MsgPack | 30.91 | 517.64 | 8 | 1 |
| uint8 | Bitbox | 10.56 | 189.31 | 0 | 0 |
| uint8 | Gob | 1008 | 1.98 | 1552 | 25 |
| uint8 | Binary | 33.26 | 60.12 | 2 | 2 |
| uint8 | MsgPack | 28.47 | 70.24 | 0 | 0 |
| uint16 | Bitbox | 16.89 | 236.87 | 2 | 1 |
| uint16 | Gob | 1012 | 3.95 | 1552 | 26 |
| uint16 | Binary | 35.47 | 112.78 | 4 | 2 |
| uint16 | MsgPack | 42.74 | 93.60 | 2 | 1 |
| uint32 | Bitbox | 18.34 | 436.19 | 4 | 1 |
| uint32 | Gob | 1023 | 7.82 | 1560 | 26 |
| uint32 | Binary | 39.05 | 204.87 | 8 | 2 |
| uint32 | MsgPack | 44.40 | 180.18 | 4 | 1 |
| uint64 | Bitbox | 17.64 | 906.79 | 8 | 1 |
| uint64 | Gob | 938.8 | 17.04 | 1568 | 26 |
| uint64 | Binary | 37.03 | 432.08 | 16 | 2 |
| uint64 | MsgPack | 29.35 | 545.20 | 8 | 1 |
| uintptr | Bitbox | 17.48 | 915.57 | 8 | 1 |
| uintptr | Gob | 1016 | 15.74 | 1584 | 26 |
| uintptr | Binary | 37.80 | 423.24 | 16 | 2 |
| uintptr | MsgPack | --- | --- | --- | --- |
| float32 | Bitbox | 18.43 | 434.06 | 4 | 1 |
| float32 | Gob | 1000 | 8.00 | 1576 | 26 |
| float32 | Binary | 35.56 | 224.96 | 8 | 2 |
| float32 | MsgPack | 28.67 | 279.03 | 4 | 1 |
| float64 | Bitbox | 17.19 | 930.51 | 8 | 1 |
| float64 | Gob | 1040 | 15.38 | 1624 | 27 |
| float64 | Binary | 36.91 | 433.47 | 16 | 2 |
| float64 | MsgPack | 29.59 | 540.66 | 8 | 1 |
| complex64 | Bitbox | 18.03 | 887.41 | 8 | 1 |
| complex64 | Gob | 1028 | 15.56 | 1624 | 27 |
| complex64 | Binary | 58.23 | 274.79 | 16 | 2 |
| complex64 | MsgPack | --- | --- | --- | --- |
| complex128 | Bitbox | 22.07 | 1450.16 | 16 | 1 |
| complex128 | Gob | 1047 | 30.55 | 1640 | 27 |
| complex128 | Binary | 69.22 | 462.27 | 32 | 2 |
| complex128 | MsgPack | --- | --- | --- | --- |

# Benchmark Results (Named Types, Arrays, Slices)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bool_pointer | Bitbox | 177.3 | 11.28 | 0 | 0 |
| named_bool_pointer | Gob | 1293 | 1.55 | 1616 | 26 |
| named_bool_pointer | Binary | 198.5 | 10.08 | 2 | 2 |
| named_bool_pointer | MsgPack | 194.4 | 10.29 | 0 | 0 |
| named_int8_pointer | Bitbox | 170.9 | 11.70 | 0 | 0 |
| named_int8_pointer | Gob | 1274 | 1.57 | 1616 | 26 |
| named_int8_pointer | Binary | 197.1 | 10.15 | 2 | 2 |
| named_int8_pointer | MsgPack | 198.0 | 10.10 | 0 | 0 |
| named_int16_pointer | Bitbox | 214.7 | 18.63 | 4 | 2 |
| named_int16_pointer | Gob | 1318 | 3.03 | 1621 | 28 |
| named_int16_pointer | Binary | 214.2 | 18.68 | 8 | 4 |
| named_int16_pointer | MsgPack | 222.0 | 18.02 | 4 | 2 |
| named_int32_pointer | Bitbox | 189.7 | 42.18 | 8 | 2 |
| named_int32_pointer | Gob | 1295 | 6.18 | 1632 | 28 |
| named_int32_pointer | Binary | 216.6 | 36.93 | 16 | 4 |
| named_int32_pointer | MsgPack | 226.1 | 35.38 | 8 | 2 |
| named_int64_pointer | Bitbox | 193.2 | 82.83 | 16 | 2 |
| named_int64_pointer | Gob | 1295 | 12.36 | 1632 | 28 |
| named_int64_pointer | Binary | 224.9 | 71.14 | 32 | 4 |
| named_int64_pointer | MsgPack | 240.6 | 66.50 | 16 | 2 |
| named_uint8_pointer | Bitbox | 166.3 | 12.03 | 0 | 0 |
| named_uint8_pointer | Gob | 1221 | 1.64 | 1616 | 26 |
| named_uint8_pointer | Binary | 197.7 | 10.12 | 2 | 2 |
| named_uint8_pointer | MsgPack | 199.2 | 10.04 | 0 | 0 |
| named_uint16_pointer | Bitbox | 186.2 | 21.48 | 4 | 2 |
| named_uint16_pointer | Gob | 1285 | 3.11 | 1621 | 28 |
| named_uint16_pointer | Binary | 215.1 | 18.60 | 8 | 4 |
| named_uint16_pointer | MsgPack | 227.0 | 17.62 | 4 | 2 |
| named_uint32_pointer | Bitbox | 188.7 | 42.39 | 8 | 2 |
| named_uint32_pointer | Gob | 1297 | 6.17 | 1632 | 28 |
| named_uint32_pointer | Binary | 226.3 | 35.34 | 16 | 4 |
| named_uint32_pointer | MsgPack | 227.3 | 35.20 | 8 | 2 |
| named_uint64_pointer | Bitbox | 188.8 | 84.75 | 16 | 2 |
| named_uint64_pointer | Gob | 1313 | 12.18 | 1632 | 28 |
| named_uint64_pointer | Binary | 231.6 | 69.08 | 32 | 4 |
| named_uint64_pointer | MsgPack | 224.2 | 71.38 | 16 | 2 |
| named_float32_pointer | Bitbox | 200.9 | 39.82 | 8 | 2 |
| named_float32_pointer | Gob | 1327 | 6.03 | 1664 | 29 |
| named_float32_pointer | Binary | 280.1 | 28.57 | 16 | 4 |
| named_float32_pointer | MsgPack | 227.1 | 35.23 | 8 | 2 |
| named_float64_pointer | Bitbox | 190.1 | 84.19 | 16 | 2 |
| named_float64_pointer | Gob | 1321 | 12.11 | 1680 | 29 |
| named_float64_pointer | Binary | 223.7 | 71.52 | 32 | 4 |
| named_float64_pointer | MsgPack | 230.9 | 69.31 | 16 | 2 |
| named_complex64_pointer | Bitbox | 204.9 | 78.10 | 16 | 2 |
| named_complex64_pointer | Gob | 1281 | 12.49 | 1672 | 29 |
| named_complex64_pointer | Binary | 243.3 | 65.77 | 32 | 4 |
| named_complex64_pointer | MsgPack | --- | --- | --- | --- |
| named_complex128_pointer | Bitbox | 200.4 | 159.72 | 32 | 2 |
| named_complex128_pointer | Gob | 1315 | 24.34 | 1688 | 29 |
| named_complex128_pointer | Binary | 241.4 | 132.57 | 64 | 4 |
| named_complex128_pointer | MsgPack | --- | --- | --- | --- |
| named_string_value | Bitbox | 251.3 | 95.52 | 48 | 4 |
| named_string_value | Gob | 1451 | 16.54 | 1712 | 30 |
| named_string_value | Binary | 355.2 | 67.56 | 560 | 4 |
| named_string_value | MsgPack | 258.1 | 92.97 | 48 | 3 |
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

# Benchmark Results (Aligned Struct)

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
