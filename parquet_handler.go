package main

import (
	"github.com/xitongsys/parquet-go-source/local"
)

type SrParquetHandlerIf interface {
	Write(resp *SrRefResponse) error
	GetParquetWriter() *ParquetWriter
}

type ParquetHandler struct {
	Header        interface{}
	ParquetWriter *ParquetWriter
	File          *local.LocalFile
}

func NewParquetHandler(header interface{}, output string) (*ParquetHandler, error) {
	pw, err := NewParquetWriter(header, output)
	if err != nil {
		return nil, err
	}
	return &ParquetHandler{
		ParquetWriter: pw,
	}, nil
}

func (ph *ParquetHandler) GetParquetWriter() *ParquetWriter { return ph.ParquetWriter }
