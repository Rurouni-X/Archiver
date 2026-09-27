package vlc

import (
	"Archiver/pkg/compression/vlc/table"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

func (ed EncoderDecoder) Encode(str string) []byte {
	table := ed.tblGenerator.NewTable(str)
	encoded := encodeBin(str, table)
	return buildEncodedFile(table, encoded)
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)
	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	tableSizeBinary, data := data[:4], data[4:]
	dataSizeBinary, data := data[:4], data[4:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)
	tblBinary, data := data[:tableSize], data[tableSize:]
	tbl := decodeTable(tblBinary)
	body := NewBinChunks(data).Join()

	return tbl, body[:dataSize]
}

func buildEncodedFile(table table.EncodingTable, data string) []byte {
	encodedTable := encodeTable(table)
	var buf bytes.Buffer
	buf.Write(encodeInt(len(encodedTable)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTable)
	buf.Write(splitByChunks(data, ChunkSize).Bytes())

	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}

func encodeTable(table table.EncodingTable) []byte {
	var tablesBuf bytes.Buffer
	if err := gob.NewEncoder(&tablesBuf).Encode(table); err != nil {
		log.Fatal("can't serialization table: ", err)
	}
	return tablesBuf.Bytes()
}

func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable
	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("can't decode table", err)
	} 
	return tbl
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

func encodeBin(str string, table table.EncodingTable) string {

	var buff strings.Builder

	for _, ch := range str {
		buff.WriteString(bin(ch, table))
	}

	return buff.String()
}
func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}

	return res
}

func splitByChunks(binStr string, chunkSize int) BinaryChunks {

	strLen := utf8.RuneCountInString(binStr)
	chunksCount := strLen / chunkSize

	if strLen/chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)
	var buff strings.Builder

	for idx, ch := range binStr {

		buff.WriteString(string(ch))

		if (idx+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buff.String()))
			buff.Reset()
		}
	}

	if buff.Len() != 0 {
		lastChunk := buff.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}
