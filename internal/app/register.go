package app

import (
	"github.com/faustikle/go-ponto/internal/errorhandler"
	"github.com/faustikle/go-ponto/internal/fileconfig"
	"github.com/faustikle/go-ponto/internal/list"
	"github.com/faustikle/go-ponto/internal/pontotel"
	"github.com/faustikle/go-ponto/internal/register"
	"github.com/spf13/cobra"
)

func Register() *cobra.Command {
	var kind string
	command := &cobra.Command{
		Use:   "put",
		Short: "Bate o ponto",
		Long: "Uso: ponto put -t Entrada\nCaso o tipo não seja passado, será calculado o próximo ponto a ser batido.",
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

			newTimes, err := register.Register(client, kind, cfg.Lat, cfg.Lon)
			if err != nil {
				errorhandler.Handler(err)
				return
			}

			list.Print(newTimes)
		},
	}

	command.Flags().StringVarP(&kind, "tipo", "t", "", "Tipo do registro: entrada, pausa, retorno ou saida. (Opcional)")

	return command
}