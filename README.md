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
| slice_128B | Bitbox | 72.24 | 3543.66 | 156 | 3 |
| slice_128B | Gob | 1145 | 223.60 | 2056 | 30 |
| slice_128B | Binary | 106.9 | 2395.29 | 512 | 1 |
| slice_128B | MsgPack | 73.68 | 3474.65 | 152 | 2 |
| slice_4KB | Bitbox | 646.5 | 12670.76 | 4124 | 3 |
| slice_4KB | Gob | 2626 | 3119.20 | 15464 | 30 |
| slice_4KB | Binary | 2101 | 3898.40 | 17408 | 7 |
| slice_4KB | MsgPack | 634.9 | 12902.46 | 4120 | 2 |
| slice_64KB | Bitbox | 7258 | 18058.85 | 65564 | 3 |
| slice_64KB | Gob | 26415 | 4962.12 | 214633 | 30 |
| slice_64KB | Binary | 55310 | 2369.78 | 284929 | 16 |
| slice_64KB | MsgPack | 8572 | 15291.04 | 65560 | 2 |
| bool | Bitbox | 11.31 | 176.88 | 0 | 0 |
| bool | Gob | 1070 | 1.87 | 1552 | 25 |
| bool | Binary | 32.80 | 60.98 | 2 | 2 |
| bool | MsgPack | 8.199 | 243.93 | 0 | 0 |
| string | Bitbox | 58.20 | 549.79 | 36 | 3 |
| string | Gob | 1064 | 30.07 | 1640 | 28 |
| string | Binary | 127.3 | 251.47 | 528 | 2 |
| string | MsgPack | 52.18 | 613.20 | 32 | 2 |
| int8 | Bitbox | 10.18 | 196.37 | 0 | 0 |
| int8 | Gob | 933.5 | 2.14 | 1552 | 25 |
| int8 | Binary | 32.78 | 61.01 | 2 | 2 |
| int8 | MsgPack | 26.98 | 74.14 | 0 | 0 |
| int16 | Bitbox | 16.43 | 243.39 | 2 | 1 |
| int16 | Gob | 1000 | 4.00 | 1552 | 26 |
| int16 | Binary | 33.86 | 118.12 | 4 | 2 |
| int16 | MsgPack | 40.42 | 98.97 | 2 | 1 |
| int32 | Bitbox | 18.12 | 441.61 | 4 | 1 |
| int32 | Gob | 1012 | 7.91 | 1560 | 26 |
| int32 | Binary | 34.49 | 231.95 | 8 | 2 |
| int32 | MsgPack | 43.25 | 184.97 | 4 | 1 |
| int64 | Bitbox | 17.42 | 918.64 | 8 | 1 |
| int64 | Gob | 935.0 | 17.11 | 1568 | 26 |
| int64 | Binary | 37.21 | 430.04 | 16 | 2 |
| int64 | MsgPack | 29.90 | 535.17 | 8 | 1 |
| uint8 | Bitbox | 10.49 | 190.71 | 0 | 0 |
| uint8 | Gob | 940.3 | 2.13 | 1552 | 25 |
| uint8 | Binary | 32.27 | 61.98 | 2 | 2 |
| uint8 | MsgPack | 28.72 | 69.64 | 0 | 0 |
| uint16 | Bitbox | 16.83 | 237.63 | 2 | 1 |
| uint16 | Gob | 1000 | 4.00 | 1552 | 26 |
| uint16 | Binary | 34.33 | 116.50 | 4 | 2 |
| uint16 | MsgPack | 42.29 | 94.58 | 2 | 1 |
| uint32 | Bitbox | 18.03 | 443.59 | 4 | 1 |
| uint32 | Gob | 1022 | 7.83 | 1560 | 26 |
| uint32 | Binary | 35.89 | 222.90 | 8 | 2 |
| uint32 | MsgPack | 45.86 | 174.43 | 4 | 1 |
| uint64 | Bitbox | 23.52 | 680.19 | 8 | 1 |
| uint64 | Gob | 1243 | 12.87 | 1568 | 26 |
| uint64 | Binary | 38.41 | 416.55 | 16 | 2 |
| uint64 | MsgPack | 38.20 | 418.81 | 8 | 1 |
| uintptr | Bitbox | 21.52 | 743.44 | 8 | 1 |
| uintptr | Gob | 950.8 | 16.83 | 1584 | 26 |
| uintptr | Binary | 36.92 | 433.42 | 16 | 2 |
| uintptr | MsgPack | --- | --- | --- | --- |
| float32 | Bitbox | 19.08 | 419.33 | 4 | 1 |
| float32 | Gob | 990.2 | 8.08 | 1576 | 26 |
| float32 | Binary | 34.70 | 230.54 | 8 | 2 |
| float32 | MsgPack | 28.21 | 283.55 | 4 | 1 |
| float64 | Bitbox | 18.33 | 872.69 | 8 | 1 |
| float64 | Gob | 1093 | 14.63 | 1624 | 27 |
| float64 | Binary | 38.56 | 414.90 | 16 | 2 |
| float64 | MsgPack | 37.89 | 422.28 | 8 | 1 |
| complex64 | Bitbox | 21.33 | 750.28 | 8 | 1 |
| complex64 | Gob | 1274 | 12.56 | 1624 | 27 |
| complex64 | Binary | 75.10 | 213.04 | 16 | 2 |
| complex64 | MsgPack | --- | --- | --- | --- |
| complex128 | Bitbox | 23.18 | 1380.72 | 16 | 1 |
| complex128 | Gob | 1130 | 28.31 | 1640 | 27 |
| complex128 | Binary | 64.00 | 499.97 | 32 | 2 |
| complex128 | MsgPack | --- | --- | --- | --- |
# Benchmark Results (Named Types, Arrays, Slices)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bool_pointer | Bitbox | 166.7 | 12.00 | 0 | 0 |
| named_bool_pointer | Gob | 1233 | 1.62 | 1616 | 26 |
| named_bool_pointer | Binary | 191.2 | 10.46 | 2 | 2 |
| named_bool_pointer | MsgPack | 194.0 | 10.31 | 0 | 0 |
| named_int8_pointer | Bitbox | 184.7 | 10.83 | 0 | 0 |
| named_int8_pointer | Gob | 1160 | 1.72 | 1616 | 26 |
| named_int8_pointer | Binary | 199.4 | 10.03 | 2 | 2 |
| named_int8_pointer | MsgPack | 205.5 | 9.73 | 0 | 0 |
| named_int16_pointer | Bitbox | 426.4 | 9.38 | 4 | 2 |
| named_int16_pointer | Gob | 1238 | 3.23 | 1621 | 28 |
| named_int16_pointer | Binary | 217.1 | 18.42 | 8 | 4 |
| named_int16_pointer | MsgPack | 217.4 | 18.40 | 4 | 2 |
| named_int32_pointer | Bitbox | 219.7 | 36.42 | 8 | 2 |
| named_int32_pointer | Gob | 1200 | 6.66 | 1632 | 28 |
| named_int32_pointer | Binary | 214.0 | 37.38 | 16 | 4 |
| named_int32_pointer | MsgPack | 224.4 | 35.66 | 8 | 2 |
| named_int64_pointer | Bitbox | 190.6 | 83.96 | 16 | 2 |
| named_int64_pointer | Gob | 1203 | 13.29 | 1632 | 28 |
| named_int64_pointer | Binary | 228.4 | 70.06 | 32 | 4 |
| named_int64_pointer | MsgPack | 226.0 | 70.80 | 16 | 2 |
| named_uint8_pointer | Bitbox | 173.1 | 11.55 | 0 | 0 |
| named_uint8_pointer | Gob | 1164 | 1.72 | 1616 | 26 |
| named_uint8_pointer | Binary | 191.4 | 10.45 | 2 | 2 |
| named_uint8_pointer | MsgPack | 197.4 | 10.13 | 0 | 0 |
| named_uint16_pointer | Bitbox | 197.3 | 20.27 | 4 | 2 |
| named_uint16_pointer | Gob | 1222 | 3.27 | 1621 | 28 |
| named_uint16_pointer | Binary | 210.3 | 19.02 | 8 | 4 |
| named_uint16_pointer | MsgPack | 217.8 | 18.37 | 4 | 2 |
| named_uint32_pointer | Bitbox | 192.5 | 41.57 | 8 | 2 |
| named_uint32_pointer | Gob | 1175 | 6.81 | 1632 | 28 |
| named_uint32_pointer | Binary | 212.2 | 37.70 | 16 | 4 |
| named_uint32_pointer | MsgPack | 227.2 | 35.22 | 8 | 2 |
| named_uint64_pointer | Bitbox | 190.3 | 84.06 | 16 | 2 |
| named_uint64_pointer | Gob | 1194 | 13.40 | 1632 | 28 |
| named_uint64_pointer | Binary | 219.6 | 72.85 | 32 | 4 |
| named_uint64_pointer | MsgPack | 223.6 | 71.55 | 16 | 2 |
| named_float32_pointer | Bitbox | 184.3 | 43.41 | 8 | 2 |
| named_float32_pointer | Gob | 1203 | 6.65 | 1664 | 29 |
| named_float32_pointer | Binary | 213.3 | 37.50 | 16 | 4 |
| named_float32_pointer | MsgPack | 224.2 | 35.68 | 8 | 2 |
| named_float64_pointer | Bitbox | 192.0 | 83.32 | 16 | 2 |
| named_float64_pointer | Gob | 1233 | 12.97 | 1680 | 29 |
| named_float64_pointer | Binary | 219.1 | 73.04 | 32 | 4 |
| named_float64_pointer | MsgPack | 224.8 | 71.18 | 16 | 2 |
| named_complex64_pointer | Bitbox | 191.2 | 83.68 | 16 | 2 |
| named_complex64_pointer | Gob | 1254 | 12.76 | 1672 | 29 |
| named_complex64_pointer | Binary | 226.4 | 70.67 | 32 | 4 |
| named_complex64_pointer | MsgPack | --- | --- | --- | --- |
| named_complex128_pointer | Bitbox | 209.8 | 152.56 | 32 | 2 |
| named_complex128_pointer | Gob | 1226 | 26.10 | 1688 | 29 |
| named_complex128_pointer | Binary | 246.5 | 129.82 | 64 | 4 |
| named_complex128_pointer | MsgPack | --- | --- | --- | --- |
| named_string_value | Bitbox | 238.4 | 100.69 | 48 | 4 |
| named_string_value | Gob | 1248 | 19.23 | 1712 | 30 |
| named_string_value | Binary | 320.5 | 74.88 | 560 | 4 |
| named_string_value | MsgPack | 271.5 | 88.40 | 48 | 3 |
| named_bytes_value | Bitbox | 344.8 | 46.41 | 76 | 4 |
| named_bytes_value | Gob | 1387 | 11.54 | 1760 | 30 |
| named_bytes_value | Binary | 362.0 | 44.20 | 560 | 3 |
| named_bytes_value | MsgPack | 396.6 | 40.34 | 48 | 2 |
| named_uint64_slice_value | Bitbox | 467.1 | 274.04 | 76 | 4 |
| named_uint64_slice_value | Gob | 9280 | 13.79 | 7752 | 189 |
| named_uint64_slice_value | Binary | 466.6 | 274.32 | 176 | 4 |
| named_uint64_slice_value | MsgPack | 1424 | 89.89 | 72 | 3 |
| named_byte_array_pointer | Bitbox | 11317 | 176.73 | 2048 | 2 |
| named_byte_array_pointer | Gob | 33149 | 60.33 | 16537 | 194 |
| named_byte_array_pointer | Binary | 20185 | 99.08 | 4096 | 4 |
| named_byte_array_pointer | MsgPack | 16951 | 117.99 | 2096 | 4 |
| named_uint32_array_pointer | Bitbox | 263.0 | 121.68 | 32 | 2 |
| named_uint32_array_pointer | Gob | 9271 | 3.45 | 7728 | 189 |
| named_uint32_array_pointer | Binary | 345.5 | 92.63 | 64 | 4 |
| named_uint32_array_pointer | MsgPack | 491.9 | 65.06 | 32 | 2 |
# Benchmark Results (Aligned Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 52.46 | 609.97 | 0 | 0 |
| EncodeDecodeStruct | Gob | 10740 | 2.98 | 8112 | 204 |
| EncodeDecodeStruct | BinaryWriteRead | 137.7 | 232.40 | 32 | 2 |
| EncodeDecodeStruct | MsgPack | 357.0 | 89.63 | 0 | 0 |
# Benchmark Results (Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeTx | Bitbox | 1264 | 151.85 | 200 | 4 |
| EncodeDecodeTx | Gob | 16561 | 11.59 | 10160 | 262 |
| EncodeDecodeTx | Binary | --- | --- | --- | --- |
| EncodeDecodeTx | MsgPack | 1994 | 96.27 | 264 | 5 |
