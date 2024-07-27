package templates

import "embed"

//go:embed template_register.html
//go:embed template_resend.html
var TemplateFS embed.FS
