package usecase

import "testing"

func TestGenerateSeats(t *testing.T) {
	seats := generateSeats(2, 3)

	expected := 6

	if len(seats) != expected {
		t.Errorf("expected %d seats, got %d", expected, len(seats))
	}

	if seats[0].Row != "A" || seats[0].Number != 1 {
		t.Errorf(
			"expected first seat to be row A, number 1; got row %s and number seat %d",
			seats[0].Row,
			seats[0].Number,
		)
	}

	lastSeat := seats[len(seats)-1]

	if lastSeat.Row != "B" || lastSeat.Number != 3 {
		t.Errorf(
			"expected last seat to be row B, number 3; got row %s and number seat %d",
			lastSeat.Row,
			lastSeat.Number,
		)
	}
}
