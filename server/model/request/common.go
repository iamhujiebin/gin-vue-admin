package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Order    string `json:"order" form:"order"`
	OrderBy  string `json:"orderBy" form:"orderBy"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
