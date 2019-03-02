package user

type Representation struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	Username  string `json:"username"`
	Biography string `json:"bio"`
	Image     string `json:"image"`
}

type LoginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterCommand struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateCommand struct {
	Email     string `json:"email"`
	Biography string `json:"bio"`
	Image     string `json:"image"`
}

type FollowCommand struct {
	Username string `json:"username"`
}

type UnfollowCommand struct {
	Username string `json:"username"`
}
