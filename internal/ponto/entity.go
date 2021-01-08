package ponto

import (
	"fmt"
	"time"
)

type BestTimes struct {
	Entrada time.Time
	Pausa   time.Time
	Retorno time.Time
	Saida   time.Time
}

type Entity struct {
	Entrada   []time.Time
	Pausa     []time.Time
	Retorno   []time.Time
	Saida     []time.Time
	Employee  string
	BestTimes BestTimes
}

func NewEntity(employee string, entrada, pausa, retorno, saida []string) Entity {
	var e []time.Time
	var p []time.Time
	var r []time.Time
	var s []time.Time

	for _, v := range entrada {
		t, err := parseTime(v)
		if err == nil {
			e = append(e, t)
		}
	}

	for _, v := range pausa {
		t, err := parseTime(v)
		if err == nil {
			p = append(p, t)
		}
	}

	for _, v := range retorno {
		t, err := parseTime(v)
		if err == nil {
			r = append(r, t)
		}
	}

	for _, v := range saida {
		t, err := parseTime(v)
		if err == nil {
			s = append(s, t)
		}
	}

	entity := Entity{
		Entrada:  e,
		Pausa:    p,
		Retorno:  r,
		Saida:    s,
		Employee: employee,
	}

	entity.BestTimes = newBestTimes(entity)

	return entity
}

func (e Entity) NextKind() string {
	if len(e.Entrada) == 0 {
		return KindEntrada
	}

	if len(e.Pausa) == 0 {
		return KindPausa
	}

	if len(e.Pausa) != len(e.Retorno) {
		return KindRetorno
	}

	if len(e.Saida) == 0 {
		return KindSaida
	}

	fmt.Println(e)

	return ""
}

// TODO: Implement multiple times calculations.
func newBestTimes(e Entity) BestTimes {
	if len(e.Entrada) > 1 || len(e.Pausa) > 1 || len(e.Retorno) > 1 || len(e.Saida) > 1 {
		return BestTimes{}
	}

	startTime, _ := time.Parse(time.RFC3339Nano, "2006-01-02T09:00:00.999999999Z07:00")
	pauseTime, _ := time.Parse(time.RFC3339Nano, "2006-01-02T12:00:00.999999999Z07:00")
	bestRetorno := e.Pausa[0].Add(time.Hour)
	lauchTime := bestRetorno.Sub(e.Pausa[0])

	if e.Retorno != nil && e.Retorno[0] != (time.Time{}) {
		lauchTime = e.Retorno[0].Sub(e.Pausa[0])
	}

	saida := e.Entrada[0].Add((time.Hour*8) + lauchTime)

	return BestTimes{
		Entrada: startTime,
		Pausa:   pauseTime,
		Retorno: bestRetorno,
		Saida:   saida,
	}
}

func parseTime(v string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, v+"Z")
}
