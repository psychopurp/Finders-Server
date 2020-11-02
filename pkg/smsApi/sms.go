package smsApi

import (
	"finders-server/global"
	"finders-server/st"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendSMS(phone, code string) (err error) {
	var (
		client   *dysmsapi.Client
		response *dysmsapi.SendSmsResponse
	)
	client, err = dysmsapi.NewClientWithAccessKey("cn-hangzhou", global.CONFIG.SMSConfig.AccessKey, global.CONFIG.SMSConfig.Secret)
	if err != nil {
		st.Debug("sms get client error", err.Error())
		return
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = global.CONFIG.SMSConfig.SignName
	request.TemplateCode = global.CONFIG.SMSConfig.TemplateCode
	//request.TemplateParam = "{\"code\": \"111\"}"
	request.TemplateParam = fmt.Sprintf("{\"code\": \"%s\"}", code)
	response, err = client.SendSms(request)
	if err != nil {
		st.Debug("sms error", err.Error())
		st.Debug("sms error and response is", response)
		return
	}

	//st.Debug("sms response is %#v", response)
	if response.Code != "OK" {
		if response.Code == "isv.DOMESTIC_NUMBER_NOT_SUPPORTED" {
			return fmt.Errorf("国际/港澳台消息模板不支持发送境内号码")
		} else if response.Code == "isv.OUT_OF_SERVICE	" {
			global.LOG.Errorf("sms 余额不足")
			return fmt.Errorf("短信发送失败")
		} else if response.Code == "isv.MOBILE_NUMBER_ILLEGAL" {
			return fmt.Errorf("	非法手机号")
		} else {
			return fmt.Errorf("短信发送失败")
		}
	}
	//fmt.Printf("response is %#v\n", response)
	return
}
