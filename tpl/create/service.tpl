package service

import (
    "context"
    "github.com/jinzhu/copier"
    v1 "{{ .ProjectName }}/api/v1"
	"{{ .ProjectName }}/internal/model"
	"{{ .ProjectName }}/internal/repository"
)

type {{ .StructName }}Service interface {
	Get{{ .StructName }}(ctx context.Context, id int64) (*model.{{ .StructName }}, error)
	Get{{ .StructNamePlural }}(ctx context.Context, condition ...any) ([]model.{{ .StructName }}, error)
	Create{{ .StructName }}(ctx context.Context, req *v1.{{ .StructName }}Request) error
	Delete{{ .StructName }}(ctx context.Context, ids []int64) error
	Update{{ .StructName }}(ctx context.Context, req *v1.{{ .StructName }}Request) error
}

func New{{ .StructName }}Service(
	service *Service,
	{{ .StructNameLowerFirst }}Repo repository.{{ .StructName }}Repo,
) {{ .StructName }}Service {
	return &{{ .StructNameLowerFirst }}Service{
		Service:    service,
		{{ .StructNameLowerFirst }}Repo: {{ .StructNameLowerFirst }}Repo,
	}
}

type {{ .StructNameLowerFirst }}Service struct {
	*Service
	{{ .StructNameLowerFirst }}Repo repository.{{ .StructName }}Repo
}

func (s *{{ .StructNameLowerFirst }}Service) Get{{ .StructName }}(ctx context.Context, id int64) (*model.{{ .StructName }}, error) {
	return s.{{ .StructNameLowerFirst }}Repo.Get{{ .StructName }}(ctx, id)
}

func (s *{{ .StructNameLowerFirst }}Service) Get{{ .StructNamePlural }}(ctx context.Context, condition ...any) ([]model.{{ .StructName }}, error) {
	return s.{{ .StructNameLowerFirst }}Repo.Get{{ .StructNamePlural }}(ctx, condition...)
}

func (s *{{ .StructNameLowerFirst }}Service) Create{{ .StructName }}(ctx context.Context, req *v1.{{ .StructName }}Request) error {
	var err error
	var {{ .StructNameLowerFirst }} = new(model.{{ .StructName }})
	if err = copier.Copy({{ .StructNameLowerFirst }}, req); err != nil {
		return err
	}
	if err = s.{{ .StructNameLowerFirst }}Repo.Create{{ .StructName }}(ctx, {{ .StructNameLowerFirst }}); err != nil {
		return err
	}
	return nil
}

func (s *{{ .StructNameLowerFirst }}Service) Update{{ .StructName }}(ctx context.Context, req *v1.{{ .StructName }}Request) error {
	var err error

	var {{ .StructNameLowerFirst }} = new(model.{{ .StructName }})
	if err = copier.Copy({{ .StructNameLowerFirst }}, req); err != nil {
		return err
	}
	if err = s.{{ .StructNameLowerFirst }}Repo.Update{{ .StructName }}(ctx, {{ .StructNameLowerFirst }}); err != nil {
		return err
	}
	return nil
}

func (s *{{ .StructNameLowerFirst }}Service) Delete{{ .StructName }}(ctx context.Context, ids []int64) error {
	var err error
	if err = s.{{ .StructNameLowerFirst }}Repo.Delete{{ .StructName }}(ctx, ids); err != nil {
		return err
	}
	return err
}
