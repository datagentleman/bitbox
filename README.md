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
| slice_128B | Bitbox | 74.61 | 3431.06 | 156 | 3 |
| slice_128B | Gob | 1147 | 223.10 | 2056 | 30 |
| slice_128B | Binary | 106.6 | 2402.24 | 512 | 1 |
| slice_128B | MsgPack | 73.77 | 3470.11 | 152 | 2 |
| slice_4KB | Bitbox | 610.7 | 13414.14 | 4124 | 3 |
| slice_4KB | Gob | 2748 | 2981.55 | 15464 | 30 |
| slice_4KB | Binary | 2076 | 3946.64 | 17408 | 7 |
| slice_4KB | MsgPack | 885.2 | 9254.23 | 4120 | 2 |
| slice_64KB | Bitbox | 8476 | 15464.18 | 65564 | 3 |
| slice_64KB | Gob | 25729 | 5094.32 | 214633 | 30 |
| slice_64KB | Binary | 28187 | 4650.02 | 284930 | 16 |
| slice_64KB | MsgPack | 7302 | 17949.33 | 65560 | 2 |
| bool | Bitbox | 10.53 | 189.92 | 0 | 0 |
| bool | Gob | 929.1 | 2.15 | 1552 | 25 |
| bool | Binary | 32.89 | 60.82 | 2 | 2 |
| bool | MsgPack | 8.172 | 244.73 | 0 | 0 |
| string | Bitbox | 56.36 | 567.75 | 36 | 3 |
| string | Gob | 1012 | 31.62 | 1640 | 28 |
| string | Binary | 120.0 | 266.60 | 528 | 2 |
| string | MsgPack | 52.14 | 613.78 | 32 | 2 |
| int8 | Bitbox | 9.974 | 200.52 | 0 | 0 |
| int8 | Gob | 959.3 | 2.08 | 1552 | 25 |
| int8 | Binary | 32.25 | 62.02 | 2 | 2 |
| int8 | MsgPack | 27.02 | 74.01 | 0 | 0 |
| int16 | Bitbox | 16.34 | 244.80 | 2 | 1 |
| int16 | Gob | 931.2 | 4.30 | 1552 | 26 |
| int16 | Binary | 33.65 | 118.89 | 4 | 2 |
| int16 | MsgPack | 41.73 | 95.86 | 2 | 1 |
| int32 | Bitbox | 17.66 | 452.97 | 4 | 1 |
| int32 | Gob | 937.4 | 8.53 | 1560 | 26 |
| int32 | Binary | 33.39 | 239.57 | 8 | 2 |
| int32 | MsgPack | 43.05 | 185.85 | 4 | 1 |
| int64 | Bitbox | 17.35 | 922.02 | 8 | 1 |
| int64 | Gob | 938.3 | 17.05 | 1568 | 26 |
| int64 | Binary | 36.84 | 434.30 | 16 | 2 |
| int64 | MsgPack | 29.46 | 543.17 | 8 | 1 |
| uint8 | Bitbox | 10.22 | 195.75 | 0 | 0 |
| uint8 | Gob | 927.1 | 2.16 | 1552 | 25 |
| uint8 | Binary | 32.87 | 60.84 | 2 | 2 |
| uint8 | MsgPack | 28.21 | 70.89 | 0 | 0 |
| uint16 | Bitbox | 16.66 | 240.06 | 2 | 1 |
| uint16 | Gob | 940.0 | 4.26 | 1552 | 26 |
| uint16 | Binary | 34.34 | 116.49 | 4 | 2 |
| uint16 | MsgPack | 41.02 | 97.51 | 2 | 1 |
| uint32 | Bitbox | 17.78 | 449.94 | 4 | 1 |
| uint32 | Gob | 935.6 | 8.55 | 1560 | 26 |
| uint32 | Binary | 33.84 | 236.42 | 8 | 2 |
| uint32 | MsgPack | 44.15 | 181.19 | 4 | 1 |
| uint64 | Bitbox | 21.21 | 754.37 | 8 | 1 |
| uint64 | Gob | 1009 | 15.85 | 1568 | 26 |
| uint64 | Binary | 35.19 | 454.62 | 16 | 2 |
| uint64 | MsgPack | 28.96 | 552.49 | 8 | 1 |
| uintptr | Bitbox | 16.75 | 955.21 | 8 | 1 |
| uintptr | Gob | 942.4 | 16.98 | 1584 | 26 |
| uintptr | Binary | 36.84 | 434.36 | 16 | 2 |
| uintptr | MsgPack | --- | --- | --- | --- |
| float32 | Bitbox | 17.24 | 464.15 | 4 | 1 |
| float32 | Gob | 937.3 | 8.54 | 1576 | 26 |
| float32 | Binary | 35.06 | 228.17 | 8 | 2 |
| float32 | MsgPack | 27.72 | 288.64 | 4 | 1 |
| float64 | Bitbox | 17.25 | 927.58 | 8 | 1 |
| float64 | Gob | 973.0 | 16.44 | 1624 | 27 |
| float64 | Binary | 36.00 | 444.47 | 16 | 2 |
| float64 | MsgPack | 28.51 | 561.26 | 8 | 1 |
| complex64 | Bitbox | 17.72 | 902.72 | 8 | 1 |
| complex64 | Gob | 994.0 | 16.10 | 1624 | 27 |
| complex64 | Binary | 55.72 | 287.17 | 16 | 2 |
| complex64 | MsgPack | --- | --- | --- | --- |
| complex128 | Bitbox | 21.94 | 1458.34 | 16 | 1 |
| complex128 | Gob | 1004 | 31.88 | 1640 | 27 |
| complex128 | Binary | 62.16 | 514.84 | 32 | 2 |
| complex128 | MsgPack | --- | --- | --- | --- |

# Benchmark Results (Named Types, Arrays, Slices)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| named_bool_pointer | Bitbox | 351.1 | 5.70 | 144 | 4 |
| named_bool_pointer | Gob | 1188 | 1.68 | 1616 | 26 |
| named_bool_pointer | Binary | 196.7 | 10.17 | 2 | 2 |
| named_bool_pointer | MsgPack | 201.6 | 9.92 | 0 | 0 |
| named_int8_pointer | Bitbox | 346.5 | 5.77 | 144 | 4 |
| named_int8_pointer | Gob | 1332 | 1.50 | 1616 | 26 |
| named_int8_pointer | Binary | 199.0 | 10.05 | 2 | 2 |
| named_int8_pointer | MsgPack | 203.1 | 9.85 | 0 | 0 |
| named_int16_pointer | Bitbox | 365.7 | 10.94 | 148 | 6 |
| named_int16_pointer | Gob | 1176 | 3.40 | 1621 | 28 |
| named_int16_pointer | Binary | 213.8 | 18.71 | 8 | 4 |
| named_int16_pointer | MsgPack | 220.9 | 18.11 | 4 | 2 |
| named_int32_pointer | Bitbox | 369.1 | 21.68 | 152 | 6 |
| named_int32_pointer | Gob | 1158 | 6.91 | 1632 | 28 |
| named_int32_pointer | Binary | 215.9 | 37.05 | 16 | 4 |
| named_int32_pointer | MsgPack | 231.3 | 34.58 | 8 | 2 |
| named_int64_pointer | Bitbox | 367.6 | 43.53 | 160 | 6 |
| named_int64_pointer | Gob | 1156 | 13.84 | 1632 | 28 |
| named_int64_pointer | Binary | 223.1 | 71.70 | 32 | 4 |
| named_int64_pointer | MsgPack | 234.5 | 68.23 | 16 | 2 |
| named_uint8_pointer | Bitbox | 352.6 | 5.67 | 144 | 4 |
| named_uint8_pointer | Gob | 1148 | 1.74 | 1616 | 26 |
| named_uint8_pointer | Binary | 200.0 | 10.00 | 2 | 2 |
| named_uint8_pointer | MsgPack | 203.9 | 9.81 | 0 | 0 |
| named_uint16_pointer | Bitbox | 368.0 | 10.87 | 148 | 6 |
| named_uint16_pointer | Gob | 1169 | 3.42 | 1621 | 28 |
| named_uint16_pointer | Binary | 222.1 | 18.01 | 8 | 4 |
| named_uint16_pointer | MsgPack | 232.9 | 17.17 | 4 | 2 |
| named_uint32_pointer | Bitbox | 365.7 | 21.88 | 152 | 6 |
| named_uint32_pointer | Gob | 1172 | 6.83 | 1632 | 28 |
| named_uint32_pointer | Binary | 244.7 | 32.69 | 16 | 4 |
| named_uint32_pointer | MsgPack | 239.0 | 33.48 | 8 | 2 |
| named_uint64_pointer | Bitbox | 378.6 | 42.26 | 160 | 6 |
| named_uint64_pointer | Gob | 1189 | 13.46 | 1632 | 28 |
| named_uint64_pointer | Binary | 318.4 | 50.24 | 32 | 4 |
| named_uint64_pointer | MsgPack | 232.6 | 68.77 | 16 | 2 |
| named_float32_pointer | Bitbox | 372.2 | 21.50 | 152 | 6 |
| named_float32_pointer | Gob | 1189 | 6.73 | 1664 | 29 |
| named_float32_pointer | Binary | 255.5 | 31.32 | 16 | 4 |
| named_float32_pointer | MsgPack | 224.3 | 35.67 | 8 | 2 |
| named_float64_pointer | Bitbox | 369.3 | 43.33 | 160 | 6 |
| named_float64_pointer | Gob | 1205 | 13.28 | 1680 | 29 |
| named_float64_pointer | Binary | 229.2 | 69.80 | 32 | 4 |
| named_float64_pointer | MsgPack | 223.6 | 71.57 | 16 | 2 |
| named_complex64_pointer | Bitbox | 369.2 | 43.34 | 160 | 6 |
| named_complex64_pointer | Gob | 1213 | 13.20 | 1672 | 29 |
| named_complex64_pointer | Binary | 231.2 | 69.22 | 32 | 4 |
| named_complex64_pointer | MsgPack | --- | --- | --- | --- |
| named_complex128_pointer | Bitbox | 380.1 | 84.20 | 176 | 6 |
| named_complex128_pointer | Gob | 1244 | 25.73 | 1688 | 29 |
| named_complex128_pointer | Binary | 249.0 | 128.53 | 64 | 4 |
| named_complex128_pointer | MsgPack | --- | --- | --- | --- |
| named_string_value | Bitbox | 385.2 | 62.31 | 144 | 7 |
| named_string_value | Gob | 1227 | 19.56 | 1712 | 30 |
| named_string_value | Binary | 337.5 | 71.11 | 560 | 4 |
| named_string_value | MsgPack | 261.6 | 91.76 | 48 | 3 |
| named_bytes_value | Bitbox | 530.6 | 30.16 | 220 | 8 |
| named_bytes_value | Gob | 1368 | 11.70 | 1760 | 30 |
| named_bytes_value | Binary | 387.4 | 41.30 | 560 | 3 |
| named_bytes_value | MsgPack | 312.3 | 51.24 | 48 | 2 |
| named_uint64_slice_value | Bitbox | 614.6 | 208.26 | 220 | 8 |
| named_uint64_slice_value | Gob | 9139 | 14.01 | 7752 | 189 |
| named_uint64_slice_value | Binary | 486.5 | 263.13 | 176 | 4 |
| named_uint64_slice_value | MsgPack | 868.1 | 147.45 | 72 | 3 |
| named_byte_array_pointer | Bitbox | 10985 | 182.06 | 2193 | 6 |
| named_byte_array_pointer | Gob | 33280 | 60.10 | 16537 | 194 |
| named_byte_array_pointer | Binary | 20451 | 97.79 | 4096 | 4 |
| named_byte_array_pointer | MsgPack | 11012 | 181.62 | 2096 | 4 |
| named_uint32_array_pointer | Bitbox | 446.4 | 71.68 | 176 | 6 |
| named_uint32_array_pointer | Gob | 8990 | 3.56 | 7728 | 189 |
| named_uint32_array_pointer | Binary | 351.6 | 91.02 | 64 | 4 |
| named_uint32_array_pointer | MsgPack | 494.2 | 64.75 | 32 | 2 |

# Benchmark Results (Nested Slices and Arrays)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| array_100x32_byte | Bitbox | 7378 | 867.46 | 6524 | 105 |
| array_100x32_byte | Gob | 74611 | 85.78 | 40736 | 302 |
| array_100x32_byte | Binary | 33048 | 193.65 | 6400 | 2 |
| array_100x32_byte | MsgPack | 10040 | 637.44 | 5600 | 101 |
| slice_2d_bytes | Bitbox | 5922 | 2161.50 | 9642 | 207 |
| slice_2d_bytes | Gob | 26273 | 487.19 | 53449 | 496 |
| slice_2d_bytes | Binary | --- | --- | --- | --- |
| slice_2d_bytes | MsgPack | 8891 | 1439.64 | 11872 | 106 |
| slice_2d_uint64 | Bitbox | 13158 | 7782.59 | 54463 | 207 |
| slice_2d_uint64 | Gob | 112891 | 907.07 | 175393 | 609 |
| slice_2d_uint64 | Binary | --- | --- | --- | --- |
| slice_2d_uint64 | MsgPack | 370308 | 276.53 | 115072 | 506 |
| slice_2d_tx | Bitbox | 198995 | 385.94 | 94692 | 4128 |
| slice_2d_tx | Gob | 518204 | 148.20 | 326207 | 5620 |
| slice_2d_tx | Binary | --- | --- | --- | --- |
| slice_2d_tx | MsgPack | 515220 | 149.06 | 166944 | 5006 |
| array_100x100_tx | Bitbox | 5125074 | 374.63 | 3207476 | 110006 |
| array_100x100_tx | Gob | 11391768 | 168.54 | 8783964 | 120412 |
| array_100x100_tx | Binary | --- | --- | --- | --- |
| array_100x100_tx | MsgPack | 12180628 | 157.63 | 2966657 | 120001 |

# Benchmark Results (Aligned Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 51.94 | 616.05 | 0 | 0 |
| EncodeDecodeStruct | Gob | 10473 | 3.06 | 8112 | 204 |
| EncodeDecodeStruct | BinaryWriteRead | 135.2 | 236.69 | 32 | 2 |
| EncodeDecodeStruct | MsgPack | 353.4 | 90.56 | 0 | 0 |

# Benchmark Results (Struct)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeTx | Bitbox | 1467 | 130.87 | 320 | 8 |
| EncodeDecodeTx | Gob | 16374 | 11.73 | 10160 | 262 |
| EncodeDecodeTx | Binary | --- | --- | --- | --- |
| EncodeDecodeTx | MsgPack | 2094 | 91.70 | 264 | 5 |
