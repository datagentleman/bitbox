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

# ⚠️ Endianness warning

Bitbox writes/read all data using the machine’s native endianness (typically little-endian).
To ensure correct decoding, both encoder and decoder must run on platforms with the same endianness.

# Things to do

* add support for Maps
* add support for nested slices
* add support for nested arrays
* add generic functions
* examples
* clean code
* clean tests
* clean benchmarks

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
| slice_128B | Bitbox | 104.5 | 2449.80 | 172 | 4 |
| slice_128B | Gob | 1313 | 195.03 | 2056 | 30 |
| slice_128B | Binary | 121.7 | 2104.12 | 512 | 1 |
| slice_128B | MsgPack | 255.2 | 1003.05 | 456 | 6 |
| slice_4KB | Bitbox | 780.4 | 10496.88 | 4140 | 4 |
| slice_4KB | Gob | 3611 | 2268.66 | 15464 | 30 |
| slice_4KB | Binary | 2454 | 3338.09 | 17408 | 7 |
| slice_4KB | MsgPack | 1386 | 5909.25 | 9155 | 6 |
| slice_64KB | Bitbox | 7835 | 16729.73 | 65580 | 4 |
| slice_64KB | Gob | 25215 | 5198.18 | 214633 | 30 |
| slice_64KB | Binary | 31866 | 4113.26 | 284930 | 16 |
| slice_64KB | MsgPack | 14464 | 9061.64 | 139571 | 6 |
| bool | Bitbox | 11.71 | 170.87 | 0 | 0 |
| bool | Gob | 1076 | 1.86 | 1552 | 25 |
| bool | Binary | 34.23 | 58.43 | 2 | 2 |
| bool | MsgPack | 101.4 | 19.72 | 160 | 3 |
| string | Bitbox | 83.31 | 384.12 | 52 | 4 |
| string | Gob | 1254 | 25.52 | 1640 | 28 |
| string | Binary | 146.7 | 218.13 | 528 | 2 |
| string | MsgPack | 147.4 | 217.07 | 192 | 5 |
| int8 | Bitbox | 9.648 | 207.30 | 0 | 0 |
| int8 | Gob | 1046 | 1.91 | 1552 | 25 |
| int8 | Binary | 34.63 | 57.75 | 2 | 2 |
| int8 | MsgPack | 129.7 | 15.42 | 160 | 3 |
| int16 | Bitbox | 16.58 | 241.25 | 2 | 1 |
| int16 | Gob | 1030 | 3.88 | 1552 | 26 |
| int16 | Binary | 34.61 | 115.56 | 4 | 2 |
| int16 | MsgPack | 144.7 | 27.64 | 162 | 4 |
| int32 | Bitbox | 18.55 | 431.35 | 4 | 1 |
| int32 | Gob | 1055 | 7.58 | 1560 | 26 |
| int32 | Binary | 34.49 | 231.98 | 8 | 2 |
| int32 | MsgPack | 141.2 | 56.64 | 164 | 4 |
| int64 | Bitbox | 17.03 | 939.66 | 8 | 1 |
| int64 | Gob | 1091 | 14.66 | 1568 | 26 |
| int64 | Binary | 39.15 | 408.69 | 16 | 2 |
| int64 | MsgPack | 185.3 | 86.34 | 168 | 4 |
| uint8 | Bitbox | 13.76 | 145.37 | 0 | 0 |
| uint8 | Gob | 1108 | 1.81 | 1552 | 25 |
| uint8 | Binary | 45.49 | 43.97 | 2 | 2 |
| uint8 | MsgPack | 124.6 | 16.05 | 160 | 3 |
| uint16 | Bitbox | 17.72 | 225.78 | 2 | 1 |
| uint16 | Gob | 1100 | 3.64 | 1552 | 26 |
| uint16 | Binary | 40.06 | 99.84 | 4 | 2 |
| uint16 | MsgPack | 156.2 | 25.60 | 162 | 4 |
| uint32 | Bitbox | 19.61 | 408.01 | 4 | 1 |
| uint32 | Gob | 1372 | 5.83 | 1560 | 26 |
| uint32 | Binary | 36.58 | 218.70 | 8 | 2 |
| uint32 | MsgPack | 152.1 | 52.59 | 164 | 4 |
| uint64 | Bitbox | 19.29 | 829.35 | 8 | 1 |
| uint64 | Gob | 1412 | 11.33 | 1568 | 26 |
| uint64 | Binary | 50.73 | 315.40 | 16 | 2 |
| uint64 | MsgPack | 137.0 | 116.75 | 168 | 4 |
| uintptr | Bitbox | 17.06 | 937.83 | 8 | 1 |
| uintptr | Gob | 1079 | 14.83 | 1584 | 26 |
| uintptr | Binary | 37.79 | 423.41 | 16 | 2 |
| uintptr | MsgPack | --- | --- | --- | --- |
| float32 | Bitbox | 17.69 | 452.32 | 4 | 1 |
| float32 | Gob | 1052 | 7.60 | 1576 | 26 |
| float32 | Binary | 35.25 | 226.96 | 8 | 2 |
| float32 | MsgPack | 119.4 | 66.98 | 164 | 4 |
| float64 | Bitbox | 16.99 | 941.69 | 8 | 1 |
| float64 | Gob | 1053 | 15.19 | 1624 | 27 |
| float64 | Binary | 37.72 | 424.15 | 16 | 2 |
| float64 | MsgPack | 120.0 | 133.35 | 168 | 4 |
| complex64 | Bitbox | 17.52 | 913.43 | 8 | 1 |
| complex64 | Gob | 1039 | 15.40 | 1624 | 27 |
| complex64 | Binary | 56.99 | 280.73 | 16 | 2 |
| complex64 | MsgPack | --- | --- | --- | --- |
| complex128 | Bitbox | 22.33 | 1433.24 | 16 | 1 |
| complex128 | Gob | 1065 | 30.04 | 1640 | 27 |
| complex128 | Binary | 63.42 | 504.54 | 32 | 2 |
| complex128 | MsgPack | --- | --- | --- | --- |
# Benchmark Results (Named Types, Arrays, Slices)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bool_pointer | Bitbox | 37.57 | 53.23 | 0 | 0 |
| named_bool_pointer | Gob | 1028 | 1.94 | 1616 | 26 |
| named_bool_pointer | Binary | 44.90 | 44.54 | 2 | 2 |
| named_bool_pointer | MsgPack | 140.0 | 14.28 | 160 | 3 |
| named_int8_pointer | Bitbox | 37.75 | 52.98 | 0 | 0 |
| named_int8_pointer | Gob | 1061 | 1.88 | 1616 | 26 |
| named_int8_pointer | Binary | 48.52 | 41.22 | 2 | 2 |
| named_int8_pointer | MsgPack | 140.5 | 14.24 | 160 | 3 |
| named_int16_pointer | Bitbox | 36.96 | 108.22 | 0 | 0 |
| named_int16_pointer | Gob | 1027 | 3.89 | 1616 | 26 |
| named_int16_pointer | Binary | 53.53 | 74.73 | 4 | 2 |
| named_int16_pointer | MsgPack | 158.8 | 25.20 | 160 | 3 |
| named_int32_pointer | Bitbox | 36.91 | 216.75 | 0 | 0 |
| named_int32_pointer | Gob | 1078 | 7.42 | 1616 | 26 |
| named_int32_pointer | Binary | 57.25 | 139.74 | 8 | 2 |
| named_int32_pointer | MsgPack | 149.9 | 53.37 | 160 | 3 |
| named_int64_pointer | Bitbox | 36.42 | 439.27 | 0 | 0 |
| named_int64_pointer | Gob | 1105 | 14.48 | 1616 | 26 |
| named_int64_pointer | Binary | 58.94 | 271.46 | 16 | 2 |
| named_int64_pointer | MsgPack | 155.6 | 102.81 | 160 | 3 |
| named_uint8_pointer | Bitbox | 38.28 | 52.25 | 0 | 0 |
| named_uint8_pointer | Gob | 1076 | 1.86 | 1616 | 26 |
| named_uint8_pointer | Binary | 47.88 | 41.77 | 2 | 2 |
| named_uint8_pointer | MsgPack | 150.3 | 13.30 | 160 | 3 |
| named_uint16_pointer | Bitbox | 37.07 | 107.89 | 0 | 0 |
| named_uint16_pointer | Gob | 1119 | 3.58 | 1616 | 26 |
| named_uint16_pointer | Binary | 49.96 | 80.06 | 4 | 2 |
| named_uint16_pointer | MsgPack | 160.1 | 24.99 | 160 | 3 |
| named_uint32_pointer | Bitbox | 36.93 | 216.62 | 0 | 0 |
| named_uint32_pointer | Gob | 1049 | 7.63 | 1616 | 26 |
| named_uint32_pointer | Binary | 54.14 | 147.76 | 8 | 2 |
| named_uint32_pointer | MsgPack | 154.8 | 51.68 | 160 | 3 |
| named_uint64_pointer | Bitbox | 35.95 | 445.08 | 0 | 0 |
| named_uint64_pointer | Gob | 1023 | 15.65 | 1616 | 26 |
| named_uint64_pointer | Binary | 54.12 | 295.66 | 16 | 2 |
| named_uint64_pointer | MsgPack | 155.1 | 103.18 | 160 | 3 |
| named_float32_pointer | Bitbox | 37.83 | 211.45 | 0 | 0 |
| named_float32_pointer | Gob | 1065 | 7.51 | 1656 | 27 |
| named_float32_pointer | Binary | 51.81 | 154.42 | 8 | 2 |
| named_float32_pointer | MsgPack | 151.5 | 52.80 | 160 | 3 |
| named_float64_pointer | Bitbox | 35.82 | 446.72 | 0 | 0 |
| named_float64_pointer | Gob | 1132 | 14.14 | 1664 | 27 |
| named_float64_pointer | Binary | 53.16 | 300.96 | 16 | 2 |
| named_float64_pointer | MsgPack | 157.3 | 101.74 | 160 | 3 |
| named_complex64_pointer | Bitbox | 36.00 | 444.44 | 0 | 0 |
| named_complex64_pointer | Gob | 1204 | 13.29 | 1656 | 27 |
| named_complex64_pointer | Binary | 54.98 | 291.01 | 16 | 2 |
| named_complex64_pointer | MsgPack | --- | --- | --- | --- |
| named_complex128_pointer | Bitbox | 36.42 | 878.52 | 0 | 0 |
| named_complex128_pointer | Gob | 1300 | 24.61 | 1656 | 27 |
| named_complex128_pointer | Binary | 64.02 | 499.88 | 32 | 2 |
| named_complex128_pointer | MsgPack | --- | --- | --- | --- |
| named_string_value | Bitbox | 59.88 | 400.82 | 16 | 2 |
| named_string_value | Gob | 1391 | 17.26 | 1680 | 28 |
| named_string_value | Binary | 138.7 | 173.02 | 528 | 2 |
| named_string_value | MsgPack | 166.6 | 144.09 | 176 | 4 |
| named_bytes_value | Bitbox | 46.07 | 347.27 | 4 | 1 |
| named_bytes_value | Gob | 1235 | 12.95 | 1712 | 28 |
| named_bytes_value | Binary | 135.2 | 118.35 | 512 | 1 |
| named_bytes_value | MsgPack | 166.1 | 96.31 | 160 | 3 |
| named_uint64_slice_value | Bitbox | 47.03 | 2721.78 | 4 | 1 |
| named_uint64_slice_value | Gob | 9809 | 13.05 | 7704 | 187 |
| named_uint64_slice_value | Binary | 137.9 | 928.21 | 128 | 2 |
| named_uint64_slice_value | MsgPack | 753.1 | 169.96 | 312 | 5 |
| named_byte_array_pointer | Bitbox | 48.88 | 40919.67 | 0 | 0 |
| named_byte_array_pointer | Gob | 28349 | 70.55 | 14488 | 192 |
| named_byte_array_pointer | Binary | 9604 | 208.24 | 2048 | 2 |
| named_byte_array_pointer | MsgPack | 514.8 | 3885.27 | 1233 | 6 |
| named_uint32_array_pointer | Bitbox | 26.22 | 1220.64 | 0 | 0 |
| named_uint32_array_pointer | Gob | 9885 | 3.24 | 7696 | 187 |
| named_uint32_array_pointer | Binary | 121.8 | 262.69 | 32 | 2 |
| named_uint32_array_pointer | MsgPack | 388.4 | 82.40 | 160 | 3 |
# Benchmark Results (Aligned Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 23.44 | 1365.15 | 0 | 0 |
| EncodeDecodeStruct | Gob | 11382 | 2.81 | 8112 | 204 |
| EncodeDecodeStruct | BinaryWriteRead | 139.7 | 229.04 | 32 | 2 |
| EncodeDecodeStruct | MsgPack | 490.9 | 65.19 | 160 | 3 |
# Benchmark Results (Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeTx | Bitbox | 22.19 | 8650.75 | 0 | 0 |
| EncodeDecodeTx | Gob | 17358 | 11.06 | 9968 | 260 |
| EncodeDecodeTx | Binary | --- | --- | --- | --- |
| EncodeDecodeTx | MsgPack | 1235 | 155.50 | 616 | 8 |
