package service

import (
	"vtm-go-bot/model"
	"vtm-go-bot/repository"
)

func GetClanDisciplinesByIdService(clanId string) []model.Discipline {
	id := ConvertStringToUint(clanId)
	disciplines := repository.GetClanDisciplinesByIdRepository(id)
	return disciplines
}
