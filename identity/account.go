package identity

type Account struct {
	ID         string
	Addr       Address
	privFields *RawAddress
}
