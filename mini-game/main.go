package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string
	Lvl  int
	Exp  float64
}

type Stat struct {
	Str  int
	Hp   int
	Def  int
	Mana int
}

type Paladin struct {
	Person
	Stat Stat
}

type Monster struct {
	Person
	Stat Stat
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) updateEtatPersoGame(fileName string) {
	b, error := json.Marshal(p)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(string(b))
	file, err := os.Create("./files/" + fileName + ".json")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString(string(b))
}

func main() {
	fmt.Println("Bonjour")
	fmt.Println("Voulez-vous lancer le jeux ? (o/n)")
	var inputLaunchGame string
	var inputExitGame bool = true
	fmt.Scanf("%s", &inputLaunchGame)
	for inputLaunchGame != "n" && inputExitGame {
		fmt.Println("Voulez-vous lancer une nouvelle partie ? (1)")
		fmt.Println("Voulez-vous lancer une continué une partie ? (2)")
		fmt.Println("Voulez-vous quittez la partie ? (3)")
		var inputMenuGame int
		fmt.Scanf("%d", &inputMenuGame)
		switch inputMenuGame {
		case 1:
			newPartie()
		case 2:
			fmt.Println("Two")
		default:
			inputExitGame = false
		}
	}
}

func newPartie() {
	fmt.Println("Veuillez saisir votre nom : ")
	var inputName string
	fmt.Scanf("%s", &inputName)
	personnage := Paladin{Person{Name: inputName, Lvl: 1, Exp: 0}, Stat{Str: 5, Hp: 9, Def: 9, Mana: 3}}
	b, error := json.Marshal(personnage)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(string(b))
	file, err := os.Create("./files/myperson.json")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString(string(b))
}

func continuePartie() {
	fmt.Println("Choisir qui atk : ")

	file, e := ioutil.ReadFile("./files/ennemy.json")
	if e != nil {
		fmt.Println("File error:", e)
		os.Exit(1)
	}
	keysBody := []byte(string(file))
	var data []Monster
	err := json.Unmarshal(keysBody, &data)
	if err != nil {
		panic(err)
	}
	// for _, value := range data {
	// }
}
