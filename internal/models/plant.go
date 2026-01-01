package models
type Plant struct {
	ID int `json:"id"`
	Name string `json:"name"`
	RoomID int `json:"room_id"`
	WaterFreq int `json:"water_frequency"`
}