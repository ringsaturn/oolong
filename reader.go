package oolong

import (
	"bytes"
	"errors"
	"io"
)

var (
	_GRIB_HEADER        = [4]byte{0x47, 0x52, 0x49, 0x42} // "GRIB"
	_GRIB_EDITION uint8 = 2                               // GRIB Edition number

	ErrInvalidHeader        = errors.New("oolong: invalid header")         // Error message for invalid header
	ErrInvalidEditionNumber = errors.New("oolong: invalid edition number") // Error message for invalid edition number
)

type FileReader interface {
	io.Reader
	io.ReaderAt
	io.Closer
}

type GRIBReader struct {
	File FileReader
}

func NewGRIBReader(f FileReader) *GRIBReader {
	return &GRIBReader{File: f}
}

func (g *GRIBReader) ValidaHeader() error {
	header := make([]byte, 4)
	_, err := g.File.ReadAt(header, 0)
	if err != nil {
		return err
	}

	if !bytes.Equal(header, _GRIB_HEADER[:]) {
		return ErrInvalidHeader
	}

	return nil
}

func (g *GRIBReader) ReadAt(p []byte, off int64) (n int, err error) {
	return g.File.ReadAt(p, off)
}

func (g *GRIBReader) Read(p []byte) (n int, err error) {
	return g.File.Read(p)
}

func (g *GRIBReader) Close() error {
	return g.File.Close()
}
