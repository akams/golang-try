package main

func (m *Monster) getStr() float64 {
	return float64(m.Stat.Str)
}

func (m *Monster) getMana() float64 {
	return float64(m.Stat.Mana)
}

func (m *Monster) getHp() float64 {
	return float64(m.Stat.Hp)
}

func (m *Monster) getDef() float64 {
	return float64(m.Stat.Def)
}

func (m *Monster) setStr(value float64) {
	m.Stat.Str = int(value)
}

func (m *Monster) setMana(value float64) {
	m.Stat.Mana = int(value)
}

func (m *Monster) setHp(value float64) {
	m.Stat.Hp = int(value)
}

func (m *Monster) setDef(value float64) {
	m.Stat.Def = int(value)
}
