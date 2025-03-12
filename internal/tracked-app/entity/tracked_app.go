package entity

type TrackedApps struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	AppName   string `json:"app_name"`
	Status    bool   `json:"status" default:"false"`
	CreatedAt string `json:"created_at"`
}
type TrackedAppsRequest struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	AppName string `json:"app_name"`
	Status  bool   `json:"status" default:"false"`
}

type TrackedAppsResponse struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	AppName string `json:"app_name"`
	Status  bool   `json:"status" default:"false"`
}

type TrackedAppsListResponse struct {
	Apps []TrackedApps `json:"apps"`
}
