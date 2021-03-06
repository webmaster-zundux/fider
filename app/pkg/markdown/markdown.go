package markdown

import (
	"html/template"
	"strings"

	"github.com/russross/blackfriday"
)

var mdExtns = 0 |
	blackfriday.EXTENSION_TABLES |
	blackfriday.EXTENSION_AUTOLINK |
	blackfriday.EXTENSION_FENCED_CODE |
	blackfriday.EXTENSION_TITLEBLOCK |
	blackfriday.EXTENSION_STRIKETHROUGH |
	blackfriday.EXTENSION_DEFINITION_LISTS |
	blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
	blackfriday.EXTENSION_HARD_LINE_BREAK

var simpleHTMLExtensions = 0 |
	blackfriday.HTML_USE_XHTML |
	blackfriday.HTML_USE_SMARTYPANTS |
	blackfriday.HTML_SMARTYPANTS_FRACTIONS |
	blackfriday.HTML_SMARTYPANTS_DASHES |
	blackfriday.HTML_SKIP_IMAGES |
	blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

var simpleRenderer = SimpleRenderer(simpleHTMLExtensions)

// Simple turns a markdown into HTML using few rules
func Simple(input string) template.HTML {
	output := blackfriday.Markdown([]byte(input), simpleRenderer, mdExtns)

	return template.HTML(strings.TrimSpace(string(output)))
}

var fullHTMLExtensions = 0 |
	blackfriday.HTML_USE_XHTML |
	blackfriday.HTML_USE_SMARTYPANTS |
	blackfriday.HTML_SMARTYPANTS_FRACTIONS |
	blackfriday.HTML_SMARTYPANTS_DASHES |
	blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

var fullRenderer = blackfriday.HtmlRenderer(fullHTMLExtensions, "", "")

// Full turns a markdown into HTML using all rules
func Full(input string) template.HTML {
	output := blackfriday.Markdown([]byte(input), fullRenderer, mdExtns)

	return template.HTML(strings.TrimSpace(string(output)))
}

var textRenderer = TextRenderer()

// PlainText parses given markdown input and return only the text
func PlainText(input string) string {
	output := blackfriday.Markdown([]byte(input), textRenderer, mdExtns)
	return strings.TrimSpace(string(output))
}
