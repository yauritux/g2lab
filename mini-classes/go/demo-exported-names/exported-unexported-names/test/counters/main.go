package counters

// AlertCounter is an unexported type that
// contains an integer counter for alerts.
type alertCounter int

// NewAlertCounter creates and returns objects of
// the unexported type alertCounter.
func NewAlertCounter(value int) alertCounter {
	return alertCounter(value)
}
