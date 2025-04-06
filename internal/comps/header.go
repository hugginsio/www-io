package comps

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Navbar() gomponents.Node {
	return html.Nav(html.Class("bg-gray-800 flex items-center space-x-4"),
		html.A(html.Href("/"), gomponents.Text("Home")),
		html.A(html.Href("/about"), gomponents.Text("About")),
	)
}
