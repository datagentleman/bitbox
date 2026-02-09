```
 _______  ___   _______  _______  _______  __   __ 
|  _    ||   | |       ||  _    ||       ||  |_|  |
| |_|   ||   | |_     _|| |_|   ||   _   ||       |
|       ||   |   |   |  |       ||  | |  ||       |
|  _   | |   |   |   |  |  _   | |  |_|  | |     | 
| |_|   ||   |   |   |  | |_|   ||       ||   _   |
|_______||___|   |___|  |_______||_______||__| |__|
```

# bitbox

Bitbox is a tiny, extremely fast, low-overhead binary encoding/decoding package for Go.
It provides a universal byte format where data can be easily moved across different languages and platforms,
as long as all systems use the same endianness.

# ⚠️ Endianness warning

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
| slice_128B | Bitbox | 88.00 | 2909.20 | 172 | 4 |
| slice_128B | Gob | 1048 | 244.21 | 1880 | 27 |
| slice_128B | Binary | 105.7 | 2421.27 | 512 | 1 |
| slice_4KB | Bitbox | 575.5 | 14234.08 | 4140 | 4 |
| slice_4KB | Gob | 2170 | 3775.81 | 11320 | 27 |
| slice_4KB | Binary | 1996 | 4104.71 | 17408 | 7 |
| slice_64KB | Bitbox | 6980 | 18777.15 | 65580 | 4 |
| slice_64KB | Gob | 18405 | 7121.57 | 149049 | 27 |
| slice_64KB | Binary | 30213 | 4338.24 | 284930 | 16 |
| bool | Bitbox | 13.77 | 145.23 | 1 | 1 |
| bool | Gob | 909.5 | 2.20 | 1552 | 25 |
| bool | Binary | 32.87 | 60.85 | 2 | 2 |
| string | Bitbox | 67.68 | 472.85 | 52 | 4 |
| string | Gob | 981.4 | 32.61 | 1640 | 28 |
| string | Binary | 123.7 | 258.78 | 528 | 2 |
| int | Bitbox | 14.52 | 1101.79 | 8 | 1 |
| int | Gob | 927.0 | 17.26 | 1568 | 26 |
| int | Binary | 37.95 | 421.56 | 16 | 2 |
| int8 | Bitbox | 14.01 | 142.74 | 1 | 1 |
| int8 | Gob | 955.4 | 2.09 | 1552 | 25 |
| int8 | Binary | 32.26 | 62.00 | 2 | 2 |
| int16 | Bitbox | 14.30 | 279.74 | 2 | 1 |
| int16 | Gob | 934.8 | 4.28 | 1552 | 26 |
| int16 | Binary | 34.46 | 116.08 | 4 | 2 |
| int32 | Bitbox | 15.94 | 501.86 | 4 | 1 |
| int32 | Gob | 947.4 | 8.44 | 1560 | 26 |
| int32 | Binary | 34.44 | 232.30 | 8 | 2 |
| int64 | Bitbox | 14.95 | 1070.53 | 8 | 1 |
| int64 | Gob | 969.6 | 16.50 | 1568 | 26 |
| int64 | Binary | 37.00 | 432.48 | 16 | 2 |
| uint | Bitbox | 15.09 | 1060.18 | 8 | 1 |
| uint | Gob | 949.1 | 16.86 | 1568 | 26 |
| uint | Binary | 39.02 | 410.01 | 16 | 2 |
| uint8 | Bitbox | 13.84 | 144.49 | 1 | 1 |
| uint8 | Gob | 957.3 | 2.09 | 1552 | 25 |
| uint8 | Binary | 32.80 | 60.98 | 2 | 2 |
| uint16 | Bitbox | 14.34 | 278.98 | 2 | 1 |
| uint16 | Gob | 1003 | 3.99 | 1552 | 26 |
| uint16 | Binary | 34.95 | 114.45 | 4 | 2 |
| uint32 | Bitbox | 15.58 | 513.62 | 4 | 1 |
| uint32 | Gob | 946.1 | 8.46 | 1560 | 26 |
| uint32 | Binary | 34.22 | 233.77 | 8 | 2 |
| uint64 | Bitbox | 14.52 | 1101.69 | 8 | 1 |
| uint64 | Gob | 933.4 | 17.14 | 1568 | 26 |
| uint64 | Binary | 35.99 | 444.56 | 16 | 2 |
| uintptr | Bitbox | 14.72 | 1087.30 | 8 | 1 |
| uintptr | Gob | 1080 | 14.82 | 1584 | 26 |
| uintptr | Binary | 37.57 | 425.90 | 16 | 2 |
| float32 | Bitbox | 15.62 | 512.27 | 4 | 1 |
| float32 | Gob | 941.7 | 8.50 | 1576 | 26 |
| float32 | Binary | 34.61 | 231.17 | 8 | 2 |
| float64 | Bitbox | 15.55 | 1028.98 | 8 | 1 |
| float64 | Gob | 979.0 | 16.34 | 1624 | 27 |
| float64 | Binary | 38.58 | 414.72 | 16 | 2 |
| complex64 | Bitbox | 14.67 | 1090.29 | 8 | 1 |
| complex64 | Gob | 992.4 | 16.12 | 1624 | 27 |
| complex64 | Binary | 56.77 | 281.86 | 16 | 2 |
| complex128 | Bitbox | 19.88 | 1609.49 | 16 | 1 |
| complex128 | Gob | 1005 | 31.84 | 1640 | 27 |
| complex128 | Binary | 62.95 | 508.31 | 32 | 2 |
# Benchmark Results (Aligned Struct, 4 Fields)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStruct | Bitbox | 17.89 | 1788.97 | 0 | 0 |
| EncodeDecodeStruct | Gob | 10096 | 3.17 | 8128 | 205 |
| EncodeDecodeStruct | BinaryWriteRead | 226.0 | 141.57 | 64 | 6 |
