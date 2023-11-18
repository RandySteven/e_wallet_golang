package utils

import (
	"math/rand"
	"time"
)

func GenerateWalletNumber() string {
	numb := "0123456789"
	randNumber := make([]byte, 9)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 9; i++ {
		randNumber[i] = numb[rand.Intn(len(numb))]
	}
	return "WR" + string(randNumber)
}
