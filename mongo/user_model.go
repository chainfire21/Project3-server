package mongo

// import()

type UserModel struct {
	Name 			string				`json:"name"`
	Email 			string				`json:"confirmedEmail"`
	UserType		string				`json:"user"`
	Results 		*Results  			`json:"results"`
	Matches 		[]string			`json:"matches"`
	Appts 			[]*Appt				`json:"appts"`
}

type Results struct {
	Location		string				`json:"location"`
	NewClients		string				`json:"new"`
	Types			[]string			`json:"types"`
	GorOne			[]string			`json:"gorone"`
	Virtual			[]string			`json:"virtual"`
	Gender			[]string			`json:"gender"`
	Topics			[]string			`json:"topics"`
	Traits			[]string			`json:"traits"`
}

type Appt struct {
	Date			string `json:"date"`
	Client			string `json:"client"`
}