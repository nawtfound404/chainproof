package ipfs

import (
	"bytes"
	"testing"
)

func TestIPFSUpload(t *testing.T) {
	client := New("http://localhost:5001")

	data := []byte(`{"a":1,"b":2}`)

	cid, err := client.Upload(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("CID:", cid)
}
func TestIPFSRoundTrip(t *testing.T) {
	client := New("http://localhost:5001")

	data := []byte(`{"a":1,"b":2}`)

	cid, err := client.Upload(data)
	if err != nil {
		t.Fatal(err)
	}

	fetched, err := client.Fetch(cid)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, fetched) {
		t.Fatal("IPFS roundtrip mismatch")
	}

	t.Log("Original:", string(data))
	t.Log("Fetched :", string(fetched))

}
