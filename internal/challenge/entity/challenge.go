package entity

type ChallengeStatus string

const (
	StatusSudah ChallengeStatus = "sudah"
	StatusBelum ChallengeStatus = "belum"
)

type Challenge struct {
	ID          int             `json:"id"`
	Title       string          `json:"title" validate:"required"`
	UserID      int             `json:"user_id" validate:"required"`
	Description string          `json:"description" validate:"required"`
	Status      ChallengeStatus `json:"status" validate:"oneof=sudah belum" default:"belum"`
	Points      int             `json:"points" default:"0"`
}

type ChallengeResponse struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Status      ChallengeStatus `json:"status"`
	Points      int             `json:"points"`
}

type ChallengeListResponse struct {
	Challenges []ChallengeResponse `json:"challenges"`
}
