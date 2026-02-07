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

# Benchmark Results (Basic Types, Slices, Arrays)

| Type | Codec | ns/op | MB/s | B/op | allocs/op |
|-----:|:------|------:|-----:|-----:|----------:|
| slice_128B | Bitbox | 148.44 | 1724.61 | 232 | 6 |
| slice_128B | Gob | 1176.40 | 217.65 | 2144 | 31 |
| slice_128B | Binary | 194.04 | 1319.34 | 760 | 5 |
| slice_4KB | Bitbox | 646.08 | 12681.92 | 4952 | 6 |
| slice_4KB | Gob | 2735.00 | 2995.28 | 16304 | 31 |
| slice_4KB | Binary | 2586.00 | 3168.08 | 21624 | 11 |
| slice_64KB | Bitbox | 9450.20 | 14260.74 | 73816 | 6 |
| slice_64KB | Gob | 28145.20 | 4684.82 | 222897 | 31 |
| slice_64KB | Binary | 35835.60 | 3662.31 | 350586 | 20 |
| array_4xuint32 | Bitbox | 72.40 | 442.02 | 64 | 3 |
| array_4xuint32 | Gob | 8986.40 | 3.56 | 7656 | 187 |
| array_4xuint32 | Binary | 196.96 | 162.49 | 208 | 6 |
| bool | Bitbox | 54.55 | 36.67 | 40 | 2 |
| bool | Gob | 1016.60 | 1.97 | 1712 | 28 |
| bool | Binary | 94.70 | 21.12 | 162 | 5 |
| string | Bitbox | 136.76 | 233.99 | 120 | 7 |
| string | Gob | 1517.60 | 22.26 | 1816 | 32 |
| string | Binary | 251.28 | 127.58 | 704 | 6 |
| int | Bitbox | 63.25 | 253.17 | 48 | 3 |
| int | Gob | 1199.00 | 13.35 | 1728 | 30 |
| int | Binary | 121.48 | 131.75 | 184 | 6 |
| int8 | Bitbox | 74.13 | 28.01 | 40 | 2 |
| int8 | Gob | 1202.00 | 1.68 | 1712 | 28 |
| int8 | Binary | 104.82 | 19.15 | 162 | 5 |
| int16 | Bitbox | 70.40 | 57.45 | 48 | 3 |
| int16 | Gob | 1999.60 | 2.86 | 1717 | 30 |
| int16 | Binary | 115.12 | 34.87 | 166 | 6 |
| int32 | Bitbox | 62.39 | 128.28 | 48 | 3 |
| int32 | Gob | 1202.00 | 6.67 | 1728 | 30 |
| int32 | Binary | 152.60 | 55.74 | 172 | 6 |
| int64 | Bitbox | 68.30 | 236.23 | 48 | 3 |
| int64 | Gob | 1171.40 | 13.67 | 1728 | 30 |
| int64 | Binary | 124.22 | 129.48 | 184 | 6 |
| uint | Bitbox | 62.48 | 256.15 | 48 | 3 |
| uint | Gob | 1115.40 | 14.35 | 1728 | 30 |
| uint | Binary | 113.82 | 140.59 | 184 | 6 |
| uint8 | Bitbox | 57.58 | 34.79 | 40 | 2 |
| uint8 | Gob | 1099.60 | 1.82 | 1712 | 28 |
| uint8 | Binary | 97.23 | 20.58 | 162 | 5 |
| uint16 | Bitbox | 64.29 | 62.31 | 48 | 3 |
| uint16 | Gob | 1123.20 | 3.56 | 1717 | 30 |
| uint16 | Binary | 108.24 | 36.96 | 166 | 6 |
| uint32 | Bitbox | 60.59 | 132.08 | 48 | 3 |
| uint32 | Gob | 1091.60 | 7.33 | 1728 | 30 |
| uint32 | Binary | 117.82 | 68.00 | 172 | 6 |
| uint64 | Bitbox | 65.44 | 244.99 | 48 | 3 |
| uint64 | Gob | 1176.20 | 13.62 | 1728 | 30 |
| uint64 | Binary | 118.32 | 135.38 | 184 | 6 |
| uintptr | Bitbox | 62.29 | 256.87 | 48 | 3 |
| uintptr | Gob | 1155.00 | 13.86 | 1744 | 30 |
| uintptr | Binary | 115.66 | 138.40 | 184 | 6 |
| float32 | Bitbox | 62.97 | 127.06 | 48 | 3 |
| float32 | Gob | 1161.80 | 6.89 | 1744 | 30 |
| float32 | Binary | 113.30 | 70.64 | 172 | 6 |
| float64 | Bitbox | 62.49 | 256.16 | 48 | 3 |
| float64 | Gob | 1398.00 | 12.02 | 1792 | 31 |
| float64 | Binary | 113.34 | 141.19 | 184 | 6 |
| complex64 | Bitbox | 61.51 | 260.15 | 48 | 3 |
| complex64 | Gob | 1147.20 | 13.95 | 1792 | 31 |
| complex64 | Binary | 137.62 | 116.25 | 184 | 6 |
| complex128 | Bitbox | 67.95 | 471.04 | 64 | 3 |
| complex128 | Gob | 1137.80 | 28.13 | 1816 | 31 |
| complex128 | Binary | 148.42 | 215.64 | 208 | 6 |

# Benchmark Results (Aligned Struct, 4 Fields)

| Benchmark | Codec | ns/op | MB/s | B/op | allocs/op |
|:----------|:------|------:|-----:|-----:|----------:|
| EncodeDecodeStructAligned4Fields | Bitbox | 55.57 | 575.85 | 48 | 2 |
| EncodeDecodeStructAligned4Fields | Gob | 10646.40 | 3.01 | 8416 | 209 |
| EncodeDecodeStructAligned4Fields | BinaryWriteRead | 300.22 | 106.60 | 224 | 9 |
