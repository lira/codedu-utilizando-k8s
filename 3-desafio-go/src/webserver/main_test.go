package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("Teste")
	response := "<b>Teste</b>"

	if result != response {
		t.Errorf("Invalid result! :( return %s, wanted %s", result, response)
	}
}
