package ride_booking

type Ride struct {
	id        string
	riderId   string
	departure string
	arrival   string
	price     float64
	status    string
}

type RideSnapshot struct {
	ID        string
	RiderId   string
	Departure string
	Arrival   string
	Price     float64
	Status    string
}

func Book(id string, riderId string, departure string, arrival string, price float64) *Ride {
	return &Ride{
		id:        id,
		riderId:   riderId,
		departure: departure,
		arrival:   arrival,
		price:     price,
		status:    "WAITING_FOR_DRIVER",
	}
}

func (r *Ride) ToSnapshot() RideSnapshot {
	return RideSnapshot{
		ID:        r.id,
		RiderId:   r.riderId,
		Departure: r.departure,
		Arrival:   r.arrival,
		Price:     r.price,
		Status:    r.status,
	}
}
