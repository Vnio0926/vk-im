package request

type MessageRequest struct {
	MessageType int32  `json:"messageType"`
	FromId      int32  `json:"from_id" binding:"required"`
	ToId        int32  `json:"to_id" binding:"required"`
	Uuid        string `json:"uuid" binding:"required"`
}

//binding 字段用于指定请求参数的验证规则。Gin 框架会根据这些规则自动验证请求参数，并在参数不符合规则时返回错误响应
