package request

type GoodsAdd struct {
	Name  string  `form:"name" binding:"required"`
	Price float32 `form:"price"  binding:"required"`
	Num   int64   `form:"num"  binding:"required"`
	Color string  `form:"color"  binding:""`
	Sign  string  `form:"sign" binding:""`
}
type Login struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type GetGoods struct {
	Id int `form:"id" json:"id" xml:"id" binding:"required"`
}
