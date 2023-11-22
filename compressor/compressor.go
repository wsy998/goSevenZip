package compressor

type ISevenZipCompressor interface {
	Compress(data []byte) []byte
	Flush()
}

type ISevenZipDecompressor interface {
	Decompress(data []byte, maxLength ...uint64) []byte
}
