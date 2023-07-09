package node

import (
	"fmt"
	"os"
	"time"
)

var timestamp int64

func GetID() (*string, error) {
	if timestamp == 0 {
		timestamp = time.Now().UTC().UnixMicro()
	}

	content, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		return nil, err
	}
	result := string(content)
	result = result + fmt.Sprintf("-%d", timestamp)
	return &result, nil
}
