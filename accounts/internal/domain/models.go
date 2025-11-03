package domain

import "time"

// Role definitions
const (
	RoleVendor   = 1
	RoleCustomer = 2
)

var AnonymousUser = User{}

type User struct {
	ID           string
	Name         string
	Username     string
	Email        string
	PhoneNumber  string
	Role         int
	IsAdmin      bool
	IsActive     bool
	LastLogin    time.Time
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Optional relationships
	Profile *UserProfile
	Vendor  *Vendor
}

// UserProfile stores customer or vendor profile information
type UserProfile struct {
	ID             string
	UserID         string
	ProfilePicture string
	CoverPhoto     string
	AddressLine1   string
	AddressLine2   string
	Country        string
	State          string
	City           string
	PinCode        string
	Longitude      string
	Latitude       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Vendor represents a restaurant only information
type Vendor struct {
	ID            string
	UserID        string
	VendorName    string
	VendorLicense string
	IsApproved    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
