package store

type Notification interface {
	Add(item string, newAlertDef alert)
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

// Add - add new alert
func (a AlertsActiveType) Add(item string, newAlertDef alert) {
	a[item] = newAlertDef
}
