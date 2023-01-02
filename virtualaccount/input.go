package virtualaccount

// / POST DATA & JSON Validation & Show Error
type Post struct {
	Title string      `json:"title" binding:"required"`
	Harga interface{} `json:"harga" binding:"required,number"`
}
