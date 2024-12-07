package auth

type (
	Handler struct {
		uc AuthUsecase
	}

	AuthUsecase interface{
		GetUserRoleFromToken(jwtToken string) (string, error) 
		IsUserAdmin(userRole string) bool
	}
)

func New(uc AuthUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
