package entity

//Book entity
type Book struct {
	Title  string  `json:"title"`
	Gender string  `json:"gender"`
	Pages  int16   `json:"pages"`
	Price  float64 `json:"price"`
}
