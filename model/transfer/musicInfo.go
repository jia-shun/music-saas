package transfer

type MusicInfo struct {
	ID           uint    `json:"id"`
	UserID       uint    `json:"userId"`
	MusicName    string  `json:"musicName"`
	CustomerName string  `json:"customerName"`
	Price        float64 `json:"price"`
	PayStatus    bool    `json:"payStatus"`
	BeganAt      string  `json:"beganAt"`
	FinishedAt   string  `json:"finishedAt"`
	FinishStatus bool    `json:"finishStatus"`
}
