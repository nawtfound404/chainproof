package canonical

import (
	"bytes"
	"testing"
)

func TestCanonicalize_orderIndependence(t *testing.T) {
	json1 := []byte(`{"b":2,"a":1}`)
	json2 := []byte(`{"a":1,"b":2}`)

	c1, _ := CanonicalizeJSON(json1)
	c2, _ := CanonicalizeJSON(json2)

	t.Log("C1:", string(c1))
	t.Log("C2:", string(c2))

	if !bytes.Equal(c1, c2) {
		t.Fatal("canonicalization mismatched")
	}
}
