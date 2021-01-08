package list

import (
	"fmt"
	"strings"
	"time"

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

func Print(times ponto.Entity) {
	fmt.Println("\nAs melhores horas são geradas com base numa jornada de 8 horas com 1 hora de almoço.")
	fmt.Printf("\nEntrada: %s\n", formatPonto(times.Entrada, times.BestTimes.Entrada))
	fmt.Printf("Pausa: %s\n", formatPonto(times.Pausa, times.BestTimes.Pausa))
	fmt.Printf("Retorno: %s\n", formatPonto(times.Retorno, times.BestTimes.Retorno))
	fmt.Printf("Saida: %s\n", formatPonto(times.Saida, times.BestTimes.Saida))
}


func formatPonto(t []time.Time, best time.Time) string {
	var f []string

	if len(t) == 0 && best != (time.Time{}) {
		return fmt.Sprintf("(Melhor hora: %s)", best.Format("15:04"))
	}

	if len(t) == 0 {
		return "(Ponto não batido)"
	}

	for _, tt := range t {
		f = append(f, tt.Format("15:04"))
	}

	return strings.Join(f, "  ")
}


func extractTime(p []pontotel.Pts) [] string {
	var t []string

	for _, v := range p {
		t = append(t, v.Time)
	}

	return t
}
