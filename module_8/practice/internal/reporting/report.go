package reporting

import (
	"fmt"
	"strings"
)

// 1. Интерфейс IReport
type IReport interface {
	Generate() string
}

// --- 2. Классы базовых отчетов ---

// SalesReport
type SalesReport struct{}

func (r *SalesReport) Generate() string {
	// Фиктивные данные
	return "--- Sales Report ---\n" +
		"Date: 2025-11-01, Amount: 100.50\n" +
		"Date: 2025-11-02, Amount: 150.00\n" +
		"Date: 2025-11-03, Amount: 80.25\n"
}

// UserReport
type UserReport struct{}

func (r *UserReport) Generate() string {
	// Фиктивные данные
	return "--- User Report ---\n" +
		"ID: 1, Name: Alice, Region: US\n" +
		"ID: 2, Name: Bob, Region: EU\n" +
		"ID: 3, Name: Charlie, Region: US\n"
}

/*
 * 3. Абстрактный декоратор (в Go реализуется через композицию)
 * Каждый декоратор будет реализовывать IReport и "оборачивать" IReport.
 */

// --- 4. Декораторы ---

// DateFilterDecorator
type DateFilterDecorator struct {
	wrapped   IReport
	startDate string
	endDate   string
}

func NewDateFilterDecorator(r IReport, start string, end string) *DateFilterDecorator {
	return &DateFilterDecorator{wrapped: r, startDate: start, endDate: end}
}

func (d *DateFilterDecorator) Generate() string {
	data := d.wrapped.Generate()
	// В реальном приложении здесь была бы логика парсинга и фильтрации.
	// Мы просто добавляем информацию о примененном фильтре.
	filterInfo := fmt.Sprintf("[Filtered by Date from %s to %s]\n", d.startDate, d.endDate)
	return data + filterInfo
}

// SortingDecorator
type SortingDecorator struct {
	wrapped IReport
	sortBy  string
}

func NewSortingDecorator(r IReport, by string) *SortingDecorator {
	return &SortingDecorator{wrapped: r, sortBy: by}
}

func (d *SortingDecorator) Generate() string {
	data := d.wrapped.Generate()
	// Моделируем сортировку
	sortInfo := fmt.Sprintf("[Sorted by %s]\n", d.sortBy)
	return data + sortInfo
}

// CsvExportDecorator
type CsvExportDecorator struct {
	wrapped IReport
}

func NewCsvExportDecorator(r IReport) *CsvExportDecorator {
	return &CsvExportDecorator{wrapped: r}
}

func (d *CsvExportDecorator) Generate() string {
	data := d.wrapped.Generate()
	// Моделируем экспорт в CSV
	csvData := strings.ReplaceAll(data, "---", "")
	csvData = strings.ReplaceAll(csvData, "\n", ", ") // Упрощенная конвертация
	return fmt.Sprintf("<CSV>\n%s\n</CSV>\n", csvData)
}

// PdfExportDecorator
type PdfExportDecorator struct {
	wrapped IReport
}

func NewPdfExportDecorator(r IReport) *PdfExportDecorator {
	return &PdfExportDecorator{wrapped: r}
}

func (d *PdfExportDecorator) Generate() string {
	data := d.wrapped.Generate()
	// Моделируем экспорт в PDF
	return fmt.Sprintf("<PDF>\n%s\n</PDF>\n", data)
}

// --- Задание 2: Новые декораторы ---

// AmountFilterDecorator
type AmountFilterDecorator struct {
	wrapped   IReport
	minAmount float64
}

func NewAmountFilterDecorator(r IReport, min float64) *AmountFilterDecorator {
	return &AmountFilterDecorator{wrapped: r, minAmount: min}
}

func (d *AmountFilterDecorator) Generate() string {
	data := d.wrapped.Generate()
	filterInfo := fmt.Sprintf("[Filtered by Amount > %.2f]\n", d.minAmount)
	return data + filterInfo
}

// --- Задание 3: Механизм динамического выбора (Фабрика/Конструктор) ---

// ReportRequest описывает пользовательский запрос
type ReportRequest struct {
	ReportType string // "sales" or "user"
	StartDate  string
	EndDate    string
	SortBy     string
	ExportAs   string // "csv", "pdf", or ""
	MinAmount  float64
}

// BuildReport динамически создает отчет на основе запроса
func BuildReport(req ReportRequest) IReport {
	var baseReport IReport

	// 1. Выбираем базовый отчет
	switch req.ReportType {
	case "sales":
		baseReport = &SalesReport{}
	case "user":
		baseReport = &UserReport{}
	default:
		// В реальном коде здесь должна быть обработка ошибки
		fmt.Println("Error: Unknown report type")
		return nil
	}

	// 2. Динамически "оборачиваем" в декораторы
	if req.StartDate != "" && req.EndDate != "" {
		baseReport = NewDateFilterDecorator(baseReport, req.StartDate, req.EndDate)
	}

	if req.MinAmount > 0 {
		baseReport = NewAmountFilterDecorator(baseReport, req.MinAmount)
	}

	if req.SortBy != "" {
		baseReport = NewSortingDecorator(baseReport, req.SortBy)
	}

	// 3. Декораторы экспорта (обычно применяются последними)
	switch req.ExportAs {
	case "csv":
		baseReport = NewCsvExportDecorator(baseReport)
	case "pdf":
		baseReport = NewPdfExportDecorator(baseReport)
	}

	return baseReport
}
