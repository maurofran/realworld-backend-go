package profile

type Representation struct {
	Username  string `json:"username"`
	Biography string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}
