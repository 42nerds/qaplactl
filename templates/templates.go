package tpl

// MainTemplate ...
func MainTemplate() []byte {
	return []byte(`
example generated main.go
`)
}

// ApplicationTemplate ...
func ApplicationTemplate() []byte {
	return []byte(`apiVersion: qapla.42nerds.com/v1alpha1
kind: Application
metadata:
  name: {{ if .DisplayName }} {{ .DisplayName }} {{ else }} standard-name {{ end }}
spec:
  displayName: {{ if .DisplayName }} {{ .DisplayName }} {{ else }} standard-name {{ end }}
  iconSrc: {{ .IconSrc }}
	menuItems:
    - text: Contacts
      items:
        - text: Edit
          href: /edit
    - text: Settings
      href: /settings
`)
}