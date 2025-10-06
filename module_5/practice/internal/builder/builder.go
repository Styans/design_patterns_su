package builder

import (
	"fmt"
	"os"
	"strings"
)

type ReportStyle struct {
	BgColor   string
	TextColor string
	FontSize  int
}

type Section struct {
	Title   string
	Content string
}

type Report struct {
	Header   string
	Content  string
	Footer   string
	Sections []Section
	Style    ReportStyle
}

func (r Report) ExportText(path string) error {
	b := strings.Builder{}
	b.WriteString(r.Header + "\n\n")
	for _, s := range r.Sections {
		b.WriteString(s.Title + "\n")
		b.WriteString(s.Content + "\n\n")
	}
	b.WriteString(r.Content + "\n\n")
	b.WriteString(r.Footer + "\n")
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func (r Report) ExportHTML(path string) error {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("<html><head><style>body{background:%s;color:%s;font-size:%dpx}</style></head><body>", r.Style.BgColor, r.Style.TextColor, r.Style.FontSize))
	b.WriteString(fmt.Sprintf("<header>%s</header>", r.Header))
	for _, s := range r.Sections {
		b.WriteString(fmt.Sprintf("<h2>%s</h2><div>%s</div>", s.Title, s.Content))
	}
	b.WriteString(fmt.Sprintf("<main>%s</main>", r.Content))
	b.WriteString(fmt.Sprintf("<footer>%s</footer>", r.Footer))
	b.WriteString("</body></html>")
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func (r Report) ExportPDF(path string) error {
	content := "PDF REPORT\n\n" + r.Header + "\n\n"
	for _, s := range r.Sections {
		content += s.Title + "\n" + s.Content + "\n\n"
	}
	content += r.Content + "\n\n" + r.Footer + "\n"
	return os.WriteFile(path, []byte(content), 0644)
}

type IReportBuilder interface {
	SetHeader(string)
	SetContent(string)
	SetFooter(string)
	AddSection(string, string)
	SetStyle(ReportStyle)
	GetReport() Report
}

type TextReportBuilder struct {
	report Report
}

func (b *TextReportBuilder) SetHeader(h string) {
	b.report.Header = h
}

func (b *TextReportBuilder) SetContent(c string) {
	b.report.Content = c
}

func (b *TextReportBuilder) SetFooter(f string) {
	b.report.Footer = f
}

func (b *TextReportBuilder) AddSection(title, content string) {
	b.report.Sections = append(b.report.Sections, Section{Title: title, Content: content})
}

func (b *TextReportBuilder) SetStyle(s ReportStyle) {
	b.report.Style = s
}

func (b *TextReportBuilder) GetReport() Report {
	return b.report
}

type HtmlReportBuilder struct {
	report Report
}

func (b *HtmlReportBuilder) SetHeader(h string) {
	b.report.Header = "<h1>" + h + "</h1>"
}

func (b *HtmlReportBuilder) SetContent(c string) {
	b.report.Content = c
}

func (b *HtmlReportBuilder) SetFooter(f string) {
	b.report.Footer = "<small>" + f + "</small>"
}

func (b *HtmlReportBuilder) AddSection(title, content string) {
	b.report.Sections = append(b.report.Sections, Section{Title: title, Content: content})
}

func (b *HtmlReportBuilder) SetStyle(s ReportStyle) {
	b.report.Style = s
}

func (b *HtmlReportBuilder) GetReport() Report {
	return b.report
}

type PdfReportBuilder struct {
	report Report
}

func (b *PdfReportBuilder) SetHeader(h string) {
	b.report.Header = h
}

func (b *PdfReportBuilder) SetContent(c string) {
	b.report.Content = c
}

func (b *PdfReportBuilder) SetFooter(f string) {
	b.report.Footer = f
}

func (b *PdfReportBuilder) AddSection(title, content string) {
	b.report.Sections = append(b.report.Sections, Section{Title: title, Content: content})
}

func (b *PdfReportBuilder) SetStyle(s ReportStyle) {
	b.report.Style = s
}

func (b *PdfReportBuilder) GetReport() Report {
	return b.report
}

type ReportDirector struct{}

func (d ReportDirector) Construct(b IReportBuilder, header, content, footer string, sections []Section, style ReportStyle) Report {
	b.SetHeader(header)
	for _, s := range sections {
		b.AddSection(s.Title, s.Content)
	}
	b.SetContent(content)
	b.SetFooter(footer)
	b.SetStyle(style)
	return b.GetReport()
}
