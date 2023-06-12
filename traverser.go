package main

import "fmt"

// Point defines a 3D coordinate system
type Point struct {
	X, Y, Z                                      float64
	CoordinateSystem, ellipsoid, Datum, epsgCode string
}

// Station defines Traverse station attributes in a traverse
type Station struct {
	stationName string
	HCRs        [4]float64
	Coordinates *Point

	Distance float64
	unit     string
	next     *Station
	prev     *Station
}

func (s *Station) String() string {
	return fmt.Sprintf("|HCRs: %v, |X: %v, |Y:  %v, |Distance: %v, |Unit: %v \n⬇\n", s.HCRs, s.Coordinates.X, s.Coordinates.Y, s.Distance, s.unit)
}

// Traverse defines the start and end of the traverse
type Traverse struct {
	head *Station
	tail *Station
	size int64
}

// IsEmpty Checks if Traverse Link is empty
func (t *Traverse) IsEmpty() bool {
	return t.head == nil

}

// PushStationFront inserts a Station in front of a traverse link
func (t *Traverse) PushStationFront(data Station) {
	newStation := &data
	if t.IsEmpty() {
		t.head = newStation
		t.tail = newStation
		t.size++
		return
	} else {
		newStation.next = t.head
		t.head.prev = newStation
		t.tail.prev = t.head
		t.head = newStation
		t.size++
	}
}

// PushStationBack inserts a Station at the back of a traverse link
func (t *Traverse) PushStationBack(data Station) {
	newStation := &data
	if t.IsEmpty() {
		t.head = newStation
		t.tail = newStation
		t.size++
		return
	}
	newStation.prev = t.tail
	t.tail.next = newStation
	t.tail = newStation
	t.size++
}

// PopStationBack removes a Station from the end of a traverse link and returns it
func (t *Traverse) PopStationBack() *Station {
	if t.IsEmpty() {
		return nil
	}
	StationToPop := t.tail
	if StationToPop == t.head {
		t.head = nil
		t.tail = nil
		t.size--

	} else {
		t.tail = StationToPop.prev
		t.tail.next = nil
		t.size--
	}
	StationToPop.next = nil
	StationToPop.prev = nil

	return StationToPop

}

// PopStationFront removes a Station from the beginning of a traverse link and returns it
func (t *Traverse) PopStationFront() *Station {
	if t.IsEmpty() {
		return nil
	}
	StationToPop := t.head
	if StationToPop == t.tail {
		t.head = nil
		t.tail = nil
		t.size--

	} else {
		t.head = StationToPop.next
		t.head.prev = nil
		t.size--
	}
	StationToPop.next = nil
	StationToPop.prev = nil

	return StationToPop
}

// traverse is used to traverse the entire link from start to end
func (t *Traverse) traverse() {
	if t.IsEmpty() {
		return
	}
	current := t.head

	for current != nil {
		//fmt.Printf("%#v\n", current)
		fmt.Println(current)
		current = current.next
	}

}

// Front returns the first Station in the link
func (t *Traverse) Front() *Station {
	if t.IsEmpty() {
		return nil
	}
	return t.head
}

// Back returns the end Station in the link
func (t *Traverse) Back() *Station {
	if t.IsEmpty() {
		return nil
	}
	return t.head
}

func (t *Traverse) ComputeBearing() {

}
