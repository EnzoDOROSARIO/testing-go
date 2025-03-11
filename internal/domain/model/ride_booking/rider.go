package ride_booking

type Rider struct {
	id          string
	currentRide Ride
}

type RiderSnapshot struct {
	ID          string
	CurrentRide RideSnapshot
}

func NewRider(id string) *Rider {
	return &Rider{
		id:          id,
		currentRide: Ride{},
	}
}

func (r *Rider) ToSnapshot() RiderSnapshot {
	return RiderSnapshot{
		ID:          r.id,
		CurrentRide: r.currentRide.ToSnapshot(),
	}
}

func (r *Rider) Book(
	rideId string,
	departure string,
	arrival string,
	price float64,
) {
	ri := BookRide(rideId, r.id, departure, arrival, price)
	r.currentRide = *ri
}

func (r *Rider) ID() string {
	return r.id
}
