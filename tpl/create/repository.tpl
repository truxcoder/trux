package repository

import (
    "context"
	"{{ .ProjectName }}/internal/model"
)

type {{ .StructName }}Repo interface {
	Get{{ .StructName }}(ctx context.Context, id int64) (*model.{{ .StructName }}, error)
	Get{{ .StructNamePlural }}(ctx context.Context, condition ...any) ([]model.{{ .StructName }}, error)
	Create{{ .StructName }}(ctx context.Context, {{ .StructName }} *model.{{ .StructName }}) error
	Delete{{ .StructName }}(ctx context.Context, ids []int64) error
	Update{{ .StructName }}(ctx context.Context, {{ .StructName }} *model.{{ .StructName }}) error
}

func New{{ .StructName }}Repo(
	repository *Repository,
) {{ .StructName }}Repo {
	return &{{ .StructNameLowerFirst }}Repo{
		Repository: repository,
	}
}

type {{ .StructNameLowerFirst }}Repo struct {
	*Repository
}

func (r *{{ .StructNameLowerFirst }}Repo) Get{{ .StructName }}(ctx context.Context, id int64) (*model.{{ .StructName }}, error) {
	var {{ .StructNameLowerFirst }} model.{{ .StructName }}

	return &{{ .StructNameLowerFirst }}, nil
}

func (r *{{ .StructNameLowerFirst }}Repo) Get{{ .StructNamePlural }}(ctx context.Context, condition ...any) ([]model.{{ .StructName }}, error) {
	var {{ .StructNamePlural }} []model.{{ .StructName }}
	var err error
	if len(condition) > 0 {
		err = r.DB(ctx).Where(condition[0], condition[1:]...).Find(&{{ .StructNamePlural }}).Error
	} else {
		err = r.DB(ctx).Find(&{{ .StructNamePlural }}).Error
	}
	return {{ .StructNamePlural }}, err
}

func (r *{{ .StructNameLowerFirst }}Repo) Create{{ .StructName }}(ctx context.Context, {{ .StructName }} *model.{{ .StructName }}) error {
	if err := r.DB(ctx).Create({{ .StructName }}).Error; err != nil {
		return err
	}
	return nil
}

func (r *{{ .StructNameLowerFirst }}Repo) Update{{ .StructName }}(ctx context.Context, {{ .StructName }} *model.{{ .StructName }}) error {
	if err := r.DB(ctx).Updates({{ .StructName }}).Error; err != nil {
		return err
	}
	return nil
}

func (r *{{ .StructNameLowerFirst }}Repo) Delete{{ .StructName }}(ctx context.Context, ids []int64) error {
	if err := r.DB(ctx).Delete(model.{{ .StructName }}{}, ids).Error; err != nil {
		return err
	}
	return nil
}
