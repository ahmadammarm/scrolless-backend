package entity

type Challenge struct {
	Title       string `json:"title" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Difficulty  string `json:"difficulty" validate:"required"`
	Points      int    `json:"points" validate:"required"`
}

type ChallengeResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Difficulty  string `json:"difficulty"`
	Points      int    `json:"points"`
}

type ChallengeListResponse struct {
	Challenges []ChallengeResponse `json:"challenges"`
}
