package ch1

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestEcho2(t *testing.T) {
	os.Args = []string{"program", "arg1", "arg2", "arg3"}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	echo2()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	got := <-outC
	want := "program arg1 arg2 arg3\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestEcho3(t *testing.T) {
	os.Args = []string{"program", "arg1", "arg2", "arg3"}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	echo3()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	got := <-outC
	want := "arg1 arg2 arg3\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestEcho4(t *testing.T) {
	os.Args = []string{"program", "arg1", "arg2", "arg3"}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	echo4()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	got := <-outC
	want := "program arg1 arg2 arg3\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestEcho5(t *testing.T) {
	os.Args = []string{"program", "arg1", "arg2", "arg3"}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	echo5()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	got := <-outC
	want := "0 program\n1 arg1\n2 arg2\n3 arg3\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestEcho6Optimal(t *testing.T) {
	os.Args = []string{"program", "arg1", "arg2", "arg3"}

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	echo6()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	got := <-outC
	want := "programarg1arg2arg3\n"

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
