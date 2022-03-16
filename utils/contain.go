package utils
type Universe map[string]int

func newUniverse(sss *[]string) *Universe{
	uni := make(Universe)
	for _, w := range *sss {
		uni[w] += 1
	}
	
	return &uni
}

func Contain(sss *[]string, ss *[]string) bool{
	uni_sss := newUniverse(sss)
	uni_ss  := newUniverse(ss)

	for w := range *uni_ss {
		if (*uni_sss)[w] < (*uni_ss)[w] {return false}
	}
	return true
}