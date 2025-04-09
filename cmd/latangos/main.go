package latangos

type Flight struct {
	FlightNumber string
	Origin       string
	Destination  string
	Departure    string
	Arrival      string
	BasePrice    float64
}

type DayMultipliers struct {
	Monday    float64
	Tuesday   float64
	Wednesday float64
	Thursday  float64
	Friday    float64
	Saturday  float64
	Sunday    float64
}

func NewDayMultipliers() DayMultipliers {
	return DayMultipliers{
		Monday:    1.0,
		Tuesday:   1.0,
		Wednesday: 1.0,
		Thursday:  1.0,
		Friday:    1.0,
		Saturday:  1.0,
		Sunday:    1.0,
	}
}

func Run() {
}
