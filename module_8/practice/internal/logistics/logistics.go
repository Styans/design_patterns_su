package logistics

import (
	"fmt"
	"log"
	"strconv"
)

// --- Задание 4: Логгирование ---
// Простой внутренний логгер
func logDelivery(message string) {
	log.Printf("[Logistics] %s\n", message)
}

// --- 1. Основные компоненты (Целевой интерфейс) ---

type IInternalDeliveryService interface {
	DeliverOrder(orderId string) string      // Доставить заказ
	GetDeliveryStatus(orderId string) string // Получить статус
	// Задание 5: Расчет стоимости
	CalculateDeliveryCost(orderId string) float64
}

// Класс InternalDeliveryService
type InternalDeliveryService struct{}

func (s *InternalDeliveryService) DeliverOrder(orderId string) string {
	status := fmt.Sprintf("Internal: Order %s is out for delivery.", orderId)
	logDelivery(status)
	return status
}

func (s *InternalDeliveryService) GetDeliveryStatus(orderId string) string {
	status := fmt.Sprintf("Internal: Order %s is near your location.", orderId)
	logDelivery(status)
	return status
}

func (s *InternalDeliveryService) CalculateDeliveryCost(orderId string) float64 {
	cost := 10.0 // Фиксированная внутренняя стоимость
	logDelivery(fmt.Sprintf("Internal: Cost for %s is %.2f", orderId, cost))
	return cost
}

// --- 2. Сторонние службы логистики (Адаптируемые классы) ---

// ExternalLogisticsServiceA
type ExternalLogisticsServiceA struct{}

func (s *ExternalLogisticsServiceA) ShipItem(itemId int) (int, error) {
	logDelivery(fmt.Sprintf("ServiceA: Shipping item %d.", itemId))
	if itemId <= 0 { // Задание 4: Обработка ошибок
		return 0, fmt.Errorf("invalid item ID")
	}
	return 12345, nil // Фиктивный ID отправки
}
func (s *ExternalLogisticsServiceA) TrackShipment(shipmentId int) string {
	return fmt.Sprintf("ServiceA: Shipment %d is in transit.", shipmentId)
}
func (s *ExternalLogisticsServiceA) GetShippingQuote(itemId int) float64 {
	return 25.50 // Стоимость
}

// ExternalLogisticsServiceB
type ExternalLogisticsServiceB struct{}

func (s *ExternalLogisticsServiceB) SendPackage(packageInfo string) (string, error) {
	logDelivery(fmt.Sprintf("ServiceB: Sending package: %s", packageInfo))
	if packageInfo == "" { // Задание 4: Обработка ошибок
		return "", fmt.Errorf("package info cannot be empty")
	}
	return "B-XYZ-789", nil // Фиктивный код отслеживания
}
func (s *ExternalLogisticsServiceB) CheckPackageStatus(trackingCode string) string {
	return fmt.Sprintf("ServiceB: Package %s is at local depot.", trackingCode)
}
func (s *ExternalLogisticsServiceB) GetPrice(packageInfo string) float64 {
	return 18.75 // Стоимость
}

// --- Задание 2: Дополнительная сторонняя служба ---
// ExternalLogisticsServiceC
type ExternalLogisticsServiceC struct{}

func (s *ExternalLogisticsServiceC) CreateJob(destination string) (string, error) {
	logDelivery(fmt.Sprintf("ServiceC: Creating job for %s", destination))
	return "C-JOB-999", nil
}
func (s *ExternalLogisticsServiceC) DispatchJob(jobId string) bool {
	logDelivery(fmt.Sprintf("ServiceC: Job %s dispatched.", jobId))
	return true
}
func (s *ExternalLogisticsServiceC) PollJob(jobId string) string {
	return fmt.Sprintf("ServiceC: Job %s is en route.", jobId)
}
func (s *ExternalLogisticsServiceC) GetJobCost(jobId string) float64 {
	return 42.00 // Стоимость
}

// --- 3. Адаптеры ---

// LogisticsAdapterA
type LogisticsAdapterA struct {
	service     *ExternalLogisticsServiceA
	shipmentMap map[string]int // Карта [orderId -> shipmentId]
}

func NewLogisticsAdapterA(s *ExternalLogisticsServiceA) *LogisticsAdapterA {
	return &LogisticsAdapterA{service: s, shipmentMap: make(map[string]int)}
}

func (a *LogisticsAdapterA) DeliverOrder(orderId string) string {
	logDelivery("AdapterA: Received DeliverOrder request.")
	// Адаптация: конвертируем orderId (string) в itemId (int)
	itemId, err := strconv.Atoi(orderId)
	if err != nil { // Задание 4: Обработка ошибок
		logDelivery(fmt.Sprintf("AdapterA: Error converting orderId %s: %v", orderId, err))
		return "AdapterA: Error - Invalid Order ID format"
	}
	shipmentId, err := a.service.ShipItem(itemId)
	if err != nil { // Задание 4: Обработка ошибок
		logDelivery(fmt.Sprintf("AdapterA: Error from ServiceA: %v", err))
		return fmt.Sprintf("AdapterA: Error - %v", err)
	}
	a.shipmentMap[orderId] = shipmentId // Сохраняем связь
	return fmt.Sprintf("AdapterA: Shipment %d created for order %s.", shipmentId, orderId)
}

func (a *LogisticsAdapterA) GetDeliveryStatus(orderId string) string {
	shipmentId, ok := a.shipmentMap[orderId]
	if !ok { // Задание 4: Обработка ошибок
		return "AdapterA: Error - No shipment found for this order."
	}
	return a.service.TrackShipment(shipmentId)
}

func (a *LogisticsAdapterA) CalculateDeliveryCost(orderId string) float64 {
	itemId, _ := strconv.Atoi(orderId) // Упрощаем, опуская ошибку
	cost := a.service.GetShippingQuote(itemId)
	logDelivery(fmt.Sprintf("AdapterA: Cost for order %s is %.2f", orderId, cost))
	return cost
}

// LogisticsAdapterB
type LogisticsAdapterB struct {
	service     *ExternalLogisticsServiceB
	trackingMap map[string]string // Карта [orderId -> trackingCode]
}

func NewLogisticsAdapterB(s *ExternalLogisticsServiceB) *LogisticsAdapterB {
	return &LogisticsAdapterB{service: s, trackingMap: make(map[string]string)}
}

func (a *LogisticsAdapterB) DeliverOrder(orderId string) string {
	logDelivery("AdapterB: Received DeliverOrder request.")
	// Адаптация: используем orderId как packageInfo
	packageInfo := fmt.Sprintf("Order #%s, Priority", orderId)
	trackingCode, err := a.service.SendPackage(packageInfo)
	if err != nil { // Задание 4
		logDelivery(fmt.Sprintf("AdapterB: Error from ServiceB: %v", err))
		return fmt.Sprintf("AdapterB: Error - %v", err)
	}
	a.trackingMap[orderId] = trackingCode
	return fmt.Sprintf("AdapterB: Package %s sent for order %s.", trackingCode, orderId)
}

func (a *LogisticsAdapterB) GetDeliveryStatus(orderId string) string {
	trackingCode, ok := a.trackingMap[orderId]
	if !ok { // Задание 4
		return "AdapterB: Error - No package found for this order."
	}
	return a.service.CheckPackageStatus(trackingCode)
}

func (a *LogisticsAdapterB) CalculateDeliveryCost(orderId string) float64 {
	packageInfo := fmt.Sprintf("Order #%s, Priority", orderId)
	cost := a.service.GetPrice(packageInfo)
	logDelivery(fmt.Sprintf("AdapterB: Cost for order %s is %.2f", orderId, cost))
	return cost
}

// LogisticsAdapterC (для Задания 2)
type LogisticsAdapterC struct {
	service *ExternalLogisticsServiceC
	jobMap  map[string]string // Карта [orderId -> jobId]
}

func NewLogisticsAdapterC(s *ExternalLogisticsServiceC) *LogisticsAdapterC {
	return &LogisticsAdapterC{service: s, jobMap: make(map[string]string)}
}

func (a *LogisticsAdapterC) DeliverOrder(orderId string) string {
	logDelivery("AdapterC: Received DeliverOrder request.")
	// Адаптация: извлекаем адрес (фиктивный)
	destination := "123 Main St, Order " + orderId
	jobId, err := a.service.CreateJob(destination)
	if err != nil { // Задание 4
		logDelivery(fmt.Sprintf("AdapterC: Error from ServiceC: %v", err))
		return fmt.Sprintf("AdapterC: Error - %v", err)
	}
	a.service.DispatchJob(jobId)
	a.jobMap[orderId] = jobId
	return fmt.Sprintf("AdapterC: Job %s dispatched for order %s.", jobId, orderId)
}

func (a *LogisticsAdapterC) GetDeliveryStatus(orderId string) string {
	jobId, ok := a.jobMap[orderId]
	if !ok { // Задание 4
		return "AdapterC: Error - No job found for this order."
	}
	return a.service.PollJob(jobId)
}

func (a *LogisticsAdapterC) CalculateDeliveryCost(orderId string) float64 {
	// Для C нам нужен jobId, если его нет - создадим временный для расчета
	jobId, ok := a.jobMap[orderId]
	if !ok {
		jobId, _ = a.service.CreateJob("Temporary Quote Destination")
	}
	cost := a.service.GetJobCost(jobId)
	logDelivery(fmt.Sprintf("AdapterC: Cost for order %s is %.2f", orderId, cost))
	return cost
}

// --- 4. Фабрика ---

// DeliveryServiceFactory (Задание 3: расширенная)
func GetDeliveryService(provider string) (IInternalDeliveryService, error) {
	logDelivery(fmt.Sprintf("Factory: Request for provider '%s'", provider))
	switch provider {
	case "internal":
		return &InternalDeliveryService{}, nil
	case "serviceA":
		// Фабрика скрывает создание адаптера и сервиса
		return NewLogisticsAdapterA(&ExternalLogisticsServiceA{}), nil
	case "serviceB":
		return NewLogisticsAdapterB(&ExternalLogisticsServiceB{}), nil
	case "serviceC": // Задание 3
		return NewLogisticsAdapterC(&ExternalLogisticsServiceC{}), nil
	default:
		// Задание 4: Обработка ошибок
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}
