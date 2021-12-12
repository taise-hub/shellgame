package controller

type Context interface {
	Param(key string) string
	PostForm(key string) string
	HTML(code int, name string, obj interface{})
	JSON(code int, obj interface{})
	Redirect(code int, name string)
	MustGet(key string) interface{}
}
