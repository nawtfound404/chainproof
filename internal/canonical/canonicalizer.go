package canonical

import (
	"bytes"
	"encoding/json"
	"sort"
)

func CanonicalizeJSON(input []byte) ([]byte, error) {
	var data interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := encodeCanonical(buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func encodeCanonical(buf *bytes.Buffer, v interface{}) error {
	switch val := v.(type) {

	case map[string]interface{}:
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		buf.WriteByte('{')
		for i, k := range keys {
			if i > 0 {
				buf.WriteByte(',')
			}

			keyBytes, _ := json.Marshal(k)
			buf.Write(keyBytes)
			buf.WriteByte(':')

			if err := encodeCanonical(buf, val[k]); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	case []interface{}:
		buf.WriteByte('[')
		for i, item := range val {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encodeCanonical(buf, item); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	default:
		valueBytes, err := json.Marshal(val)
		if err != nil {
			return err
		}
		buf.Write(valueBytes)
	}

	return nil
}
