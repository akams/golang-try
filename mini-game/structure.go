package main

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

type Wizard struct {
	Person
	Stat
}

type Warrior struct {
	Person
	Stat
}
