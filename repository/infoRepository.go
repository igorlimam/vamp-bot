package repository

import "vtm-go-bot/model"

func AddInfo(name string, kind string, info string) map[string]string {
	infoToBeInserted := model.Info{
		Name: name,
		Kind: kind,
		Info: info,
	}
	var status string
	err := InsertIntoTable(&infoToBeInserted)
	if err == nil {
		status = "Informação adicionada com sucesso!"
	} else {
		status = "Erro ao adicionar informação: " + err.Error()
	}
	return map[string]string{"status": status}
}

func UpdateInfo(id uint, name string, kind string, info string) map[string]string {
	infoToBeUpdated := model.Info{
		ID:   id,
		Name: name,
		Kind: kind,
		Info: info,
	}
	var status string
	err := UpdateTable(&infoToBeUpdated)
	if err == nil {
		status = "Informação atualizada com sucesso!"
	} else {
		status = "Erro ao atualizar informação: " + err.Error()
	}
	return map[string]string{"status": status}
}

func GetInfoByID(id uint) model.Info {
	var info model.Info
	GetByID(&info, id)
	return info
}

func GetAllInfos() []model.Info {
	var infos []model.Info
	GetAll(&infos)
	return infos
}

func DeleteInfo(id uint) map[string]string {
	var info model.Info
	GetByID(&info, id)
	err := DeleteFromTable(&info)
	var status string
	if err == nil {
		status = "Informação deletada com sucesso!"
	} else {
		status = "Erro ao deletar informação: " + err.Error()
	}
	return map[string]string{"status": status}
}
