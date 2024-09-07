package main



type iGun interface {
	getName() string
	setName(newName string)
	getPower() int
	setPower(newPower int)
}

type Gun struct {
	power int
	name  string
}

type Ak47 struct {
	Gun
}

type Musket struct {
	Gun
}

func (gun *Gun) setName(newName string) {
	gun.name = newName
}

func (gun *Gun) setPower(newPower int) {
	gun.power = newPower
}

func (gun *Gun) getName() string {
	return gun.name
}

func (gun *Gun) getPower() int {
	return gun.power
}

func newAk47(gunPower int) iGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: gunPower,
		},
	}
}

func newMusket(gunPower int) iGun {
	return &Ak47{
		Gun: Gun{
			name:  "Musket gun",
			power: gunPower,
		},
	}
}

func gunFactory(gunType string) iGun {
	if gunType == "ak47" {
		return newAk47(12)
	} else if gunType == "musket" {
		return newMusket(14)
	} else {
		return newMusket(0)
	}
}
