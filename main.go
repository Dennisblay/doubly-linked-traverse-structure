package main

func main() {

	stn := Station{
		Coordinates: &Point{
			X:                2000,
			Y:                4000,
			Z:                288,
			CoordinateSystem: "Ghana National Grid",
		},
		Distance: 1000,
		unit:     "m",
		next:     nil,
		prev:     nil,
	}
	stn2 := Station{
		Coordinates: &Point{
			X:                2500,
			Y:                7000,
			Z:                300,
			CoordinateSystem: "Ghana National Grid",
		},
		Distance: 1500,
		unit:     "m",
		next:     nil,
		prev:     nil,
	}

	t := Traverse{}
	t.PushStationBack(stn)
	t.PushStationBack(stn2)
	t.traverse()

}
