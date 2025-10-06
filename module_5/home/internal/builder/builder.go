package builder

type Report struct {
	Header  string
	Content string
	Footer  string
}

func (r Report) String() string {
	return r.Header + "\n" + r.Content + "\n" + r.Footer
}

type IReportBuilder interface {
	SetHeader(string)
	SetContent(string)
	SetFooter(string)
	GetReport() Report
}

type TextReportBuilder struct {
	report Report
}

func (b *TextReportBuilder) SetHeader(header string) {
	b.report.Header = "TEXT HEADER: " + header
}

func (b *TextReportBuilder) SetContent(content string) {
	b.report.Content = content
}

func (b *TextReportBuilder) SetFooter(footer string) {
	b.report.Footer = "TEXT FOOTER: " + footer
}

func (b *TextReportBuilder) GetReport() Report {
	return b.report
}

type HtmlReportBuilder struct {
	report Report
}

func (b *HtmlReportBuilder) SetHeader(header string) {
	b.report.Header = "<h1>" + header + "</h1>"
}

func (b *HtmlReportBuilder) SetContent(content string) {
	b.report.Content = "<p>" + content + "</p>"
}

func (b *HtmlReportBuilder) SetFooter(footer string) {
	b.report.Footer = "<footer>" + footer + "</footer>"
}

func (b *HtmlReportBuilder) GetReport() Report {
	return b.report
}

type ReportDirector struct{}

func (r ReportDirector) ConstructReport(b IReportBuilder, header, content, footer string) Report {
	b.SetHeader(header)
	b.SetContent(content)
	b.SetFooter(footer)
	return b.GetReport()
}
