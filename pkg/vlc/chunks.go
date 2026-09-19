package vlc

import (
	"fmt"
	"strconv"
	"strings"
)

const ChunkSize = 8
const sep = " "
type EncodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk
type HexChunk string
type HexChunks []HexChunk


func (hc HexChunks) Tostring() string {

	switch len(hc) {
		case 0:
			return ""
		case 1:
			return string(hc[0])
	}

	var buff strings.Builder

	buff.WriteString(string(hc[0]))
	for _, v := range hc[1:] {
		buff.WriteString(sep)
		buff.WriteString(string(v))
	}

	return buff.String()
}

func (hc HexChunks) ToBinary() BinaryChunks {

	res := make(BinaryChunks, 0, len(hc))

	for _, v := range hc {
		res = append(res, v.ToBinary())
	}
	return res
}

func (hc HexChunk) ToBinary() BinaryChunk {

	num, err := strconv.ParseUint(string(hc), 16, ChunkSize)

	if err != nil {
		panic("can't parse hex chunk" + err.Error())
	}

	res := fmt.Sprintf("%08b", num)

	return BinaryChunk(res)
}

func (bc BinaryChunks) ToHex() HexChunks {
	 
	res := make(HexChunks, 0, len(bc))

	for _, chunk := range bc {
		res = append(res, chunk.ToHex())
	}
	return res
}

func (bc BinaryChunks) Join() string {

	var buff strings.Builder

	for _, v := range bc {
		buff.WriteString(string(v))
	}

	return buff.String()
}

func (bc BinaryChunk) ToHex() HexChunk {

	num, err := strconv.ParseUint(string(bc), 2, ChunkSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func NewhexChunks(str string) HexChunks {

	parts := strings.Split(str, sep)
	res := make(HexChunks, 0, len(parts))

	for _, v := range parts {
		res = append(res, HexChunk(v))
	}
	return res
}