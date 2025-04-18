package db

type UserRepository interface {
	AddRecord(user *User) error
	GetRecord(guid string) (User, error)
	CheckUniqGuid(guid string) bool
	UpdateReftokenLiveTokenUnicCode(guid, refreshTokenHash, LiveToken, unicCode string)
}

func AddRecord(repo UserRepository, user *User) error {
	err := repo.AddRecord(user)
	return err
}

func GetRecord(repo UserRepository, guid string) (User, error) {
	return repo.GetRecord(guid)
}

func CheckUniqGuid(repo UserRepository, guid string) bool {
	return repo.CheckUniqGuid(guid)
}

func UpdateReftokenLiveTokenUnicCode(repo UserRepository, guid, refreshTokenHash, LiveToken, unicCode string) {
	repo.UpdateReftokenLiveTokenUnicCode(guid, refreshTokenHash, LiveToken, unicCode)
}
