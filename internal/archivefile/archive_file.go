package archivefile

type ArchiveFile struct {
	Id       int
	fileInfo map[string]any
}

func NewArchiveFile(id int, m map[string]any) *ArchiveFile {
	return &ArchiveFile{
		Id:       id,
		fileInfo: m,
	}
}
func (receiver *ArchiveFile) getProperty(k string) any {
	if v, found := receiver.fileInfo[k]; found {
		return v
	}
	return nil
}

func (receiver *ArchiveFile) testAttribute(targetBit int) bool {
	attributes := receiver.getProperty("attributes")
	if attributes == nil {
		return false
	}
	return (attributes.(int) & targetBit) == targetBit
}
