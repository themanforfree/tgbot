package jwglxt

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req"
)

func getCsrftoken() (string, error) {
	res, err := req.Get(Config.LoginUrl, Config.Header)
	if err != nil {
		return "", err
	}
	if res.Response().StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.Response().StatusCode, res.Response().Status)
	}
	dom, err := goquery.NewDocumentFromReader(res.Response().Body)
	if err != nil {
		return "", err
	}
	var csrftoken string
	dom.Find("#csrftoken").Each(func(i int, s *goquery.Selection) {
		csrftoken = s.AttrOr("value", "")
	})
	if csrftoken == "" {
		return "", fmt.Errorf("csrftoken is empty")
	}
	return csrftoken, nil
}

func getRsa(pwd, n, e string) (string, error) {
	message := []byte(pwd)
	rsa_n, err := base64.StdEncoding.DecodeString(n)
	if err != nil {
		return "", err
	}
	rsa_e, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		return "", err
	}
	int_e := big.NewInt(0).SetBytes(rsa_e)
	int_n := big.NewInt(0).SetBytes(rsa_n)
	key := rsa.PublicKey{N: int_n, E: int(int_e.Int64())}
	encropy_pwd, err := rsa.EncryptPKCS1v15(rand.Reader, &key, message)
	if err != nil {
		return "", err
	}
	result := base64.StdEncoding.EncodeToString(encropy_pwd)
	return result, nil
}

func getXnmXqm() (int, int) {
	time_now := time.Now()
	if time_now.Month() < time.March {
		return time_now.Year() - 1, 1
	} else if time_now.Month() < time.September {
		return time_now.Year() - 1, 2
	} else {
		return time_now.Year(), 1
	}
}
