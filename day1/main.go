package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GO language is strongly typed, you can even create new types and define a variable with it
// You can also add a tag to each parameter to define the way they appear in json
type Carro struct {
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano    int    `json:"-"`
}

/*We can also reference the structure and create a method */
func (c Carro) Parar() {
	fmt.Println("O carro, " + c.Nome + " está parando")
}

func (c Carro) Andar() {
	fmt.Println("O carro, " + c.Nome + " está andando")
}

func main() {
	carro1 := Carro{Nome: "Ford", Modelo: "Focus", Ano: 2012}
	carro2 := Carro{Nome: "Toyota", Modelo: "Supra", Ano: 2019}

	fmt.Println(carro1.Nome)
	fmt.Println(carro2.Nome)

	carro1.Andar()
	carro2.Andar()

	carro1.Parar()
	carro2.Parar()

	//Create handler while using an anonymous function:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Convert struct to json
		json.NewEncoder(w).Encode(carro1)
	})

	//Create webserver:
	http.ListenAndServe(":3333", nil)

}
