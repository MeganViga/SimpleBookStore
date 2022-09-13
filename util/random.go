package util

import (
	"math/rand"
	"strings"
	"time"
)


func randInt(min,max int)int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max - min +1 ) + min)
}


func randString(length int)string{
	var sb strings.Builder
	letterBytes := "abcdefghijklmnopqrstABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i:=0;i<length;i++{
		rand.Seed(time.Now().UnixNano())
		c := letterBytes[rand.Intn(len(letterBytes))]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomPrice()float64{
	return float64(randInt(0,1000))
}

func RandomBookName()string{
	return randString(10)
}

func RandomUserName()string{
	return randString(32)
}
func RandomUserEmail()string{
	name := randString(5)
	return name + "@gmail.com"
}
func RandomQuantity()int32{
	return int32(rand.Intn(10))
}