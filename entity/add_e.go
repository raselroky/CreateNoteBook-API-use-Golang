package entity

type Add_e struct {
	Id      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Purpose string `json:"purpose" binding:"required"`
	//Id      int    `json:"id" binding:"required"`
}

type Naam struct {
	Id int `json:"id"`
}
