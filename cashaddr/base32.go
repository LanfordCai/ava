package cashaddr

type base32Encoder struct {
	alphabet string
}

func newBase32Encoder(alphabet string) *base32Encoder {
	return &base32Encoder{alphabet: alphabet}
}

func (e *base32Encoder) numToCode() map[byte]rune {
	result := make(map[byte]rune)
	for i, r := range e.alphabet {
		result[byte(i)] = r
	}
	return result
}

func (e *base32Encoder) codeToNum() map[rune]byte {
	result := make(map[rune]byte)
	for i, r := range e.alphabet {
		result[r] = byte(i)
	}
	return result
}

func (e *base32Encoder) encode(payload []byte) string {
	result := []rune{}
	set := e.numToCode()
	for _, b := range payload {
		result = append(result, set[b])
	}
	return string(result)
}

func (e *base32Encoder) decode(str string) []byte {
	result := []byte{}
	set := e.codeToNum()
	for _, r := range str {
		result = append(result, set[r])
	}
	return result
}
