package pipeprotocol

import (
	"bytes"
	"compress/gzip"
	"io"
)

func dataCompressForGzip(content []byte) []byte {
	var data bytes.Buffer
	gzipWrite, err := gzip.NewWriterLevel(&data, gzip.BestCompression)
	if err != nil {
		return nil
	}

	writeLen := 0
	for writeLen < len(content) {
		n, err := gzipWrite.Write(content[writeLen:])
		if err != nil {
			return nil
		}
		writeLen += n
	}

	err = gzipWrite.Flush()
	if err != nil {
		return nil
	}
	err = gzipWrite.Close()
	if err != nil {
		return nil
	}
	return data.Bytes()
}

func dataDecompressForGzip(compressedData []byte) []byte {
	gzipReader, err := gzip.NewReader(bytes.NewBuffer(compressedData))
	if err != nil {
		return nil
	}

	content, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil
	}

	return content
}
