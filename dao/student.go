package dao

//Student - student dao
type Student struct {
	StudentID int      `json:"studentID,omitempty"`
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//Address - address dao
type Address struct {
	StreetAddress  string `json:"streetAddress,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	ZipCode        string `json:"zipCode,omitempty"`
	Country        string `json:"country,omitempty"`
}
