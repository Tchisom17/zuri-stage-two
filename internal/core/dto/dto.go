package dto

type OperationRequest struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
}
type OperationResponse struct {
	SlackUsername string `json:"slackUsername"`
	OperationType string `json:"operation_type"`
	Result        int    `json:"result"`
}
