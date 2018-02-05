package logring

import (
	"log"
	"testing"
)

func TestCount(t *testing.T) {
	want := 10
	log.SetOutput(Writer(3))
	for i := 0; i < want; i++ {
		log.Println(i)
	}
	got := Count()
	if got != want {
		t.Errorf("Count() wanted %d but got %d", want, got)
	}
}

func TestRecent(t *testing.T) {
	want := 10
	log.SetOutput(Writer(100))
	for i := 0; i < want; i++ {
		log.Println(i)
	}
	got := len(Recent())
	if got != want {
		t.Errorf("Recent() wanted %d but got %d", want, got)
	}

	want = 10
	log.SetOutput(Writer(want))
	for i := 0; i < 100; i++ {
		log.Println(i)
	}
	got = len(Recent())
	if got != want {
		t.Errorf("Recent() wanted %d but got %d", want, got)
	}
}
