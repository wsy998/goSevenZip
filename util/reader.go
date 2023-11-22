package util

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"unicode/utf16"
)

func ReadBytesU64(r io.Reader, count uint64) ([]byte, error) {
	b := make([]byte, count)
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func ReadBytes(r io.Reader, count int) ([]byte, error) {
	if count < 0 {
		return nil, errors.New("count not be negative")
	}
	return ReadBytesU64(r, uint64(count))
}
func ReadByte(r io.Reader) (byte, error) {
	readBytes, err := ReadBytes(r, 1)
	if err != nil {
		return 0, err
	}
	return readBytes[0], nil
}
func ReadRealUint64(r io.Reader) (uint64, error) {
	var s uint64
	err := binary.Read(r, binary.LittleEndian, &s)
	if err != nil {
		return 0, err
	}
	return s, nil
}
func ReadUint32(r io.Reader) (uint32, error) {
	var i uint32
	err := binary.Read(r, binary.LittleEndian, &i)
	if err != nil {
		return 0, err
	}

	return i, nil
}
func ReadUtf16String(r io.Reader) (string, error) {
	b := make([]uint16, 0)
	for i := 0; i < 65535; i++ {
		if data, err := ReadBytes(r, 2); err != nil {
			return "", err
		} else {
			if bytes.Equal(data, []byte{0x00, 0x00}) {
				break
			}
			b = append(b, uint16(data[1])>>8|uint16(data[0]))
		}
	}
	return string(utf16.Decode(b)), nil
}
func ReadBooleans(r io.Reader, count int, checkAll ...bool) ([]bool, error) {
	_checkAll := DefaultOrGet(false, checkAll...)
	if _checkAll {
		if allDefined, err := ReadByte(r); err != nil {
			return nil, err
		} else {
			if allDefined != 0x00 {
				return Repeat(count, true), nil
			}
		}
	}
	result := make([]bool, count)
	b := 0
	mask := 0
	for i := 0; i < count; i++ {
		if mask == 0 {

			if readByte, err := ReadByte(r); err != nil {
				return nil, err
			} else {
				b = int(readByte)
			}
			mask = 0x80
		}
		result[i] = b&mask != 0
		mask >>= 1
	}
	return result, nil
}
func ReadUint64(r io.Reader) (uint64, error) {
	var b int
	if readByte, err := ReadByte(r); err != nil {
		return 0, err
	} else {
		b = int(readByte)
		if b == 255 {
			return ReadRealUint64(r)
		}
	}
	mask := 0x80
	vlen := 8
	blen := [][]int{}
	for _, d := range blen {
		v, l := d[0], d[1]
		if b <= v {
			vlen = l
			break
		}
		mask >>= 1
	}
	if vlen == 0 {
		return uint64(b & (mask - 1)), nil
	}
	var value uint64
	if val, err := ReadBytes(r, vlen); err != nil {
		return 0, err
	} else {
		value, _ = binary.Uvarint(val)
	}
	highPart := b & (mask - 1)
	return value + (highPart << (vlen * 8)), nil
}
func ReadCrcs(r io.Reader, count int) ([]uint32, error) {
	if data, err := ReadBytes(r, count*4); err != nil {
		return nil, err
	} else {
		uint32s := make([]uint32, count)
		for i := 0; i < count; i++ {
			s := data[i*4 : i*4+4]
			var u uint32
			if err := binary.Read(bytes.NewBuffer(s), binary.LittleEndian, &u); err != nil {
				return nil, err
			}
			uint32s[i] = u
		}
		return uint32s, nil
	}

}
