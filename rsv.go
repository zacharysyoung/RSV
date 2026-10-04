package rsv

import (
	"bytes"
	"io"
)

const (
	EOV byte = 0xFE // End of Value
	EOR      = 0xFF // End of Row
)

func Read(data []byte) (rows [][]string) {
	_rows := bytes.Split(data, []byte{EOR})
	for i := 0; i < len(_rows)-1; i++ {
		row := []string{}
		values := bytes.Split(_rows[i], []byte{EOV})
		for j := 0; j < len(values)-1; j++ {
			row = append(row, string(values[j]))
		}
		rows = append(rows, row)
	}
	return
}

func Write(rows [][]string) (data []byte, err error) {
	buf := &bytes.Buffer{}
	for _, row := range rows {
		for _, value := range row {
			buf.WriteString(value)
			buf.Write([]byte{EOV})
		}
		_, err := buf.Write([]byte{EOR})
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

type Reader struct{ r io.Reader }

func NewReader(r io.Reader) Reader { return Reader{r} }

func (r Reader) ReadAll() ([][]string, error) {
	p, err := io.ReadAll(r.r)
	if err != nil {
		return nil, err
	}
	rows := Read(p)
	return rows, nil
}

type Writer struct{ w io.Writer }

func NewWriter(w io.Writer) Writer { return Writer{w} }

func (w Writer) WriteAll(rows [][]string) error {
	data, err := Write(rows)
	if err != nil {
		return err
	}
	_, err = w.w.Write(data)
	return err
}
