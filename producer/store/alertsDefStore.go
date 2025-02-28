package store

type Notification interface {
	AddAlert(item string, newAlertDef alert)
}

// alert type definition
type alert struct {
	alertType    string //"Price Alert - Low Price", "Price Alert - High Price"
	priceTrigger int
}

// AlertsActiveType - used to store the active alert definitions
type AlertsActiveType map[string]alert

// Instances of alerts
var alertsActive = AlertsActiveType{
	"Hulk Figure":     {alertType: "Price Alert - High Price", priceTrigger: 865},
	"Deadpool Figure": {alertType: "Price Alert - Low Price", priceTrigger: 1045},
}

// inAlertsActive - type to store the active alerts privately.  This encapsulated map (within the struct) is used to
// facilitate easier unit testing.  Because inAlertsActive implements the Notification interface, mocks can be injected more easily
// into the code
type inAlertsActive struct {
	data map[string]alert
}

// AddAlert - adds a new alert to the alerts active - i.e. the global alert store
func (a AlertsActiveType) AddAlert(item string, newAlertDef alert) {
	a[item] = newAlertDef
}

// AddAlert - adds a new alert to the alerts active - i.e. the private memory store used to facilitate easier testing
func (a inAlertsActive) AddAlert(item string, newAlertDef alert) {
	a.data[item] = newAlertDef
}
