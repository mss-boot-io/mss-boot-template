/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 9:24
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 9:24
 */

package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/bsonx"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/enum"
)

func init() {
	mongodb.AppendTable(&{{.model | ToCamel}}{})
}

// {{.model | ToCamel}} {{.comment}}
type {{.model | ToCamel}} struct {
	{{ range .columns }}
	{{.name | ToCamel}} {{.type}} `bson:"{{.name | LeftLower}}" json:"{{.name | LeftLower}}"`
	{{end}}
	CreatedAt   time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt" bson:"updatedAt"`
}

func ({{.model | ToCamel}}) TableName() string {
	return "{{.model}}"
}

func (e *{{.model | ToCamel}}) Make() {
	// ops := options.Index()
	// ops.SetName("name")
	// ops.SetUnique(true)
	// _, err := e.C().Indexes().CreateOne(context.TODO(), mongo.IndexModel{
	// 	Keys:    bsonx.Doc{
	//	{"name", bsonx.Int32(1)},
	//},
	// 	Options: ops,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func (e *{{.model | ToCamel}}) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}
