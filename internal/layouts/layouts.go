package layouts

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

type Layout struct {
	// TODO: name
	// TODO: extra files
}

// TODO: what am I doing
type LayoutProvider interface {
	Layout() gomponents.Node
}

func Base(title string, description string, head []gomponents.Node, body []gomponents.Node) gomponents.Node {
	return components.HTML5(components.HTML5Props{
		Language:    "en",
		Title:       title,
		Description: description,
		Head: append(
			head,
			html.Link(
				html.Rel("preconnect"),
				html.Href("https://fonts.googleapis.com"),
			),
			html.Link(
				html.Rel("preconnect"),
				html.Href("https://fonts.gstatic.com"),
				gomponents.Attr("crossorigin"),
			),
			html.Link(
				html.Href("https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400&display=swap"),
				html.Rel("stylesheet"),
			),
			html.Link(
				html.Href("https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@700&display=swap"),
				html.Rel("stylesheet"),
			),
			// TODO: replace with tailwind CLI before merge
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
			),
		),
		Body: body,
	})
}
