package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/faustikle/go-ponto/internal/errorhandler"
	"github.com/faustikle/go-ponto/internal/fileconfig"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: fmt.Sprintf("Inicializa as configurações na sua HOME"),
		Run: func(cmd *cobra.Command, args []string) {
			var cfg fileconfig.Config
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Sua senha para acessar o ponto: ")
			pass, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
			cfg.Password = string(pass)

			fmt.Println("\n\nAbra o 'Local Storage' na Devtools do seu navegador no site 'https://web.pontotel.com.br/#/password', e preencha os valores solicitados")

			fmt.Print("ngStorage-compMan.id: ")
			compMan, _ := reader.ReadString('\n')
			compMan = strings.ReplaceAll(compMan, "\n", "")
			cfg.CompMan = strings.ReplaceAll(compMan, "\r\n", "")

			fmt.Print("ngStorage-id_browser.id: ")
			fingerprint, _ := reader.ReadString('\n')
			fingerprint = strings.ReplaceAll(fingerprint, "\n", "")
			cfg.Fingerprint = strings.ReplaceAll(fingerprint, "\r\n", "")

			fmt.Println("\nAgora será solicitado sua geolocalização. Utilize o Google Maps para achar sua posição.")

			fmt.Print("Latitude (ex: -23.5420731): ")
			lat, _ := reader.ReadString('\n')
			lat = strings.ReplaceAll(lat, "\n", "")
			cfg.Lat = strings.ReplaceAll(lat, "\r\n", "")

			fmt.Print("Longitude (ex: -46.699522599999995): ")
			lon, _ := reader.ReadString('\n')
			lon = strings.ReplaceAll(lon, "\n", "")
			cfg.Lon = strings.ReplaceAll(lon, "\r\n", "")

			err := fileconfig.Save(cfg)
			if err != nil {
				errorhandler.Handler(err)
				return
			}

			fmt.Println("\nConfigurações registradas com sucesso!")
		},
	}
}
