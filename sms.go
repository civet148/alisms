package alisms

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/civet148/log"
	"strings"
)


type AliSms struct {
	c *dysmsapi20170525.Client
}

type SmsConfig struct {
	Domain          string
	AccessKeyID     string
	AccessKeySecret string
}

type TemplateConfig struct {
	SignName      string //signature name
	TemplateCode  string //template code
	TemplateParam string //template parameters
}

func NewAliSms(cfg *SmsConfig) (sms *AliSms, err error) {
	config := &openapi.Config{
		Endpoint: tea.String(cfg.Domain),
		AccessKeyId: tea.String(cfg.AccessKeyID),
		AccessKeySecret: tea.String(cfg.AccessKeySecret),
	}
	c := &dysmsapi20170525.Client{}
	c, err = dysmsapi20170525.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &AliSms{
		c: c,
	}, nil
}


func (s *AliSms) Send(strPhone string, tc *TemplateConfig, args...interface{}) (err error) {
	strTemplateParam := fmt.Sprintf(tc.TemplateParam, args...)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(tc.SignName),
		TemplateCode:  tea.String(tc.TemplateCode),
		PhoneNumbers:  tea.String(strPhone),
		TemplateParam: tea.String(strTemplateParam),
	}
	r, err := s.c.SendSms(sendSmsRequest)
	if err != nil {
		return log.Errorf("sms response error [%s]", err)
	}
	if !strings.EqualFold(*r.Body.Code, "OK") {
		return log.Errorf("sms response code [%s] error [%s]", *r.Body.Code, r.Body.Message)
	}
	return
}

func (s *AliSms) Close() {
}