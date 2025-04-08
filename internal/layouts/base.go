package layouts

import (
	"maragu.dev/gomponents"
)

type BaseLayout struct{}

// Node satisfies [layouts.LayoutProvider].
func (l *Layout) Node() gomponents.Node {
	return nil
}
