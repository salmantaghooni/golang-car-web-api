package dto

type CreateColorRequest struct {
	Name    string `json:"name" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode" binding:"min=7,max=7"`
}

type UpdateColorRequest struct {
	Name    string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode,omitempty" binding:"min=7,max=7"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}
