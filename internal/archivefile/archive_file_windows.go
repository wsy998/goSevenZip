//go:build windows

package archivefile

import (
	"golang.org/x/sys/windows"
)

func (receiver *ArchiveFile) Readonly() bool {
	return receiver.testAttribute(windows.FILE_ATTRIBUTE_READONLY)
}
