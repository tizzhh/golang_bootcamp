package reader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type DBReader interface {
	ReadData() error
	OutputData() error
}

type Cakes struct {
	XMLName xml.Name `json:"-"  xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	// XMLName xml.Name `xml:"cake"`
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingridients []Ingridient `json:"ingredients" xml:"ingredients>item"`
}

type Ingridient struct {
	// XMLName xml.Name `xml:"item"`
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

type JSONReader struct {
	Path  string
	Cakes Cakes
}

func (jr *JSONReader) ReadData() error {
	jsonFile, err := os.Open(jr.Path)
	if err != nil {
		return fmt.Errorf("error during file reading: %s", err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jr.Cakes)

	return nil
}

func (jr *JSONReader) OutputData() error {
	out, err := json.MarshalIndent(jr.Cakes, "", "    ")
	if err != nil {
		return fmt.Errorf("error during OutputData: %s", err)
	}

	fmt.Println(string(out))

	return nil
}

type XMLReader struct {
	Path  string
	Cakes Cakes
}

func (xr *XMLReader) ReadData() error {
	xmlFile, err := os.Open(xr.Path)
	if err != nil {
		return fmt.Errorf("error during file reading: %s", err)
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &xr.Cakes)

	return nil
}

func (xr *XMLReader) OutputData() error {
	out, err := xml.MarshalIndent(xr.Cakes, "", "    ")
	if err != nil {
		return fmt.Errorf("error during OutputData: %s", err)
	}

	fmt.Println(string(out))

	return nil
}

func ParseInput() (string, error) {
	var path string
	if len(os.Args) != 3 || os.Args[1] != "-f" {
		return "", fmt.Errorf("usage: ./DBReader -f [filename xml or json]")
	}
	path = os.Args[2]
	if !strings.HasSuffix(path, ".xml") && !strings.HasSuffix(path, ".json") {
		return "", fmt.Errorf("unknown file format")
	}
	return path, nil
}

func CreateReader(path string) DBReader {
	switch {
	case strings.HasSuffix(path, ".xml"):
		return &XMLReader{path, Cakes{}}
	case strings.HasSuffix(path, ".json"):
		return &JSONReader{path, Cakes{}}
	}
	return nil
}
