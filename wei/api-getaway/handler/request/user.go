package request

type User struct {
	Name    string `form:"name"  binding:"required"`
	Age     int    `form:"age"  binding:"required"`
	Address string `form:"address"  binding:"required"`
}
type Login struct {
	Name string `form:"name"  binding:"required"`
	Age  int    `form:"age"  binding:"required"`
}
type Update struct {
	Id      int    `form:"id"  binding:"required"`
	Address string `form:"address"  binding:"required"`
}
