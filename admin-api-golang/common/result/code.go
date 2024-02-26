// 状态码

package result

// Codes 定义状态
type Codes struct {
	SUCCESS         uint
	FAILED          uint
	Message         map[uint]string
	NOAUTH          uint
	AUTHFORMATERROR uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:         200,
	FAILED:          501,
	NOAUTH:          403,
	AUTHFORMATERROR: 405,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:         "成功",
		ApiCode.FAILED:          "失败",
		ApiCode.NOAUTH:          "请求头中的auth为空",
		ApiCode.AUTHFORMATERROR: "请求头中auth格式有误",
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
