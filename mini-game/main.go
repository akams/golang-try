package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Bonjour")
	fmt.Println("Voulez-vous lancer le jeux ? (o/n)")
	var inputLaunchGame string
	var inputExitGame bool = true
	fmt.Scanf("%s", &inputLaunchGame)
	for inputLaunchGame != "n" && inputExitGame {
		fmt.Println("Voulez-vous lancer une nouvelle partie ? (1)")
		fmt.Println("Voulez-vous lancer une continué une partie ? (2)")
		fmt.Println("Voulez-vous quittez le jeux ? (3)")
		var inputMenuGame int
		fmt.Scanf("%d", &inputMenuGame)
		switch inputMenuGame {
		case 1:
			newPartie()
		case 2:
			continuePartie()
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
	menuInGame()
}

func menuInGame() {
	var inputExitMenu bool = true
	for inputExitMenu {
		person := loadPerson()
		monsters := loadMonster()
		fmt.Println("Lancer une attaque ? (1)")
		fmt.Println("Prendre des soins ? (2)")
		fmt.Println("Voir mes stat (3)")
		fmt.Println("Quitter la partie ? (4)")
		var input int
		fmt.Scanf("%d", &input)
		switch input {
		case 1:
			launchFight(person, monsters)
		case 2:
			healPerso(person)
		case 3:
			displayStat(person)
		default:
			inputExitMenu = false
		}
	}
}

func launchFight(person Paladin, monsters []Monster) {
	fmt.Println("Choisir qui atk : ")
	for _, value := range monsters {
		fmt.Println("Name: ", value.Name, "Lvl: ", value.Lvl, "HP:", value.Stat.Hp)
	}
	fmt.Println("saisir le nom de votre adversaire: ")
	var inputName string
	fmt.Scanf("%s", &inputName)
	en := search(inputName, monsters)
	person.Attaque(en, monsters)
}

func healPerso(person Paladin) {
	fmt.Println("==personnage==", person.Name, " - ", person.getHp(), "HP")
	person.setHp(person.getHp() + 3)
	fmt.Println("==Le personnage==", person.Name, "récupère:", 3, "HP")
	fmt.Println("==Le personnage==", person.Name, "à", person.getHp(), "HP")
	person.UpdateEtatPersoGame()
}

func displayStat(person Paladin) {
	fmt.Println("Nom:", person.Name, " - Lvl:", person.getLvl(), " - Exp:", person.getExp())
	fmt.Println("str:", person.getStr(), " - mana:", person.getMana())
	fmt.Println("hp:", person.getHp(), " - def:", person.getDef())
}

func loadPerson() Paladin {
	fmt.Println("==Chargement du personnage==")
	filePerson, e := ioutil.ReadFile("./files/myperson.json")
	if e != nil {
		fmt.Println("File error:", e)
		os.Exit(1)
	}
	keysBodyPerso := []byte(string(filePerson))
	var dataPerso Paladin
	error := json.Unmarshal(keysBodyPerso, &dataPerso)
	if error != nil {
		panic(error)
	}
	fmt.Println("==Fin Chargement du personnage==")
	return dataPerso
}

func loadMonster() []Monster {
	fmt.Println("==Chargement des ennemies==")
	file, e := ioutil.ReadFile("./files/ennemy.json")
	if e != nil {
		fmt.Println("File error:", e)
		os.Exit(1)
	}
	keysBody := []byte(string(file))
	var dataMonster []Monster
	err := json.Unmarshal(keysBody, &dataMonster)
	if err != nil {
		panic(err)
	}
	fmt.Println("==Fin Chargement des ennemies==")
	return dataMonster
}
