package tools

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strings"
)

func GenerateSignature(method, relativePath string, body []byte, timestamp, secretKey string) (string, error) {
	minified, err := MinifyJSON(body)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(minified))
	bodyHashHex := strings.ToLower(hex.EncodeToString(hash[:]))

	raw := method + ":" + relativePath + ":" + bodyHashHex + ":" + timestamp + ":" + secretKey
	finalHash := sha256.Sum256([]byte(raw))

	return hex.EncodeToString(finalHash[:]), nil
}

func MinifyJSON(data []byte) (string, error) {
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return "", err
	}

	sorted := make(map[string]interface{})
	keys := make([]string, 0, len(temp))
	for k := range temp {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sorted[k] = temp[k]
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "") // minify

	if err := enc.Encode(sorted); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}
