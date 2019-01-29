package mongo

// import()

type UserModel struct {
	Name 			string				`json:"name"`
	Password		string				`json:"pword"`
	Email 			string				`json:"confirmedEmail"`
	UserType		string				`json:"userType"`
	Results 		map[string]string   `json:"results"`
	Matches 		[]string			`json:"matches"`
	Appts 			[]*Appt				`json:"appts"`
}

type Appt struct {
	Date			string `json:"date"`
	Client			string `json:"client"`
}