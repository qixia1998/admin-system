// 状态码

package result

// Codes 定义状态
type Codes struct {
	SUCCESS uint
	FAILED  uint
	Message map[uint]string
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS: 200,
	FAILED:  501,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS: "成功",
		ApiCode.FAILED:  "失败",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
