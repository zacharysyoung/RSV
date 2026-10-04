package rsv

import (
	"bytes"
	"slices"
	"testing"
)

var (
	rows = [][]string{
		{"aaa", "", "ccc"},
		{},
		{"zzz", "yyy"},
	}

	data = []byte{
		'a', 'a', 'a', EOV, EOV, 'c', 'c', 'c', EOV, EOR,
		EOR,
		'z', 'z', 'z', EOV, 'y', 'y', 'y', EOV, EOR,
	}
)

func TestRead(t *testing.T) {
	got := Read(data)
	for i, x := range got {
		if !slices.Equal(x, rows[i]) {
			t.Errorf("for row %d, got %q; want %q", i+1, x, rows[i])
		}
	}
}

func TestWrite(t *testing.T) {
	got, err := Write(rows)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(got, data) {
		t.Fatalf("\n got %q;\nwant %q", got, data)
	}
}

func TestReader(t *testing.T) {
	reader := NewReader(bytes.NewReader(data))
	got, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for i, x := range got {
		if !slices.Equal(x, rows[i]) {
			t.Errorf("for row %d, got %q; want %q", i+1, x, rows[i])
		}
	}
}

func TestWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	writer := NewWriter(buf)
	err := writer.WriteAll(rows)
	if err != nil {
		t.Error(err)
	}
	got := buf.Bytes()
	if !slices.Equal(got, data) {
		t.Fatalf("\n got %q;\nwant %q", got, data)
	}
}
