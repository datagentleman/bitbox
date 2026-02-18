```
 _______  ___  _______   _______  _______  __   __ 
|  _    ||   ||       | |  _    ||       ||  |_|  |
| |_|   ||   ||_     _| | |_|   ||   _   ||       |
|       ||   |  |   |   |       ||  | |  ||       |
|  _   | |   |  |   |   |  _   | |  |_|  | |     | 
| |_|   ||   |  |   |   | |_|   ||       ||   _   |
|_______||___|  |___|   |_______||_______||__| |__|
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

# Things to do

* add support for nested slices
* add support for nested arrays
* add support for slices with non fixed values
* add support for arrays with non fixed values
* add support for Maps
* add examples
* clean code
* clean tests
* clean benchmarks (0/5)

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
| slice_128B | Bitbox | 75.82 | 3376.52 | 156 | 3 |
| slice_128B | Gob | 1290 | 198.49 | 2056 | 30 |
| slice_128B | Binary | 106.5 | 2403.85 | 512 | 1 |
| slice_128B | MsgPack | 75.61 | 3385.87 | 152 | 2 |
| slice_4KB | Bitbox | 591.5 | 13848.72 | 4124 | 3 |
| slice_4KB | Gob | 2794 | 2931.94 | 15464 | 30 |
| slice_4KB | Binary | 4548 | 1801.29 | 17408 | 7 |
| slice_4KB | MsgPack | 893.0 | 9173.86 | 4120 | 2 |
| slice_64KB | Bitbox | 6867 | 19085.87 | 65564 | 3 |
| slice_64KB | Gob | 23786 | 5510.44 | 214633 | 30 |
| slice_64KB | Binary | 31277 | 4190.72 | 284930 | 16 |
| slice_64KB | MsgPack | 8660 | 15135.33 | 65560 | 2 |
| bool | Bitbox | 11.11 | 180.09 | 0 | 0 |
| bool | Gob | 947.6 | 2.11 | 1552 | 25 |
| bool | Binary | 33.07 | 60.48 | 2 | 2 |
| bool | MsgPack | 8.199 | 243.93 | 0 | 0 |
| string | Bitbox | 54.66 | 585.48 | 36 | 3 |
| string | Gob | 998.7 | 32.04 | 1640 | 28 |
| string | Binary | 128.7 | 248.73 | 528 | 2 |
| string | MsgPack | 58.94 | 542.94 | 32 | 2 |
| int8 | Bitbox | 10.35 | 193.23 | 0 | 0 |
| int8 | Gob | 1449 | 1.38 | 1552 | 25 |
| int8 | Binary | 37.99 | 52.64 | 2 | 2 |
| int8 | MsgPack | 29.12 | 68.69 | 0 | 0 |
| int16 | Bitbox | 18.47 | 216.55 | 2 | 1 |
| int16 | Gob | 1152 | 3.47 | 1552 | 26 |
| int16 | Binary | 36.54 | 109.46 | 4 | 2 |
| int16 | MsgPack | 47.50 | 84.20 | 2 | 1 |
| int32 | Bitbox | 19.77 | 404.64 | 4 | 1 |
| int32 | Gob | 1076 | 7.43 | 1560 | 26 |
| int32 | Binary | 35.78 | 223.57 | 8 | 2 |
| int32 | MsgPack | 44.99 | 177.82 | 4 | 1 |
| int64 | Bitbox | 18.22 | 878.27 | 8 | 1 |
| int64 | Gob | 1015 | 15.76 | 1568 | 26 |
| int64 | Binary | 38.14 | 419.52 | 16 | 2 |
| int64 | MsgPack | 30.74 | 520.42 | 8 | 1 |
| uint8 | Bitbox | 11.04 | 181.14 | 0 | 0 |
| uint8 | Gob | 1109 | 1.80 | 1552 | 25 |
| uint8 | Binary | 36.87 | 54.24 | 2 | 2 |
| uint8 | MsgPack | 29.43 | 67.95 | 0 | 0 |
| uint16 | Bitbox | 20.99 | 190.53 | 2 | 1 |
| uint16 | Gob | 1684 | 2.37 | 1552 | 26 |
| uint16 | Binary | 44.90 | 89.09 | 4 | 2 |
| uint16 | MsgPack | 44.47 | 89.94 | 2 | 1 |
| uint32 | Bitbox | 22.79 | 351.09 | 4 | 1 |
| uint32 | Gob | 1011 | 7.91 | 1560 | 26 |
| uint32 | Binary | 34.84 | 229.63 | 8 | 2 |
| uint32 | MsgPack | 43.68 | 183.16 | 4 | 1 |
| uint64 | Bitbox | 17.68 | 905.05 | 8 | 1 |
| uint64 | Gob | 1030 | 15.53 | 1568 | 26 |
| uint64 | Binary | 37.29 | 429.02 | 16 | 2 |
| uint64 | MsgPack | 29.91 | 534.99 | 8 | 1 |
| uintptr | Bitbox | 17.98 | 889.74 | 8 | 1 |
| uintptr | Gob | 1074 | 14.90 | 1584 | 26 |
| uintptr | Binary | 39.55 | 404.53 | 16 | 2 |
| uintptr | MsgPack | --- | --- | --- | --- |
| float32 | Bitbox | 21.87 | 365.80 | 4 | 1 |
| float32 | Gob | 1164 | 6.87 | 1576 | 26 |
| float32 | Binary | 40.02 | 199.90 | 8 | 2 |
| float32 | MsgPack | 28.88 | 277.00 | 4 | 1 |
| float64 | Bitbox | 18.12 | 883.10 | 8 | 1 |
| float64 | Gob | 1093 | 14.64 | 1624 | 27 |
| float64 | Binary | 37.77 | 423.66 | 16 | 2 |
| float64 | MsgPack | 30.36 | 526.98 | 8 | 1 |
| complex64 | Bitbox | 20.60 | 776.84 | 8 | 1 |
| complex64 | Gob | 1156 | 13.84 | 1624 | 27 |
| complex64 | Binary | 63.22 | 253.09 | 16 | 2 |
| complex64 | MsgPack | --- | --- | --- | --- |
| complex128 | Bitbox | 28.35 | 1128.76 | 16 | 1 |
| complex128 | Gob | 1087 | 29.44 | 1640 | 27 |
| complex128 | Binary | 66.86 | 478.58 | 32 | 2 |
| complex128 | MsgPack | --- | --- | --- | --- |
# Benchmark Results (Named Types, Arrays, Slices)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bool_pointer | Bitbox | 22.78 | 87.78 | 0 | 0 |
| named_bool_pointer | Gob | 1091 | 1.83 | 1616 | 26 |
| named_bool_pointer | Binary | 50.21 | 39.83 | 2 | 2 |
| named_bool_pointer | MsgPack | 41.72 | 47.93 | 0 | 0 |
| named_int8_pointer | Bitbox | 26.39 | 75.80 | 0 | 0 |
| named_int8_pointer | Gob | 1140 | 1.75 | 1616 | 26 |
| named_int8_pointer | Binary | 52.33 | 38.22 | 2 | 2 |
| named_int8_pointer | MsgPack | 51.93 | 38.51 | 0 | 0 |
| named_int16_pointer | Bitbox | 24.96 | 160.26 | 0 | 0 |
| named_int16_pointer | Gob | 1109 | 3.61 | 1616 | 26 |
| named_int16_pointer | Binary | 60.17 | 66.48 | 4 | 2 |
| named_int16_pointer | MsgPack | 56.89 | 70.32 | 0 | 0 |
| named_int32_pointer | Bitbox | 20.76 | 385.31 | 0 | 0 |
| named_int32_pointer | Gob | 1120 | 7.14 | 1616 | 26 |
| named_int32_pointer | Binary | 71.14 | 112.46 | 8 | 2 |
| named_int32_pointer | MsgPack | 59.04 | 135.51 | 0 | 0 |
| named_int64_pointer | Bitbox | 20.27 | 789.49 | 0 | 0 |
| named_int64_pointer | Gob | 1141 | 14.02 | 1616 | 26 |
| named_int64_pointer | Binary | 56.27 | 284.34 | 16 | 2 |
| named_int64_pointer | MsgPack | 58.78 | 272.21 | 0 | 0 |
| named_uint8_pointer | Bitbox | 21.92 | 91.22 | 0 | 0 |
| named_uint8_pointer | Gob | 1075 | 1.86 | 1616 | 26 |
| named_uint8_pointer | Binary | 47.00 | 42.56 | 2 | 2 |
| named_uint8_pointer | MsgPack | 50.06 | 39.95 | 0 | 0 |
| named_uint16_pointer | Bitbox | 21.48 | 186.22 | 0 | 0 |
| named_uint16_pointer | Gob | 1042 | 3.84 | 1616 | 26 |
| named_uint16_pointer | Binary | 52.34 | 76.42 | 4 | 2 |
| named_uint16_pointer | MsgPack | 58.22 | 68.70 | 0 | 0 |
| named_uint32_pointer | Bitbox | 21.90 | 365.34 | 0 | 0 |
| named_uint32_pointer | Gob | 1045 | 7.66 | 1616 | 26 |
| named_uint32_pointer | Binary | 59.51 | 134.43 | 8 | 2 |
| named_uint32_pointer | MsgPack | 55.55 | 144.03 | 0 | 0 |
| named_uint64_pointer | Bitbox | 21.27 | 752.38 | 0 | 0 |
| named_uint64_pointer | Gob | 1061 | 15.08 | 1616 | 26 |
| named_uint64_pointer | Binary | 55.54 | 288.06 | 16 | 2 |
| named_uint64_pointer | MsgPack | 60.10 | 266.20 | 0 | 0 |
| named_float32_pointer | Bitbox | 21.48 | 372.49 | 0 | 0 |
| named_float32_pointer | Gob | 1173 | 6.82 | 1656 | 27 |
| named_float32_pointer | Binary | 52.85 | 151.37 | 8 | 2 |
| named_float32_pointer | MsgPack | 70.30 | 113.80 | 0 | 0 |
| named_float64_pointer | Bitbox | 27.22 | 587.77 | 0 | 0 |
| named_float64_pointer | Gob | 1305 | 12.26 | 1664 | 27 |
| named_float64_pointer | Binary | 53.37 | 299.79 | 16 | 2 |
| named_float64_pointer | MsgPack | 63.97 | 250.10 | 0 | 0 |
| named_complex64_pointer | Bitbox | 20.93 | 764.40 | 0 | 0 |
| named_complex64_pointer | Gob | 1157 | 13.83 | 1656 | 27 |
| named_complex64_pointer | Binary | 56.18 | 284.82 | 16 | 2 |
| named_complex64_pointer | MsgPack | --- | --- | --- | --- |
| named_complex128_pointer | Bitbox | 20.48 | 1562.42 | 0 | 0 |
| named_complex128_pointer | Gob | 1117 | 28.66 | 1656 | 27 |
| named_complex128_pointer | Binary | 73.68 | 434.31 | 32 | 2 |
| named_complex128_pointer | MsgPack | --- | --- | --- | --- |
| named_string_value | Bitbox | 46.71 | 513.83 | 16 | 2 |
| named_string_value | Gob | 1154 | 20.80 | 1680 | 28 |
| named_string_value | Binary | 177.2 | 135.42 | 528 | 2 |
| named_string_value | MsgPack | 75.97 | 315.90 | 16 | 1 |
| named_bytes_value | Bitbox | 52.08 | 307.22 | 4 | 1 |
| named_bytes_value | Gob | 1174 | 13.63 | 1712 | 28 |
| named_bytes_value | Binary | 159.6 | 100.24 | 512 | 1 |
| named_bytes_value | MsgPack | 63.33 | 252.66 | 0 | 0 |
| named_uint64_slice_value | Bitbox | 65.03 | 1968.23 | 4 | 1 |
| named_uint64_slice_value | Gob | 9743 | 13.14 | 7704 | 187 |
| named_uint64_slice_value | Binary | 141.0 | 907.72 | 128 | 2 |
| named_uint64_slice_value | MsgPack | 527.3 | 242.75 | 24 | 1 |
| named_byte_array_pointer | Bitbox | 52.85 | 37841.77 | 0 | 0 |
| named_byte_array_pointer | Gob | 24015 | 83.28 | 14488 | 192 |
| named_byte_array_pointer | Binary | 9687 | 206.46 | 2048 | 2 |
| named_byte_array_pointer | MsgPack | 138.7 | 14414.95 | 48 | 2 |
| named_uint32_array_pointer | Bitbox | 29.38 | 1089.12 | 0 | 0 |
| named_uint32_array_pointer | Gob | 9970 | 3.21 | 7696 | 187 |
| named_uint32_array_pointer | Binary | 120.7 | 265.17 | 32 | 2 |
| named_uint32_array_pointer | MsgPack | 275.1 | 116.31 | 0 | 0 |
# Benchmark Results (Aligned Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 21.61 | 1480.79 | 0 | 0 |
| EncodeDecodeStruct | Gob | 12198 | 2.62 | 8112 | 204 |
| EncodeDecodeStruct | BinaryWriteRead | 207.4 | 154.30 | 32 | 2 |
| EncodeDecodeStruct | MsgPack | 367.1 | 87.18 | 0 | 0 |
# Benchmark Results (Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeTx | Bitbox | 546.6 | 351.24 | 192 | 2 |
| EncodeDecodeTx | Gob | 22600 | 8.50 | 10160 | 262 |
| EncodeDecodeTx | Binary | --- | --- | --- | --- |
| EncodeDecodeTx | MsgPack | 3062 | 62.70 | 264 | 5 |
