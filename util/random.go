package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano()) //纳秒
}

//随机Int64
func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//随机string

func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}



func RandOwner() string {
	return RandString(6)
}

//随机金额
func RandMoney() int64 {
	return RandInt(0, 1000)
}

func RandCurrency() string {
	currencies := []string{"USD", "EUR", "RMB", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
