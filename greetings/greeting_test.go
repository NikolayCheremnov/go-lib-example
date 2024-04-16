package greetings

import "testing"

func TestGetGreeting(t *testing.T) {
	t.Log(GetGreeting("John Doe"))
}
