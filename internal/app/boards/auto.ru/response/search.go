package response

type Search struct {
	Reviews []review `json:"reviews"`
}

type review struct {
	Id        string    `json:"id"`
	Item      item      `json:"item"`
	Reviewer  reviewer  `json:"reviewer"`
	Published string    `json:"published"`
	Updated   string    `json:"updated"`
	Title     string    `json:"title"`
	Content   []content `json:"content"`
}

type item struct {
	Auto auto `json:"auto"`
}

type auto struct {
	Category   string `json:"category"`
	Mark       string `json:"mark"`
	Model      string `json:"model"`
	SuperGenId string `json:"super_gen"`
	Year       int64  `json:"year"`
}

type reviewer struct {
	Id               string    `json:"id"`
	Fn               string    `json:"fn"`
	Login            string    `json:"login"`
	SignPrivacy      string    `json:"sign_privacy"`
	AvatarUrl        avatarPic `json:"avatar_url"`
	RegistrationDate string    `json:"registration_date"`
}

type avatarPic struct {
	Pic24x24   string `json:"24x24"`
	Pic100x100 string `json:"100x100"`
	Pic430x600 string `json:"430x600"`
	Pic48x48   string `json:"48x48"`
	Pic200x200 string `json:"200x200"`
}

type content struct {
	Key         string         `json:"key"`
	Type        string         `json:"type"`
	Value       []contentValue `json:"content_value"`
	CreatedTime string         `json:"created_time"`
}

type contentValue struct {
	Value string `json:"value"`
	Image image  `json:"image"`
}

type image struct {
	Pic320x240  string `json:"320x240"`
	Pic640x480  string `json:"640x480"`
	Pic540x400  string `json:"540x400"`
	Pic1200x900 string `json:"1200x900"`
}
