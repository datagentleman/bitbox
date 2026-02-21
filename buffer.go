package bitbox

// Buffer class for encoding/decoding data.
type Buffer struct {
	data []byte
	off  int
}

// Create new Buffer.
func NewBuffer(data []byte) *Buffer {
	return &Buffer{data: data, off: 0}
}

// Decode data from buffer into objects.
func (b *Buffer) Decode(objects ...any) error {
	return Decode(b, objects...)
}

// Return remaining buffer length.
func (b *Buffer) Len() int {
	return len(b.data[b.off:])
}

// Read data from buffer into dst.
func (b *Buffer) Read(dst []byte) int {
	n := copy(dst, b.data[b.off:])
	b.off += n

	return n
}

// Write data from src into buffer.
func (b *Buffer) Write(src []byte) {
	b.data = append(b.data, src...)
}

// Take next N bytes from buffer.
// This will advance offset.
func (b *Buffer) Next(num int) ([]byte, error) {
	if b.off+num > len(b.data) {
		return nil, outOfBounds(b.off+num, len(b.data))
	}

	off := b.off
	b.off += num

	return b.data[off:b.off], nil
}

// Return remaining bytes from buffer.
func (b *Buffer) Data() []byte {
	return b.data[b.off:]
}

// Clear all bytes in buffer.
func (b *Buffer) Clear() {
	b.data = b.data[:0]
	b.off = 0
}
