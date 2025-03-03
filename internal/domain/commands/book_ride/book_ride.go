package book_ride

func NewRideBooker() *RideBooker {
	return &RideBooker{}
}

type RideBooker struct{}

func (b *RideBooker) Execute() string {
	return "ride"
}
