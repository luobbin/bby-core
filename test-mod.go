package testmod

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"crypto/rand"
	"github.com/google/uuid"
)

// UniqueId2 返回一个"unix时间和n*2长度的唯一码"字符串
func UniqueId2(n uint16) string {
	b := make([]byte, n) //8 character
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	t := time.Now().Unix()
	return strconv.FormatInt(t, 10) + "-" + hex.EncodeToString(b)
}

func UniqueId(n uint16) string {
	b := make([]byte, n) //8 character
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	t := time.Now().Unix()
	return strconv.FormatInt(t, 10) + "-" + hex.EncodeToString(b)
}

func GetUuid(s string) string {
	fmt.Printf("获取请求字符串：%s", s)
	requestId := uuid.New().String()
	return requestId
}
