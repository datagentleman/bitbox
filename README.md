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

# Things to do

* ~~add support for nested slices~~
* ~~add support for nested arrays~~
* add support for POD structs
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
| slice_128B | Bitbox | 76.82 | 3332.55 | 156 | 3 |
| slice_128B | Gob | 1253 | 204.38 | 2056 | 30 |
| slice_128B | Binary | 105.9 | 2417.83 | 512 | 1 |
| slice_128B | MsgPack | 75.87 | 3374.24 | 152 | 2 |
| slice_4KB | Bitbox | 610.2 | 13425.02 | 4124 | 3 |
| slice_4KB | Gob | 2692 | 3042.86 | 15464 | 30 |
| slice_4KB | Binary | 2042 | 4012.30 | 17408 | 7 |
| slice_4KB | MsgPack | 674.7 | 12140.99 | 4120 | 2 |
| slice_64KB | Bitbox | 7250 | 18078.13 | 65564 | 3 |
| slice_64KB | Gob | 26350 | 4974.27 | 214633 | 30 |
| slice_64KB | Binary | 34751 | 3771.71 | 284930 | 16 |
| slice_64KB | MsgPack | 10160 | 12900.25 | 65560 | 2 |
| bool | Bitbox | 10.70 | 186.85 | 0 | 0 |
| bool | Gob | 1004 | 1.99 | 1552 | 25 |
| bool | Binary | 32.70 | 61.17 | 2 | 2 |
| bool | MsgPack | 8.084 | 247.40 | 0 | 0 |
| string | Bitbox | 58.03 | 551.47 | 36 | 3 |
| string | Gob | 1019 | 31.41 | 1640 | 28 |
| string | Binary | 120.9 | 264.60 | 528 | 2 |
| string | MsgPack | 51.57 | 620.55 | 32 | 2 |
| int8 | Bitbox | 9.857 | 202.89 | 0 | 0 |
| int8 | Gob | 937.7 | 2.13 | 1552 | 25 |
| int8 | Binary | 31.55 | 63.39 | 2 | 2 |
| int8 | MsgPack | 26.69 | 74.95 | 0 | 0 |
| int16 | Bitbox | 16.09 | 248.57 | 2 | 1 |
| int16 | Gob | 954.8 | 4.19 | 1552 | 26 |
| int16 | Binary | 34.66 | 115.39 | 4 | 2 |
| int16 | MsgPack | 40.87 | 97.86 | 2 | 1 |
| int32 | Bitbox | 17.65 | 453.38 | 4 | 1 |
| int32 | Gob | 963.2 | 8.31 | 1560 | 26 |
| int32 | Binary | 33.32 | 240.12 | 8 | 2 |
| int32 | MsgPack | 43.44 | 184.16 | 4 | 1 |
| int64 | Bitbox | 17.93 | 892.41 | 8 | 1 |
| int64 | Gob | 944.5 | 16.94 | 1568 | 26 |
| int64 | Binary | 36.27 | 441.19 | 16 | 2 |
| int64 | MsgPack | 29.37 | 544.77 | 8 | 1 |
| uint8 | Bitbox | 10.19 | 196.27 | 0 | 0 |
| uint8 | Gob | 913.5 | 2.19 | 1552 | 25 |
| uint8 | Binary | 32.14 | 62.24 | 2 | 2 |
| uint8 | MsgPack | 28.07 | 71.25 | 0 | 0 |
| uint16 | Bitbox | 16.82 | 237.82 | 2 | 1 |
| uint16 | Gob | 1021 | 3.92 | 1552 | 26 |
| uint16 | Binary | 34.59 | 115.63 | 4 | 2 |
| uint16 | MsgPack | 45.09 | 88.71 | 2 | 1 |
| uint32 | Bitbox | 18.07 | 442.68 | 4 | 1 |
| uint32 | Gob | 940.7 | 8.50 | 1560 | 26 |
| uint32 | Binary | 33.87 | 236.21 | 8 | 2 |
| uint32 | MsgPack | 42.67 | 187.47 | 4 | 1 |
| uint64 | Bitbox | 18.75 | 853.52 | 8 | 1 |
| uint64 | Gob | 986.5 | 16.22 | 1568 | 26 |
| uint64 | Binary | 35.52 | 450.41 | 16 | 2 |
| uint64 | MsgPack | 36.80 | 434.76 | 8 | 1 |
| uintptr | Bitbox | 17.06 | 937.92 | 8 | 1 |
| uintptr | Gob | 1031 | 15.52 | 1584 | 26 |
| uintptr | Binary | 35.93 | 445.28 | 16 | 2 |
| uintptr | MsgPack | --- | --- | --- | --- |
| float32 | Bitbox | 17.69 | 452.27 | 4 | 1 |
| float32 | Gob | 1077 | 7.43 | 1576 | 26 |
| float32 | Binary | 35.25 | 226.94 | 8 | 2 |
| float32 | MsgPack | 28.16 | 284.08 | 4 | 1 |
| float64 | Bitbox | 17.78 | 900.01 | 8 | 1 |
| float64 | Gob | 1023 | 15.64 | 1624 | 27 |
| float64 | Binary | 36.62 | 436.94 | 16 | 2 |
| float64 | MsgPack | 28.89 | 553.77 | 8 | 1 |
| complex64 | Bitbox | 18.26 | 876.07 | 8 | 1 |
| complex64 | Gob | 1066 | 15.01 | 1624 | 27 |
| complex64 | Binary | 55.88 | 286.32 | 16 | 2 |
| complex64 | MsgPack | --- | --- | --- | --- |
| complex128 | Bitbox | 22.45 | 1425.65 | 16 | 1 |
| complex128 | Gob | 1014 | 31.55 | 1640 | 27 |
| complex128 | Binary | 63.42 | 504.54 | 32 | 2 |
| complex128 | MsgPack | --- | --- | --- | --- |
# Benchmark Results (Named Types, Arrays, Slices)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bool_pointer | Bitbox | 167.2 | 11.96 | 0 | 0 |
| named_bool_pointer | Gob | 1241 | 1.61 | 1616 | 26 |
| named_bool_pointer | Binary | 197.2 | 10.14 | 2 | 2 |
| named_bool_pointer | MsgPack | 195.4 | 10.23 | 0 | 0 |
| named_int8_pointer | Bitbox | 173.6 | 11.52 | 0 | 0 |
| named_int8_pointer | Gob | 1223 | 1.64 | 1616 | 26 |
| named_int8_pointer | Binary | 196.5 | 10.18 | 2 | 2 |
| named_int8_pointer | MsgPack | 197.7 | 10.12 | 0 | 0 |
| named_int16_pointer | Bitbox | 182.6 | 21.91 | 4 | 2 |
| named_int16_pointer | Gob | 1275 | 3.14 | 1621 | 28 |
| named_int16_pointer | Binary | 215.5 | 18.56 | 8 | 4 |
| named_int16_pointer | MsgPack | 219.8 | 18.20 | 4 | 2 |
| named_int32_pointer | Bitbox | 200.0 | 40.00 | 8 | 2 |
| named_int32_pointer | Gob | 1230 | 6.50 | 1632 | 28 |
| named_int32_pointer | Binary | 228.4 | 35.03 | 16 | 4 |
| named_int32_pointer | MsgPack | 223.9 | 35.74 | 8 | 2 |
| named_int64_pointer | Bitbox | 190.6 | 83.97 | 16 | 2 |
| named_int64_pointer | Gob | 1217 | 13.15 | 1632 | 28 |
| named_int64_pointer | Binary | 232.7 | 68.75 | 32 | 4 |
| named_int64_pointer | MsgPack | 226.4 | 70.68 | 16 | 2 |
| named_uint8_pointer | Bitbox | 173.5 | 11.52 | 0 | 0 |
| named_uint8_pointer | Gob | 1372 | 1.46 | 1616 | 26 |
| named_uint8_pointer | Binary | 198.7 | 10.07 | 2 | 2 |
| named_uint8_pointer | MsgPack | 195.0 | 10.26 | 0 | 0 |
| named_uint16_pointer | Bitbox | 186.0 | 21.51 | 4 | 2 |
| named_uint16_pointer | Gob | 1402 | 2.85 | 1621 | 28 |
| named_uint16_pointer | Binary | 214.3 | 18.67 | 8 | 4 |
| named_uint16_pointer | MsgPack | 220.0 | 18.18 | 4 | 2 |
| named_uint32_pointer | Bitbox | 189.4 | 42.23 | 8 | 2 |
| named_uint32_pointer | Gob | 1354 | 5.91 | 1632 | 28 |
| named_uint32_pointer | Binary | 214.5 | 37.30 | 16 | 4 |
| named_uint32_pointer | MsgPack | 230.3 | 34.74 | 8 | 2 |
| named_uint64_pointer | Bitbox | 188.8 | 84.75 | 16 | 2 |
| named_uint64_pointer | Gob | 1213 | 13.19 | 1632 | 28 |
| named_uint64_pointer | Binary | 221.9 | 72.09 | 32 | 4 |
| named_uint64_pointer | MsgPack | 225.7 | 70.88 | 16 | 2 |
| named_float32_pointer | Bitbox | 189.5 | 42.22 | 8 | 2 |
| named_float32_pointer | Gob | 1403 | 5.70 | 1664 | 29 |
| named_float32_pointer | Binary | 219.6 | 36.43 | 16 | 4 |
| named_float32_pointer | MsgPack | 224.8 | 35.58 | 8 | 2 |
| named_float64_pointer | Bitbox | 190.4 | 84.03 | 16 | 2 |
| named_float64_pointer | Gob | 1674 | 9.56 | 1680 | 29 |
| named_float64_pointer | Binary | 220.2 | 72.65 | 32 | 4 |
| named_float64_pointer | MsgPack | 225.2 | 71.05 | 16 | 2 |
| named_complex64_pointer | Bitbox | 190.8 | 83.86 | 16 | 2 |
| named_complex64_pointer | Gob | 1347 | 11.88 | 1672 | 29 |
| named_complex64_pointer | Binary | 229.4 | 69.74 | 32 | 4 |
| named_complex64_pointer | MsgPack | --- | --- | --- | --- |
| named_complex128_pointer | Bitbox | 206.7 | 154.78 | 32 | 2 |
| named_complex128_pointer | Gob | 1241 | 25.79 | 1688 | 29 |
| named_complex128_pointer | Binary | 246.0 | 130.10 | 64 | 4 |
| named_complex128_pointer | MsgPack | --- | --- | --- | --- |
| named_string_value | Bitbox | 250.3 | 95.88 | 48 | 4 |
| named_string_value | Gob | 1338 | 17.94 | 1712 | 30 |
| named_string_value | Binary | 333.3 | 72.00 | 560 | 4 |
| named_string_value | MsgPack | 260.7 | 92.04 | 48 | 3 |
| named_bytes_value | Bitbox | 340.3 | 47.01 | 76 | 4 |
| named_bytes_value | Gob | 1488 | 10.75 | 1760 | 30 |
| named_bytes_value | Binary | 373.4 | 42.85 | 560 | 3 |
| named_bytes_value | MsgPack | 321.0 | 49.84 | 48 | 2 |
| named_uint64_slice_value | Bitbox | 427.1 | 299.68 | 76 | 4 |
| named_uint64_slice_value | Gob | 9870 | 12.97 | 7752 | 189 |
| named_uint64_slice_value | Binary | 470.1 | 272.29 | 176 | 4 |
| named_uint64_slice_value | MsgPack | 871.9 | 146.81 | 72 | 3 |
| named_byte_array_pointer | Bitbox | 11037 | 181.21 | 2048 | 2 |
| named_byte_array_pointer | Gob | 34900 | 57.31 | 16536 | 194 |
| named_byte_array_pointer | Binary | 20235 | 98.84 | 4096 | 4 |
| named_byte_array_pointer | MsgPack | 14516 | 137.78 | 2096 | 4 |
| named_uint32_array_pointer | Bitbox | 256.0 | 125.02 | 32 | 2 |
| named_uint32_array_pointer | Gob | 9925 | 3.22 | 7728 | 189 |
| named_uint32_array_pointer | Binary | 342.4 | 93.46 | 64 | 4 |
| named_uint32_array_pointer | MsgPack | 511.4 | 62.57 | 32 | 2 |
# Benchmark Results (Aligned Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 50.30 | 636.12 | 0 | 0 |
| EncodeDecodeStruct | Gob | 10169 | 3.15 | 8112 | 204 |
| EncodeDecodeStruct | BinaryWriteRead | 136.3 | 234.82 | 32 | 2 |
| EncodeDecodeStruct | MsgPack | 347.7 | 92.04 | 0 | 0 |
# Benchmark Results (Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeTx | Bitbox | 1244 | 154.29 | 200 | 4 |
| EncodeDecodeTx | Gob | 16177 | 11.87 | 10160 | 262 |
| EncodeDecodeTx | Binary | --- | --- | --- | --- |
| EncodeDecodeTx | MsgPack | 1987 | 96.62 | 264 | 5 |
