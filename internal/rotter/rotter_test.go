package rotter

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	orig := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	Run()

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = orig

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	if got, want := buf.String(), "rotter\n"; got != want {
		t.Fatalf("Run() wrote %q, want %q", got, want)
	}
}
