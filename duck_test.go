package duck

import (
	"net/http"
	"testing"
)

func TestDuckSet(t *testing.T) {
	req, _ := http.NewRequest("GET", "localhost", nil)
	SetContext(req, "test", "value")
	value := GetContext(req, "test")
	if value.(string) != "value" {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "localhost", nil)
	for i := 0; i < 10; i++ {
		SetContext(req, i, i)
	}

	values := GetAllContext(req)
	c := values[1]

	if c != 1 {
		t.Fail()
	}
}
