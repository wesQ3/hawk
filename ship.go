package hawk

import (
	"time"
)

type Ship struct {
	Symbol        string      `json:"symbol"`
	Registration  Registration `json:"registration"`
	Nav           Nav          `json:"nav"`
	Crew          Crew         `json:"crew"`
	Frame         Frame        `json:"frame"`
	Reactor       Reactor      `json:"reactor"`
	Engine        Engine       `json:"engine"`
	Modules       []Module     `json:"modules"`
	Mounts        []Mount      `json:"mounts"`
	Cargo         Cargo        `json:"cargo"`
	Fuel          Fuel         `json:"fuel"`
}

type Registration struct {
	Name         string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role         string `json:"role"`
}

type Nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          Route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type Route struct {
	Destination   Location    `json:"destination"`
	Departure     Location    `json:"departure"`
	DepartureTime time.Time   `json:"departureTime"`
	Arrival       time.Time   `json:"arrival"`
}

type Location struct {
	Symbol       string  `json:"symbol"`
	Type         string  `json:"type"`
	SystemSymbol string  `json:"systemSymbol"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
}

type Crew struct {
	Current    int     `json:"current"`
	Required   int     `json:"required"`
	Capacity   int     `json:"capacity"`
	Rotation   string  `json:"rotation"`
	Morale     float64 `json:"morale"`
	Wages      float64 `json:"wages"`
}

type Frame struct {
	Symbol        string       `json:"symbol"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Condition     float64      `json:"condition"`
	ModuleSlots   int          `json:"moduleSlots"`
	MountingPoints int         `json:"mountingPoints"`
	FuelCapacity  float64      `json:"fuelCapacity"`
	Requirements  Requirements `json:"requirements"`
}

type Reactor struct {
	Symbol        string       `json:"symbol"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Condition     float64      `json:"condition"`
	PowerOutput   float64      `json:"powerOutput"`
	Requirements  Requirements `json:"requirements"`
}

type Engine struct {
	Symbol        string       `json:"symbol"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Condition     float64      `json:"condition"`
	Speed         float64      `json:"speed"`
	Requirements  Requirements `json:"requirements"`
}

type Module struct {
	Symbol        string       `json:"symbol"`
	Capacity      float64      `json:"capacity"`
	Range         float64      `json:"range"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Requirements  Requirements `json:"requirements"`
}

type Mount struct {
	Symbol        string       `json:"symbol"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Strength      float64      `json:"strength"`
	Deposits      []string     `json:"deposits"`
	Requirements  Requirements `json:"requirements"`
}

type Cargo struct {
	Capacity  int          `json:"capacity"`
	Units     int          `json:"units"`
	Inventory []InventoryItem `json:"inventory"`
}

type InventoryItem struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int    `json:"units"`
}

type Fuel struct {
	Current  float64     `json:"current"`
	Capacity float64     `json:"capacity"`
	Consumed ConsumedFuel `json:"consumed"`
}

type ConsumedFuel struct {
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type Requirements struct {
	Power float64 `json:"power"`
	Crew  int     `json:"crew"`
	Slots int     `json:"slots"`
}
