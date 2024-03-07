package model

type ExpressionRequest struct {
	Expression string `json:"expression" binding:"required"`
	Uuid       string `json:"uuid" binding:"required"`
}

type ExpressionResultRequest struct {
	RequestId string `json:"request_id" binding:"required"`
}
