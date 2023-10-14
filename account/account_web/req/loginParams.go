package req

type LoginByPassword struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
