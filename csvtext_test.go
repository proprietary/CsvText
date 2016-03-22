package csvtext

import (
	"testing"
)

func TestSetArr(t *testing.T) {
	a := new(CsvText).SetArr([]string{
		"lol", "wut", "hi",
	})
	e := "lol,wut,hi"
	if a.Csv != e {
		t.Errorf("Expected\t%s\nGot\t\t%s\n", e, a.Csv)
	}
}

func TestSetCsv(t *testing.T) {
	b := new(CsvText).SetCsv("lol,wut, hi, loser")
	e := []string{"lol", "wut", "hi", "loser"}
	for i, _ := range e {
		if e[i] != b.Arr[i] {
			t.Errorf("Expected\t%s\nGot\t\t%s\n", e, b.Arr)
		}
	}
}
