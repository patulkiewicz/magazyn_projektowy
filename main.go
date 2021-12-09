package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var baza []Product

func main() {
	s := `Wybierz co chcesz zrobić
	1. Dodaj produkt
	2. Usuń produkt
	3. Edytuj produkt
	4. Wyświetl produkty
	5. wyjscie z programu
	6. Importowanie
	7. Exportowanie
        `
	for {
		fmt.Println(s)
		var wybór int
		fmt.Scanln(&wybór)
		switch wybór {
		case 1:
			dodaj()
		case 2:
			usun()
		case 3:
			edytuj()
		case 4:
			wypisz()
		case 5:
			os.Exit(0)
		case 6:
			importowanie()
		case 7:
			eksportowanie()
		}
	}
}

func dodaj() {
	var name string
	fmt.Println("podaj nazwę produktu")
	scanln(&name)
	for name == "" {
		fmt.Println("podaj nazwę produktu")
		scanln(&name)
	}

	var price string
	fmt.Println("podaj cenę")
	scanln(&price)
	price_val, err := strconv.ParseFloat(price, 64)
	for price == "" || err != nil {
		fmt.Println("podaj prawidłową cenę używając liczb i kropki aby wprowadzić części setne")
		scanln(&price)
		price_val, err = strconv.ParseFloat(price, 64)
	}

	id := len(baza) + 1
	p := Product{Id: id, Name: name, Price: price_val}
	baza = append(baza, p)
}

func scanln(val *string) {
	in := bufio.NewReader(os.Stdin)
	*val, _ = in.ReadString('\n')
	*val = strings.TrimSpace(*val)
}

func usun() {
	fmt.Println("wybierz id przedmiotu do usunięcia")
	var id int
	fmt.Scanln(&id)
	var usunięty bool
	n := 0
	for _, product := range baza {
		if product.Id != id {
			baza[n] = product
			n += 1
		} else {
			usunięty = true
		}
	}
	baza = baza[:n]
	if !usunięty {
		fmt.Println("wybrany produkt nie istnieje")
	} else {
		fmt.Println("produkt usunięty pomyślnie")
	}
}

func edytuj() {
	fmt.Println("Podaj id produktu do edycji")
	var id int
	var edytowany bool
	fmt.Scanln(&id)
	for idx, product := range baza {
		if product.Id == id {
			edytowany = true
			fmt.Printf("name [%s]: ", product.Name)
			var name string
			fmt.Scanln(&name)
			fmt.Println(name)
			if name != "" {
				baza[idx].Name = name
			}

			fmt.Printf("price [%.2f]:", product.Price)
			var price string
			fmt.Scanln(&price)
			if price != "" {
				var err error
				baza[idx].Price, err = strconv.ParseFloat(price, 64)
				for err != nil {
					fmt.Println("podaj prawidłową cenę używając liczb i kropki aby wprowadzić części setne")
					fmt.Scanln(&price)
					baza[idx].Price, err = strconv.ParseFloat(price, 64)

				}
			}
		}
	}
	if !edytowany {
		fmt.Println("wybrany produkt nie istnieje")
	} else {
		fmt.Println("produkt edytowany pomyślnie")
	}
}

func wypisz() {
	for _, product := range baza {
		fmt.Printf("%+v\n", product)
		fmt.Printf("nazwa produktu: %s cena: %.2fzł\n", product.Name, product.Price)
	}
}

type Product struct {
	Id    int
	Name  string
	Price float64
}

func importowanie() {
	err := readGob("./baza.gob", &baza)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("baza została pomyslnie zaimportowana")
}

func eksportowanie() {
	err := writeGob("./baza.gob", baza)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("baza została pomyslnie wyeksportowana do pliku baza.gob")
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
