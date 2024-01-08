package afrogpocs

import "embed"

//go:embed CNVD
//go:embed CVE
//go:embed default-login
//go:embed disclosure
//go:embed fingerprint
//go:embed unauthorized
//go:embed vulnerability
var templates embed.FS
