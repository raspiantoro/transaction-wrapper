package payload

type GetPersonResponses struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       uint64 `json:"age"`
}

type GetPersonsResponses struct {
	Users []GetPersonResponses
}
