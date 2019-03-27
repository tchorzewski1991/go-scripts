package main

import "fmt"

// Document defines common interface for new documents.
// CreatePages() is a factory method that concrete document
// type will implement.
type Document interface {
	CreatePages()
}

// DefaultDocument provides common behavior for each of document
// subtypes. Concrete document will embed DefaultDocument to get
// reference to Pages attribute.
type DefaultDocument struct {
	Pages []string
}

// NewDefaultDocument is a constructor. It returns pointer to
// fresh instance of DefaultDocument.
func NewDefaultDocument() *DefaultDocument {
	return &DefaultDocument{}
}

// String returns document pages concatenated with comma.
func (doc *DefaultDocument) String() string {
	var result string

	for _, page := range doc.Pages {
		result += fmt.Sprintf("%v, ", page)
	}

	return result
}

// AddPage is a setter. It appends new page to the current
// Pages attribute.
func (doc *DefaultDocument) AddPage(page string) {
	doc.Pages = append(doc.Pages, page)
}

// Resume is a document subtype. It implements Document interface
// and embeds DefaultDocument.
type Resume struct {
	*DefaultDocument
}

// NewResume is a constructor. It returns pointer to fresh
// instance of Resume.
func NewResume() *Resume {
	return &Resume{NewDefaultDocument()}
}

// CreatePages implements Document interface on Resume instance.
// CreatePages will add two new pages to Pages member attribute.
func (r *Resume) CreatePages() {
	r.AddPage("Resume Page 1")
	r.AddPage("Resume Page 2")
}

// Report is a document subtype. It implements Document interface
// and embeds DefaultDocument.
type Report struct {
	*DefaultDocument
}

// NewReport is a constructor. It returns pointer to fresh
// instance of Report.
func NewReport() *Report {
	return &Report{NewDefaultDocument()}
}

// CreatePages implements Document interface on Report instance.
// CreatePages will add three new pages to Pages member attribute.
func (r *Report) CreatePages() {
	r.AddPage("Report Page 1")
	r.AddPage("Report Page 2")
	r.AddPage("Report Page 3")
}

func main() {
	// This is important to remember that our doc variable needs to be
	// declared with Document type. Document is an interface type so
	// it can be reassigned easily with anything that implements that
	// type. Decision about type of assigned document, and the behavior
	// related to process of creating pages will be resolved during the
	// runtime.
	var doc Document

	doc = NewResume()
	doc.CreatePages()
	fmt.Println(doc)

	doc = NewReport()
	doc.CreatePages()
	fmt.Println(doc)
}
