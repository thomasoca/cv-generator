package main

// User ...
type User struct {
	PersonalInfo struct {
		Name      string `json:"name"`
		Headline  string `json:"headline"`
		Picture   string `json:"picture"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Github    string `json:"github"`
		Linkedin  string `json:"linkedin"`
		Twitter   string `json:"twitter"`
		Location1 string `json:"location_1"`
		Location2 string `json:"location_2"`
	} `json:"personal_info"`
	MainSection struct {
		AboutMe struct {
			Label        string `json:"label"`
			Descriptions string `json:"descriptions"`
		} `json:"about_me"`
		WorkExperience struct {
			Label string `json:"label"`
			Lists []struct {
				Company      string   `json:"company"`
				Position     string   `json:"position"`
				StartPeriod  string   `json:"start_period"`
				EndPeriod    string   `json:"end_period"`
				Location     string   `json:"location"`
				Descriptions []string `json:"descriptions"`
			} `json:"lists"`
		} `json:"work_experience"`
		Education struct {
			Label string `json:"label"`
			Lists []struct {
				Institution  string   `json:"institution"`
				Major        string   `json:"major"`
				Level        string   `json:"level"`
				Gpa          string   `json:"gpa"`
				StartPeriod  string   `json:"start_period"`
				EndPeriod    string   `json:"end_period"`
				Location     string   `json:"location"`
				Descriptions []string `json:"descriptions"`
			} `json:"lists"`
		} `json:"education"`
		Skills struct {
			Label        string   `json:"label"`
			Descriptions []string `json:"descriptions"`
		} `json:"skills"`
		Languages struct {
			Label        string `json:"label"`
			Descriptions []struct {
				Language string `json:"language"`
				Fluency  int    `json:"fluency"`
			} `json:"descriptions"`
		} `json:"languages"`
	} `json:"main_section"`
}
