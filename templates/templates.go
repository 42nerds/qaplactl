package tpl

// ApplicationTemplate ...
const ApplicationTemplate string = `apiVersion: qapla.42nerds.com/v1alpha1
kind: Application
metadata:
  name: {{ .ObjectMeta.Name }}
  namespace: qapla
spec:
  displayName: {{ .Spec.DisplayName }}
  iconSrc: {{ .Spec.IconSrc }}
  menuList:
  - text: First
    menuItems:
    - text: Edit First
      href: /edit
  - text: Second
    href: /second
`
