package shared

type Airport struct {
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Airports returns a list of airports with their codes, names, cities, countries, and coordinates.
func Airports() []Airport {
	return []Airport{
		{Code: "GRU", Name: "São Paulo-Guarulhos International Airport", City: "São Paulo", Country: "Brazil", Latitude: -23.4356, Longitude: -46.4731},
		{Code: "EZE", Name: "Ministro Pistarini International Airport", City: "Buenos Aires", Country: "Argentina", Latitude: -34.8222, Longitude: -58.5358},
		{Code: "SCL", Name: "Arturo Merino Benitez International Airport", City: "Santiago", Country: "Chile", Latitude: -33.3930, Longitude: -70.7858},

		{Code: "LHR", Name: "London Heathrow Airport", City: "London", Country: "United Kingdom", Latitude: 51.4700, Longitude: -0.4543},
		{Code: "FCO", Name: "Leonardo da Vinci International Airport", City: "Rome", Country: "Italy", Latitude: 41.8003, Longitude: 12.2389},
		{Code: "WAW", Name: "Warsaw Chopin Airport", City: "Warsaw", Country: "Poland", Latitude: 52.1657, Longitude: 20.9671},

		{Code: "DXB", Name: "Dubai International Airport", City: "Dubai", Country: "United Arab Emirates", Latitude: 25.2532, Longitude: 55.3657},
		{Code: "DOH", Name: "Hamad International Airport", City: "Doha", Country: "Qatar", Latitude: 25.2736, Longitude: 51.6080},
		{Code: "DEL", Name: "Indira Gandhi International Airport", City: "New Delhi", Country: "India", Latitude: 28.5562, Longitude: 77.1000},

		{Code: "MIA", Name: "Miami International Airport", City: "Miami", Country: "United States", Latitude: 25.7959, Longitude: -80.2870},
		{Code: "ONT", Name: "Ontario International Airport", City: "Ontario", Country: "United States", Latitude: 34.0559, Longitude: -117.6005},
		{Code: "MEX", Name: "Mexico City International Airport Benito Juárez", City: "Mexico City", Country: "Mexico", Latitude: 19.4361, Longitude: -99.0719},

		{Code: "HND", Name: "Tokyo Haneda Airport", City: "Tokyo", Country: "Japan", Latitude: 35.5494, Longitude: 139.7798},
		{Code: "ICN", Name: "Incheon International Airport", City: "Incheon", Country: "South Korea", Latitude: 37.4602, Longitude: 126.4407},
		{Code: "PEK", Name: "Beijing Capital International Airport", City: "Beijing", Country: "China", Latitude: 40.0801, Longitude: 116.5846},

		{Code: "JNB", Name: "O. R. Tambo International Airport", City: "Johannesburg", Country: "South Africa", Latitude: -26.1392, Longitude: 28.2460},
		{Code: "CAI", Name: "Cairo International Airport", City: "Cairo", Country: "Egypt", Latitude: 30.1120, Longitude: 31.4000},
		{Code: "RBA", Name: "Rabat–Salé Airport", City: "Rabat", Country: "Morocco", Latitude: 34.0515, Longitude: -6.7515},

		{Code: "SYD", Name: "Sydney Kingsford Smith Airport", City: "Sydney", Country: "Australia", Latitude: -33.9399, Longitude: 151.1753},
		{Code: "WLG", Name: "Wellington International Airport", City: "Wellington", Country: "New Zealand", Latitude: -41.3272, Longitude: 174.8050},
	}
}
