package models

type Item struct {
	OrderId uint
	Id      uint
	Amount  uint
}
type Order struct {
	Id          int
	PaymentType int
	Items       []Item
}
