package consts

const (
	FilterLzma     uint64 = 4611686018427387905
	FilterLzma2           = 33
	FilterDelta           = 3
	FilterArm             = 7
	FilterArmthumb        = 8
	FilterIa64            = 6
	FilterPowerpc         = 5
	FilterSparc           = 9
	FilterX86             = 4

	FilterCryptoAes256Sha256 = 0x06F10701
)
const (
	PresetDefault = 6
	PresetExtreme = 2147483648
)

var (
	ArchiveFilter = []map[string]int{
		{"id": FilterX86},
		{"id": FilterLzma2, "preset": 7 | PresetDefault},
	}
	EncodedHeaderFilter = []map[string]int{
		{"id": FilterLzma2, "preset": 7 | PresetDefault},
	}
	EncryptedHeaderFilter  = []map[string]int{{"id": FilterCryptoAes256Sha256}}
	EncryptedArchiveFilter = []map[string]int{{"id": FilterLzma2, "preset": 7 | PresetDefault}, {"id": FilterCryptoAes256Sha256}}
)
