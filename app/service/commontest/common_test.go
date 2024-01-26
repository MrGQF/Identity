package commontest

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gitee.com/QunXiongZhuLu/KongMing/crypto"
	"math/big"
	"net/url"
	"strings"
	"testing"
)

type Data struct {
	Val string
}

func TestRandNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(1))
		fmt.Println(result)
	}
}

func TestPost(t *testing.T) {
	var (
		val url.Values
	)

	val = make(map[string][]string)
	val.Add("param1", "123")
	val.Add("param2", "true")
	_ = strings.NewReader(val.Encode())
	fmt.Println(val.Encode())
}

func TestRsaDecrop(t *testing.T) {
	privateKeyByte, publicKeyByte, err := crypto.CreateKeys(1024)
	if err != nil {
		return
	}

	// 加密

	cip, err := crypto.RsaEncrypt("test300033", publicKeyByte)
	val := base64.StdEncoding.EncodeToString(cip)
	fmt.Println("Encode cip", val)

	byte, err := base64.StdEncoding.DecodeString(val)
	plaint, err := crypto.RsaDecrypto(byte, privateKeyByte)
	if err != nil {
		fmt.Println("RsaDecrypto err", err)
	}
	fmt.Println(plaint)
}
