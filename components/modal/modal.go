package modal

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Type int

const (
	confirmModalType Type = iota
)

type Renderer interface {
	Render(echo.Context) error
}

// Maps the modal types to the appropriate templ function
var modalMap = map[Type]func(modalDetail) templ.Component{
	confirmModalType: confirmModal,
}

// Render renders the modal
func (m modalDetail) Render(c echo.Context) error {
	componentFunc, ok := modalMap[m.Type]
	if !ok {
		panic("unknown modal type")
	}

	modal := componentFunc(m)
	c.Response().Header().Set("HX-Reswap", "innerHTML")
	c.Response().Header().Set("HX-Retarget", "#modal-container")
	return modal.Render(c.Request().Context(), c.Response())
}

type modalDetail struct {
	Type          Type
	Heading       string
	Text          string
	CancelText    string
	ConfirmText   string
	ConfirmMethod string
	ConfirmAction string
}

func newModalDetail(
	modalType Type,
	heading string,
	text string,
	cancelText string,
	confirmText string,
	confirmMethod string,
	confirmAction string,
) modalDetail {
	if cancelText == "" {
		cancelText = "Cancel" // Default cancel button text
	}
	if confirmText == "" {
		confirmText = "OK" // Default confirm button text
	}
	if confirmMethod == "" {
		confirmMethod = "hx-post" // Default confirm method
	}
	return modalDetail{
		Type:          modalType,
		Heading:       heading,
		Text:          text,
		CancelText:    cancelText,
		ConfirmText:   confirmText,
		ConfirmMethod: confirmMethod,
		ConfirmAction: confirmAction,
	}
}

// Example

// <div id="modal-container"></div>

// func (ep *settingsEndpoint) deleteAllQuestionsHandler() echo.HandlerFunc {
//     return func(c echo.Context) error {
//         // Determine if this action has been confirmed
//         //   confirmed is true of the param is ?confirmed=true
//         //   confirmed is false if the param is ?confirmed=true or not present
//         confirmed, _ := strconv.ParseBool(c.Request().URL.Query().Get("confirmed"))

//         if !confirmed {
//             // Return a confirmation modal if the user has not confirmed the action
//             m := modal.NewConfirmModal(modal.ConfirmModalState{
//                 Heading:       "Delete all questions",
//                 Text:          "Are you sure you want to delete all of your questions? This action cannot be undone.",
//                 ConfirmText:   "Delete",    // text for the action button
//                 ConfirmMethod: "hx-delete", // The metod to be used for the action button
//                 ConfirmAction: c.Request().URL.Path + "?confirmed=true", // The current URL concatenated with the confirmed query param
//             })
//             return m.Render(c)
//         }

//         // Perform the deletion action - logic not implemented in this example
//         err := deleteAllQuestions()
//         return err
//     }
// }
