package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateUserResponse struct {
	*ApiResponse
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
}

type DeleteUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}

type UpdateUserResponse struct {
	*ApiResponse
}

type UpdateUserRequest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updatedAge" binding:"required"`
}
