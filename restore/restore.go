package restore

import (
	"cos_upload/template"
)

// StringRESTORE restore string
func StringRESTORE(s string) string {
	var str []string
	str = template.StringCHANGE(template.StringTOSHUZU(s))
	str = template.StringCHANGE2(str)
	return template.SHUZUtoSTring(str)
}
