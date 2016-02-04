//go:generate go run generate.go

// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.
package elem

import (
	"github.com/neelance/dom"
)

// The HTML Anchor Element (<a>) defines a hyperlink to a location on the same page or any other page on the Web. It can also be used (in an obsolete way) to create an anchor point—a destination for hyperlinks within the content of a page, so that links aren't limited to connecting simply to the top of a page.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "a"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Abbreviation element (<abbr>) represents an abbreviation and optionally provides a full description for it. If present, the title attribute must contain this full description and nothing else.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "abbr"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Address Element (<address>) should be used by authors to supply contact information for its nearest <article> or <body> ancestor; in the latter case, it applies to the whole document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "address"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <area> element defines a hot-spot region on an image, and optionally associates it with a hypertext link. This element is used only within a <map> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "area"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Article Element (<article>) represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable (e.g., in syndication). This could be a forum post, a magazine or newspaper article, a blog entry, an object, or any other independent item of content. Each <article> should be identified, typically by including a heading (h1-h6 element) as a child of the <article> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "article"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <aside> element represents a section of the page with content connected tangentially to the rest, which could be considered separate from that content. These sections are often represented as sidebars or inserts. They often contain the definitions on the sidebars, such as definitions from the glossary; there may also be other types of information, such as related advertisements; the biography of the author; web applications; profile information or related links on the blog.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "aside"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <audio> element is used to embed sound content in documents. It may contain several audio sources, represented using the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "audio"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <b> Element represents a span of text stylistically different from normal text, without conveying any special importance or relevance. It is typically used for keywords in a summary, product names in a review, or other spans of text whose typical presentation would be boldfaced. Another example of its use is to mark the lead sentence of each paragraph of an article.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "b"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Base Element (<base>) specifies the base URL to use for all relative URLs contained within a document. There can be only one <base> element in a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "base"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <bdi> Element (or Bi-Directional Isolation Element) isolates a span of text that might be formatted in a different direction from other text outside it.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "bdi"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <bdo> Element (or HTML bidirectional override element) is used to override the current directionality of text. It causes the directionality of the characters to be ignored in favor of the specified directionality.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "bdo"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <blockquote> Element (or HTML Block Quotation Element) indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see Notes for how to change it). A URL for the source of the quotation may be given using the cite attribute, while a text representation of the source can be given using the <cite> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "blockquote"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <br> Element (or HTML Line Break Element) produces a line break in text (carriage-return). It is useful for writing a poem or an address, where the division of lines is significant.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "br"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <button> Element represents a clickable button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "button"}
	dom.List(markup).Apply(e)
	return e
}

// Editorial review completed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "canvas"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <caption> Element (or HTML Table Caption Element) represents the title of a table. Though it is always the first descendant of a <table>, its styling, using CSS, may place it elsewhere, relative to the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "caption"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Citation Element (<cite>) represents a reference to a creative work. It must include the title of a work or a URL reference, which may be in an abbreviated form according to the conventions used for the addition of citation metadata.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "cite"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Code Element (<code>) represents a fragment of computer code. By default, it is displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "code"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Column Element (<col>) defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "col"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Column Group Element (<colgroup>) defines a group of columns within a table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "colgroup"}
	dom.List(markup).Apply(e)
	return e
}

// The command element represents a command which the user can invoke.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/command
func Command(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "command"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <content> element is used inside of Shadow DOM as an insertion point. It is not intended to be used in ordinary HTML. It is used with Web Components.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/content
func Content(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "content"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <data> Element links a given content with a machine-readable translation. If the content is time- or date-related, the <time> must be used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "data"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Datalist Element (<datalist>) contains a set of <option> elements that represent the values available for other controls.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "datalist"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Description Element (<dd>) indicates the description of a term in a description list (<dl>) element. This element can occur only as a child element of a definition list and it must follow a <dt> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "dd"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Deleted Text Element (<del>) represents a range of text that has been deleted from a document. This element is often (but need not be) rendered with strike-through text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "del"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Details Element (<details>) is used as a disclosure widget from which the user can retrieve additional information.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "details"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Definition Element (<dfn>) represents the defining instance of a term.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "dfn"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <dialog> element represents a dialog box or other interactive component, such as an inspector or window. <form> elements can be integrated within a dialog by specifying them with the attribute method="dialog". When such a form is submitted, the dialog is closed with a returnValue attribute set to the value of the submit button used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "dialog"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <div> element (or HTML Document Division Element) is the generic container for flow content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element (such as <article> or <nav>) is appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "div"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <dl> Element (or HTML Description List Element) encloses a list of pairs of terms and descriptions. Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "dl"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <dt> element (or HTML Definition Term Element) identifies a term in a definition list. This element can occur only as a child element of a <dl>. It is usually followed by a <dd> element; however, multiple <dt> elements in a row indicate several terms that are all defined by the immediate next <dd> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "dt"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <element> element is used to define new custom DOM elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/element
func Element(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "element"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Emphasis Element (<em>) marks text that has stress emphasis. The <em> element can be nested, with each level of nesting indicating a greater degree of emphasis.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "em"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <embed> Element represents an integration point for an external application or interactive content (in other words, a plug-in).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "embed"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <fieldset> element is used to group several controls as well as labels (<label>) within a web form.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "fieldset"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <figcaption> Element represents a caption or a legend associated with a figure or an illustration described by the rest of the data of the <figure> element which is its immediate ancestor which means <figcaption> can be the first or last element inside a <figure> block. Also, the HTML Figcaption Element is optional; if not provided, then the parent figure element will have no caption.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "figcaption"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <figure> Element represents self-contained content, frequently with a caption (<figcaption>), and is typically referenced as a single unit. While it is related to the main flow, its position is independent of the main flow. Usually this is an image, an illustration, a diagram, a code snippet, or a schema that is referenced in the main text, but that can be moved to another page or to an appendix without affecting the main flow.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "figure"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Font Element (<font>) defines the font size, color and face for its content.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/font
func Font(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "font"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Footer Element (<footer>) represents a footer for its nearest sectioning content or sectioning root element. A footer typically contains information about the author of the section, copyright data or links to related documents.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "footer"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <form> element represents a document section that contains interactive controls to submit information to a web server.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "form"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <header> Element represents a group of introductory or navigational aids. It may contain some heading elements but also other elements like a logo, wrapped section's header, a search form, and so on.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "header"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <hgroup> Element (HTML Headings Group Element) represents the heading of a section. It defines a single title that participates in the outline of the document as the heading of the implicit or explicit section that it belongs to.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "hgroup"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <hr> element represents a thematic break between paragraph-level elements (for example, a change of scene in a story, or a shift of topic with a section). In previous versions of HTML, it represented a horizontal rule. It may still be displayed as a horizontal rule in visual browsers, but is now defined in semantic terms, rather than presentational terms.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "hr"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <i> Element represents a range of text that is set off from the normal text for some reason, for example, technical terms, foreign language phrases, or fictional character thoughts. It is typically displayed in italic type.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "i"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Inline Frame Element (<iframe>) represents a nested browsing context, effectively embedding another HTML page into the current page. In HTML 4.01, a document may contain a head and a body or a head and a frame-set, but not both a body and a frame-set. However, an <iframe> can be used within a normal document body. Each browsing context has its own session history and active document. The browsing context that contains the embedded content is called the parent browsing context. The top-level browsing context (which has no parent) is typically the browser window.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "iframe"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Image Element (<img>) represents an image in the document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "img"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <input> element is used to create interactive controls for web-based forms in order to accept data from the user. How an <input> works varies considerably depending on the value of its type attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "input"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <ins> Element (or HTML Inserted Text) HTML represents a range of text that has been added to a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "ins"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Keyboard Input Element (<kbd>) represents user input and produces an inline element displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "kbd"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <keygen> element exists to facilitate generation of key material, and submission of the public key as part of an HTML form. This mechanism is designed for use with Web-based certificate management systems. It is expected that the <keygen> element will be used in an HTML form along with other information needed to construct a certificate request, and that the result of the process will be a signed certificate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/keygen
func KeyGen(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "keygen"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Label Element (<label>) represents a caption for an item in a user interface. It can be associated with a control either by placing the control element inside the <label> element, or by using the for attribute. Such a control is called the labeled control of the label element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "label"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <legend> Element (or HTML Legend Field Element) represents a caption for the content of its parent <fieldset>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "legend"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML List Item Element (<li>) is used to represent an item in a list. It must be contained in a parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>). In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are usually displayed with an ascending counter on the left, such as a number or letter.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "li"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Link Element (<link>) specifies relationships between the current document and an external resource. Possible uses for this element include defining a relational framework for navigation. This Element is most used to link to style sheets.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "link"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <main> element represents the main content of  the <body> of a document or application. The main content area consists of content that is directly related to, or expands upon the central topic of a document or the central functionality of an application. This content should be unique to the document, excluding any content that is repeated across a set of documents such as sidebars, navigation links, copyright information, site logos, and search forms (unless the document's main function is as a search form).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "main"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <map> element is used with <area> elements to define an image map (a clickable link area).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "map"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Mark Element (<mark>) represents highlighted text, i.e., a run of text marked for reference purpose, due to its relevance in a particular context. For example it can be used in a page showing search results to highlight every instance of the searched-for word.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "mark"}
	dom.List(markup).Apply(e)
	return e
}

// Technical review completed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "menu"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <menuitem> element represents a command that a user is able to invoke through a popup menu. This includes context menus, as well as menus that might be attached to a menu button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem
func MenuItem(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "menuitem"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Meta Element (<meta>) represents any metadata information that cannot be represented by one of the other HTML meta-related elements (<base>, <link>, <script>, <style> or <title>).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "meta"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <meter> Element represents either a scalar value within a known range or a fractional value.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "meter"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Navigation Element (<nav>) represents a section of a page that links to other pages or to parts within the page: a section with navigation links.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "nav"}
	dom.List(markup).Apply(e)
	return e
}

// <noframes> is an HTML element which is used to supporting browsers which are not able to support <frame> elements or configured to do so.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes
func NoFrames(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "noframes"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <noscript> Element defines a section of html to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "noscript"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Embedded Object Element (<object>) represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "object"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <ol> Element (or HTML Ordered List Element) represents an ordered list of items. Typically, ordered-list items are displayed with a preceding numbering, which can be of any form, like numerals, letters or Romans numerals or even simple bullets. This numbered style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "ol"}
	dom.List(markup).Apply(e)
	return e
}

// In a Web form, the HTML <optgroup> element  creates a grouping of options within a <select> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "optgroup"}
	dom.List(markup).Apply(e)
	return e
}

// In a Web form, the HTML <option> element is used to create a control representing an item within a <select>, an <optgroup> or a <datalist> HTML5 element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "option"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <output> element represents the result of a calculation or user action.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "output"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <p> element (or HTML Paragraph Element) represents a paragraph of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "p"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <param> Element (or HTML Parameter Element) defines parameters for <object>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "param"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <picture> element is a container used to specify multiple <source>s for a specific <img> contained in it. The browser will choose the most suitable source according to the current layout of the page (the constraints of the box the image will appear in) and the device it will be displayed on (e.g. a normal or hiDPI device).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "picture"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Preformatted Text (<pre>) represents preformatted text. Text within this element is typically displayed in a non-proportional font exactly as it is laid out in the file. Whitespaces inside this element are displayed as typed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "pre"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <progress> Element is used to view the completion progress of a task. While the specifics of how it's displayed is left up to the browser developer, it's typically displayed as a progress bar. Javascript can be used to manipulate the value of progress bar.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "progress"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Quote Element (<q>) indicates that the enclosed text is a short inline quotation. This element is intended for short quotations that don't require paragraph breaks; for long quotations use <blockquote> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "q"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <rp> element is used to provide fall-back parenthesis for browsers non-supporting ruby annotations. Ruby annotations are for showing pronounciation of East Asian characters, like using Japanese furigana or Taiwainese bopomofo characters. The <rp> element is used in the case of lack of <ruby> element support its content has what should be displayed in order to indicate the presence of a ruby annotation, usually parentheses.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "rp"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <rt> Element embraces pronunciation of characters presented in a ruby annotations, which are used to describe the pronunciation of East Asian characters. This element is always used inside a <ruby> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "rt"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <rtc> Element embraces semantic annotations of characters presented in a ruby of <rb> elements used inside of <ruby> element. <rb> elements can have both pronunciation (<rt> and semantic (<rtc>) annotations.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func Rtc(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "rtc"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <ruby> Element represents a ruby annotation. Ruby annotations are for showing pronunciation of East Asian characters.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "ruby"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Strikethrough Element (<s>) renders text with a strikethrough, or a line through it. Use the <s> element to represent things that are no longer relevant or no longer accurate. However, <s> is not appropriate when indicating document edits; for that, use the <del> and <ins> elements, as appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "s"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <samp> element is an element intended to identify sample output from a computer program. It is usually displayed in the browser's default monotype font (such as Lucida Console).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "samp"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Script Element (<script>) is used to embed or reference an executable script within an HTML or XHTML document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "script"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Section Element (<section>) represents a generic section of a document, i.e., a thematic grouping of content, typically with a heading. Each <section> should be identified, typically by including a heading (<h1>-<h6> element) as a child of the <section> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "section"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML select (<select>) element represents a control that presents a menu of options. The options within the menu are represented by <option> elements, which can be grouped by <optgroup> elements. Options can be pre-selected for the user.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "select"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <shadow> element is used as a shadow DOM insertion point. You might use it if you have created multiple shadow roots under a shadow host. It is not useful in ordinary HTML. It is used with Web Components.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Shadow
func Shadow(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "shadow"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Small Element (<small>) makes the text font size one size smaller (for example, from large to medium, or from small to x-small) down to the browser's minimum font size.  In HTML5, this element is repurposed to represent side-comments and small print, including copyright and legal text, independent of its styled presentation.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "small"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <source> element is used to specify multiple media resources for <picture>, <audio> and <video> elements. It is an empty element. It is commonly used to serve the same media in multiple formats supported by different browsers.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "source"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <span> element is a generic inline container for phrasing content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element is appropriate. <span> is very much like a <div> element, but <div> is a block-level element whereas a <span> is an inline element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "span"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Strong Element (<strong>) gives text strong importance, and is typically displayed in bold.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "strong"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Style Element (<style>) contains style information for a document, or part of a document. By default, the style instructions written inside that element are expected to be CSS.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "style"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Subscript Element (<sub>) defines a span of text that should be displayed, for typographic reasons, lower, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "sub"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML summary element (<summary>) is used as a summary, caption, or legend for the content of a <details> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "summary"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Superscript Element (<sup>) defines a span of text that should be displayed, for typographic reasons, higher, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "sup"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Element (<table>) represents data in two dimensions or more.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "table"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Body Element () defines one or more <tr> element data-rows to be the body of its parent <table> element (as long as no <tr> elements are immediate children of that table element.)  In conjunction with a preceding <thead> and/or <tfoot> element, <tbody> provides additional semantic information for devices such as printers and displays. Of the parent table's child elements, <tbody> represents the content which, when longer than a page, will most likely differ for each page printed; while the content of <thead> and <tfoot> will be the same or similar for each page printed. For displays, <tbody> will enable separate scrolling of the <thead>, <tfoot>, and <caption> elements of the same parent <table> element.  Note that unlike the <thead>, <tfoot>, and <caption> elements however, multiple <tbody> elements are permitted (if consecutive), allowing the data-rows in long tables to be divided into different sections, each separately formatted as needed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "tbody"}
	dom.List(markup).Apply(e)
	return e
}

// The Table cell HTML element (<td>) defines a cell of a table that contains data. It participates in the table model.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "td"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML template element <template> is a mechanism for holding client-side content that is not to be rendered when a page is loaded but may subsequently be instantiated during runtime using JavaScript.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "template"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <textarea> element represents a multi-line plain-text editing control.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "textarea"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Foot Element (<tfoot>) defines a set of rows summarizing the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "tfoot"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Header Cell Element (<th>) defines a cell that is a header for a group of cells of a table. The group of cells that the header refers to is defined by the scope and headers attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "th"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Head Element (<thead>) defines a set of rows defining the head of the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "thead"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <time> element represents either a time on a 24-hour clock or a precise date in the Gregorian calendar (with optional time and timezone information).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "time"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Title Element (<title>) defines the title of the document, shown in a browser's title bar or on the page's tab. It can only contain text and any contained tags are not interpreted.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "title"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Table Row Element (<tr>) defines a row of cells in a table. Those can be a mix of <td> and <th> elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "tr"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <track> element is used as a child of the media elements—<audio> and <video>. It lets you specify timed text tracks (or time-based data), for example to automatically handle subtitles. The tracks are formatted in WebVTT format (.vtt files) — Web Video Text Tracks.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "track"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Underline Element (<u>) renders text with an underline, a line under the baseline of its content.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "u"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML unordered list element (<ul>) represents an unordered list of items, namely a collection of items that do not have a numerical ordering, and their order in the list is meaningless. Typically, unordered-list items are displayed with a bullet, which can be of several forms, like a dot, a circle or a squared. The bullet style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "ul"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML Variable Element (<var>) represents a variable in a mathematical expression or a programming context.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "var"}
	dom.List(markup).Apply(e)
	return e
}

// The HTML <video> element is used to embed video content. It may contain several video sources, represented using the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "video"}
	dom.List(markup).Apply(e)
	return e
}

// The Word Break Opportunity (<wbr>) HTML element represents a position within text where the browser may optionally break a line, though its line-breaking rules would not otherwise create a break at that location.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "wbr"}
	dom.List(markup).Apply(e)
	return e
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header1(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "h1"}
	dom.List(markup).Apply(e)
	return e
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header2(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "h2"}
	dom.List(markup).Apply(e)
	return e
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header3(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "h3"}
	dom.List(markup).Apply(e)
	return e
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header4(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "h4"}
	dom.List(markup).Apply(e)
	return e
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header5(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "h5"}
	dom.List(markup).Apply(e)
	return e
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header6(markup ...dom.Markup) *dom.Element {
	e := &dom.Element{TagName: "h6"}
	dom.List(markup).Apply(e)
	return e
}
