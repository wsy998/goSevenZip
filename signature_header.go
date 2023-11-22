package sevenzip

import (
	"bytes"
	"errors"
	"io"
	"os"

	"sevenzip/consts"
	"sevenzip/util"
)

type OpenMode byte

const (
	ModeAdd OpenMode = iota
	ModeWrite
	ModeRead
)

type File struct {
	r    io.Writer
	mode OpenMode

	Filter            []map[string]int
	Dereference       bool
	Pwd               string
	HeaderEncryption  bool
	BlockSize         int
	Mp                bool
	PasswordProtected bool
	EncodedHeaderMode bool
	fileRefCnt        int
	dict              map[string]any
}

type FileOption struct {
	Filters          []map[string]int
	Dereference      bool
	Password         string
	HeaderEncryption bool
	BlockSize        int
	Mp               bool
}

var DefaultFileOption = &FileOption{
	Filters:          nil,
	Dereference:      false,
	Password:         "",
	HeaderEncryption: false,
	BlockSize:        0,
	Mp:               false,
}

func (f *File) checkMode() error {
	if t, ok := f.r.(io.Reader); !ok {
		return errors.New("not support current mode")
	} else {
		if !f.checkIsSevenFile(t) {
			return errors.New("this is not 7z file")
		}
	}
	return nil
}
func (f *File) needRead() bool {
	return f.mode == ModeAdd || f.mode == ModeRead
}
func (f *File) checkIsSevenFile(t io.Reader) bool {
	ks, err := util.ReadBytes(t, 6)
	if err != nil {
		return false
	}
	return bytes.Equal(ks, []byte{'7', 'z', 0xBC, 0xAF, 0x27, 0x1C})
}

func (f *File) parseVersion() ([2]byte, error) {
	readByte, err := util.ReadBytes(f.r.(io.Reader), 2)
	if err != nil {
		return [2]byte{}, err
	}
	return [2]byte(readByte), nil
}

func (f *File) parse() error {
	// if v, err := f.parseVersion(); err != nil {
	// 	return err
	// } else {
	// 	f.Version = v
	// }
	// if readUint32, err := util.ReadUint32(f.r.(io.Reader)); err != nil {
	// 	return err
	// } else {
	// 	f.StartHeaderCrC = readUint32
	// }
	return nil
}

func Open(r io.Writer, mode OpenMode, opt ...*FileOption) (*File, error) {
	p := util.DefaultOrGet(DefaultFileOption, opt...)
	if p == nil {
		p = DefaultFileOption
	}
	l := &File{
		r:                 r,
		mode:              mode,
		Mp:                p.Mp,
		PasswordProtected: p.Password != "",
		BlockSize:         util.GetDefaultBlockSize(p.BlockSize),
		EncodedHeaderMode: true,
		HeaderEncryption:  p.HeaderEncryption,
		fileRefCnt:        1,
		dict:              make(map[string]any),
		Dereference:       p.Dereference,
	}
	switch mode {
	case ModeWrite:
		l.prepareWrite(p.Filters, p.Password)
	}
	return l, nil
}
func OpenFile(f string, mode OpenMode) (*File, error) {
	file, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return Open(file, mode)
}
func (f *File) Close() error {
	if c, ok := f.r.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

func (f *File) prepareWrite(filters []map[string]int, password string) {
	var ft = filters
	if password != "" && ft != nil {
		ft = consts.EncryptedArchiveFilter
	} else if ft == nil {
		ft = consts.ArchiveFilter
	}
	files = NewArchiveFileList()

}
