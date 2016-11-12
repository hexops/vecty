//go:generate go run generate.go

// Package elem defines markup to create DOM elements.
//
// Generated from "HTML element reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under
// CC-BY-SA 2.5.
package elem

import "github.com/gopherjs/vecty"

// The HTML Anchor Element (<a>) defines a hyperlink to a location on the same
// page or any other page on the Web. It can also be used (in an obsolete way)
// to create an anchor point—a destination for hyperlinks within the content
// of a page, so that links aren't limited to connecting simply to the top of a
// page.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("a", markup...)
}

// Abbreviation (or HTML Abbreviation Element) represents an abbreviation and
// optionally provides a full description for it. If present, the title
// attribute must contain this full description and nothing else.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("abbr", markup...)
}

// Address supplies contact information for its nearest <article> or <body>
// ancestor; in the latter case, it applies to the whole document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("address", markup...)
}

// Area defines a hot-spot region on an image, and optionally associates it
// with a hypertext link. This element is used only within a <map> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("area", markup...)
}

// Article represents a self-contained composition in a document, page,
// application, or site, which is intended to be independently distributable or
// reusable (e.g., in syndication). This could be a forum post, a magazine or
// newspaper article, a blog entry, an object, or any other independent item of
// content. Each <article> should be identified, typically by including a
// heading (<h1>-<h6> element) as a child of the <article> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("article", markup...)
}

// Aside represents a section of the page with content connected tangentially
// to the rest, which could be considered separate from that content. These
// sections are often represented as sidebars or inserts. They often contain
// the definitions on the sidebars, such as definitions from the glossary;
// there may also be other types of information, such as related
// advertisements; the biography of the author; web applications; profile
// information or related links on the blog.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("aside", markup...)
}

// Audio is used to embed sound content in documents. It may contain one or
// more audio sources, represented using the src attribute or the <source>
// element; the browser will choose the most suitable one.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("audio", markup...)
}

// Bold represents a span of text stylistically different from normal text,
// without conveying any special importance or relevance. It is typically used
// for keywords in a summary, product names in a review, or other spans of text
// whose typical presentation would be boldfaced. Another example of its use is
// to mark the lead sentence of each paragraph of an article.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("b", markup...)
}

// Base specifies the base URL to use for all relative URLs contained within a
// document. There can be only one <base> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("base", markup...)
}

// BidirectionalIsolation (or Bi-Directional Isolation Element) isolates a span
// of text that might be formatted in a different direction from other text
// outside it.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("bdi", markup...)
}

// BidirectionalOverride (or HTML bidirectional override element) is used to
// override the current directionality of text. It causes the directionality of
// the characters to be ignored in favor of the specified directionality.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("bdo", markup...)
}

// BlockQuote (or HTML Block Quotation Element) indicates that the enclosed
// text is an extended quotation. Usually, this is rendered visually by
// indentation (see Notes for how to change it). A URL for the source of the
// quotation may be given using the cite attribute, while a text representation
// of the source can be given using the <cite> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("blockquote", markup...)
}

// Body represents the content of an HTML document. There can be only one
// <body> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func Body(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("body", markup...)
}

// The HTML element line break <br> produces a line break in text
// (carriage-return). It is useful for writing a poem or an address, where the
// division of lines is significant.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("br", markup...)
}

// Button represents a clickable button.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("button", markup...)
}

// Canvas can be used to draw graphics via scripting (usually JavaScript). For
// example, it can be used to draw graphs, make photo compositions or even
// perform animations. You may (and should) provide alternate content inside
// the <canvas> block. That content will be rendered both on older browsers
// that don't support canvas and in browsers with JavaScript disabled.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("canvas", markup...)
}

// Caption (or HTML Table Caption Element) represents the title of a table.
// Though it is always the first descendant of a <table>, its styling, using
// CSS, may place it elsewhere, relative to the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("caption", markup...)
}

// The HTML Citation Element (<cite>) represents a reference to a creative
// work. It must include the title of a work or a URL reference, which may be
// in an abbreviated form according to the conventions used for the addition of
// citation metadata.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("cite", markup...)
}

// The HTML Code Element (<code>) represents a fragment of computer code. By
// default, it is displayed in the browser's default monospace font.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("code", markup...)
}

// The HTML Table Column Element (<col>) defines a column within a table and is
// used for defining common semantics on all common cells. It is generally
// found within a <colgroup> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("col", markup...)
}

// The HTML Table Column Group Element (<colgroup>) defines a group of columns
// within a table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("colgroup", markup...)
}

// Data links a given content with a machine-readable translation. If the
// content is time- or date-related, the <time> must be used.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("data", markup...)
}

// The HTML Datalist Element (<datalist>) contains a set of <option> elements
// that represent the values available for other controls.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("datalist", markup...)
}

// Description (HTML Description Element) indicates the description of a term
// in a description list (<dl>) element. This element can occur only as a child
// element of a description list and it must follow a <dt> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dd", markup...)
}

// The HTML Deleted Text Element (<del>) represents a range of text that has
// been deleted from a document. This element is often (but need not be)
// rendered with strike-through text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("del", markup...)
}

// Details is used as a disclosure widget from which the user can retrieve
// additional information.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("details", markup...)
}

// The HTML Definition Element (<dfn>) represents the defining instance of a
// term.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dfn", markup...)
}

// Dialog represents a dialog box or other interactive component, such as an
// inspector or window. <form> elements can be integrated within a dialog by
// specifying them with the attribute method="dialog". When such a form is
// submitted, the dialog is closed with a returnValue attribute set to the
// value of the submit button used.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dialog", markup...)
}

// Div (or HTML Document Division Element) is the generic container for flow
// content, which does not inherently represent anything. It can be used to
// group elements for styling purposes (using the class or id attributes), or
// because they share attribute values, such as lang. It should be used only
// when no other semantic element (such as <article> or <nav>) is appropriate.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("div", markup...)
}

// DescriptionList (or HTML Description List Element) encloses a list of pairs
// of terms and descriptions. Common uses for this element are to implement a
// glossary or to display metadata (a list of key-value pairs).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dl", markup...)
}

// DefinitionTerm (or HTML Definition Term Element) identifies a term in a
// definition list. This element can occur only as a child element of a <dl>.
// It is usually followed by a <dd> element; however, multiple <dt> elements in
// a row indicate several terms that are all defined by the immediate next <dd>
// element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("dt", markup...)
}

// Element is used to define new custom DOM elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/element
func Element(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("element", markup...)
}

// The HTML element emphasis  <em> marks text that has stress emphasis. The
// <em> element can be nested, with each level of nesting indicating a greater
// degree of emphasis.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("em", markup...)
}

// Embed represents an integration point for an external application or
// interactive content (in other words, a plug-in).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("embed", markup...)
}

// FieldSet is used to group several controls as well as labels (<label>)
// within a web form.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("fieldset", markup...)
}

// FigureCaption represents a caption or a legend associated with a figure or
// an illustration described by the rest of the data of the <figure> element
// which is its immediate ancestor which means <figcaption> can be the first or
// last element inside a <figure> block. Also, the HTML Figcaption Element is
// optional; if not provided, then the parent figure element will have no
// caption.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("figcaption", markup...)
}

// Figure represents self-contained content, frequently with a caption
// (<figcaption>), and is typically referenced as a single unit. While it is
// related to the main flow, its position is independent of the main flow.
// Usually this is an image, an illustration, a diagram, a code snippet, or a
// schema that is referenced in the main text, but that can be moved to another
// page or to an appendix without affecting the main flow.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("figure", markup...)
}

// Footer represents a footer for its nearest sectioning content or sectioning
// root element. A footer typically contains information about the author of
// the section, copyright data or links to related documents.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("footer", markup...)
}

// Form represents a document section that contains interactive controls to
// submit information to a web server.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("form", markup...)
}

// Header represents a group of introductory or navigational aids. It may
// contain some heading elements but also other elements like a logo, wrapped
// section's header, a search form, and so on.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("header", markup...)
}

// HeadingsGroup (HTML Headings Group Element) represents the heading of a
// section. It defines a single title that participates in the outline of the
// document as the heading of the implicit or explicit section that it belongs
// to.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("hgroup", markup...)
}

// HorizontalRule represents a thematic break between paragraph-level elements
// (for example, a change of scene in a story, or a shift of topic with a
// section). In previous versions of HTML, it represented a horizontal rule. It
// may still be displayed as a horizontal rule in visual browsers, but is now
// defined in semantic terms, rather than presentational terms.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("hr", markup...)
}

// Italic represents a range of text that is set off from the normal text for
// some reason, for example, technical terms, foreign language phrases, or
// fictional character thoughts. It is typically displayed in italic type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("i", markup...)
}

// The HTML Inline Frame Element (<iframe>) represents a nested browsing
// context, effectively embedding another HTML page into the current page. In
// HTML 4.01, a document may contain a head and a body or a head and a
// frameset, but not both a body and a frameset. However, an <iframe> can be
// used within a normal document body. Each browsing context has its own
// session history and active document. The browsing context that contains the
// embedded content is called the parent browsing context. The top-level
// browsing context (which has no parent) is typically the browser window.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("iframe", markup...)
}

// Image represents an image in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("img", markup...)
}

// The HTML element <input> is used to create interactive controls for
// web-based forms in order to accept data from the user. How an <input> works
// varies considerably depending on the value of its type attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("input", markup...)
}

// InsertedText (or HTML Inserted Text) HTML represents a range of text that
// has been added to a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ins", markup...)
}

// The HTML Keyboard Input Element (<kbd>) represents user input and produces
// an inline element displayed in the browser's default monospace font.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("kbd", markup...)
}

// The HTML Label Element (<label>) represents a caption for an item in a user
// interface. It can be associated with a control either by placing the control
// element inside the <label> element, or by using the for attribute. Such a
// control is called the labeled control of the label element. One input can be
// associated with multiple labels.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("label", markup...)
}

// Legend (or HTML Legend Field Element) represents a caption for the content
// of its parent <fieldset>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("legend", markup...)
}

// ListItem (or HTML List Item Element) is used to represent an item in a list.
// It must be contained in a parent element: an ordered list (<ol>), an
// unordered list (<ul>), or a menu (<menu>). In menus and unordered lists,
// list items are usually displayed using bullet points. In ordered lists, they
// are usually displayed with an ascending counter on the left, such as a
// number or letter.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("li", markup...)
}

// Link specifies relationships between the current document and an external
// resource. Possible uses for this element include defining a relational
// framework for navigation. This Element is most used to link to style sheets.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("link", markup...)
}

// Main represents the main content of  the <body> of a document or
// application. The main content area consists of content that is directly
// related to, or expands upon the central topic of a document or the central
// functionality of an application. This content should be unique to the
// document, excluding any content that is repeated across a set of documents
// such as sidebars, navigation links, copyright information, site logos, and
// search forms (unless the document's main function is as a search form).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("main", markup...)
}

// Map is used with <area> elements to define an image map (a clickable link
// area).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("map", markup...)
}

// The HTML Mark Element (<mark>) represents highlighted text, i.e., a run of
// text marked for reference purpose, due to its relevance in a particular
// context. For example it can be used in a page showing search results to
// highlight every instance of the searched-for word.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("mark", markup...)
}

// Menu represents a group of commands that a user can perform or activate.
// This includes both list menus, which might appear across the top of a
// screen, as well as context menus, such as those that might appear underneath
// a button after it has been clicked.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("menu", markup...)
}

// MenuItem represents a command that a user is able to invoke through a popup
// menu. This includes context menus, as well as menus that might be attached
// to a menu button.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem
func MenuItem(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("menuitem", markup...)
}

// Meta represents any metadata information that cannot be represented by one
// of the other HTML meta-related elements (<base>, <link>, <script>, <style>
// or <title>).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("meta", markup...)
}

// Meter represents either a scalar value within a known range or a fractional
// value.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("meter", markup...)
}

// Navigation (HTML Navigation Element) represents a section of a page that
// links to other pages or to parts within the page: a section with navigation
// links.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("nav", markup...)
}

// <noframes> is an HTML element which is used to supporting browsers which are
// not able to support <frame> elements or configured to do so.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes
func NoFrames(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("noframes", markup...)
}

// NoScript defines a section of html to be inserted if a script type on the
// page is unsupported or if scripting is currently turned off in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("noscript", markup...)
}

// The HTML Embedded Object Element (<object>) represents an external resource,
// which can be treated as an image, a nested browsing context, or a resource
// to be handled by a plugin.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("object", markup...)
}

// OrderedList (or HTML Ordered List Element) represents an ordered list of
// items. Typically, ordered-list items are displayed with a preceding
// numbering, which can be of any form, like numerals, letters or Romans
// numerals or even simple bullets. This numbered style is not defined in the
// HTML description of the page, but in its associated CSS, using the
// list-style-type property.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ol", markup...)
}

// In a Web form, the HTML <optgroup> element  creates a grouping of options
// within a <select> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("optgroup", markup...)
}

// In a Web form, the HTML <option> element is used to create a control
// representing an item within a <select>, an <optgroup> or a <datalist> HTML5
// element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("option", markup...)
}

// Output represents the result of a calculation or user action.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("output", markup...)
}

// Paragraph (or HTML Paragraph Element) represents a paragraph of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("p", markup...)
}

// Parameter (or HTML Parameter Element) defines parameters for <object>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("param", markup...)
}

// Picture is a container used to specify multiple <source> elements for a
// specific <img> contained in it. The browser will choose the most suitable
// source according to the current layout of the page (the constraints of the
// box the image will appear in) and the device it will be displayed on (e.g. a
// normal or hiDPI device.)
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("picture", markup...)
}

// Preformatted (or HTML Preformatted Text) represents preformatted text. Text
// within this element is typically displayed in a non-proportional
// ("monospace") font exactly as it is laid out in the file. Whitespace inside
// this element is displayed as typed.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("pre", markup...)
}

// Progress is used to view the completion progress of a task. While the
// specifics of how it's displayed is left up to the browser developer, it's
// typically displayed as a progress bar. Javascript can be used to manipulate
// the value of progress bar.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("progress", markup...)
}

// The HTML Quote Element (<q>) indicates that the enclosed text is a short
// inline quotation. This element is intended for short quotations that don't
// require paragraph breaks; for long quotations use <blockquote> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("q", markup...)
}

// RubyParenthesis is used to provide fall-back parenthesis for browsers
// non-supporting ruby annotations. Ruby annotations are for showing
// pronunciation of East Asian characters, like using Japanese furigana or
// Taiwainese bopomofo characters. The <rp> element is used in the case of lack
// of <ruby> element support its content has what should be displayed in order
// to indicate the presence of a ruby annotation, usually parentheses.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("rp", markup...)
}

// RubyText embraces pronunciation of characters presented in a ruby
// annotations, which are used to describe the pronunciation of East Asian
// characters. This element is always used inside a <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("rt", markup...)
}

// RubyTextContainer embraces semantic annotations of characters presented in a
// ruby of <rb> elements used inside of <ruby> element. <rb> elements can have
// both pronunciation (<rt>) and semantic (<rtc>) annotations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func RubyTextContainer(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("rtc", markup...)
}

// Ruby represents a ruby annotation. Ruby annotations are for showing
// pronunciation of East Asian characters.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ruby", markup...)
}

// The HTML Strikethrough Element (<s>) renders text with a strikethrough, or a
// line through it. Use the <s> element to represent things that are no longer
// relevant or no longer accurate. However, <s> is not appropriate when
// indicating document edits; for that, use the <del> and <ins> elements, as
// appropriate.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("s", markup...)
}

// Sample is an element intended to identify sample output from a computer
// program. It is usually displayed in the browser's default monotype font
// (such as Lucida Console).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("samp", markup...)
}

// The HTML Script Element (<script>) is used to embed or reference an
// executable script within an HTML or XHTML document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("script", markup...)
}

// Section represents a generic section of a document, i.e., a thematic
// grouping of content, typically with a heading. Each <section> should be
// identified, typically by including a heading (<h1>-<h6> element) as a child
// of the <section> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("section", markup...)
}

// The HTML select (<select>) element represents a control that presents a menu
// of options. The options within the menu are represented by <option>
// elements, which can be grouped by <optgroup> elements. Options can be
// pre-selected for the user.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("select", markup...)
}

// Shadow is used as a shadow DOM insertion point. You might use it if you have
// created multiple shadow roots under a shadow host. It is not useful in
// ordinary HTML. It is used with Web Components.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Shadow
func Shadow(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("shadow", markup...)
}

// The HTML Small Element (<small>) makes the text font size one size smaller
// (for example, from large to medium, or from small to x-small) down to the
// browser's minimum font size.  In HTML5, this element is repurposed to
// represent side-comments and small print, including copyright and legal text,
// independent of its styled presentation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("small", markup...)
}

// Source specifies multiple media resources for either the <picture>, the
// <audio> or the <video> element. It is an empty element. It is commonly used
// to serve the same media content in multiple formats supported by different
// browsers.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("source", markup...)
}

// Span is a generic inline container for phrasing content, which does not
// inherently represent anything. It can be used to group elements for styling
// purposes (using the class or id attributes), or because they share attribute
// values, such as lang.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("span", markup...)
}

// The HTML Strong Element (<strong>) gives text strong importance, and is
// typically displayed in bold.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("strong", markup...)
}

// Style contains style information for a document, or part of a document. By
// default, the style instructions written inside that element are expected to
// be CSS.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("style", markup...)
}

// The HTML Subscript Element (<sub>) defines a span of text that should be
// displayed, for typographic reasons, lower, and often smaller, than the main
// span of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("sub", markup...)
}

// Summary is used as a summary, caption, or legend for the content of a
// <details> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("summary", markup...)
}

// The HTML Superscript Element (<sup>) defines a span of text that should be
// displayed, for typographic reasons, higher, and often smaller, than the main
// span of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("sup", markup...)
}

// The HTML Table Element (<table>) represents tabular data - i.e., information
// expressed via a two dimensional data table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("table", markup...)
}

// The HTML Table Body Element (<tbody>) defines one or more <tr> element
// data-rows to be the body of its parent <table> element (as long as no <tr>
// elements are immediate children of that table element.)  In conjunction with
// a preceding <thead> and/or <tfoot> element, <tbody> provides additional
// semantic information for devices such as printers and displays. Of the
// parent table's child elements, <tbody> represents the content which, when
// longer than a page, will most likely differ for each page printed; while the
// content of <thead> and <tfoot> will be the same or similar for each page
// printed. For displays, <tbody> will enable separate scrolling of the
// <thead>, <tfoot>, and <caption> elements of the same parent <table> element.
//  Note that unlike the <thead>, <tfoot>, and <caption> elements however,
// multiple <tbody> elements are permitted (if consecutive), allowing the
// data-rows in long tables to be divided into different sections, each
// separately formatted as needed.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("tbody", markup...)
}

// The Table cell HTML element (<td>) defines a cell of a table that contains
// data. It participates in the table model.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("td", markup...)
}

// The HTML template element <template> is a mechanism for holding client-side
// content that is not to be rendered when a page is loaded but may
// subsequently be instantiated during runtime using JavaScript. 
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("template", markup...)
}

// TextArea represents a multi-line plain-text editing control.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("textarea", markup...)
}

// The HTML Table Foot Element (<tfoot>) defines a set of rows summarizing the
// columns of the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("tfoot", markup...)
}

// The HTML element table header cell <th> defines a cell as header of a group
// of table cells. The exact nature of this group is defined by the scope and
// headers attributes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("th", markup...)
}

// The HTML Table Head Element (<thead>) defines a set of rows defining the
// head of the columns of the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("thead", markup...)
}

// Time represents either a time on a 24-hour clock or a precise date in the
// Gregorian calendar (with optional time and timezone information).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("time", markup...)
}

// Title defines the title of the document, shown in a browser's title bar or
// on the page's tab. It can only contain text, and any contained tags are
// ignored.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("title", markup...)
}

// The HTML element table row <tr> defines a row of cells in a table. Those can
// be a mix of <td> and <th> elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("tr", markup...)
}

// Track is used as a child of the media elements—<audio> and <video>. It
// lets you specify timed text tracks (or time-based data), for example to
// automatically handle subtitles. The tracks are formatted in WebVTT format
// (.vtt files) — Web Video Text Tracks.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("track", markup...)
}

// The HTML Underline Element (<u>) renders text with an underline, a line
// under the baseline of its content.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("u", markup...)
}

// UnorderedList (or HTML Unordered List Element) represents an unordered list
// of items, namely a collection of items that do not have a numerical
// ordering, and their order in the list is meaningless. Typically,
// unordered-list items are displayed with a bullet, which can be of several
// forms, like a dot, a circle or a squared. The bullet style is not defined in
// the HTML description of the page, but in its associated CSS, using the
// list-style-type property.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("ul", markup...)
}

// The HTML Variable Element (<var>) represents a variable in a mathematical
// expression or a programming context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("var", markup...)
}

// Use the  HTML <video> element to embed video content in a document. The
// video element contains one or more video sources. To specify a video source,
// use either the src attribute or the <source> element; the browser will
// choose the most suitable one.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("video", markup...)
}

// The HTML element word break opportunity <wbr> represents a position within
// text where the browser may optionally break a line, though its line-breaking
// rules would not otherwise create a break at that location.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("wbr", markup...)
}

// Heading elements implement six levels of document headings, <h1> is the most
// important and <h6> is the least. A heading element briefly describes the
// topic of the section it introduces. Heading information may be used by user
// agents, for example, to construct a table of contents for a document
// automatically.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header1(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h1", markup...)
}

// Heading elements implement six levels of document headings, <h1> is the most
// important and <h6> is the least. A heading element briefly describes the
// topic of the section it introduces. Heading information may be used by user
// agents, for example, to construct a table of contents for a document
// automatically.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header2(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h2", markup...)
}

// Heading elements implement six levels of document headings, <h1> is the most
// important and <h6> is the least. A heading element briefly describes the
// topic of the section it introduces. Heading information may be used by user
// agents, for example, to construct a table of contents for a document
// automatically.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header3(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h3", markup...)
}

// Heading elements implement six levels of document headings, <h1> is the most
// important and <h6> is the least. A heading element briefly describes the
// topic of the section it introduces. Heading information may be used by user
// agents, for example, to construct a table of contents for a document
// automatically.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header4(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h4", markup...)
}

// Heading elements implement six levels of document headings, <h1> is the most
// important and <h6> is the least. A heading element briefly describes the
// topic of the section it introduces. Heading information may be used by user
// agents, for example, to construct a table of contents for a document
// automatically.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header5(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h5", markup...)
}

// Heading elements implement six levels of document headings, <h1> is the most
// important and <h6> is the least. A heading element briefly describes the
// topic of the section it introduces. Heading information may be used by user
// agents, for example, to construct a table of contents for a document
// automatically.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header6(markup ...vecty.MarkupOrComponentOrHTML) *vecty.HTML {
	return vecty.Tag("h6", markup...)
}
