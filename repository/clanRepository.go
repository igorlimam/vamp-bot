package repository

import "vtm-go-bot/model"

func AddClan(clanName string, description string, bane string, compulsion string, desciplines []model.Discipline) map[string]string {
	clan := model.Clan{
		Name:        clanName,
		Description: description,
		Bane:        bane,
		Compulsion:  compulsion,
		Disciplines: desciplines,
	}
	var status string
	err := InsertIntoTable(&clan)
	if err == nil {
		status = "Clan added successfully"
	} else {
		status = "Error adding clan: " + err.Error()
	}

	return map[string]string{"status": status}
}

func UpdateClan(id uint, clanName string, description string, bane string, compulsion string, desciplines []model.Discipline) map[string]string {
	clan := model.Clan{
		ID:          id,
		Name:        clanName,
		Description: description,
		Bane:        bane,
		Compulsion:  compulsion,
		Disciplines: desciplines,
	}
	err := UpdateTable(&clan)
	DB.Model(&clan).Association("Disciplines").Clear()
	if len(desciplines) > 0 {
		DB.Model(&clan).Association("Disciplines").Append(desciplines)
	}
	if err == nil {
		return map[string]string{"status": "Clan updated successfully"}
	} else {
		return map[string]string{"status": "Error updating clan: " + err.Error()}
	}
}

func GetAllClans() []model.Clan {
	var clans []model.Clan
	GetAll(&clans)
	return clans
}

func GetClanByID(id uint) model.Clan {
	var clan model.Clan
	GetByID(&clan, id)
	return clan
}

func DeleteClan(id uint) map[string]string {
	var clan model.Clan
	GetByID(&clan, id)

	err := DeleteFromTable(&clan)
	var status string
	if err == nil {
		status = "Clan deleted successfully"
	} else {
		status = "Error deleting clan: " + err.Error()
	}
	return map[string]string{"status": status}
}
