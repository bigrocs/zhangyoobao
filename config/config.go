package config

// 服务器URL ： https://proxy.szzt.com.cn/cs-api
// accessId : ISKJF459JD9FGU34
// accessKey : KLSKF3OD4RU3CDU9M3VOM39384VD35
// productKey : a1Z0BXAK0jS

type Config struct {
	AgentId string `json:"agentId"` // 开发者ID
	Key     string `json:"key"`     // 开发者密钥
	Sandbox bool   `json:"sandbox"` // 沙盒
}
