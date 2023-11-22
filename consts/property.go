package consts

type Property byte

const (
	PropertyEnd Property = iota
	PropertyHeader
	PropertyArchiveProperties
	PropertyAdditionalStreamsInfo
	PropertyMainStreamsInfo
	PropertyFilesInfo
	PropertyPackInfo
	PropertyUnpackInfo
	PropertySubStreamsInfo
	PropertySize
	PropertyCrc
	PropertyFolder
	PropertyCodersUnpackSize
	PropertyNumUnpackStream
	PropertyEmptyStream
	PropertyEmptyFile
	PropertyAnti
	PropertyName
	PropertyCreationTime
	PropertyLastAccessTime
	PropertyLastWriteTime
	PropertyAttributes
	PropertyComment
	PropertyEncodedHeader
	PropertyStartPos
	PropertyDummy
)
