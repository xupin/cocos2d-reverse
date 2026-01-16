package main

import (
	"os"
	"testing"

	"github.com/xupin/cocos2d-reverse/utils/xor"
)

func TestXOR(t *testing.T) {
	sign := "VZ7sx4MGw4"
	key := "CeAcMKM6QveeNpymbm2v"

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
	dec := xor.Crypt(bs, []byte(key))

	if len(dec) == 0 {
		t.Fatalf("decrypt result is empty")
	}

	if err := os.WriteFile(dstPath, dec, os.ModePerm); err != nil {
		t.Fatalf("write decrypted file failed: %v", err)
	}
}
