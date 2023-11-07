package usecase

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Shortener
type Shortener interface {
	GetLinkPattern() string
	Save(originalUrl string) (string, error)
	GetOriginal(alias string) (string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=UrlRepo
type UrlRepo interface {
	Save(originalUrl string, alias string) (string, error)
	AliasForURLExists(originalUrl string) (bool, error)
	GetOriginal(alias string) (string, error)
	GetAlias(originalUrl string) (string, error)
	AliasExists(alias string) (bool, error)
}
