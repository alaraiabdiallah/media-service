package app

type (
	AppContext struct{
		version string
	}
	Map map[string]interface{}
)

func (ctx *AppContext) GetVersion() string{
	return ctx.version
}

func Ctx() *AppContext {
	return &AppContext{
		version: "0.0.1",
	}
}