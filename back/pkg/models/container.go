package models

type ContainerCreateRequest struct {
	Image 	string 		`json:"image"`
	Name 		string 		`json:"name"`
	Ports 	[]string 	`json:"ports"`
}
