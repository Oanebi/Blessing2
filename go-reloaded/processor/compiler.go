package processor

func Compiler(s string) string {
	s = Hextodec(s)
	s = Bintodec(s)
	s = Upper(s)
	s = Lower(s)
	s = Capital(s)
	s = Quote(s)
	return s
}
