package strand

// ToRNA converts DNA strand to RNA strand
func ToRNA(dna string) (rna string) {
	n := map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}
	for _, v := range dna {
		rna += n[v]
	}
	return rna
}
