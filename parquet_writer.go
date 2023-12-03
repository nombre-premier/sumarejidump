package main

import (
	"fmt"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

const (
	defaultRowGroupSize = 128 * 1024 * 1024 //128M
	defaultPageSize = 8 * 1024 //8K
	defaultConcurrency = 4
)

type ParquetWriter struct {
	ParquetFile *local.LocalFile
	ParquetWriter *writer.ParquetWriter
}

func NewParquetWriter(header interface{}, filePath string) (*ParquetWriter, error) {
	fmt.Printf("Opening file writer for path: %s\n", filePath)
	fw, err := local.NewLocalFileWriter(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open FileWriter: %w", err)
	}

	pw, err := writer.NewParquetWriter(fw, header, defaultConcurrency)
	if err != nil {
		fmt.Printf("Initializing parquet writer with header: %+v\n", header)
		return nil, fmt.Errorf("failed to initiate ParquetWriter with header: %+v\nerror: %w", header, err)
	}

	pw.RowGroupSize = defaultRowGroupSize
	pw.PageSize = defaultPageSize
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
	if err := pw.ParquetWriter.WriteStop(); err != nil {
		fmt.Errorf("WriteStop error %w", err)
	}
	pw.ParquetFile.Close()
}
