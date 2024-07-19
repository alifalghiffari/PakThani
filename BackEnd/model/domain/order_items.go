package domain

// OrderItems mendefinisikan item dalam pesanan
type OrderItems struct {
	Id        int // ID unik untuk item pesanan
	OrderId   int // ID pesanan yang berisi item ini
	ProductId int // ID produk yang dibeli
	Quantity  int // Jumlah produk yang dibeli
	Price     int // Harga produk per item
	Product   Product
}
