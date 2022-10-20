package plugins

import (
	"plugins/file"
)

func NewPlugin() func() string {
	return func() string {
		return "new a file plugin -->" + file.NewFile()
	}
}
