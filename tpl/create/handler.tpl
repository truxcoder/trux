package handler

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
    "go.uber.org/zap"
    v1 "{{ .ProjectName }}/api/v1"
    "{{ .ProjectName }}/internal/model"
    "{{ .ProjectName }}/internal/service"
    "net/http"
    "strconv"
)

type {{ .StructName }}Handler struct {
	*Handler
	{{ .StructNameLowerFirst }}Service service.{{ .StructName }}Service
}

func New{{ .StructName }}Handler(
    handler *Handler,
    {{ .StructNameLowerFirst }}Service service.{{ .StructName }}Service,
) *{{ .StructName }}Handler {
	return &{{ .StructName }}Handler{
		Handler:      handler,
		{{ .StructNameLowerFirst }}Service: {{ .StructNameLowerFirst }}Service,
	}
}

func (h *{{ .StructName }}Handler) Get{{ .StructName }}(ctx *gin.Context) {
    var (
		err    error
		id     int64
		{{ .StructNameLowerFirst }} *model.{{ .StructName }}
	)
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if {{ .StructNameLowerFirst }}, err = h.{{ .StructNameLowerFirst }}Service.Get{{ .StructName }}(ctx, id); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleSuccess(ctx, {{ .StructNameLowerFirst }})
}

func (h *{{ .StructName }}Handler) Get{{ .StructNamePlural }}(ctx *gin.Context) {
	var err error
	var req v1.{{ .StructName }}Request
	var {{ .StructNamePluralLowerFirst }} []model.{{ .StructName }}
	if err = jsoniter.Unmarshal([]byte(ctx.DefaultQuery("json", "{}")), &req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if {{ .StructNamePluralLowerFirst }}, err = h.{{ .StructNameLowerFirst }}Service.Get{{ .StructNamePlural }}(ctx, &req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleSuccess(ctx, {{ .StructNamePluralLowerFirst }})
}

func (h *{{ .StructName }}Handler) Create{{ .StructName }}(ctx *gin.Context) {
	var err error
	var req v1.{{ .StructName }}Request
	if err = ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if err = h.{{ .StructNameLowerFirst }}Service.Create{{ .StructName }}(ctx, &req); err != nil {
		h.logger.WithContext(ctx).Error("{{ .StructNameLowerFirst }}Service.Create{{ .StructName }} error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleCustomSuccess(ctx, v1.MsgCreateSuccess, nil)
}

func (h *{{ .StructName }}Handler) Update{{ .StructName }}(ctx *gin.Context) {
	var err error
	var req v1.{{ .StructName }}Request
	if err = ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if err = h.{{ .StructNameLowerFirst }}Service.Update{{ .StructName }}(ctx, &req); err != nil {
		h.logger.WithContext(ctx).Error("{{ .StructNameLowerFirst }}Service.Update{{ .StructName }} error", zap.Error(err))
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleCustomSuccess(ctx, v1.MsgUpdateSuccess, nil)
}

func (h *{{ .StructName }}Handler) Delete{{ .StructName }}(ctx *gin.Context) {
	var (
		err error
		id  int64
		ids []int64
	)
	if id, err = strconv.ParseInt(ctx.Param("id"), 10, 64); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	ids = append(ids, id)

	if err = h.{{ .StructNameLowerFirst }}Service.Delete{{ .StructName }}(ctx, ids); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	v1.HandleCustomSuccess(ctx, v1.MsgDeleteSuccess, nil)
}
