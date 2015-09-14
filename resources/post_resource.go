package resources
import "time"

type PostResource struct {
	Id        string `json:"id"`
	CarPlate  string `json:"car_plate"`
	Message   string `json:"message"`
	UserName  string `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}
