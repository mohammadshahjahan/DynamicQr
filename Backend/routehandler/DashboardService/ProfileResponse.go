package dashboardservice

type ProfileResponse struct {
	Username string                   `json:"username"`
	Email    string                   `json:"email"`
	Name     string                   `json:"name"`
	Count    int32                    `json:"count"`
	QRs      []map[string]interface{} `json:"qrs"`
}
