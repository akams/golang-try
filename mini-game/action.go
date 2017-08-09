package main

import (
	"strings"
)

func search(name string, enemy []Monster) Monster {
	var res Monster
	nameToLowerCase := strings.ToLower(name)
	for _, value := range enemy {
		enemyName := strings.ToLower(value.Name)
		if strings.Index(enemyName, nameToLowerCase) != -1 {
			res = value
		}
	}
	return res
}
