package domain

type RegisterVendorNotify struct {
	Name          string
	Email         string
	ActivationURL string
}

type RegisterCustomerNotify struct {
	Name          string
	Email         string
	ActivationURL string
}
