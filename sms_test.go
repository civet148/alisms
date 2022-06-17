package alisms

import (
	"github.com/civet148/log"
	"testing"
)

func TestSmsClient_SendAuthCode(t *testing.T) {
	var sc = &SmsConfig{
		Domain:          "dysmsapi.aliyuncs.com",
		AccessKeyID:     "gjT9Ghco68xA0r9FSPc1g",
		AccessKeySecret: "LTAI5t9Xiwbs3Qvk4Cf336h5",
	}
	var smsTmp = &TemplateConfig{
		SignName:      "ali-cloud",
		TemplateCode:  "SMS_205324219",
		TemplateParam: `{"code":"%s"}`, //您的验证码为：${code}，请勿泄露于他人
	}
	sms, err := NewAliSms(sc)
	if err != nil {
		log.Errorf("%s", err)
		return
	}
	sms.Send("18182371693", smsTmp, "123456")
}

