package domain

import "time"

// Role definitions
const (
	RoleVendor   = 1
	RoleCustomer = 2
)

type User struct {
	ID           string
	Name         string
	Username     string
	Email        string
	PhoneNumber  string
	Role         int
	IsAdmin      bool
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Optional relationships
	Profile *UserProfile
	Vendor  *Vendor
}

// UserProfile stores customer or vendor profile information
type UserProfile struct {
	ID             uint64
	UserID         uint64
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

// Vendor represents a restaurant or food vendor
type Vendor struct {
	ID            uint64
	UserID        uint64
	VendorName    string
	VendorLicense string
	IsApproved    bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
