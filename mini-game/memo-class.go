type Person struct {
	name string
	lvl  int
	exp  float64
}

type Stat struct {
	str  int
	hp   int
	def  int
	mana int
}

type Paladin struct {
	Person
	Stat
}

type Wizard struct {
	Person
	Stat
}

type Warrior struct {
	Person
	Stat
}

func factory() interface{} {
	fmt.Println("Veuillez saisir votre nom : ")
	var inputName string
	fmt.Scanf("%s", &inputName)
	fmt.Println("Pour créer un Paladin (1), Wizard (2), Warrior (3)")
	var input float64
	fmt.Scanf("%f", &input)
if input == 1 {
return Paladin{Person{name: inputName, lvl: 1, exp: 0}, Stat{str: 5, hp: 9, def: 9, mana: 3}}
} else if input == 2 {
	return Wizard{Person{name: inputName, lvl: 1, exp: 0}, Stat{str: 3, hp: 5, def: 5, mana: 9}}
} else {
	return Warrior{Person{name: inputName, lvl: 1, exp: 0}, Stat{str: 9, hp: 5, def: 5, mana: 3}}
}
}
