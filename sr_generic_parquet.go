package main

import (
	"encoding/json"
	"fmt"
)

type SrGenericParquet[T any] struct {
	*ParquetHandler
	buf []T
}

func (gp *SrGenericParquet[T]) Write(resp *SrRefResponse) error {
	for _, r := range resp.Result {
		var item T
		json.Unmarshal([]byte(r.String()), &item)
		if err := gp.ParquetHandler.ParquetWriter.Write(item); err != nil {
			return fmt.Errorf("failed to write parquet: %w", err)
		}
	}
	return nil
}

func NewSrGenericParquet[T any](output string) (*SrGenericParquet[T], error) {
	ph, err := NewParquetHandler(new(T), output)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate ParquetHandler for %T: %w", new(T), err)
	}
	return &SrGenericParquet[T]{
		ParquetHandler: ph,
		buf:            make([]T, 0),
	}, nil
}
