package repository

import "vtm-go-bot/model"

func AddPower(disciplineId uint, name string, description string,
	dicePool string, cost string, duration string, system string,
	ptype string, amalgam string, level int) map[string]string {

	powerToBeInserted := model.Power{
		DisciplineID: disciplineId,
		Name:         name,
		Description:  description,
		DicePool:     dicePool,
		Cost:         cost,
		Duration:     duration,
		System:       system,
		Kind:         ptype,
		Amalgam:      amalgam,
		Level:        level,
	}
	var status string
	err := InsertIntoTable(&powerToBeInserted)
	if err == nil {
		status = "Poder adicionado com sucesso!"
	} else {
		status = "Erro ao adicionar poder: " + err.Error()
	}
	return map[string]string{"status": status}
}

func UpdatePower(id uint, disciplineId uint, name string, description string,
	dicePool string, cost string, duration string, system string,
	ptype string, amalgam string, level int) map[string]string {

	powerToBeUpdated := model.Power{
		ID:           id,
		DisciplineID: disciplineId,
		Name:         name,
		Description:  description,
		DicePool:     dicePool,
		Cost:         cost,
		Duration:     duration,
		System:       system,
		Kind:         ptype,
		Amalgam:      amalgam,
		Level:        level,
	}
	var status string
	err := UpdateTable(&powerToBeUpdated)
	if err == nil {
		status = "Poder atualizado com sucesso!"
	} else {
		status = "Erro ao atualizar poder: " + err.Error()
	}
	return map[string]string{"status": status}
}

func GetAllPowers() []model.Power {
	var powers []model.Power
	GetAll(&powers)
	return powers
}

func GetPowersByDiciplineId(disciplineId uint) []model.Power {
	var powers []model.Power
	GetByField(&powers, "discipline_id", disciplineId)
	return powers
}

func GetPowerById(id uint) model.Power {
	var power model.Power
	GetByID(&power, id)
	return power
}

func DeletePower(id uint) map[string]string {
	var power model.Power
	GetByID(&power, id)
	var status string
	err := DeleteFromTable(&power)
	if err == nil {
		status = "Poder deletado com sucesso!"
	} else {
		status = "Erro ao deletar poder: " + err.Error()
	}
	return map[string]string{"status": status}
}
