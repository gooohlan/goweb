package callback

type ResponseType int

const (
	// 空响应
	ResponseUnKnow ResponseType = -1

	// 正常响应
	ResponseOK ResponseType = 0

	// 正常处理，但是没有数据
	ResponseNoData ResponseType = 90100

	// 正常处理，但数据库无响应
	ResponseDataNotResp ResponseType = 90400

	// 常规错误
	ResponseError ResponseType = 90000

	// 响应异常
	ResponseCrash ResponseType = 90500

	// 参数验证失败
	ResponseRequest ResponseType = 90900
)
