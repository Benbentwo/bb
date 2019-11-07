package github

const (
	GitAuthConfigFile = "gitAuth.yaml"
)
type AuthServer struct {
	URL   string      `json:"url"`
	Users []*UserAuth `json:"users"`
	Name  string      `json:"name"`
	Kind  string      `json:"kind"`

	CurrentUser string `json:"currentuser"`
}

type UserAuth struct {
	Username    string `json:"username"`
	ApiToken    string `json:"apitoken"`
	BearerToken string `json:"bearertoken"`
	Password    string `json:"password,omitempty"`
}

type AuthConfig struct {
	Servers []*AuthServer `json:"servers"`

	DefaultUsername  string `json:"defaultusername"`
	CurrentServer    string `json:"currentserver"`
}