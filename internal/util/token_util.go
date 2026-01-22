package util

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateToken(totalLength int) string {
	const tsLen = 13 // Unix millisecond length

	if totalLength < tsLen {
		return ""
	}

	// 1. 毫秒时间戳（13 位）
	ts := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 2. 需要补足的字符长度
	remain := totalLength - tsLen

	// hex 每个字节 = 2 chars
	byteLen := (remain + 1) / 2

	randBytes := make([]byte, byteLen)
	if _, err := rand.Read(randBytes); err != nil {
		return ""
	}

	randomHex := hex.EncodeToString(randBytes)[:remain]

	// 3. 拼接
	return ts + randomHex
}
