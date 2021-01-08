package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/faustikle/go-ponto/internal/errorhandler"
	"github.com/faustikle/go-ponto/internal/fileconfig"
	"github.com/faustikle/go-ponto/internal/list"
	"github.com/faustikle/go-ponto/internal/pontotel"
	"github.com/spf13/cobra"
)

func LS() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "Mostra seus pontos batidos",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := fileconfig.Load()
			if err != nil {
				errorhandler.Handler(err)
				return
			}

			client := pontotel.NewClient(pontotel.ClientConfig{
				CompMan:     cfg.CompMan,
				FingerPrint: cfg.Fingerprint,
				Password:    cfg.Password,
			})

			pontos, err := list.List(client)
			if err != nil {
				errorhandler.Handler(err)
				return
			}

			fmt.Printf("Entrada: %s\n", formatPonto(pontos.Entrada))
			fmt.Printf("Pausa: %s\n", formatPonto(pontos.Pausa))
			fmt.Printf("Retorno: %s\n", formatPonto(pontos.Retorno))
			fmt.Printf("Saida: %s\n", formatPonto(pontos.Saida))
		},
	}
}

func formatPonto(t []time.Time) string {
	var f []string

	if len(t) == 0 {
		return "(Ponto não batido)"
	}

	for _, tt := range t {
		f = append(f, tt.Format("15:04"))
	}

	return strings.Join(f, "  ")
}
