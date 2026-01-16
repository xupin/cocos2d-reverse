package main

import (
	"os"
	"testing"

	"github.com/xupin/cocos2d-reverse/utils/xxtea"
)

func TestXXTEA(t *testing.T) {
	sign := "sm!fd9%3nX4@"
	key := "gb4sZzoV/fP2P)t3"

	srcPath := "./assets/9322.png"
	dstPath := "./assets/9322_dec.png"

	bs, err := os.ReadFile(srcPath)
	if err != nil {
		t.Fatalf("read source file failed: %v", err)
	}

	if len(bs) <= len(sign) {
		t.Fatalf("file too small, invalid encrypted content")
	}

	bs = bs[len(sign):]
	dec := xxtea.Decrypt(bs, []byte(key))

	if len(dec) == 0 {
		t.Fatalf("decrypt result is empty")
	}

	if err := os.WriteFile(dstPath, dec, os.ModePerm); err != nil {
		t.Fatalf("write decrypted file failed: %v", err)
	}
}
