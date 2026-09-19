package vlc

import (
	"strings"
	"unicode"
	"unicode/utf8"
)



func Encode(str string) string {

	str = prepareText(str)
	binStr := encodeBin(str)
	chunks := splitByChunks(binStr, ChunkSize)

	return chunks.ToHex().Tostring()
}

func Decode(encodedtext string) string {

	result := NewhexChunks(encodedtext).ToBinary().Join()
	dTree := getEncodingTable().DecodingTree()

	return exportText(dTree.Decode(result))
}

func prepareText(str string) string {
	
	var buff strings.Builder

	for _, ch := range str {
		
		if unicode.IsUpper(ch) {
			buff.WriteRune('!')
			buff.WriteRune(unicode.ToLower(ch))
		} else {
			buff.WriteRune(ch)
		}
	}
	return buff.String()
}

func exportText(str string) string {

	var buff strings.Builder
	var isCapital bool

	for _, ch := range str {

		if isCapital {
			buff.WriteRune(unicode.ToUpper(ch))
			isCapital = false
			continue
		}

		if ch == '!' {
			isCapital = true
			continue
		} else {
			buff.WriteRune(ch)
		}
	}

	return buff.String()
}


func encodeBin(str string) string {

	var buff strings.Builder

	for _, ch := range str {
		buff.WriteString(bin(ch))
	}

	return buff.String()
}
func bin(ch rune) string {
	table := getEncodingTable()

	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}
func getEncodingTable() EncodingTable {

	return EncodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}



func splitByChunks(binStr string, chunkSize int) BinaryChunks {

	strLen := utf8.RuneCountInString(binStr)
	chunksCount := strLen / chunkSize

	if strLen / chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)
	var buff strings.Builder

	for idx, ch := range binStr {

		buff.WriteString(string(ch))

		if (idx+1) % chunkSize == 0 {
			res = append(res, BinaryChunk(buff.String()))
			buff.Reset()
		}
	}

	if buff.Len() != 0 {
		lastChunk := buff.String()
		lastChunk += strings.Repeat("0", chunkSize - len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}