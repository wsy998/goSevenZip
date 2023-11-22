package util

import (
	"hash/crc32"
	"runtime"
)

func DefaultOrGet[T any](b T, defined ...T) T {
	if len(defined) > 0 {
		return defined[0]
	}
	return b
}
func Repeat[T any](count int, value T) []T {
	ts := make([]T, count)
	for i := 0; i < count; i++ {
		ts[i] = value
	}
	return ts
}
func CalculateCrc32(data []byte, value uint32, blockSize ...int) uint32 {
	_blockSize := DefaultOrGet(1024*1024, blockSize...)
	l := len(data)
	var v uint32
	if l < _blockSize {
		v = crc32.Update(value, crc32.IEEETable, data)
	} else {
		pos := _blockSize
		v = crc32.Update(value, crc32.IEEETable, data[:pos])
		for pos < l {
			v = crc32.Update(v, crc32.IEEETable, data[pos:pos+_blockSize])
			pos += _blockSize
		}
	}
	return v & 0xFFFFFFFF
}
func GetDefaultBlockSize(d ...int) int {
	var s int
	if runtime.GOARCH == `amd64` {
		s = 1048576
	} else {
		s = 32768
	}
	t := DefaultOrGet(s, d...)
	if t == 0 {
		return s
	} else {
		return t
	}
}
