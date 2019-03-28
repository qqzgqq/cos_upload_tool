package upset

import (
	"cos_upload/template"
)

// StringUPSET upset string
func StringUPSET(s string) string {
	var str []string
	str = template.StringCHANGE2(template.StringTOSHUZU(s))
	str = template.StringCHANGE(str)
	return template.SHUZUtoSTring(str)
}
