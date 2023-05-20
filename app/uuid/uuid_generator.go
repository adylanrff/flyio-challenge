package uuid

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const UUIDLength = 32

func GenerateUUID(nodeID string) string {
	epoch := time.Now().UnixNano()
	rand.Seed(epoch)
	b := make([]byte, UUIDLength+2)
	rand.Read(b)
	randStr := fmt.Sprintf("%x", b)[2 : UUIDLength+2]
	hash := md5.Sum([]byte(fmt.Sprintf("%s-%s", randStr, nodeID)))
	return hex.EncodeToString(hash[:])
}
