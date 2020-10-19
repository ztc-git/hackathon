package util

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
)

func SendSms(phone string, captcha string) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", None, None)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = "岛遇"
	request.TemplateCode = "SMS_187105016"
	request.TemplateParam = "{\"code\": " + captcha +"}"

	response, err := client.SendSms(request)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("response is %#v\n", response)

}
