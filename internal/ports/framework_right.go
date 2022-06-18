package ports

type DBPort interface {
	CloseDbConnection()
	AddToHistory(anser int32, opetaion string) error
}
