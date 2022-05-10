package gui

import "strings"

type Renderer struct {
	builder strings.Builder
}

func NewRenderer() *Renderer {
	return &Renderer{
		builder: strings.Builder{},
	}
}

//TODO: rewrite
func (rd *Renderer) Append(s string) {
	rd.builder.WriteString(s + "\n")
}

func (rd *Renderer) AppendNewLine() {
	rd.builder.WriteRune('\n')
}

func (rd *Renderer) AppendRune(r rune) {
	rd.builder.WriteRune(r)
}

func (rd *Renderer) Render() string {
	result := rd.builder.String()
	if len(result) < 1 {
		return result
	}
	return result[:len(result)-1]
}
