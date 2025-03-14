package entity

type TrackedApps struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AppName   string `json:"app_name"`
	Status    bool   `json:"status" default:"false"`
	Duration  int    `json:"duration"`
	EndTime   int64  `json:"end_time"`
	CreatedAt int64  `json:"created_at"`
}

type TrackedAppsRequest struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AppName   string `json:"app_name"`
	Status    bool   `json:"status" default:"false"`
	Duration  int    `json:"duration"`
	CreatedAt int64  `json:"created_at"`
}

type TrackedAppsResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AppName   string `json:"app_name"`
	Status    bool   `json:"status" default:"false"`
	CreatedAt int64  `json:"created_at"`
}

type TrackedAppsListResponse struct {
	Apps []TrackedApps `json:"apps"`
}
