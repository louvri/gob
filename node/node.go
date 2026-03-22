package node

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	timestamp int64
	once      sync.Once
)

func GetID() (*string, error) {
	once.Do(func() {
		timestamp = time.Now().UTC().UnixMicro()
	})

	content, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%s-%d", strings.TrimSpace(string(content)), timestamp)
	return &result, nil
}
