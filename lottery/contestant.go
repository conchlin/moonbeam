package lottery

type Contestant struct {
	Name string
}

func CreateContestant(name string) *Contestant {
	player := &Contestant{
		Name: name,
	}

	return player
}
