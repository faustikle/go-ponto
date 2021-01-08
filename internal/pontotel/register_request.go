package pontotel

type RegisterRequest struct {
	Employee     string
	SessionToken string
	Fingerprint  string
	Kind         string
	Lat          string
	Lon          string
}
