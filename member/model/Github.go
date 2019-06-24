package model


type GithubMember struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}


type GithubUser struct {
	Login             string      `json:"login"`
	AvatarURL         string      `json:"avatar_url"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
}

// UserByFollowers implements sort.Interface based on the Followers field.
type GithubUserByFollowers []GithubUser

func (u GithubUserByFollowers) Len() int           { return len(u) }
func (u GithubUserByFollowers) Less(i, j int) bool { return u[i].Followers > u[j].Followers }
func (u GithubUserByFollowers) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }


type GithubGeneralError struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}
