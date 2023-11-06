package domain

type Contact struct {
	ID              int64
	PrimaryPhone    string
	AdditionalPhone string
	Email           string
}

type ContactResponse struct {
	ID              int64
	PrimaryPhone    string
	AdditionalPhone string
	Email           string
}
