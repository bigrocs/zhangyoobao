package zhangyoobao

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigrocs/zhangyoobao/requests"
	uuid "github.com/satori/go.uuid"
)

func TestSendVoiceMsg(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AgentId = os.Getenv("zhangyoobao_AgentId")
	client.Config.Key = os.Getenv("zhangyoobao_Key")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "sendVoiceMsg"
	request.BizContent = map[string]interface{}{
		"deviceName": os.Getenv("zhangyoobao_deviceName"),
		"amount":     "1922",
		"orderId":    uuid.NewV4().String(),
		"prefix":     "39",
		"random":     uuid.NewV4().String(),
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	r, err := response.GetVerifySignDataMap()
	fmt.Println("TestPlay", r, err)
	t.Log(r, err, "|||")
}

func TestBinDevice(t *testing.T) {
	// 创建连接
	client := NewClient()
	client.Config.AgentId = os.Getenv("zhangyoobao_AgentId")
	client.Config.Key = os.Getenv("zhangyoobao_Key")
	client.Config.Sandbox = false
	// 配置参数
	request := requests.NewCommonRequest()
	request.ApiName = "binDevice"
	request.BizContent = map[string]interface{}{
		"deviceType": "17",
		"deviceName": os.Getenv("zhangyoobao_deviceName"),
		"amount":     "192",
		"orderId":    "asadasas",
		"prefix":     "5",
		"random":     uuid.NewV4().String(),
	}
	// 请求
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
	}
	r, err := response.GetVerifySignDataMap()
	fmt.Println("TestPlay", r, err)
	t.Log(r, err, "|||")
}
