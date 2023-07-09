package object

import (
	"github.com/google/uuid"
	"hash/fnv"
	"strings"
)

func GenerateRunningNumbers() uint32 {
	uuidWithHyphen := uuid.New()
	uuidWithoutHyphen := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	h := fnv.New32a()
	_, _ = h.Write([]byte(uuidWithoutHyphen))
	return h.Sum32()
}
