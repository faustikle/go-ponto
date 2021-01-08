package ponto

const (
	KindEntrada = "Entrada"
	KindPausa   = "Pausa"
	KindRetorno = "Retorno Pausa"
	KindSaida   = "Saída"
)

func NewKind(kind string) string {
	switch kind {
	case "entrada":
		return KindEntrada
	case "pausa":
		return KindPausa
	case "retorno":
		return KindRetorno
	case "saida":
		return KindSaida
	default:
		return ""
	}
}

func ValidKind(kind string) bool {
	switch kind {
	case "entrada":
		return true
	case "pausa":
		return true
	case "retorno":
		return true
	case "saida":
		return true
	default:
		return false
	}
}
