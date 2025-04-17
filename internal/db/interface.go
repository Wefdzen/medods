package db

type UserRepository interface {
	AddRecord(user *User) error
	GetRecord(guid string) (User, error)
}

func AddRecord(repo UserRepository, user User) error {
	err := repo.AddRecord(&user)
	return err
}

func GetRecord(repo UserRepository, guid string) (User, error) {
	return repo.GetRecord(guid)
}
