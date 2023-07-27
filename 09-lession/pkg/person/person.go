package person

type Person interface {
	GetAges() int
}

type Customer struct {
	ages int
}

type Emploer struct {
	ages int
}

func (p *Customer) GetAges() int {
	return p.ages
}

func (p *Emploer) GetAges() int {
	return p.ages
}

func GetOlder(p ...Person) int {
	maxAge := 0

	for _, v := range p {
		if old := v.GetAges(); old > maxAge {
			maxAge = old
		}
	}

	return maxAge
}
