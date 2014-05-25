package redisurlparser_test

import (
	redisurlparser "github.com/scottmotte/redisurlparser"
	"testing"
)

const REDIS_URL = "redis://redistogo:64cfea0093507536a374ab6ad40f8263@angelfish.redistogo.com:10039/"

func TestParse(t *testing.T) {
	result, err := redisurlparser.Parse(REDIS_URL)
	if err != nil {
		t.Errorf("Error", err)
	}
	if result.Password != "64cfea0093507536a374ab6ad40f8263" {
		t.Errorf("Password incorrect")
	}
	if result.Username != "redistogo" {
		t.Errorf("Username incorrect")
	}
	if result.Host != "angelfish.redistogo.com" {
		t.Errorf("Host incorrect")
	}
	if result.Port != "10039" {
		t.Errorf("Port incorrect")
	}
}
