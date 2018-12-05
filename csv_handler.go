package main

type SrCSVHandlerIf interface {
	Write(resp *SrRefResponse) *CSVWriter
	GetCSVWriter() *CSVWriter
}

type CSVHandler struct {
	Header    interface{}
	CSVWriter *CSVWriter
}

func NewCSVHandler(header interface{}, output string) (*CSVHandler, error) {
	cw, err := NewCSVWriter(header, output)
	if err != nil {
		return nil, err
	}
	return &CSVHandler{
		Header:    header,
		CSVWriter: cw,
	}, nil
}

func (ch *CSVHandler) GetCSVWriter() *CSVWriter { return ch.CSVWriter }
