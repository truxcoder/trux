package v1

type {{ .StructName }}Request struct {
	ID int64 `json:"id,string" binding:"id"`
}
