package request

type SearchMusicParams struct {
	PageInfo
	Keyword  string `json:"keyword"`
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
	UserId   uint   `json:"userId"`
}
