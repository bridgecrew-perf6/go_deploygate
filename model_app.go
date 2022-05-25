package deploygate

// Request

type AppMemberConfig struct {
	Owner    string
	Platform string
	AppId    string
	Users    string
	Role     string
	File     string
}

type AppTeamConfig struct {
	Organizations string
	Platform      string
	AppId         string
}

type AppConfig struct {
	Owner    string
	Platform string
	AppId    string
	File     string
}

// Response

type Member struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type GetAppMemberResponse struct {
	Error   bool                        `mapstructure:"error"`
	Results *GetAppMemberResponseResult `mapstructure:"results"`
}

type GetAppMemberResponseResult struct {
	Usage *Usage    `mapstructure:"usage"`
	Users []*Member `mapstructure:"users"`
}

type Usage struct {
	Used uint `mapstructure:"used"`
	Max  uint `mapstructure:"max"`
}

type AddAppMemberResponse struct {
	Error   bool                        `mapstructure:"error"`
	Message string                      `mapstructure:"message"`
	Because string                      `mapstructure:"because"`
	Results *AddAppMemberResponseResult `mapstructure:"results"`
}

type AddAppMemberResponseResult struct {
	Invite  string     `mapstructure:"invite"`
	Added   []*Added   `mapstructure:"added"`
	Invited []*Invited `mapstructure:"invited"`
}

type Added struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type Invited struct {
	Name string `mapstructure:"name"`
	Role uint   `mapstructure:"role"`
}

type DeleteAppMemberResponse struct {
	Error   bool                           `mapstructure:"error"`
	Message string                         `mapstructure:"message"`
	Because string                         `mapstructure:"because"`
	Results *DeleteAppMemberResponseResult `mapstructure:"results"`
}

type DeleteAppMemberResponseResult struct {
	Invite string `mapstructure:"invite"`
}

type UploadAppResponse struct {
	Error   bool   `mapstructure:"error"`
	Message string `mapstructure:"message"`
	Because string `mapstructure:"because"`
}

type GetAppTeamsResponse struct {
	Error bool     `mapstructure:"error"`
	Teams []*Teams `mapstructure:"teams"`
}

type Teams struct {
	Name        string `mapstructure:"name"`
	Role        string `mapstructure:"role"`
	MemberCount uint   `mapstructure:"member_count"`
}
