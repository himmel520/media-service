package auth

type (
	Handler struct {
		uc AuthUsecase
	}

	AuthUsecase interface{}
)

func New(uc AuthUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
