package person2

type Person interface{}

type Customer struct {
	ages int
}

type Emploer struct {
	ages int
}

func GetOlder(p ...Person) Person {
	var older Person
	maxAge := 0

	for _, v := range p {
		if cust, ok := v.(Customer); ok && cust.ages > maxAge {
			maxAge = cust.ages
			older = cust
		} else if emp, ok := v.(Emploer); ok && emp.ages > maxAge {
			maxAge = emp.ages
			older = emp
		}
	}

	return older
}
