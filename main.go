package main

import "fmt"

var baza []Product
func main() {
	s := `Wybierz co chcesz zrobić
	1. Dodaj produkt
	2. Usuń produkt
	3. Edytuj produkt
	4. Wyświetl produkty
		`
	for{
		fmt.Println(s)
		var wybór int
		fmt.Scanln(&wybór)
			switch wybór {
			case 1:
				dodaj()
			case 2:
			case 3:
			case 4:
				wypisz()
		}
	}
}
func dodaj()  {
	var name string
	fmt.Println("podaj imię")
	fmt.Scanln(&name)
	var price int
	fmt.Println("podaj cenę")
	fmt.Scanln(&price)
	p:=Product{name,price}
	baza=append(baza,p)
}

func wypisz()  {
	fmt.Println(baza)
}

type Product struct {
    name  string
    price int

}