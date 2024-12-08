package auth

type (
	Handler struct {
		uc AuthUsecase
	}

	AuthUsecase interface{
		GetUserRoleFromToken(jwtToken string) (int, error) 
		IsUserAdmin(userRole int) bool
	}
)

func New(uc AuthUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
