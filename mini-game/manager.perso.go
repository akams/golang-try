package main

import (
	"math"
)

func (p *Paladin) getStr() float64 {
	return float64(p.Stat.Str)
}

func (p *Paladin) getMana() float64 {
	return float64(p.Stat.Mana)
}

func (p *Paladin) getHp() float64 {
	return float64(p.Stat.Hp)
}

func (p *Paladin) getDef() float64 {
	return float64(p.Stat.Def)
}

func (p *Paladin) getLvl() int {
	return p.Person.Lvl
}

func (p *Paladin) getExp() float64 {
	return float64(p.Person.Exp)
}

func (p *Paladin) setStr(value float64) {
	p.Stat.Str = int(value)
}

func (p *Paladin) setMana(value float64) {
	p.Stat.Mana = int(value)
}

func (p *Paladin) setHp(value float64) {
	p.Stat.Hp = int(value)
}

func (p *Paladin) setDef(value float64) {
	p.Stat.Def = int(value)
}

func (p *Paladin) setLvl(value int) {
	p.Person.Lvl = value
}

func (p *Paladin) setExp(value float64) {
	p.Person.Exp = value
}

func (p *Paladin) leveling(value float64) {
	if p.getExp()+value >= 100 {
		p.setLvl(p.getLvl() + 1)
		p.setExp(0)
	}
	p.setExp(math.Ceil(p.getExp() + value))
}
