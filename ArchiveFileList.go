package sevenzip

import (
	"errors"

	"sevenzip/util"
)

type ArchiveFileList struct {
	filesList []map[string]any
	index     int
	offset    int
}

func NewArchiveFileList(offset ...int) *ArchiveFileList {
	return &ArchiveFileList{
		index:     0,
		filesList: make([]map[string]any, 0),
		offset:    util.DefaultOrGet(0, offset...)}
}
func (a *ArchiveFileList) Append(fileInfo map[string]any) {
	a.filesList = append(a.filesList, fileInfo)
}
func (a *ArchiveFileList) Len() int {
	return len(a.filesList)
}
func (a *ArchiveFileList) Get(index int) (*ArchiveFile, error) {
	if index > len(a.filesList) {
		return nil, errors.New("sequence index out of range")
	}
	if index < 0 {
		return nil, errors.New("sequence index out of range")
	}

	return NewArchiveFile(index+a.offset, a.filesList[index]), nil
}
