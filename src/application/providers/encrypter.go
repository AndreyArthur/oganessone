package providers

type EncrypterProvider interface {
	Hash(text string) (string, error)
	Compare(text string, hash string) (bool, error)
}
