package typeform

// import (
// 	"encoding/json"
// )


type Headers struct {
	TotalItems 			int16				`json:"total_items"`
	PageCount 			int8				`json:"page_count"`
	Items				[]Item				`json:"items"`
}

type Item struct {
	LandingId			string					`json:"landing_id"`
	Token				string					`json:"token"`
	ResponseID			string					`json:"response_id"`
	LandedAt			string					`json:"landed_at"`
	SubmittedAt			string					`json:"submitted_at"`
	Metadatas			map[string]string		`json:"metadata"`
	Answers				string					`json:"answers"`
}

// type Metadata struct {
// 	UserAgent			string					`json:"user_agent"`
// 	Platform			string					`json:"platform"`
// 	Referer				string					`json:"referer"`
// 	NetworkId			string					`json:"network_id"`
// 	Browser				string					`json:"browser"`
// }
// {
// 	"field": {
// 		"id": "KHsG0b8vAu6E",
// 		"type": "email",
// 		"ref": "97f8e18ad06a02e6"
// 	},
// 	"type": "email",
// 	"email": "fed@me.com"
// },
// {
// 	"field": {
// 		"id": "qGTsWckbGLTv",
// 		"type": "multiple_choice",
// 		"ref": "9a9b61b5-415d-4113-aa3b-7bdd821761fc"
// 	},
// 	"type": "choice",
// 	"choice": {
// 		"label": "Yes"
// 	}
// },
// {
// 	"field": {
// 		"id": "d4go5hJkiDvD",
// 		"type": "multiple_choice",
// 		"ref": "f94cde7c-5b76-49fd-84f7-1dd707028982"
// 	},
// 	"type": "choices",
// 	"choices": {
// 		"labels": [
// 			"Personal Life Coaching",
// 			"Business and Entrepreneurial Coaching"
// 		]
// 	}
// },
// {
// 	"field": {
// 		"id": "py6CaN5XnhrW",
// 		"type": "multiple_choice",
// 		"ref": "6823ef5e-b31e-43a5-98fb-d2edf38b4672"
// 	},
// 	"type": "choices",
// 	"choices": {
// 		"labels": [
// 			"Group Coaching",
// 			"One-on-One Coaching"
// 		]
// 	}
// },
// {
// 	"field": {
// 		"id": "HINgdOcl1Yhz",
// 		"type": "multiple_choice",
// 		"ref": "9dbb23c8-02c2-40be-ac60-a6068dd1316d"
// 	},
// 	"type": "choices",
// 	"choices": {
// 		"labels": [
// 			"Virtual Coaching",
// 			"In-person Coaching"
// 		]
// 	}
// },
// {
// 	"field": {
// 		"id": "uljOpqXKl8d6",
// 		"type": "multiple_choice",
// 		"ref": "64b3462c-364e-46f1-9c00-36e42d61e50a"
// 	},
// 	"type": "choice",
// 	"choice": {
// 		"label": "Male"
// 	}
// },
// {
// 	"field": {
// 		"id": "BEralFyP9TlY",
// 		"type": "multiple_choice",
// 		"ref": "768b3578-909b-45ab-83d2-899bd1ec5be8"
// 	},
// 	"type": "choice",
// 	"choice": {
// 		"label": "Career/professional development"
// 	}
// },
// {
// 	"field": {
// 		"id": "TppF0REuGyX5",
// 		"type": "multiple_choice",
// 		"ref": "54ddd675-e656-43ea-8ae4-d6214961de9c"
// 	},
// 	"type": "choices",
// 	"choices": {
// 		"labels": [
// 			"Cynical",
// 			"Educated",
// 			"Independent",
// 			"Introverted"
// 		]
// 	}
// }