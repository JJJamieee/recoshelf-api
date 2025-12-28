package requests

type BatchDeleteReleaseRequest struct {
	IDs []int64 `json:"ids" validate:"required,min=1,max=100,dive,gt=0"`
}
