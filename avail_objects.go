package avail

type ArrayOfRoute struct {
	Route []Route
}

type ArrayOfPublicMessage struct {
	PublicMessage []PublicMessage
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
}

type ArrayOfStop struct {
	Stop []Stop
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

type ArrayOfStopDeparture struct {
	StopDeparture StopDeparture
}

type StopDeparture struct {
	StopId          int
	StopRecordId    int
	RouteDirections RouteDirections
}

type RouteDirections struct {
	RouteDirection []RouteDirection
}
type RouteDirection struct {
	RouteId        int
	RouteSortOrder int
	RouteRecordId  int
	Direction      string
	DirectionCode  string
	IsHeadway      bool
	IsDone         bool
	Departures     Departures
}

type Departures struct {
	Departure []Departure
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
