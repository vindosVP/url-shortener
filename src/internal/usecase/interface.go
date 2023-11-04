package usecase

type Shortener interface {
	GetLinkPattern() string
	Save(originalUrl string) (string, error)
	GetOriginal(alias string) (string, error)
}

type UrlRepo interface {
	Save(originalUrl string, alias string) (string, error)
	AliasForURLExists(originalUrl string) (bool, error)
	GetOriginal(alias string) (string, error)
	GetAlias(originalUrl string) (string, error)
	AliasExists(alias string) (bool, error)
}
