package serializers

type User struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=20"`
	Password string `json:"password" form:"password" binding:"required,min=4"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}
