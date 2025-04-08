// Content engines process files during site generation.
package engine

import "io/fs"

type Engine interface {
	Name() string
	Parse([]byte, fs.FileInfo)
	Render() ([]byte, error)
}

type DefaultEngine struct {
	content []byte
	info    fs.FileInfo
}

// The default engine implementation just copies files.
func Default() *DefaultEngine {
	return &DefaultEngine{}
}

func (e *DefaultEngine) Name() string {
	return e.info.Name()
}

func (e *DefaultEngine) Parse(content []byte, info fs.FileInfo) {
	e.content = content
	e.info = info
}

func (e *DefaultEngine) Render() ([]byte, error) {
	return e.content, nil
}
