package forms

type CreateOrderForm struct {
	Address string `form:"address" json:"address" binding:"required"`
	Name    string `form:"name" json:"name" binding:"required"`
	Mobile  string `form:"mobile" json:"mobile" binding:"required"`
	Post    string `form:"post" json:"post" binding:"required"`
}
