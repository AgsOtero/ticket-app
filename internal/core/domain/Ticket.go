package domain

type Ticket struct {
	ID      int64
	EventID int64
	SeatID  int64
	BuyID   *int64
	Price   float32
	Status  string
}
