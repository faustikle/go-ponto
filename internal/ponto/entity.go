package ponto

import (
	"fmt"
	"time"
)

type Entity struct {
	Entrada  []time.Time
	Pausa    []time.Time
	Retorno  []time.Time
	Saida    []time.Time
	Employee string
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

	return Entity{
		Entrada:  e,
		Pausa:    p,
		Retorno:  r,
		Saida:    s,
		Employee: employee,
	}
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

func parseTime(v string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, v+"Z")
}
