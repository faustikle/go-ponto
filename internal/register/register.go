package register

import (
	"fmt"

	"github.com/faustikle/go-ponto/internal/list"
	"github.com/faustikle/go-ponto/internal/ponto"
	"github.com/faustikle/go-ponto/internal/pontotel"
)

func Register(c *pontotel.Client, kind, lat, lon string) (ponto.Entity, error) {
	if kind != "" && !ponto.ValidKind(kind) {
		return ponto.Entity{}, fmt.Errorf("Tipo inválido, os tipos aceitos são: entrada, pausa, retorno e saida.")
	}

	times, err := list.List(c)
	if err != nil {
		return ponto.Entity{}, err
	}

	res, err := c.Register(times, ponto.NewKind(kind), lat, lon)
	if err != nil {
		return ponto.Entity{}, err
	}

	newTimes, err := list.List(c)
	if err != nil {
		return ponto.Entity{}, err
	}

	fmt.Println(res.Success)

	return newTimes, nil
}
