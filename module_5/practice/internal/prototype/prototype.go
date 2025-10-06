package prototype

type Weapon struct {
	Name   string
	Damage int
}

func (w Weapon) Clone() Weapon {
	return Weapon{Name: w.Name, Damage: w.Damage}
}

type Armor struct {
	Name    string
	Defense int
}

func (a Armor) Clone() Armor {
	return Armor{Name: a.Name, Defense: a.Defense}
}

type Skill struct {
	Name  string
	Power int
}

func (s Skill) Clone() Skill {
	return Skill{Name: s.Name, Power: s.Power}
}

type Character struct {
	Name      string
	Health    int
	Strength  int
	Agility   int
	Intellect int
	Weapon    Weapon
	Armor     Armor
	Skills    []Skill
}

func (c Character) Clone() Character {
	newSkills := make([]Skill, len(c.Skills))
	for i, s := range c.Skills {
		newSkills[i] = s.Clone()
	}
	return Character{
		Name:      c.Name,
		Health:    c.Health,
		Strength:  c.Strength,
		Agility:   c.Agility,
		Intellect: c.Intellect,
		Weapon:    c.Weapon.Clone(),
		Armor:     c.Armor.Clone(),
		Skills:    newSkills,
	}
}
