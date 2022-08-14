package toolkit

import "testing"

// func name - starts with Test, then the Receiver Tools, then _, then the name of the Function to Test
func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Error("Incorrect lenth for the generated random string")
	}
}
