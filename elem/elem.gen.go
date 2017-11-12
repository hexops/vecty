//go:generate go run generate.go

// Package elem defines markup to create DOM elements.
//
// Generated from "HTML element reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under
// CC-BY-SA 2.5.
package elem

import "github.com/gopherjs/vecty"

// Anchor (or anchor element) creates a hyperlink to other web pages, files,
// locations within the same page, email addresses, or any other URL.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("a", children...)
}

// Abbreviation represents an abbreviation and optionally provides a full
// description for it. If present, the title attribute must contain this full
// description and nothing else.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("abbr", children...)
}

// Address supplies contact information for its nearest <article> or <body>
// ancestor; in the latter case, it applies to the whole document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("address", children...)
}

// Area defines a hot-spot region on an image, and optionally associates it
// with a hypertext link. This element is used only within a <map> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("area", children...)
}

// Article represents a self-contained composition in a document, page,
// application, or site, which is intended to be independently distributable or
// reusable (e.g., in syndication). Examples include: a forum post, a magazine
// or newspaper article, or a blog entry.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("article", children...)
}

// Aside represents a section of a document with content connected tangentially
// to the main content of the document (often presented as a sidebar).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("aside", children...)
}

// Audio is used to embed sound content in documents. It may contain one or
// more audio sources, represented using the src attribute or the <source>
// element: the browser will choose the most suitable one. It can also be the
// destination for streamed media, using a MediaStream.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("audio", children...)
}

// Bold represents a span of text stylistically different from normal text,
// without conveying any special importance or relevance, and that is typically
// rendered in boldface.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("b", children...)
}

// Base specifies the base URL to use for all relative URLs contained within a
// document. There can be only one <base> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("base", children...)
}

// BidirectionalIsolation (bidirectional isolation) isolates a span of text
// that might be formatted in a different direction from other text outside it.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("bdi", children...)
}

// BidirectionalOverride (bidirectional override) is used to override the
// current directionality of text. It causes the directionality of the
// characters to be ignored in favor of the specified directionality.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("bdo", children...)
}

// BlockQuote (or HTML Block Quotation Element) indicates that the enclosed
// text is an extended quotation. Usually, this is rendered visually by
// indentation (see Notes for how to change it). A URL for the source of the
// quotation may be given using the cite attribute, while a text representation
// of the source can be given using the <cite> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("blockquote", children...)
}

// Body represents the content of an HTML document. There can be only one
// <body> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func Body(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("body", children...)
}

// Break produces a line break in text (carriage-return). It is useful for
// writing a poem or an address, where the division of lines is significant.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("br", children...)
}

// Button represents a clickable button.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("button", children...)
}

// Use the HTML <canvas> element with the canvas scripting API to draw graphics
// and animations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("canvas", children...)
}

// Caption represents the title of a table. Though it is always the first
// descendant of a <table>, its styling, using CSS, may place it elsewhere,
// relative to the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("caption", children...)
}

// Citation represents a reference to a creative work. It must include the
// title of a work or a URL reference, which may be in an abbreviated form
// according to the conventions used for the addition of citation metadata.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("cite", children...)
}

// Code represents a fragment of computer code. By default, it is displayed in
// the browser's default monospace font.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("code", children...)
}

// Column defines a column within a table and is used for defining common
// semantics on all common cells. It is generally found within a <colgroup>
// element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("col", children...)
}

// ColumnGroup defines a group of columns within a table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("colgroup", children...)
}

// Data links a given content with a machine-readable translation. If the
// content is time- or date-related, the <time> element must be used.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("data", children...)
}

// DataList contains a set of <option> elements that represent the values
// available for other controls.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("datalist", children...)
}

// Description indicates the description of a term in a description list
// (<dl>).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dd", children...)
}

// DeletedText represents a range of text that has been deleted from a
// document. This element is often (but need not be) rendered with
// strike-through text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("del", children...)
}

// Details is used as a disclosure widget from which the user can retrieve
// additional information.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("details", children...)
}

// Definition represents the defining instance of a term.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dfn", children...)
}

// Dialog represents a dialog box or other interactive component, such as an
// inspector or window.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dialog", children...)
}

// Div is the generic container for flow content and does not inherently
// represent anything. Use it to group elements for purposes such as styling
// (using the class or id attributes), marking a section of a document in a
// different language (using the lang attribute), and so on.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("div", children...)
}

// DescriptionList represents a description list. The element encloses a list
// of groups of terms and descriptions. Common uses for this element are to
// implement a glossary or to display metadata (a list of key-value pairs).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dl", children...)
}

// DefinitionTerm identifies a term in a description list. This element can
// occur only as a child element of a <dl>. It is usually followed by a <dd>
// element; however, multiple <dt> elements in a row indicate several terms
// that are all defined by the immediate next <dd> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dt", children...)
}

// Emphasis marks text that has stress emphasis. The <em> element can be
// nested, with each level of nesting indicating a greater degree of emphasis.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("em", children...)
}

// Embed represents an integration point for an external application or
// interactive content (in other words, a plug-in).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("embed", children...)
}

// FieldSet is used to group several controls as well as labels (<label>)
// within a web form.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("fieldset", children...)
}

// FigureCaption represents a caption or a legend associated with a figure or
// an illustration described by the rest of the data of the <figure> element
// which is its immediate ancestor.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("figcaption", children...)
}

// Figure represents self-contained content, frequently with a caption
// (<figcaption>), and is typically referenced as a single unit.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("figure", children...)
}

// Footer represents a footer for its nearest sectioning content or sectioning
// root element. A footer typically contains information about the author of
// the section, copyright data or links to related documents.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("footer", children...)
}

// Form represents a document section that contains interactive controls to
// submit information to a web server.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("form", children...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading1(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h1", children...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading2(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h2", children...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading3(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h3", children...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading4(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h4", children...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading5(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h5", children...)
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Heading6(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h6", children...)
}

// Header represents a group of introductory or navigational aids. It may
// contain some heading elements but also other elements like a logo, a search
// form, and so on.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("header", children...)
}

// HeadingsGroup represents a multi-level heading for a section of a document.
// It groups a set of <h1>–<h6> elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("hgroup", children...)
}

// HorizontalRule represents a thematic break between paragraph-level elements
// (for example, a change of scene in a story, or a shift of topic with a
// section). In previous versions of HTML, it represented a horizontal rule. It
// may still be displayed as a horizontal rule in visual browsers, but is now
// defined in semantic terms, rather than presentational terms.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("hr", children...)
}

// Italic represents a range of text that is set off from the normal text for
// some reason, for example, technical terms, foreign language phrases, or
// fictional character thoughts. It is typically displayed in italic type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("i", children...)
}

// InlineFrame represents a nested browsing context, effectively embedding
// another HTML page into the current page. In HTML 4.01, a document may
// contain a head and a body or a head and a frameset, but not both a body and
// a frameset. However, an <iframe> can be used within a normal document body.
// Each browsing context has its own session history and active document. The
// browsing context that contains the embedded content is called the parent
// browsing context. The top-level browsing context (which has no parent) is
// typically the browser window.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("iframe", children...)
}

// Image represents an image in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("img", children...)
}

// Input is used to create interactive controls for web-based forms in order to
// accept data from the user.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("input", children...)
}

// InsertedText represents a range of text that has been added to a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ins", children...)
}

// KeyboardInput represents user input and produces an inline element displayed
// in the browser's default monospace font.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("kbd", children...)
}

// Label represents a caption for an item in a user interface.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("label", children...)
}

// Legend represents a caption for the content of its parent <fieldset>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("legend", children...)
}

// ListItem is used to represent an item in a list. It must be contained in a
// parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu
// (<menu>). In menus and unordered lists, list items are usually displayed
// using bullet points. In ordered lists, they are usually displayed with an
// ascending counter on the left, such as a number or letter.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("li", children...)
}

// Link specifies relationships between the current document and an external
// resource. Possible uses for this element include defining a relational
// framework for navigation. This element is most used to link to style sheets.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("link", children...)
}

// Main represents the main content of the <body> of a document, portion of a
// document, or application. The main content area consists of content that is
// directly related to, or expands upon the central topic of, a document or the
// central functionality of an application.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("main", children...)
}

// Map is used with <area> elements to define an image map (a clickable link
// area).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("map", children...)
}

// Mark represents highlighted text, i.e., a run of text marked for reference
// purpose, due to its relevance in a particular context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("mark", children...)
}

// Menu represents a group of commands that a user can perform or activate.
// This includes both list menus, which might appear across the top of a
// screen, as well as context menus, such as those that might appear underneath
// a button after it has been clicked.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("menu", children...)
}

// MenuItem represents a command that a user is able to invoke through a popup
// menu. This includes context menus, as well as menus that might be attached
// to a menu button.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem
func MenuItem(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("menuitem", children...)
}

// Meta represents metadata that cannot be represented by other HTML
// meta-related elements, like <base>, <link>, <script>, <style> or <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("meta", children...)
}

// Meter represents either a scalar value within a known range or a fractional
// value.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("meter", children...)
}

// Navigation represents a section of a page whose purpose is to provide
// navigation links, either within the current document or to other documents.
// Common examples of navigation sections are menus, tables of contents, and
// indexes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("nav", children...)
}

// <noframes> is an HTML element which is used to support browsers which are
// not able to support <frame> elements or configured to do so.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes
func NoFrames(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("noframes", children...)
}

// NoScript defines a section of HTML to be inserted if a script type on the
// page is unsupported or if scripting is currently turned off in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("noscript", children...)
}

// Object represents an external resource, which can be treated as an image, a
// nested browsing context, or a resource to be handled by a plugin.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("object", children...)
}

// OrderedList represents an ordered list of items, typically rendered as a
// numbered list.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ol", children...)
}

// OptionsGroup creates a grouping of options within a <select> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("optgroup", children...)
}

// Option is used to define an item contained in a <select>, an <optgroup>, or
// a <datalist> element. As such, <option> can represent menu items in popups
// and other lists of items in an HTML document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("option", children...)
}

// Output represents the result of a calculation or user action.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("output", children...)
}

// Paragraph represents a paragraph of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("p", children...)
}

// Parameter defines parameters for an <object> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("param", children...)
}

// Picture is a container used to specify multiple <source> elements for a
// specific <img> contained in it. The browser will choose the most suitable
// source according to the current layout of the page (the constraints of the
// box the image will appear in) and the device it will be displayed on (e.g. a
// normal or hiDPI device.)
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("picture", children...)
}

// Preformatted represents preformatted text. Text within this element is
// typically displayed in a non-proportional ("monospace") font exactly as it
// is laid out in the file. Whitespace inside this element is displayed as
// typed.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("pre", children...)
}

// Progress represents the completion progress of a task, typically displayed
// as a progress bar.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("progress", children...)
}

// Quote indicates that the enclosed text is a short inline quotation. This
// element is intended for short quotations that don't require paragraph
// breaks; for long quotations use the <blockquote> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("q", children...)
}

// RubyParenthesis is used to provide fall-back parentheses for browsers that
// do not support display of ruby annotations using the <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("rp", children...)
}

// RubyText embraces pronunciation of characters presented in a ruby
// annotations, which are used to describe the pronunciation of East Asian
// characters. This element is always used inside a <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("rt", children...)
}

// RubyTextContainer embraces semantic annotations of characters presented in a
// ruby of <rb> elements used inside of <ruby> element. <rb> elements can have
// both pronunciation (<rt>) and semantic (<rtc>) annotations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func RubyTextContainer(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("rtc", children...)
}

// Ruby represents a ruby annotation. Ruby annotations are for showing
// pronunciation of East Asian characters.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ruby", children...)
}

// Strikethrough renders text with a strikethrough, or a line through it. Use
// the <s> element to represent things that are no longer relevant or no longer
// accurate. However, <s> is not appropriate when indicating document edits;
// for that, use the <del> and <ins> elements, as appropriate.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("s", children...)
}

// Sample is an element intended to identify sample output from a computer
// program. It is usually displayed in the browser's default monotype font
// (such as Lucida Console).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("samp", children...)
}

// Script is used to embed or reference an executable script.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("script", children...)
}

// Section represents a standalone section of functionality contained within an
// HTML document, typically with a heading, which doesn't have a more specific
// semantic element to represent it.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("section", children...)
}

// Select represents a control that provides a menu of options:
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("select", children...)
}

// Slot—part of the Web Components technology suite—is a placeholder inside
// a web component that you can fill with your own markup, which lets you
// create separate DOM trees and present them together.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot
func Slot(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("slot", children...)
}

// Small makes the text font size one size smaller (for example, from large to
// medium, or from small to x-small) down to the browser's minimum font size.
// In HTML5, this element is repurposed to represent side-comments and small
// print, including copyright and legal text, independent of its styled
// presentation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("small", children...)
}

// Source specifies multiple media resources for either the <picture>, the
// <audio> or the <video> element. It is an empty element. It is commonly used
// to serve the same media content in multiple formats supported by different
// browsers.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("source", children...)
}

// Span is a generic inline container for phrasing content, which does not
// inherently represent anything. It can be used to group elements for styling
// purposes (using the class or id attributes), or because they share attribute
// values, such as lang.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("span", children...)
}

// Strong gives text strong importance, and is typically displayed in bold.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("strong", children...)
}

// Style contains style information for a document, or part of a document. By
// default, the style instructions written inside that element are expected to
// be CSS.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("style", children...)
}

// Subscript defines a span of text that should be displayed, for typographic
// reasons, lower, and often smaller, than the main span of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("sub", children...)
}

// Summary is used as a summary, caption, or legend for the content of a
// <details> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("summary", children...)
}

// Superscript defines a span of text that should be displayed, for typographic
// reasons, higher, and often smaller, than the main span of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("sup", children...)
}

// Table represents tabular data — that is, information expressed via a
// two-dimensional data table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("table", children...)
}

// TableBody groups one or more <tr> elements as the body of a <table> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("tbody", children...)
}

// TableData defines a cell of a table that contains data. It participates in
// the table model.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("td", children...)
}

// Template is a mechanism for holding client-side content that is not to be
// rendered when a page is loaded but may subsequently be instantiated during
// runtime using JavaScript.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("template", children...)
}

// TextArea represents a multi-line plain-text editing control.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("textarea", children...)
}

// TableFoot defines a set of rows summarizing the columns of the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("tfoot", children...)
}

// TableHeader defines a cell as header of a group of table cells. The exact
// nature of this group is defined by the scope and headers attributes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("th", children...)
}

// TableHead defines a set of rows defining the head of the columns of the
// table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("thead", children...)
}

// Time represents either a time on a 24-hour clock or a precise date in the
// Gregorian calendar (with optional time and timezone information).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("time", children...)
}

// Title defines the title of the document, shown in a browser's title bar or
// on the page's tab. It can only contain text, and any contained tags are
// ignored.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("title", children...)
}

// TableRow defines a row of cells in a table. Those can be a mix of <td> and
// <th> elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("tr", children...)
}

// Track is used as a child of the media elements <audio> and <video>. It lets
// you specify timed text tracks (or time-based data), for example to
// automatically handle subtitles. The tracks are formatted in WebVTT format
// (.vtt files) — Web Video Text Tracks.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("track", children...)
}

// Underline renders text with an underline, a line under the baseline of its
// content. In HTML5, this element represents a span of text with an
// unarticulated, though explicitly rendered, non-textual annotation, such as
// labeling the text as being a proper name in Chinese text (a Chinese proper
// name mark), or labeling the text as being misspelled.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("u", children...)
}

// UnorderedList represents an unordered list of items, typically rendered as a
// bulleted list.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ul", children...)
}

// Variable represents a variable in a mathematical expression or a programming
// context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("var", children...)
}

// Use the HTML <video> element to embed video content in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("video", children...)
}

// WordBreakOpportunity represents a word break opportunity—a position within
// text where the browser may optionally break a line, though its line-breaking
// rules would not otherwise create a break at that location.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(children ...vecty.ComponentOrHTML) *vecty.HTML {
	return vecty.Tag("wbr", children...)
}
