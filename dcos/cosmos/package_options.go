package cosmos

type PackageSecurityOptions struct {
	StrictMode bool `json:"strict-mode,omitempty"`
}

type PackageServiceOptions struct {
	SecretName    string `json:"secretName,omitempty"`
	Principal     string `json:"principal,omitempty"`
	MesosProtocol string `json:"mesosProtocol,omitempty"`
}

type PackageOptions struct {
	Service PackageServiceOptions `json:"service,omitempty"`

	Security PackageSecurityOptions `json:"security,omitempty"`
}
