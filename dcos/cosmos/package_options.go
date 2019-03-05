package cosmos

// PackageSecurityOptions describes the security options available when
// installing a package
type PackageSecurityOptions struct {
	StrictMode bool `json:"strict-mode,omitempty"`
}

// PackageServiceOptions describes the service options available when
// installing a package
type PackageServiceOptions struct {
	SecretName    string `json:"secretName,omitempty"`
	Principal     string `json:"principal,omitempty"`
	MesosProtocol string `json:"mesosProtocol,omitempty"`
}

// PackageOptions describes the options available when installing a package
type PackageOptions struct {
	Service PackageServiceOptions `json:"service,omitempty"`

	Security PackageSecurityOptions `json:"security,omitempty"`
}
