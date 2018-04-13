package dac

type AccountDac interface {
	GetAccount() string
}

//Service
func GetAccountInformation(acc AccountDac) string {
	return acc.GetAccount() + " studio"
}
