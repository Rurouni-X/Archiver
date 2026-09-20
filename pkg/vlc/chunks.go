package vlc

import (
	"fmt"
	"strconv"
	"strings"
)

const ChunkSize = 8
type EncodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk

func (bc BinaryChunks) Join() string {

	var buff strings.Builder

	for _, v := range bc {
		buff.WriteString(string(v))
	}

	return buff.String()
}

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, v := range data {
		res = append(res, NewBinChunk(v))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

func (bcs BinaryChunks) Bytes() []byte {

	res := make([]byte, 0, len(bcs))

	for _, ch := range bcs {
		res = append(res, ch.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {

	num, err := strconv.ParseUint(string(bc), 2, ChunkSize)

	if err != nil {
		panic("can't parse binary chunk" + err.Error())
	}

	return byte(num)
}