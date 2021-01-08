package pontotel

type LoginRequest struct {
	CompMan     string `json:"compMan"`
	Fingerprint string `json:"fingerprint"`
	Pwd         string `json:"pwd"`
}
