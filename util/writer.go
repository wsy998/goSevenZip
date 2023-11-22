package util

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
	"unicode/utf16"
)

func Write(w io.Writer, data any) error {
	switch value := data.(type) {
	case byte:
		return WriteByte(w, value)
	case []byte:
		return WriteBytes(w, value)
	case uint64:
		return WriteUint64(w, value)
	case uint32:
		return WriteUint32(w, value)
	case []bool:
		return WriteBooleans(w, value)
	case string:
		return WriteUtf16String(w, value)
	default:
		return errors.New("not support data type")
	}
}

func WriteByte(w io.Writer, data byte) error {
	return WriteBytes(w, []byte{data})
}
func WriteBytes(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func WriteRealUint64(w io.Writer, value uint64) error {
	return binary.Write(w, binary.LittleEndian, value)
}
func WriteUint64(w io.Writer, value uint64) error {
	if value < 0x80 {
		return WriteByte(w, byte(value))
	}
	if value > 0x01FFFFFFFFFFFFFF {
		if err := WriteByte(w, 0xff); err != nil {
			return err
		}
		return binary.Write(w, binary.NativeEndian, value)
	}
	bL := (bits.Len64(value) + 7) / 8
	ba := make([]byte, bL)
	binary.PutUvarint(ba, value)
	highByte := int(ba[bL-1])
	if highByte < 2<<(8-bL-1) {
		for i := 0; i < bL-1; i++ {
			highByte |= 0x80 >> i
		}
		if err := WriteByte(w, byte(highByte)); err != nil {
			return err
		}
		if err := WriteBytes(w, ba[:bL-1]); err != nil {
			return err
		}
		return nil
	}
	mask := 0x80
	for x := 0; x < bL; x++ {
		mask |= 0x80 >> x
	}
	if err := Write(w, byte(mask)); err != nil {
		return err
	}
	if err := WriteBytes(w, ba); err != nil {
		return err
	}
	return nil
}
func WriteUint32(w io.Writer, value uint32) error {
	return binary.Write(w, binary.LittleEndian, value)
}

func WriteBooleans(w io.Writer, booleans []bool, allDefined ...bool) error {
	_allDefined := DefaultOrGet(true, allDefined...)
	t := true
	for _, boolean := range booleans {
		t = t && boolean
	}
	if _allDefined && t {
		return WriteByte(w, 0x01)
	} else if _allDefined {
		if err := WriteByte(w, 0x00); err != nil {
			return err
		}
	}
	o := make([]byte, -(-len(booleans) / 8))
	for i, boolean := range booleans {
		if boolean {
			o[i/8] |= 1 << (7 - i%8)
		}
	}
	return WriteBytes(w, o)
}
func WriteUtf16String(w io.Writer, str string) error {
	encode := utf16.Encode([]rune(str))
	if err := binary.Write(w, binary.LittleEndian, encode); err != nil {
		return err
	}

	return WriteBytes(w, []byte{0x00, 0x00})
}
func WriteCrcs(w io.Writer, crcs []uint32) error {
	for _, crc := range crcs {
		if err := WriteUint32(w, crc); err != nil {
			return err
		}
	}
	return nil
}
