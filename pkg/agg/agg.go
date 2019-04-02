package agg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type AGG struct {
	header    Header
	fileTable []File
	fileData  []byte
	fileNames []FileName

	fileMap map[string]*File
}

func Load(r io.ReadSeeker) (*AGG, error) {
	header := Header{}
	if err := binary.Read(r, binary.LittleEndian, &header); err != nil {
		return nil, errors.Wrap(err, "failed to parse AGG header")
	}

	fileTable := make([]File, header.NumFiles())
	if err := binary.Read(r, binary.LittleEndian, fileTable); err != nil {
		return nil, errors.Wrap(err, "failed to parse file table")
	}

	dataSize := 0
	for i := 0; i < len(fileTable); i++ {
		dataSize += fileTable[i].Size()
	}

	fileData := make([]byte, dataSize)
	if err := binary.Read(r, binary.LittleEndian, fileData); err != nil {
		return nil, errors.Wrap(err, "failed to parse file data")
	}

	fileNames := make([]FileName, header.NumFiles())
	if err := binary.Read(r, binary.LittleEndian, fileNames); err != nil {
		return nil, errors.Wrap(err, "failed to parse file names")
	}

	fileMap := map[string]*File{}
	for i, fileName := range fileNames {
		fileMap[fileName.String()] = &fileTable[i]
	}

	agg := &AGG{
		header:    header,
		fileTable: fileTable,
		fileData:  fileData,
		fileNames: fileNames,
		fileMap:   fileMap,
	}

	return agg, nil
}

func (agg *AGG) Open(name string) (io.ReadSeeker, error) {
	file, ok := agg.fileMap[name]
	if !ok {
		return nil, os.ErrNotExist
	}

	offset := file.Offset()
	return bytes.NewReader(agg.fileData[offset : offset+file.Size()]), nil
}

func (agg *AGG) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of files: %v\n", agg.header.NumFiles())

	for i := 0; i < len(agg.fileTable); i++ {
		fmt.Fprintf(&b, "%v: %v", agg.fileNames[i], agg.fileTable[i])
	}

	return b.String()
}
