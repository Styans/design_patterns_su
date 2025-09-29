package creater

import "practice/internal/document"

type DocumentCreator interface {
	CreateDocument(data string) document.IDocument
}

type ReportCreator struct{}

func (r *ReportCreator) CreateDocument(title string) document.IDocument {
	return &document.Report{Title: title}
}

type ResumeCreator struct{}

func (r *ResumeCreator) CreateDocument(owner string) document.IDocument {
	return &document.Resume{Owner: owner}
}

type LetterCreator struct{}

func (l *LetterCreator) CreateDocument(content string) document.IDocument {
	return &document.Letter{Content: content}
}

func GetCreator(docType string) DocumentCreator {
	switch docType {
	case "report":
		return &ReportCreator{}
	case "resume":
		return &ResumeCreator{}
	case "letter":
		return &LetterCreator{}
	default:
		return nil
	}
}
