package model

type Notification struct {
	DeviceToken []string   `json:"deviceToken"`
	Title       string     `json:"title"`
	Body        string     `json:"body"`
	Data        *DataField `json:"data"`
}

type DataField struct {
	NotificationType NotificationType  `json:"notificationType"`
	ImageUrl         *string           `json:"imageUrl"`
	Payload          map[string]string `json:"payload"`
}

type NotificationType string

const (
	Payment  NotificationType = "PAYMENT"
	Accounts NotificationType = "ACCOUNTS"
	Holiday  NotificationType = "HOLIDAY"
)
