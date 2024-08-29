package input

type Type string

const (
	textType      Type = "text"
	passwordType  Type = "password"
	emailType     Type = "email"
	numberType    Type = "number"
	checkboxType  Type = "checkbox"
	radioType     Type = "radio"
	colorType     Type = "color"
	dateType      Type = "date"
	date_timeType Type = "datetime-local"
	fileType      Type = "file"
)
