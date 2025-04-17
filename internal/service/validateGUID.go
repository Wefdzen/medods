package service

import "github.com/google/uuid"

func ValidateGuid(guid string) error {
	err := uuid.Validate(guid)
	if err != nil {
		return err
	}
	return nil
}
