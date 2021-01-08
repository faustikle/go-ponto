package pontotel

type LoginResponse struct {
	Success Success `json:"success"`
}

type Success struct {
	DayPts       DayPts   `json:"dayPts"`
	Employee     Employee `json:"emproyee"`
	SessionToken string   `json:"sessionToken"`
}

type DayPts struct {
	Entrada []Pts `json:"entrada"`
	Pausa   []Pts `json:"pausa"`
	Retorno []Pts `json:"retorno"`
	Saida   []Pts `json:"saida"`
}

type Employee struct {
	ID string `json:"id"`
}

type Pts struct {
	Time string `json:"time"`
	Fake bool   `json:"fake"`
}
