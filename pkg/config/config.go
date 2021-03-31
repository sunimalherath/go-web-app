package config

import (
	"github.com/alexedwards/scs/v2"
	"text/template"
)

// AppConfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction bool
	Session *scs.SessionManager
}
