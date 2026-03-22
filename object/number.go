package object

import (
	"hash/fnv"

	"github.com/google/uuid"
)

func GenerateRunningNumbers() uint32 {
	id := uuid.New()
	h := fnv.New32a()
	_, _ = h.Write(id[:])
	return h.Sum32()
}
