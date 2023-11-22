package sevenzip

type Bond struct {
	InCoder  uint64
	OutCoder uint64
}

func NewBond(inCoder uint64, outCoder uint64) *Bond {
	return &Bond{InCoder: inCoder, OutCoder: outCoder}
}
