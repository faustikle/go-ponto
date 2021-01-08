package app

import (
	"fmt"

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

			fmt.Println("Buscando pontos batidos...")

			client := pontotel.NewClient(pontotel.ClientConfig{
				CompMan:     cfg.CompMan,
				FingerPrint: cfg.Fingerprint,
				Password:    cfg.Password,
			})

			times, err := list.List(client)
			if err != nil {
				errorhandler.Handler(err)
				return
			}

			list.Print(times)
		},
	}
}
