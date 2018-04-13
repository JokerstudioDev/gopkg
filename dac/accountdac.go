package dac

type RealAccount struct{}

func (RealAccount) GetAccount() string {
	return "ของจริง"
}
