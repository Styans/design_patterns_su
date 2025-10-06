package yagni

type User struct {
	Name  string
	Email string
}

func (u User) SaveToDatabase() {}

type FileReader struct{}

func (fr FileReader) ReadFile(filePath string) string {
	return "file content"
}

type ReportGenerator struct{}

func (rg ReportGenerator) GeneratePdfReport() {}
