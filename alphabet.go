package gibberish

type Alphabet struct {
	Upper string
	Lower string
	Punct string
}

func (a *Alphabet) Word(n *Normal) string {
	txt := word(n)
	fill(n, a.Lower, txt)
	return string(txt)
}

func (a *Alphabet) WordUpper(n *Normal) string {
	txt := word(n)
	fill(n, a.Upper, txt[0:1])
	fill(n, a.Lower, txt[1:])
	return string(txt)
}

func word(n *Normal) []byte {
	sz := n.Sample()
	if sz == 0 {
		sz = 1
	}
	txt := make([]byte, sz)
	return txt
}

func fill(n *Normal, from string, to []byte) {
	n := len(from)
	for i := range to {
		txt[i] = from[n.Rand.Intn(n)]
	}
}
