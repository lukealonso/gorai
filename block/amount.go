package gorai

type Amount struct {
	Value [16]byte
}

func (a *Amount) MarshalText() ([]byte, error) {
	return encodeHex(a.Value[:]), nil
}

func (a *Amount) UnmarshalText(b []byte) error {
	return decodeHex(a.Value[:], b)
}

func (a *Amount) Bytes() []byte {
	return a.Value[:]
}
