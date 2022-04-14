package repository

type Database interface {
	SaveUrlMapping(shortUrl string, originalUrl string, userId string)
	RetrieveInitialUrl(shortUrl string) string
}

//Using Factory pattern here. It will provide us with flexibility to
// add various DB into our service.
func LookUpDatabase(dbName string) (Database, error) {
	switch dbName {
	case "redis":
		return InitializeRedisDB()
	default:
		return nil, &NotImplementedDatabaseError{}
	}
}
