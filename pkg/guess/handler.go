package guess

type Accuracy struct {
	Exact int
	Near  int
}

func Compare(secret string, guess string) Accuracy {
	var s []byte = []byte(secret)
	var g []byte = []byte(guess)
	len := len(secret)
	accuracy := Accuracy{}

	// First case - Exact matches
	for i := 0; i < len; i++ {
		if s[i] == g[i] {
			accuracy.Exact++
			s[i] = 'x'
		}
	}

	// Second case - Near matches
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if g[i] == s[j] {
				accuracy.Near++
				s[j] = 'x'
				break
			}
		}
	}

	return accuracy
}
