package serializers

type RegisterRequest struct {
	Usernmae string `json:"username" form:"username" binding:"required,min=1,max=20"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
}
