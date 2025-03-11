package trip_scanner

type FakeTripScanner struct {
	distance int
}

func NewFakeTripScanner(distance int) *FakeTripScanner {
	return &FakeTripScanner{distance: distance}
}

func (s *FakeTripScanner) DistanceBetween(departure string, arrival string) int {
	return s.distance
}

func (s *FakeTripScanner) InParis(address string) bool {
	return address == "PARIS_ADDRESS"
}
