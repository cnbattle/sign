package sign

import (
	"github.com/cnbattle/sign/helper"
	"sort"
	"strings"
)

// Chinaums 银联商务
type Chinaums struct {
}

// CheckSign 验签
func (c *Chinaums) CheckSign(md5Key string, params map[string]string) (bool, error) {
	originalSign := params["sign"]
	delete(params, "sign")
	sign, err := c.Sign(md5Key, params)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(originalSign, sign), nil
}

// Sign 签名
func (c *Chinaums) Sign(md5Key string, params map[string]string) (string, error) {
	//生成待签字串 和  sign
	signString, err := c.buildSignString(params)
	if err != nil {
		return "", err
	}
	newSignString := signString + md5Key
	return strings.ToUpper(helper.Sha256(newSignString)), nil
}

func (c *Chinaums) buildSignString(params map[string]string) (string, error) {
	if len(params) == 0 {
		return "", nil
	}

	var keys []string
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var str string
	for i, key := range keys {
		value := params[key]
		if i == len(keys)-1 {
			str += key + "=" + value
		} else {
			str += key + "=" + value + "&"
		}
	}
	return str, nil
}
