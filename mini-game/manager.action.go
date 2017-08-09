package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

func (p *Paladin) UpdateEtatPersoGame() {
	b, error := json.Marshal(p)
	if error != nil {
		fmt.Println(error)
		return
	}
	file, err := os.Create("./files/myperson.json")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString(string(b))
}

func UpdateEtatMonsterGame(monster Monster, monsters []Monster) {
	nameToLowerCase := strings.ToLower(monster.Name)
	for index, value := range monsters {
		enemyName := strings.ToLower(value.Name)
		if strings.Index(enemyName, nameToLowerCase) != -1 {
			monsters[index] = monster
		}
	}
	b, error := json.Marshal(monsters)
	if error != nil {
		fmt.Println(error)
		return
	}
	file, err := os.Create("./files/ennemy.json")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString(string(b))
}

func (p *Paladin) Attaque(monster Monster, monsters []Monster) {
	fmt.Println("==", p.Name, "== attaque => ", monster.Name)
	dommageP := math.Abs((p.getStr() - monster.getDef()))
	remnantLifeM := math.Abs(dommageP - monster.getHp())
	monster.setHp(remnantLifeM)
	fmt.Println("==", p.Name, "== inflige => ", dommageP, "dommage")
	fmt.Println("==", monster.Name, "== contre attaque => ", p.Name)
	dommageM := math.Abs((monster.getStr() - p.getDef()))
	remnantLifeP := math.Abs(dommageM - p.getHp())
	p.setHp(remnantLifeP)
	fmt.Println("==", monster.Name, "== inflige => ", dommageM, "dommage")
	p.leveling(7)
	p.UpdateEtatPersoGame()
	UpdateEtatMonsterGame(monster, monsters)
}
