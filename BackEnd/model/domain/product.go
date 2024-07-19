package domain

type Product struct {
	Id          int
	Name        string
	Image       string
	Description string
	Price       int
	CategoryId  int
	Category    Category
	Quantity    int
	Slug        string
	Images      []Image
}
