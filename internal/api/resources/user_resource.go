package resources

type UserResource struct {
	Id int `json:"id"`
}

func NewUserResource(id int) *UserResource {
	return &UserResource{Id: id}
}
