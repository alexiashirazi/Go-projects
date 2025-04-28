package model

type Product struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	CategoryID    string `json:"category_id"`
	DeviceType    string `json:"device_type"`
	Model         string `json:"model"`
	Color         string `json:"color"`
	Storage       string `json:"storage"`
	BatteryHealth string `json:"battery_health"`
	Processor     string `json:"processor"`
	Ram           string `json:"ram"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
}
