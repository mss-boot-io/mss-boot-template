/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 22:46
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 22:46
 */

package form

import (
	"time"

	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"github.com/mss-boot-io/mss-boot/pkg/response/curd"
)

type {{.model | ToCamel}}CreateReq struct {
{{ range .Columns }}
{{ if eq .action-create true }}
	//{{.comment}}
	{{.name | ToCamel}} {{.type}} `bson:"{{.name | LeftLower}}" json:"{{.name | LeftLower}}"`
{{ end }}
{{ end }}
	//创建时间
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

func (e *{{.model | ToCamel}}CreateReq) SetCreatedAt() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt
}

type {{.model | ToCamel}}UpdateReq struct {
	curd.OneID
{{ range .Columns }}
{{ if eq .action-update true }}
	//{{.comment}}
	{{.name | ToCamel}} {{.type}} `bson:"{{.name | LeftLower}}" json:"{{.name | LeftLower}}"`
{{ end }}
{{ end }}
	//更新时间
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

func (e *{{.model | ToCamel}}UpdateReq) SetUpdatedAt() {
	e.UpdatedAt = time.Now()
}

type {{.model | ToCamel}}GetReq struct {
	curd.OneID
}

type {{.model | ToCamel}}GetResp struct {
{{ range .Columns }}
{{ if eq .action-get true }}
	//{{.comment}}
	{{.name | ToCamel}} {{.type}} `bson:"{{.name | LeftLower}}" json:"{{.name | LeftLower}}"`
{{ end }}
{{ end }}
	//创建时间
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	//更新时间
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type {{.model | ToCamel}}DeleteReq struct {
	curd.OneID
}

type {{.model | ToCamel}}ListReq struct {
	curd.Pagination
{{ range .Columns }}
{{ if eq .action-search true }}
	//{{.comment}}
	{{.name | ToCamel}} {{.type}} `bson:"{{.name | LeftLower}}" form:"{{.name | LeftLower}} query:"{{.name | LeftLower}} search:"type:{{.search-type}};column:{{.name | LeftLower}}"`
{{ end }}
{{ end }}
}

type {{.model | ToCamel}}ListItem struct {
{{ range .Columns }}
{{ if eq .action-list true }}
	//{{.comment}}
	{{.name | ToCamel}} {{.type}} `bson:"{{.name | LeftLower}}" json:"{{.name | LeftLower}}"`
{{ end }}
{{ end }}
}