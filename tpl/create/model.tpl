package model

type {{ .StructName }} struct {
	Base
}

func (m *{{ .StructName }}) TableName() string {
    return "{{ .StructNameSnakeCase }}"
}
