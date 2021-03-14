package payload

type GetPersonResponses struct {
	ID        string  `json:"id"`
	UserName  string  `json:"user_name"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name, omitempty"`
	Age       uint64  `json:"age"`
}

type GetPersonsResponses struct {
	Users []GetPersonResponses
}

type CreatePersonRequests struct {
	UserName  string  `json:"user_name"`
	Password  string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Age       uint64  `json:"age"`
}

type CreatePersonResponses struct {
	Message string
}
