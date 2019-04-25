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
	data, err := agg.Data(name)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}

func (agg *AGG) Data(name string) ([]byte, error) {
	file, ok := agg.fileMap[name]
	if !ok {
		return nil, errors.Wrapf(os.ErrNotExist, "failed to find %v", name)
	}

	// We want position in fileData, and not in the whole file.
	// Substract the header (2 bytes) and the fileTable (12 bytes each file).
	offset := file.Offset() - 2 - 12*len(agg.fileTable)
	return agg.fileData[offset : offset+file.Size()], nil
}

func (agg *AGG) Files(ext string) []string {
	var files []string
	for filename, _ := range agg.fileMap {
		if strings.HasSuffix(strings.ToLower(filename), "."+strings.ToLower(ext)) {
			files = append(files, filename)
		}
	}
	return files
}

func (agg *AGG) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Number of files: %v\n", agg.header.NumFiles())

	for i := 0; i < len(agg.fileTable); i++ {
		fmt.Fprintf(&b, "%v: %v", agg.fileNames[i], agg.fileTable[i])
	}

	return b.String()
}
