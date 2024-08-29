package form

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Type int

const (
	genericFormType  Type = iota
	loginFormType    Type = 1
	registerFormType Type = 2
)

type Renderer interface {
	Render(echo.Context) error
}

// Maps the form types to the appropriate templ function
var formMap = map[Type]func(formDetail) templ.Component{
	genericFormType: form,
}

type formDetail struct {
	Type      Type
	Heading   string
	Method    string
	ActionURL string
	Class     string
}

func (f formDetail) Render(c echo.Context) error {
	componentFunc, ok := formMap[f.Type]
	if !ok {
		panic("unknown form type")
	}
	frm := componentFunc(f)

	return frm.Render(c.Request().Context(), c.Response())
}

func newFormDetail(t Type, heading, method, actionUrl, class string) formDetail {
	if method == "" {
		method = "post"
	}

	return formDetail{
		Type:      t,
		Heading:   heading,
		Method:    method,
		ActionURL: actionUrl,
		Class:     class,
	}
}
