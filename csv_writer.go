package main

import (
	"bufio"
	"os"

	"github.com/gocarina/gocsv"
)

type CSVWriter struct {
	file     *os.File
	buffer   *bufio.Writer
	csvWrite *gocsv.SafeCSVWriter
	empData  interface{}
}

func NewCSVWriter(empData interface{}, name string) (*CSVWriter, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	cw := CSVWriter{}
	cw.file = file
	cw.buffer = bufio.NewWriter(cw.file)
	cw.csvWrite = gocsv.DefaultCSVWriter(cw.buffer)
	cw.empData = empData
	gocsv.MarshalCSV(empData, cw.csvWrite)
	return &cw, nil

}

func (cw *CSVWriter) Write(i interface{}) error {
	err := gocsv.MarshalCSVWithoutHeaders(i, cw.csvWrite)
	if err != nil {
		return err
	}
	return nil
}

func (cw *CSVWriter) Close() {
	cw.buffer.Flush()
	cw.file.Close()
}
