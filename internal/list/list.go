package list

import (
	"github.com/faustikle/go-ponto/internal/ponto"
	"github.com/faustikle/go-ponto/internal/pontotel"
)

func List(c *pontotel.Client) (ponto.Entity, error) {
	res, err := c.Login()
	if err != nil {
		return ponto.Entity{}, err
	}

	return ponto.NewEntity(
		res.Success.Employee.ID,
		extractTime(res.Success.DayPts.Entrada),
		extractTime(res.Success.DayPts.Pausa),
		extractTime(res.Success.DayPts.Retorno),
		extractTime(res.Success.DayPts.Saida),
	), nil
}

func extractTime(p []pontotel.Pts) [] string {
	var t []string

	for _, v := range p {
		t = append(t, v.Time)
	}

	return t
}
