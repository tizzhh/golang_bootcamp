package comparer

import (
	"compareDB/reader"
	"fmt"
	"strings"
)

func CompareDB(r1, r2 reader.DBReader) {
	r1.ReadData()
	r2.ReadData()
	cakeData1, cakeData2 := r1.GetCakeData(), r2.GetCakeData()
	cake1Map := make(map[string]reader.Cake, len(cakeData1))
	for _, cake := range cakeData1 {
		cake1Map[cake.Name] = cake
	}

	for _, cake2 := range cakeData2 {
		if cake1, ok := cake1Map[cake2.Name]; ok {
			compareCakesAndPrint(cake1, cake2)
			delete(cake1Map, cake2.Name)
		} else {
			fmt.Printf("ADDED cake \"%s\"\n", cake2.Name)
		}
	}

	for _, cake1 := range cake1Map {
		fmt.Printf("REMOVED cake \"%s\"\n", cake1.Name)
	}
}

func compareCakesAndPrint(cake1, cake2 reader.Cake) {
	if cake2.Time != cake1.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake2.Name, cake2.Time, cake1.Time)
	}
	compareIngredients(cake1.Ingridients, cake2.Ingridients, cake1.Name)
}

func compareIngredients(ingredientData1, ingredientData2 []reader.Ingredient, cakeName string) {
	ingr1Map := make(map[string]reader.Ingredient, len(ingredientData1))
	for _, ingr := range ingredientData1 {
		ingr1Map[ingr.Name] = ingr
	}

	for _, ingr2 := range ingredientData2 {
		if ingr1, ok := ingr1Map[ingr2.Name]; ok {
			compareIngredientsAndPrint(ingr1, ingr2, cakeName)
			delete(ingr1Map, ingr2.Name)
		} else {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingr2.Name, cakeName)
		}
	}

	for _, ingr1 := range ingr1Map {
		fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ingr1.Name, cakeName)
	}
}

func compareIngredientsAndPrint(ingr1, ingr2 reader.Ingredient, cakeName string) {
	if (ingr2.Unit != "" && ingr1.Unit != "") && (ingr2.Unit != ingr1.Unit) {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ingr1.Name, cakeName, ingr2.Unit, ingr1.Unit)
	} else if ingr2.Count != ingr1.Count {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ingr1.Name, cakeName, ingr2.Count, ingr1.Count)
	}
	if ingr2.Unit == "" && ingr1.Unit != "" {
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", ingr1.Unit, ingr1.Name, cakeName)
	}
	if ingr1.Unit == "" && ingr2.Unit != "" {
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", ingr2.Unit, ingr1.Name, cakeName)
	}
}

func MakeComparer(path1, path2 string) (dbcomparer1, dbcomparer2 reader.DBReader) {
	if strings.HasSuffix(path1, ".xml") {
		dbcomparer1 = &reader.XMLReader{Path: path1}
	} else {
		dbcomparer1 = &reader.JSONReader{Path: path1}
	}

	if strings.HasSuffix(path2, ".xml") {
		dbcomparer2 = &reader.XMLReader{Path: path2}
	} else {
		dbcomparer2 = &reader.JSONReader{Path: path2}
	}
	return
}
