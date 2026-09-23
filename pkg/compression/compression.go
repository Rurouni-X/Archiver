package compression

type Encoder interface {
	Encode(string) []byte
}

type Decoder interface {
	Decode([]byte) string
}
