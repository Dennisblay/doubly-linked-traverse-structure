package main

import (
	"flag"
	"fmt"
)

func main__() {

	// Define flags
	_filename := flag.String("filename", "data/traverse_data.csv", "file path")
	filename := *_filename

	// Parse command-line arguments
	flag.Parse()

	read := CsvReader{filename: filename}
	data := read.ReadFile()
	fmt.Println(data.Names())
	fmt.Println(data.Select([]string{"Ins."}))
	//dataParser := DataParser{data:data}
	//dataParser.parse()

	stn := Station{
		stationName: "TP1",
		Coordinates: &Point{
			X:                2000,
			Y:                4000,
			Z:                288,
			CoordinateSystem: "Ghana National Grid",
			ellipsoid:        "War Office",
			Datum:            "Accra",
			epsgCode:         "EPSG:2500",
		},
		Distance: 1000,
		unit:     "ft",
		next:     nil,
		prev:     nil,
	}
	stn2 := Station{
		stationName: "TP6",
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

}
