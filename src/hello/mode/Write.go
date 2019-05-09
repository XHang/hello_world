package mode

type Write interface {
	WriteByte(b byte)
	Fresh()
}
