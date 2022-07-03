package main

import (
	"go/constant"
	"go/token"
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/mongo_plugin"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Symbols = map[string]map[string]reflect.Value{}

//go:generate go install github.com/traefik/yaegi/cmd/yaegi@v0.10.0
//go:generate yaegi extract gofiber-demo/plugins/env_plugin
//go:generate yaegi extract gofiber-demo/plugins/app_plugin
//go:generate yaegi extract gofiber-demo/plugins/mongo_plugin
//go:generate yaegi extract go.mongodb.org/mongo-driver/bson/primitive
func init() {
	Symbols["gofiber-demo/plugins/env_plugin/env_plugin"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ENV":              reflect.ValueOf(&env_plugin.ENV).Elem(),
		"GetEnv":           reflect.ValueOf(env_plugin.GetEnv),
		"GetEnvInt":        reflect.ValueOf(env_plugin.GetEnvInt),
		"GetEnvTimeSecond": reflect.ValueOf(env_plugin.GetEnvTimeSecond),
		"LoadEnv":          reflect.ValueOf(env_plugin.LoadEnv),
		"Register":         reflect.ValueOf(env_plugin.Register),
	}
	Symbols["gofiber-demo/plugins/app_plugin/app_plugin"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New":      reflect.ValueOf(app_plugin.New),
		"Register": reflect.ValueOf(app_plugin.Register),

		// type definitions
		"Application":  reflect.ValueOf((*app_plugin.Application)(nil)),
		"CallbackFunc": reflect.ValueOf((*app_plugin.CallbackFunc)(nil)),
	}
	Symbols["gofiber-demo/plugins/mongo_plugin/mongo_plugin"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Register": reflect.ValueOf(mongo_plugin.Register),
	}
	Symbols["go.mongodb.org/mongo-driver/bson/primitive/primitive"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CompareTimestamp":          reflect.ValueOf(primitive.CompareTimestamp),
		"ErrInvalidHex":             reflect.ValueOf(&primitive.ErrInvalidHex).Elem(),
		"ErrParseInf":               reflect.ValueOf(&primitive.ErrParseInf).Elem(),
		"ErrParseNaN":               reflect.ValueOf(&primitive.ErrParseNaN).Elem(),
		"ErrParseNegInf":            reflect.ValueOf(&primitive.ErrParseNegInf).Elem(),
		"IsValidObjectID":           reflect.ValueOf(primitive.IsValidObjectID),
		"MaxDecimal128Exp":          reflect.ValueOf(constant.MakeFromLiteral("6111", token.INT, 0)),
		"MinDecimal128Exp":          reflect.ValueOf(constant.MakeFromLiteral("-6176", token.INT, 0)),
		"NewDateTimeFromTime":       reflect.ValueOf(primitive.NewDateTimeFromTime),
		"NewDecimal128":             reflect.ValueOf(primitive.NewDecimal128),
		"NewObjectID":               reflect.ValueOf(primitive.NewObjectID),
		"NewObjectIDFromTimestamp":  reflect.ValueOf(primitive.NewObjectIDFromTimestamp),
		"NilObjectID":               reflect.ValueOf(&primitive.NilObjectID).Elem(),
		"ObjectIDFromHex":           reflect.ValueOf(primitive.ObjectIDFromHex),
		"ParseDecimal128":           reflect.ValueOf(primitive.ParseDecimal128),
		"ParseDecimal128FromBigInt": reflect.ValueOf(primitive.ParseDecimal128FromBigInt),

		// type definitions
		"A":             reflect.ValueOf((*primitive.A)(nil)),
		"Binary":        reflect.ValueOf((*primitive.Binary)(nil)),
		"CodeWithScope": reflect.ValueOf((*primitive.CodeWithScope)(nil)),
		"D":             reflect.ValueOf((*primitive.D)(nil)),
		"DBPointer":     reflect.ValueOf((*primitive.DBPointer)(nil)),
		"DateTime":      reflect.ValueOf((*primitive.DateTime)(nil)),
		"Decimal128":    reflect.ValueOf((*primitive.Decimal128)(nil)),
		"E":             reflect.ValueOf((*primitive.E)(nil)),
		"JavaScript":    reflect.ValueOf((*primitive.JavaScript)(nil)),
		"M":             reflect.ValueOf((*primitive.M)(nil)),
		"MaxKey":        reflect.ValueOf((*primitive.MaxKey)(nil)),
		"MinKey":        reflect.ValueOf((*primitive.MinKey)(nil)),
		"Null":          reflect.ValueOf((*primitive.Null)(nil)),
		"ObjectID":      reflect.ValueOf((*primitive.ObjectID)(nil)),
		"Regex":         reflect.ValueOf((*primitive.Regex)(nil)),
		"Symbol":        reflect.ValueOf((*primitive.Symbol)(nil)),
		"Timestamp":     reflect.ValueOf((*primitive.Timestamp)(nil)),
		"Undefined":     reflect.ValueOf((*primitive.Undefined)(nil)),
	}
}

func main() {
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)
	i.Use(Symbols)

	i.EvalPath("./console/main.go")

	i.REPL()
}
