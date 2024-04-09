package main

import (
	"os"

	"github.com/xupin/cocos2d-reverse/utils/xxtea"
)

func main() {
	sign := "sm!fd9%3nX4@"
	key := "gb4sZzoV/fP2P)t3"

	bs, err := os.ReadFile("./assets/res/9322.png")
	if err != nil {
		panic(err)
	}
	bs = bs[len(sign):]
	nBs := xxtea.Decrypt(bs, key)
	os.WriteFile("./assets/res/9322_decry.png", nBs, os.ModePerm)
}
