package repository

import (
	"vtm-go-bot/model"
)

func AddMerit(name string, description string, kind string, levelsInfo string) map[string]string {

	meritToBeInserted := model.Merit{
		Name:        name,
		Description: description,
		Kind:        kind,
		LevelsInfo:  levelsInfo,
	}
	var status string
	err := InsertIntoTable(&meritToBeInserted)
	if err == nil {
		status = "Mérito adicionado com sucesso!"
	} else {
		status = "Erro ao adicionar mérito: " + err.Error()
	}
	return map[string]string{"status": status}
}

func GetMeritByID(id uint) model.Merit {
	var merit model.Merit
	GetByID(&merit, id)
	return merit
}

func GetMeritsByKind(kind string) []model.Merit {
	var merits []model.Merit
	GetByField(&merits, "kind", kind)
	return merits
}

func UpdateMerit(id uint, name string, description string, kind string, levelsInfo string) map[string]string {

	meritToBeUpdated := model.Merit{
		ID:          id,
		Name:        name,
		Description: description,
		Kind:        kind,
		LevelsInfo:  levelsInfo,
	}
	var status string
	err := UpdateTable(&meritToBeUpdated)
	if err == nil {
		status = "Mérito atualizado com sucesso!"
	} else {
		status = "Erro ao atualizar mérito: " + err.Error()
	}
	return map[string]string{"status": status}
}

func DeleteMerit(id uint) map[string]string {
	var merit model.Merit
	GetByID(&merit, id)
	var status string
	err := DeleteFromTable(&merit)
	if err == nil {
		status = "Mérito deletado com sucesso!"
	} else {
		status = "Erro ao deletar mérito: " + err.Error()
	}
	return map[string]string{"status": status}
}
