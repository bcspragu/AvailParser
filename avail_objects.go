package avail

type ArrayOfRoute struct {
	Route []Route
}

type ArrayOfPublicMessage struct {
	PublicMessage []PublicMessage
}

type ArrayOfStop struct {
	Stop []Stop
}

type ArrayOfStopDeparture struct {
	StopDeparture StopDeparture
}

type ArrayOfVehicleLocation struct {
	VehicleLocation []VehicleLocation
}

type PublicMessage struct {
	DaysOfWeek   int
	FromDate     AvailTime
	FromTime     AvailTime
	Message      string
	MessageId    int
	Priority     int
	PublicAccess int
	Published    bool
	Routes       []Route
	ToDate       AvailTime
	ToTime       AvailTime
}

type Route struct {
	ServerId           int
	RouteId            int
	RouteRecordId      int
	ShortName          string
	LongName           string
	RouteAbbreviation  string
	Color              string
	TextColor          string
	IsVisible          bool
	SortOrder          int
	RouteTraceFilename string
	IsHeadway          bool
	IncludeInGoogle    bool
	Stops              []Stop            `xml:"Stops>Stop"`
	RouteStops         []RouteStop       `xml:"RouteStops>RouteStop"`
	Directions         []Direction       `xml:"Directions>Direction"`
	Vehicles           []VehicleLocation `xml:"VehicleLocations>VehicleLocation"`
}

type RouteStop struct {
	RouteId   int
	StopId    int
	SortOrder int
	Direction string
}

type Direction struct {
	Dir string
}

type VehicleLocation struct {
	VehicleId      int
	Name           string
	Latitude       float64
	Longitude      float64
	RouteId        int
	TripId         int
	RunId          int
	Direction      string
	DirectionLong  string
	Heading        int
	Deviation      float64
	OpStatus       string
	CommStatus     string
	GPSStatus      int
	LastStop       string
	LastUpdated    AvailTime
	DisplayStatus  string
	BlockFareboxId int
}

type Stop struct {
	StopId       int
	StopRecordId int
	Name         string
	Description  string
	Latitude     float64
	Longitude    float64
	IsTimePoint  bool
}

type StopDeparture struct {
	StopId          int
	StopRecordId    int
	RouteDirections []RouteDirection `xml:"RouteDirections>RouteDirection"`
}

type RouteDirection struct {
	RouteId        int
	RouteSortOrder int
	RouteRecordId  int
	Direction      string
	DirectionCode  string
	IsHeadway      bool
	IsDone         bool
	Departures     []Departure `xml:"Departures>Departure"`
}

type Departure struct {
	EDT         AvailTime
	ETA         AvailTime
	LastUpdated AvailTime
	SDT         AvailTime
	STA         AvailTime
	Trip        Trip
}

type Trip struct {
	TripId               int
	RunId                int
	BlockFareboxId       int
	TripRecordId         int
	RunRecordId          int
	ServiceLevelRecordId int
	StopSequence         int
	PatternRecordId      int
	TripStartTime        string
	InternalSignDesc     string
	InternetServiceDesc  string
	IVRServiceDesc       string
	TripDirection        string
}
