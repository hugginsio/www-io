package layouts

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
)

func Default(title string) gomponents.Node {
	return components.HTML5(components.HTML5Props{
		Language: "en-US",
	})
}
