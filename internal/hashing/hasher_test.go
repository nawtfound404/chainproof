package hashing

import "testing"

func TestHash_Stable(t *testing.T) {
	data := []byte(`{"a":1,"b":2}`)

	hash1 := HashSHA256(data)
	hash2 := HashSHA256(data)

	if hash1 != hash2 {
		t.Fatal("hash not stable")
	}
}
