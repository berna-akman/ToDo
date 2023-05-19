package board

type Board struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Column      []string `json:"columns,default:['To Do','In Progress','In Test','Done']"`
}
