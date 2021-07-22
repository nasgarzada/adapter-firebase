package model

type Notification struct {
	DeviceToken []string   `json:"deviceToken"`
	Title       string     `json:"title"`
	Body        string     `json:"body"`
	Data        *DataField `json:"data"`
}

type DataField struct {
	NotificationType string            `json:"notificationType"`
	ImageUrl         *string           `json:"imageUrl"`
	Payload          map[string]string `json:"payload"`
}
