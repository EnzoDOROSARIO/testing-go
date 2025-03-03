package ride_booking

type Ride struct {
	id string
}

type RideSnapshot struct {
	ID string
}

func NewRide(id string) *Ride {
	return &Ride{id: id}
}

func (r *Ride) ToSnapshot() RideSnapshot {
	return RideSnapshot{ID: r.id}
}
