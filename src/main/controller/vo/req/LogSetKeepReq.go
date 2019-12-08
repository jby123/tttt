package req

type LogSetKeepReq struct {
	Value int `json:"value" binding:"required"`
}
