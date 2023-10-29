package main

import (
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

type ParquetWriter struct {
	ParquetFile *local.LocalFile
	ParquetWriter *writer.ParquetWriter
}

func NewParquetWriter(header interface{}, output string) (*ParquetWriter, error) {
	filePath := output + ".parquet"
	fw, err := local.NewLocalFileWriter(filePath)
	if err != nil {
		return nil, err
	}

	pw, err := writer.NewParquetWriter(fw, header, 2)
	if err != nil {
		return nil, err
	}
	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	return &ParquetWriter{
		ParquetFile:    fw.(*local.LocalFile),
		ParquetWriter:  pw,
	}, nil
}

func (pw *ParquetWriter) Write(i interface{}) error {
	if err := pw.ParquetWriter.Write(i); err != nil {
		return err
	}
	return nil
}

func (pw *ParquetWriter) Close() {
	pw.ParquetWriter.WriteStop()
	pw.ParquetFile.Close()
}
