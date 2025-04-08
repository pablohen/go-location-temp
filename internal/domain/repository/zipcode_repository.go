package repository

type Location struct {
	City    string
	State   string
	Country string
	ZipCode string
}

type ZipCodeRepository interface {
	GetLocationByZipCode(zipCode string) (*Location, error)
}
