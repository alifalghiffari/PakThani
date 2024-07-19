package domain

// Order mendefinisikan sebuah pesanan
type Order struct {
	Id            int           // ID unik untuk pesanan
	UserId        int           // ID pengguna yang membuat pesanan
	TotalItems    int           // Jumlah total item dalam pesanan
	TotalPrice    int           // Harga total dari pesanan
	OrderStatus   OrderStatus   // Status pesanan (misalnya: pending, diproses, selesai, dll.)
	PaymentStatus PaymentStatus // Status pembayaran pesanan (misalnya: pending, berhasil, gagal, dll.)
	OrderItems    []OrderItems  // Item dalam pesanan
	User          User
	Address       Address
}

// OrderStatus mendefinisikan status pesanan
type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
	Cancelled  OrderStatus = "Cancelled"
)

// PaymentStatus mendefinisikan status pembayaran pesanan
type PaymentStatus string

const (
	PaymentPending PaymentStatus = "Pending"
	PaymentSuccess PaymentStatus = "Success"
	PaymentFailed  PaymentStatus = "Failed"
)
