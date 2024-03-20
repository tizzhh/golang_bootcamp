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
	GetCakeData() []Cake
}

type Cakes struct {
	XMLName xml.Name `json:"-"  xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingridients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
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
		return fmt.Errorf("error during file reading: %v", err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jr.Cakes)

	return nil
}

func (jr *JSONReader) OutputData() error {
	out, err := json.MarshalIndent(jr.Cakes, "", "    ")
	if err != nil {
		return fmt.Errorf("error during OutputData: %v", err)
	}

	fmt.Println(string(out))

	return nil
}

func (jr *JSONReader) GetCakeData() []Cake {
	return jr.Cakes.Cakes
}

type XMLReader struct {
	Path  string
	Cakes Cakes
}

func (xr *XMLReader) ReadData() error {
	xmlFile, err := os.Open(xr.Path)
	if err != nil {
		return fmt.Errorf("error during file reading: %v", err)
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &xr.Cakes)

	return nil
}

func (xr *XMLReader) OutputData() error {
	out, err := xml.MarshalIndent(xr.Cakes, "", "    ")
	if err != nil {
		return fmt.Errorf("error during OutputData: %v", err)
	}

	fmt.Println(string(out))

	return nil
}

func (xr *XMLReader) GetCakeData() []Cake {
	return xr.Cakes.Cakes
}

func OutputJsonXml(reader DBReader) {
	switch r := reader.(type) {
	case (*JSONReader):
		xr := &XMLReader{Path: r.Path, Cakes: r.Cakes}
		xr.OutputData()
	case (*XMLReader):
		jr := &JSONReader{Path: r.Path, Cakes: r.Cakes}
		jr.OutputData()
	}
}

func ParseInput() (string, error) {
	mode := "reader"

	if ((len(os.Args) != 3) && len(os.Args) != 5) || !((os.Args[1] == "--old" && os.Args[3] == "--new") || (os.Args[1] == "--new" && os.Args[3] == "--old")) || (os.Args[1] == os.Args[3]) {
		return "", fmt.Errorf(`usage: ./compareDB -f [filename xml or json]
		or ./compareDB --[old/new] original_database.[xml/json] --[new/old] stolen_database.[json/xml]`)
	}

	path := os.Args[2]
	if !strings.HasSuffix(path, ".xml") && !strings.HasSuffix(path, ".json") {
		return "", fmt.Errorf("unknown file format")
	}
	if len(os.Args) == 5 {
		path = os.Args[4]
		if !strings.HasSuffix(path, ".xml") && !strings.HasSuffix(path, ".json") {
			return "", fmt.Errorf("unknown file format")
		}
		if os.Args[1] == "--old" {
			mode = "compare2to1"
		} else {
			mode = "compare1to2"
		}
	}
	return mode, nil
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
