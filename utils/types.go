package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type EncodingTable map[rune]string
type BinaryChunk string
type BinaryChunks []BinaryChunk
type HexChunk string
type HexChunks []HexChunk


func (hc HexChunks) Tostring() string {

	const sep = " "

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

func (bc BinaryChunks) ToHex() HexChunks {
	 
	res := make(HexChunks, 0, len(bc))

	for _, chunk := range bc {
		res = append(res, chunk.ToHex())
	}
	return res
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