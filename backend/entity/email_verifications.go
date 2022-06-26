package entity

type EmailVerification struct {
	Id        int
	Email     string
	Signature string
	Expired   int
}
