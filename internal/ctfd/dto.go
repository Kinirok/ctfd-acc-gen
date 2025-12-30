package ctfd

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Success bool `json:"success"`
	Data    struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		CTFDUser string `json:"name"`
		CTFDPass string `json:"pass"`
		TeamID   int    `json:"team_id"`
	} `json:"data"`
	StatusCode int `json:"status_code"`
}

type CreateTeamRequest struct {
	TeamName string `json:"name"`
}

type CreateTeamResponse struct {
	Success bool `json:"success"`
	Data    struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	StatusCode int `json:"status_code"`
}

type TeamResponse struct {
	Success bool                 `json:"success"`
	Team    string               `json:"team"`
	Members []CreateUserResponse `json:"members"`
}

type ExistenceResponse struct {
	Success bool `json:"success"`
}
