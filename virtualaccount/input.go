package virtualaccount

type Post struct {
	Title string      `json:"title" binding:"required"`
	Harga interface{} `json:"harga" binding:"required,number"`
}
