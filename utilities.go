package main

import (
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

type CsvReader struct {
	filename string
}

func IsError(err error) {
	if err != nil {
		panic(err)
	}
}
func (c CsvReader) ReadFile() dataframe.DataFrame {
	file, err := os.Open(c.filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)

	//fmt.Println(df)
	return df
}

func (c CsvReader) WriteToCSV() {

}

type DataParser struct {
	data dataframe.DataFrame
}

func (dp *DataParser) parse() {

}
func (dp *DataParser) Columns() []string {
	return dp.data.Names()
}

func (dp *DataParser) NCols() int {
	return dp.data.Ncol()
}
