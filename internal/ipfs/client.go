package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)



type Client struct{
	Endpoint string
}

func New(endpoint string) *Client {
	return &Client{Endpoint: endpoint}

}

func (c *Client) Upload(data []byte) (string, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	
	part, err := writer.CreateFormFile("file", "proof.json")
	if err != nil {
		return "", err
	}

	_, err = part.Write(data)
	if err != nil {
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", c.Endpoint+"/api/v0/add", &body)
	if err != nil {
		return "", nil
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var result struct{
		Hash string
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("Invalid IPFS response: %s", string(respBody))

	}
	return result.Hash, nil
}

func (c *Client) Fetch(cid string) ([]byte, error) {
	resp, err := http.Post(
		c.Endpoint+"/api/v0/cat?arg="+cid,
		"",
		nil,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
