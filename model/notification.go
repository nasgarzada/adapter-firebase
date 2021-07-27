package model

type Notification struct {
	DeviceToken []string   `json:"deviceToken"`
	Title       string     `json:"title"`
	Body        string     `json:"body"`
	Data        *DataField `json:"data"`
}

type DataField struct {
	NotificationGroup NotificationGroup `json:"notificationGroup"`
	NotificationType  NotificationType  `json:"notificationType"`
	ImageUrl          *string           `json:"imageUrl"`
	Payload           map[string]string `json:"payload"`
}

type NotificationGroup string
type NotificationType string

const (
	Broadcasting NotificationGroup = "BROADCASTING"
	ByUser       NotificationGroup = "BY_USER"
	ByCustomer   NotificationGroup = "BY_CUSTOMER"
)

const (
	Payment  NotificationType = "PAYMENT"
	Accounts NotificationType = "ACCOUNTS"
	Holiday  NotificationType = "HOLIDAY"
)
