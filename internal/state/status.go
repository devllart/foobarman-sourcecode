package state

var Status = "Norm" // Status barman

func HandlerStatus() {
	if Money < 10 {
		SetStatusBad()
	} else {
		SetStatusNorm()
	}
}

func SetStatusBad() {
	Status = "Bad"
}

func SetStatusNorm() {
	Status = "Norm"
}
