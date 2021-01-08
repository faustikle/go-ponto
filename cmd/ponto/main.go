package main

import (
	"github.com/faustikle/go-ponto/internal/app"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ponto",
		Short: "Um CLI para registrar e consultar pontos batidos no Pontotel.",
	}
)

func init() {
	rootCmd.AddCommand(app.LS())
	rootCmd.AddCommand(app.Register())
	rootCmd.AddCommand(app.Init())
}

func main() {
	rootCmd.Execute()
}


