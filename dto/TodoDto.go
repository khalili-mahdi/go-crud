package dto

type TodoDto struct {
	Title string `form:"Title" json:"Title" xml:"Title"  binding:"required"`
	Body  string `form:"Body" json:"Body" xml:"Body"  binding:"required"`
}
