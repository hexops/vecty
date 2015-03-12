// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.
package elem

import (
	"github.com/neelance/dom"
)

// The HTML <a> Element (or the HTML Anchor Element) defines a hyperlink, the named target destination for a hyperlink, or both.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("a", aspects...)
}

// The HTML <abbr> Element (or HTML Abbreviation Element) represents an abbreviation and optionally provides a full description for it. If present, the title attribute must contain this full description and nothing else.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("abbr", aspects...)
}

// The HTML <address> Element may be used by authors to supply contact information for its nearest <article> or <body> ancestor; in the latter case, it applies to the whole document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("address", aspects...)
}

// The HTML <area> element defines a hot-spot region on an image, and optionally associates it with a hypertext link. This element is used only within a <map> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("area", aspects...)
}

// The HTML <article> Element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable, e.g., in syndication. This could be a forum post, a magazine or newspaper article, a blog entry, a user-submitted comment, an interactive widget or gadget, or any other independent item of content. Each <article> should be identified, typically by including a heading (h1-h6 element) as a child of the <article> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("article", aspects...)
}

// The HTML <aside> element represents a section of the page with content connected tangentially to the rest, which could be considered separate from that content. These sections are often represented as sidebars or inserts. They often contain the definitions on the sidebars, such as definitions from the glossary; there may also be other types of information, such as related advertisements; the biography of the author; web applications; profile information or related links on the blog.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("aside", aspects...)
}

// The HTML <audio> element is used to embed sound content in documents. It may contain several audio sources, represented using the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("audio", aspects...)
}

// The HTML <b> Element represents a span of text stylistically different from normal text, without conveying any special importance or relevance. It is typically used for keywords in a summary, product names in a review, or other spans of text whose typical presentation would be boldfaced. Another example of its use is to mark the lead sentence of each paragraph of an article.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("b", aspects...)
}

// The HTML Base Element (<base>) specifies the base URL to use for all relative URLs contained within a document. There can be only one <base> element in a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("base", aspects...)
}

// The HTML <bdi> Element (or Bi-Directional Isolation Element) isolates a span of text that might be formatted in a different direction from other text outside it.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("bdi", aspects...)
}

// The HTML <bdo> Element (or HTML bidirectional override element) is used to override the current directionality of text. It causes the directionality of the characters to be ignored in favor of the specified directionality.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("bdo", aspects...)
}

// The HTML <blockquote> Element (or HTML Block Quotation Element) indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see Notes for how to change it). A URL for the source of the quotation may be given using the cite attribute, while a text representation of the source can be given using the <cite> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("blockquote", aspects...)
}

// The HTML <br> Element (or HTML Line Break Element) produces a line break in text (carriage-return). It is useful for writing a poem or an address, where the division of lines is significant.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("br", aspects...)
}

// The HTML <button> Element represents a clickable button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("button", aspects...)
}

// The HTML <canvas> Element can be used to draw graphics via scripting (usually JavaScript). For example, it can be used to draw graphs, make photo compositions or even perform animations. You may (and should) provide alternate content inside the <canvas> block. That content will be rendered both on older browsers that don't support canvas and in browsers with JavaScript disabled.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("canvas", aspects...)
}

// The HTML <caption> Element (or HTML Table Caption Element) represents the title of a table. Though it is always the first descendant of a <table>, its styling, using CSS, may place it elsewhere, relative to the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("caption", aspects...)
}

// The HTML <cite> Element (or HTML Citation Element) represents a reference to a creative work. It must include the title of a work or a URL reference, which may be in an abbreviated form according to the conventions used for the addition of citation metadata.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("cite", aspects...)
}

// The HTML <code> Element represents a fragment of computer code. By default, it is displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("code", aspects...)
}

// The HTML <col> Element (or HTML Table Column Element) defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("col", aspects...)
}

// The HTML <colgroup> Element (or HTML Table Column Group Element) defines a group of columns within a table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("colgroup", aspects...)
}

// The command element represents a command which the user can invoke.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/command
func Command(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("command", aspects...)
}

// The HTML <content> element is used inside of Shadow DOM as an insertion point. It is not intended to be used in ordinary HTML. It is used with Web Components.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/content
func Content(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("content", aspects...)
}

// The HTML <data> Element links a given content with a machine-readable translation. If the content is time- or date-related, the <time> must be used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("data", aspects...)
}

// The HTML <datalist> Element contains a set of <option> elements that represent the values available for other controls.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("datalist", aspects...)
}

// The HTML <dd> Element (or HTML Description Element) indicates the description of a term in a description list (<dl>) element. This element can occur only as a child element of a definition list and it must follow a <dt> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("dd", aspects...)
}

// The HTML <del> element (or HTML Deleted Text Element) represents a range of text that has been deleted from a document. This element is often (but need not be) rendered with strike-through text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("del", aspects...)
}

// The HTML <details> Element is used as a disclosure widget from which the user can retrieve additional information.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("details", aspects...)
}

// The HTML <dfn> element (or HTML Definition Element) represents the defining instance of a term.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("dfn", aspects...)
}

// The HTML <dialog> element represents a dialog box or other interactive component, such as an inspector or window. <form> elements can be integrated within a dialog by specifying them with the attribute method="dialog". When such a form is submitted, the dialog is closed with a returnValue attribute set to the value of the submit button used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("dialog", aspects...)
}

// The HTML <div> element (or HTML Document Division Element) is the generic container for flow content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element (such as <article> or <nav>) is appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("div", aspects...)
}

// The HTML <dl> Element (or HTML Description List Element) encloses a list of pairs of terms and descriptions. Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("dl", aspects...)
}

// The HTML <dt> element (or HTML Definition Term Element) identifies a term in a definition list. This element can occur only as a child element of a <dl>. It is usually followed by a <dd> element; however, multiple <dt> elements in a row indicate several terms that are all defined by the immediate next <dd> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("dt", aspects...)
}

// The HTML <element> element is used to define new custom DOM elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/element
func Element(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("element", aspects...)
}

// The HTML <em> element (or HTML Emphasis Element) marks text that has stress emphasis. The <em> element can be nested, with each level of nesting indicating a greater degree of emphasis.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("em", aspects...)
}

// The HTML <embed> Element represents an integration point for an external application or interactive content (in other words, a plug-in).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("embed", aspects...)
}

// The HTML <fieldset> element is used to group several controls as well as labels (<label>) within a web form.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("fieldset", aspects...)
}

// The HTML <figcaption> Element represents a caption or a legend associated with a figure or an illustration described by the rest of the data of the <figure> element which is its immediate ancestor which means <figcaption> can be the first or last element inside a <figure> block. Also, the HTML Figcaption Element is optional; if not provided, then the parent figure element will have no caption.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("figcaption", aspects...)
}

// The HTML <figure> Element represents self-contained content, frequently with a caption (<figcaption>), and is typically referenced as a single unit. While it is related to the main flow, its position is independent of the main flow. Usually this is an image, an illustration, a diagram, a code snippet, or a schema that is referenced in the main text, but that can be moved to another page or to an appendix without affecting the main flow.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("figure", aspects...)
}

// The HTML Font Element (<font>) defines the font size, color and face for its content.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/font
func Font(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("font", aspects...)
}

// The HTML <footer> Element represents a footer for its nearest sectioning content or sectioning root element (i.e, its nearest parent <article>, <aside>, <nav>, <section>, <blockquote>, <body>, <details>, <fieldset>, <figure>, <td>). A footer typically contains information about the author of the section, copyright data or links to related documents.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("footer", aspects...)
}

// The HTML <form> element represents a document section that contains interactive controls to submit information to a web server.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("form", aspects...)
}

// The HTML <header> Element represents a group of introductory or navigational aids. It may contain some heading elements but also other elements like a logo, wrapped section's header, a search form, and so on.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("header", aspects...)
}

// The HTML <hgroup> Element (HTML Headings Group Element) represents the heading of a section. It defines a single title that participates in the outline of the document as the heading of the implicit or explicit section that it belongs to.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("hgroup", aspects...)
}

// The HTML <hr> element represents a thematic break between paragraph-level elements (for example, a change of scene in a story, or a shift of topic with a section). In previous versions of HTML, it represented a horizontal rule. It may still be displayed as a horizontal rule in visual browsers, but is now defined in semantic terms, rather than presentational terms.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("hr", aspects...)
}

// The HTML <i> Element represents a range of text that is set off from the normal text for some reason, for example, technical terms, foreign language phrases, or fictional character thoughts. It is typically displayed in italic type.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("i", aspects...)
}

// The HTML <iframe> Element (or HTML inline frame element) represents a nested browsing context, effectively embedding another HTML page into the current page. In HTML 4.01, a document may contain a head and a body or a head and a frame-set, but not both a body and a frame-set. However, an <iframe> can be used within a normal document body. Each browsing context has its own session history and active document. The browsing context that contains the embedded content is called the parent browsing context. The top-level browsing context (which has no parent) is typically the browser window.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("iframe", aspects...)
}

// The HTML <img> Element (or HTML Image Element) represents an image of the document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("img", aspects...)
}

// The HTML <input> element is used to create interactive controls for web-based forms in order to accept data from the user. The semantics of an <input> varies considerably depending on the value of its type attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Input
func Input(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("input", aspects...)
}

// The HTML <ins> Element (or HTML Inserted Text) HTML represents a range of text that has been added to a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("ins", aspects...)
}

// The HTML <kbd> Element (or HTML Keyboard Input Element) represents user input and produces an inline element displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("kbd", aspects...)
}

// The HTML <keygen> element exists to facilitate generation of key material, and submission of the public key as part of an HTML form. This mechanism is designed for use with Web-based certificate management systems. It is expected that the <keygen> element will be used in an HTML form along with other information needed to construct a certificate request, and that the result of the process will be a signed certificate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/keygen
func KeyGen(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("keygen", aspects...)
}

// The HTML <label> Element represents a caption for an item in a user interface. It can be associated with a control either by placing the control element inside the label element, or by using the for attribute. Such a control is called the labeled control of the label element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("label", aspects...)
}

// The HTML <legend> Element (or HTML Legend Field Element) represents a caption for the content of its parent <fieldset>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("legend", aspects...)
}

// The HTML List item element (<li>) is used to represent a list item. It should be contained in an ordered list (<ol>), an unordered list (<ul>) or a menu (<menu>), where it represents a single entity in that list. In menus and unordered lists, list items are ordinarily displayed using bullet points. In ordered lists, they are usually displayed with some ascending counter on the left such as a number or letter
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("li", aspects...)
}

// The HTML Link Element (<link>) specifies relationships between the current document and an external resource. Possible uses for this element include defining a relational framework for navigation. This Element is most used to link to style sheets.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("link", aspects...)
}

// The HTML <main> element represents the main content of  the <body> of a document or application. The main content area consists of content that is directly related to, or expands upon the central topic of a document or the central functionality of an application. This content should be unique to the document, excluding any content that is repeated across a set of documents such as sidebars, navigation links, copyright information, site logos, and search forms (unless, of course, the document's main function is as a search form).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("main", aspects...)
}

// The HTML <map> element is used with <area> elements to define an image map (a clickable link area).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("map", aspects...)
}

// The HTML <mark> Element represents highlighted text, i.e., a run of text marked for reference purpose, due to its relevance in a particular context. For example it can be used in a page showing search results to highlight every instance of the searched-for word.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("mark", aspects...)
}

// The HTML <menu> element represents a group of commands that a user can perform or activate. This includes both toolbar menus, which might appear across the top of a screen, as well as popup menus, such as those that might appear underneath a button after it has been clicked.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("menu", aspects...)
}

// The HTML <menuitem> element represents a command that a user is able to invoke through a popup menu. This includes context menus, as well as menus that might be attached to a menu button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem
func MenuItem(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("menuitem", aspects...)
}

// The HTML Meta Element (<meta>) represents any metadata information that cannot be represented by one of the other HTML meta-related elements (<base>, <link>, <script>, <style> or <title>).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("meta", aspects...)
}

// The HTML <meter> Element represents either a scalar value within a known range or a fractional value.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("meter", aspects...)
}

// The HTML Navigation Element (<nav>) represents a section of a page that links to other pages or to parts within the page: a section with navigation links.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("nav", aspects...)
}

// <noframes> is an HTML element which is used to supporting browsers which are not able to support <frame> elements or configured to do so.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes
func NoFrames(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("noframes", aspects...)
}

// The HTML <noscript> Element defines a section of html to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("noscript", aspects...)
}

// The HTML <object> Element (or HTML Embedded Object Element) represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("object", aspects...)
}

// The HTML <ol> Element (or HTML Ordered List Element) represents an ordered list of items. Typically, ordered-list items are displayed with a preceding numbering, which can be of any form, like numerals, letters or Romans numerals or even simple bullets. This numbered style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("ol", aspects...)
}

// In a Web form, the HTML <optgroup> element  creates a grouping of options within a <select> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("optgroup", aspects...)
}

// In a Web form, the HTML <option> element is used to create a control representing an item within a <select>, an <optgroup> or a <datalist> HTML5 element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("option", aspects...)
}

// The HTML <output> element represents the result of a calculation or user action.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("output", aspects...)
}

// The HTML <p> element (or HTML Paragraph Element) represents a paragraph of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("p", aspects...)
}

// The HTML <param> Element (or HTML Parameter Element) defines parameters for <object>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("param", aspects...)
}

// The HTML <picture> element is a container used to specified multiple <source> for a specific <img> contained in it. The browser will choose the most suitable source according to the current layout (constraints of the box the image will appear in) of the page and the device it will be displayed on (like a normal or hiDPI device).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("picture", aspects...)
}

// The HTML <pre> Element (or HTML Preformatted Text) represents preformatted text. Text within this element is typically displayed in a non-proportional font exactly as it is laid out in the file. Whitespaces inside this element are displayed as typed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("pre", aspects...)
}

// The HTML <progress> Element is used to view the completion progress of a task. While the specifics of how it's displayed is left up to the browser developer, it's typically displayed as a progress bar. Javascript can be used to manipulate the value of progress bar.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("progress", aspects...)
}

// The HTML <q> Element (or HTML Quote Element) indicates that the enclosed text is a short inline quotation. This element is intended for short quotations that don't require paragraph breaks; for long quotations use <blockquote> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("q", aspects...)
}

// The HTML <rp> element is used to provide fall-back parenthesis for browsers non-supporting ruby annotations. Ruby annotations are for showing pronounciation of East Asian characters, like using Japanese furigana or Taiwainese bopomofo characters. The <rp> element is used in the case of lack of <ruby> element support its content has what should be displayed in order to indicate the presence of a ruby annotation, usually parentheses.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("rp", aspects...)
}

// The HTML <rt> Element embraces pronunciation of character presented in a ruby annotations, which are for showing pronunciation of East Asian characters and the <rt> element is used inside of <ruby> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("rt", aspects...)
}

// The HTML <rtc> Element embraces semantic annotations of characters presented in a ruby of <rb> elements used inside of <ruby> element. <rb> elements can have both pronunciation (<rt> and semantic (<rtc>) annotations.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func Rtc(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("rtc", aspects...)
}

// The HTML <ruby> Element represents a ruby annotation. Ruby annotations are for showing pronounciation of East Asian characters.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("ruby", aspects...)
}

// The HTML Strikethrough Element (<s>) renders text with a strikethrough, or a line through it. Use the <s> element to represent things that are no longer relevant or no longer accurate. However, <s> is not appropriate when indicating document edits; for that, use the <del> and <ins> elements, as appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("s", aspects...)
}

// The HTML <samp> element is an element intended to identify sample output from a computer program. It is usually displayed in the browser's default monotype font (such as Lucida Console).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("samp", aspects...)
}

// The HTML <script> element is used to embed or reference an executable script within an HTML or XHTML document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("script", aspects...)
}

// Technical review completed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("section", aspects...)
}

// The HTML select (<select>) element represents a control that presents a menu of options. The options within the menu are represented by <option> elements, which can be grouped by <optgroup> elements. Options can be pre-selected for the user.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("select", aspects...)
}

// The HTML <shadow> element is used as a shadow DOM insertion point. You might use it if you have created multiple shadow roots under a shadow host. It is not useful in ordinary HTML. It is used with Web Components.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Shadow
func Shadow(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("shadow", aspects...)
}

// The HTML Small Element (<small>) makes the text font size one size smaller (for example, from large to medium, or from small to x-small) down to the browser's minimum font size.  In HTML5, this element is repurposed to represent side-comments and small print, including copyright and legal text, independent of its styled presentation.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("small", aspects...)
}

// The HTML <source> element is used to specify multiple media resources for <picture>, <audio> and <video> elements. It is an empty element. It is commonly used to serve the same media in multiple formats supported by different browsers.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("source", aspects...)
}

// The HTML <span> element is a generic inline container for phrasing content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element is appropriate. <span> is very much like a <div> element, but <div> is a block-level element whereas a <span> is an inline element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("span", aspects...)
}

// The HTML Strong Element (<strong>) gives text strong importance, and is typically displayed in bold.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("strong", aspects...)
}

// The HTML <style> element contains style information for a document, or a part of document. The specific style information is contained inside of this element, usually in the CSS.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("style", aspects...)
}

// The HTML Subscript Element (<sub>) defines a span of text that should be displayed, for typographic reasons, lower, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("sub", aspects...)
}

// The HTML summary element (<summary>) is used as a summary, caption, or legend for the content of a <details> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("summary", aspects...)
}

// The HTML Superscript Element (<sup>) defines a span of text that should be displayed, for typographic reasons, higher, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("sup", aspects...)
}

// The HTML Table Element (<table>) represents data in two dimensions or more.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("table", aspects...)
}

// The HTML Table Body Element (<tbody>) defines one or more rows as the body of its parent <table> element when no <tr> elements are children of the parent.  In conjunction with a preceding <tfoot> and <thead> element, (<tbody>) provides additional semantic information for devices such as printers and displays. Of the parent table's child elements, (<tbody>) will represent the content, if longer than a page, that will most likely differ for each page printed. The <tfoot> and <thead> elements' content will be consistent for each page printed. For displays, (<tbody>) will enable separate scrolling from the <tfoot>, <thead> and <caption> elements of the same parent <table> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("tbody", aspects...)
}

// The Table cell HTML element (<td>) defines a cell of a table that contains data. It participates in the table model.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("td", aspects...)
}

// The HTML template element <template> is a mechanism for holding client-side content that is not to be rendered when a page is loaded but may subsequently be instantiated during runtime using JavaScript. 
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("template", aspects...)
}

// The HTML <textarea> element represents a multi-line plain-text editing control.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("textarea", aspects...)
}

// The HTML Table Foot Element (<tfoot>) defines a set of rows summarizing the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("tfoot", aspects...)
}

// The HTML Table Header Cell Element (<th>) defines a cell that is a header for a group of cells of a table. The group of cells that the header refers to is defined by the scope and headers attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("th", aspects...)
}

// The HTML Table Head Element (<thead>) defines a set of rows defining the head of the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("thead", aspects...)
}

// The HTML <time> element represents either a time on a 24-hour clock or a precise date in the Gregorian calendar (with optional time and timezone information).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("time", aspects...)
}

// The HTML <title> element (HTML Title Element) defines the title of the document, shown in a browser's title bar or on the page's tab. It can only contain text and any contained tags are not interpreted.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("title", aspects...)
}

// The HTML Table Row Element (<tr>) defines a row of cells in a table. Those can be a mix of <td> and <th> elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("tr", aspects...)
}

// The HTML <track> element is used as a child of the media elements—<audio> and <video>. It lets you specify timed text tracks (or time-based data), for example to automaticaly handle subtitles.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("track", aspects...)
}

// The HTML Underline Element (<u>) renders text with an underline, a line under the baseline of its content.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("u", aspects...)
}

// The HTML unordered list element (<ul>) represents an unordered list of items, namely a collection of items that do not have a numerical ordering, and their order in the list is meaningless. Typically, unordered-list items are displayed with a bullet, which can be of several forms, like a dot, a circle or a squared. The bullet style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("ul", aspects...)
}

// The HTML Variable Element (<var>) represents a variable in a mathematical expression or a programming context.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("var", aspects...)
}

// The HTML <video> element is used to embed video content. It may contain several video sources, represented using the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("video", aspects...)
}

// The Word Break Opportunity (<wbr>) HTML element represents a position within text where the browser may optionally break a line, though its line-breaking rules would not otherwise create a break at that location.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("wbr", aspects...)
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header1(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("h1", aspects...)
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header2(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("h2", aspects...)
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header3(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("h3", aspects...)
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header4(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("h4", aspects...)
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header5(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("h5", aspects...)
}

// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header6(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("h6", aspects...)
}
