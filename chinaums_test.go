package sign

import (
	"strings"
	"testing"
)

func TestChinaums_CheckSign(t *testing.T) {
	md5Key := "cnbattle"
	params := map[string]string{
		"walletOption":     "SINGLE",
		"billNo":           "31940000201700002",
		"billDate":         "2017-06-26",
		"requestTimestamp": "2017-06-26 17:28:02",
		"instMid":          "QRPAYDEFAULT",
		"msgSrc":           "WWW.TEST.COM",
		"totalAmount":      "1",
		"msgType":          "bills.getQRCode",
		"mid":              "898340149000005",
		"tid":              "88880001",
		"signType":         "SHA256",
		"sign":             "7F799A10933D94BDEF6675F5DA30C58B8EADF1E8FB4ACA7A18B0E1384BBD7D7F",
	}

	c := &Chinaums{}
	got, err := c.CheckSign(md5Key, params)
	if err != nil {
		t.Fatal(err)
	}
	if !got {
		t.Fatal("验证失败")
	}
}

func TestChinaums_Sign(t *testing.T) {
	md5Key := "cnbattle"
	sign := "7F799A10933D94BDEF6675F5DA30C58B8EADF1E8FB4ACA7A18B0E1384BBD7D7F"

	params := map[string]string{
		"walletOption":     "SINGLE",
		"billNo":           "31940000201700002",
		"billDate":         "2017-06-26",
		"requestTimestamp": "2017-06-26 17:28:02",
		"instMid":          "QRPAYDEFAULT",
		"msgSrc":           "WWW.TEST.COM",
		"totalAmount":      "1",
		"msgType":          "bills.getQRCode",
		"mid":              "898340149000005",
		"tid":              "88880001",
		"signType":         "SHA256",
	}

	c := &Chinaums{}
	got, err := c.Sign(md5Key, params)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.EqualFold(got, sign) {
		t.Fatal("签名失败")
	}
}
