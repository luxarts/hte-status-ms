package domain

type Coordinates struct {
	Latitude  *float64 `json:"lat"`
	Longitude *float64 `json:"lon"`
}

type Status struct {
	DeviceID    string      `json:"device_id" bson:"device_id"`
	Timestamp   int64       `json:"ts" bson:"ts"`
	Battery     int64       `json:"bat" bson:"bat"`
	Coordinates Coordinates `json:"coords" bson:"coords"`
}

func (p *Status) IsValid() bool {
	return p.DeviceID != "" && p.Timestamp > 0 && p.Battery >= 0 && p.Coordinates.Latitude != nil && p.Coordinates.Longitude != nil
}
