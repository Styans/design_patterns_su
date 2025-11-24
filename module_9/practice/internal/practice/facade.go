package practice

import "fmt"

// -------------------- Подсистема: бронирование номеров --------------------

type RoomBookingSystem struct {
	nextID int
}

func NewRoomBookingSystem() *RoomBookingSystem {
	return &RoomBookingSystem{}
}

func (r *RoomBookingSystem) CheckAvailability(roomType string) bool {
	fmt.Printf("[RoomBooking] Checking availability for room type: %s\n", roomType)
	// Для лабораторной работы просто считаем, что номера всегда есть
	return true
}

func (r *RoomBookingSystem) BookRoom(guestName, roomType string, nights, roomNumber int) string {
	r.nextID++
	id := fmt.Sprintf("RB-%03d", r.nextID)
	fmt.Printf("[RoomBooking] Room booked: id=%s, guest=%s, type=%s, nights=%d, room=%d\n",
		id, guestName, roomType, nights, roomNumber)
	return id
}

func (r *RoomBookingSystem) CancelBooking(bookingID string) {
	fmt.Printf("[RoomBooking] Booking %s is cancelled\n", bookingID)
}

// -------------------- Подсистема: ресторан --------------------

type RestaurantSystem struct {
	nextID int
}

func NewRestaurantSystem() *RestaurantSystem {
	return &RestaurantSystem{}
}

func (r *RestaurantSystem) BookTable(guestName string, persons int, time string) string {
	r.nextID++
	id := fmt.Sprintf("RS-%03d", r.nextID)
	fmt.Printf("[Restaurant] Table booked: id=%s, guest=%s, persons=%d, time=%s\n",
		id, guestName, persons, time)
	return id
}

func (r *RestaurantSystem) OrderFood(guestName string, items []string) {
	fmt.Printf("[Restaurant] Food ordered for %s: %v\n", guestName, items)
}

func (r *RestaurantSystem) CancelTable(bookingID string) {
	fmt.Printf("[Restaurant] Table booking %s is cancelled\n", bookingID)
}

// -------------------- Подсистема: мероприятия --------------------

type EventManagementSystem struct {
	nextID int
}

func NewEventManagementSystem() *EventManagementSystem {
	return &EventManagementSystem{}
}

func (e *EventManagementSystem) BookEventHall(eventName, date string, participants int) string {
	e.nextID++
	id := fmt.Sprintf("EV-%03d", e.nextID)
	fmt.Printf("[Event] Event hall booked: id=%s, name=%s, date=%s, participants=%d\n",
		id, eventName, date, participants)
	return id
}

func (e *EventManagementSystem) OrderEquipment(eventID string, equipment []string) {
	fmt.Printf("[Event] Equipment ordered for event %s: %v\n", eventID, equipment)
}

func (e *EventManagementSystem) CancelEvent(eventID string) {
	fmt.Printf("[Event] Event %s is cancelled\n", eventID)
}

// -------------------- Подсистема: уборка --------------------

type CleaningService struct{}

func NewCleaningService() *CleaningService {
	return &CleaningService{}
}

func (c *CleaningService) ScheduleCleaning(roomNumber int, time string) {
	fmt.Printf("[Cleaning] Cleaning scheduled: room=%d, time=%s\n", roomNumber, time)
}

func (c *CleaningService) PerformCleaning(roomNumber int) {
	fmt.Printf("[Cleaning] Cleaning performed in room %d\n", roomNumber)
}

func (c *CleaningService) RequestUrgentCleaning(roomNumber int) {
	fmt.Printf("[Cleaning] URGENT cleaning requested for room %d\n", roomNumber)
}

// -------------------- Доп. подсистема: такси --------------------

type TaxiService struct{}

func NewTaxiService() *TaxiService {
	return &TaxiService{}
}

func (t *TaxiService) OrderTaxi(guestName, time string) {
	fmt.Printf("[Taxi] Taxi ordered for %s at %s\n", guestName, time)
}

// -------------------- Фасад: HotelFacade --------------------

type HotelFacade struct {
	roomBooking *RoomBookingSystem
	restaurant  *RestaurantSystem
	events      *EventManagementSystem
	cleaning    *CleaningService
	taxi        *TaxiService
}

func NewHotelFacade() *HotelFacade {
	return &HotelFacade{
		roomBooking: NewRoomBookingSystem(),
		restaurant:  NewRestaurantSystem(),
		events:      NewEventManagementSystem(),
		cleaning:    NewCleaningService(),
		taxi:        NewTaxiService(),
	}
}

// Сценарий 1: Бронирование номера + ресторан + уборка
func (h *HotelFacade) BookRoomWithRestaurantAndCleaning(
	guestName, roomType string,
	nights, roomNumber int,
	dinnerTime string,
) (roomBookingID, tableBookingID string) {

	fmt.Println("=== Scenario: Book room with restaurant and cleaning ===")

	if !h.roomBooking.CheckAvailability(roomType) {
		fmt.Println("[HotelFacade] No rooms available")
		return "", ""
	}

	roomBookingID = h.roomBooking.BookRoom(guestName, roomType, nights, roomNumber)
	h.cleaning.ScheduleCleaning(roomNumber, "10:00")
	tableBookingID = h.restaurant.BookTable(guestName, 2, dinnerTime)
	h.restaurant.OrderFood(guestName, []string{"Steak", "Salad", "Juice"})

	return roomBookingID, tableBookingID
}

// Сценарий 2: Организация мероприятия + номера + оборудование
func (h *HotelFacade) OrganizeEventWithRoomsAndEquipment(
	eventName, date string,
	participants int,
) (eventID, roomBookingID string) {

	fmt.Println("=== Scenario: Organize event with rooms and equipment ===")

	eventID = h.events.BookEventHall(eventName, date, participants)
	h.events.OrderEquipment(eventID, []string{"Projector", "Microphones", "Speakers"})

	// Допустим, бронируем блок номеров для участников
	roomBookingID = h.roomBooking.BookRoom(
		fmt.Sprintf("Group-%s", eventName),
		"Standard", (participants / 2), 500)

	return eventID, roomBookingID
}

// Сценарий 3: Бронирование стола + автоматический заказ такси
func (h *HotelFacade) BookRestaurantTableWithTaxi(
	guestName string,
	persons int,
	time string,
) string {

	fmt.Println("=== Scenario: Book restaurant table with taxi ===")

	tableID := h.restaurant.BookTable(guestName, persons, time)
	h.taxi.OrderTaxi(guestName, time)
	return tableID
}

// Доп. методы фасада: отмена бронирований и уборка по запросу

func (h *HotelFacade) CancelRoomBooking(bookingID string) {
	fmt.Println("=== Cancel room booking via facade ===")
	h.roomBooking.CancelBooking(bookingID)
}

func (h *HotelFacade) CancelRestaurantBooking(bookingID string) {
	fmt.Println("=== Cancel restaurant booking via facade ===")
	h.restaurant.CancelTable(bookingID)
}

func (h *HotelFacade) CancelEventBooking(eventID string) {
	fmt.Println("=== Cancel event booking via facade ===")
	h.events.CancelEvent(eventID)
}

func (h *HotelFacade) RequestCleaningOnDemand(roomNumber int) {
	fmt.Println("=== On-demand cleaning via facade ===")
	h.cleaning.RequestUrgentCleaning(roomNumber)
}
