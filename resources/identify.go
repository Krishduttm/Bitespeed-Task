package resources

type GetContactDetails struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type ContactDetailsOutput struct {
	PrimaryContactID    int      `json:"primary_contact_id"`
	Emails              []string `json:"emails"`
	PhoneNumbers        []string `json:"phone_numbers"`
	SecondaryContactIds []int    `json:"secondary_contact_ids"`
}
