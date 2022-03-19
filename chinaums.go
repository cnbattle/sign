package sign

import (
	"errors"
	"github.com/cnbattle/sign/helper"
	"sort"
	"strings"
)

// Chinaums 银联商务
type Chinaums struct {
}

// CheckSign 验签
func (c *Chinaums) CheckSign(signKey string, params map[string]string) (bool, error) {
	originalSign := params["sign"]
	delete(params, "sign")
	sign, err := c.Sign(signKey, params)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(originalSign, sign), nil
}

// Sign 签名
func (c *Chinaums) Sign(signKey string, params map[string]string) (string, error) {
	//生成待签字串 和  sign
	signString, err := c.buildSignString(params)
	if err != nil {
		return "", err
	}
	newSignString := signString + signKey
	return strings.ToUpper(helper.Sha256(newSignString)), nil
}

func (c *Chinaums) buildSignString(params map[string]string) (string, error) {
	if len(params) == 0 {
		return "", errors.New("params为空")
	}

	_, ok := params["sign"]
	if ok {
		delete(params, "sign")
	}

	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var str string
	for i, key := range keys {
		value := params[key]
		if len(value) == 0 {
			continue
		}

		if i == len(keys)-1 {
			str += key + "=" + value
		} else {
			str += key + "=" + value + "&"
		}
	}
	return str, nil
}
