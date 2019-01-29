package mongo

// import()

type UserModel struct {
	Name 			string
	Email 			string
	UserType		string
	Results 		map[string]string   `json:"results"`
	Matches 		[]string
	Appts 			[]*Appt
}

type Appt struct {
	Date			string `json:"date"`
	Client			string `json:"client"`
}