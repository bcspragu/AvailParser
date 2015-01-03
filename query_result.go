package avail

type QueryResult map[Stop]DeparturesByRoute

type DeparturesByRoute map[Route][]Departure
