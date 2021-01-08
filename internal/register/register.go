package register

import (
	"fmt"

	"github.com/faustikle/go-ponto/internal/list"
	"github.com/faustikle/go-ponto/internal/ponto"
	"github.com/faustikle/go-ponto/internal/pontotel"
)

func Register(c *pontotel.Client, kind, lat, lon string) (string, error) {
	if kind != "" && !ponto.ValidKind(kind) {
		return "", fmt.Errorf("Tipo inválido, os tipos aceitos são: entrada, pausa, retorno e saida.")
	}

	times, err := list.List(c)
	if err != nil {
		return "", err
	}

	res, err := c.Register(times, ponto.NewKind(kind), lat, lon)
	if err != nil {
		return "", err
	}

	return res.Success, nil
}
