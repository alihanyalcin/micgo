package interfaces

type DBClient interface {
	CloseSession()

	InsertSomethingToTest() error
}
