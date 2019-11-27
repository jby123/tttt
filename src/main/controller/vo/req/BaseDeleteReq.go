package req

type BaseDeleteReq struct {
	Ids string `json:"ids" binding:"required"`
}
