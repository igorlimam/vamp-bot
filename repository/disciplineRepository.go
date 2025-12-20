package repository

import (
	"log"
	"vtm-go-bot/model"
)

func AddDiscipline(name string, dtype string, resonance string, threat string, description string) map[string]string {
	disciplineToBeInserted := model.Discipline{
		Name:        name,
		Dtype:       dtype,
		Resonance:   resonance,
		Threat:      threat,
		Description: description,
	}
	var status string
	err := InsertIntoTable(&disciplineToBeInserted)
	if err == nil {
		status = "Disciplina adicionada com sucesso!"
	} else {
		status = "Erro ao adicionar disciplina: " + err.Error()
	}
	log.Printf("Inserted Discipline: %s with ID: %d\n", name, disciplineToBeInserted.ID)
	return map[string]string{"status": status}
}

func UpdateDiscipline(id uint, name string, dtype string, resonance string, threat string, description string) map[string]string {
	disciplineToBeUpdated := model.Discipline{
		ID:          id,
		Name:        name,
		Dtype:       dtype,
		Resonance:   resonance,
		Threat:      threat,
		Description: description,
	}
	err := UpdateTable(&disciplineToBeUpdated)
	var status string
	if err == nil {
		status = "Disciplina atualizada com sucesso!"
	} else {
		status = "Erro ao atualizar disciplina: " + err.Error()
	}
	log.Printf("Updated Discipline ID %d: %s\n", id, name)
	return map[string]string{"status": status}
}

func GetAllDisciplines() []model.Discipline {
	var disciplines []model.Discipline
	GetAll(&disciplines)
	return disciplines
}

func GetDisciplineById(id uint) model.Discipline {
	var discipline model.Discipline
	GetByID(&discipline, id)
	return discipline
}

func GetDisciplineByName(name string) model.Discipline {
	var disciplines []model.Discipline
	GetByField(&disciplines, "name", name)
	return disciplines[0]
}

func DeleteDiscipline(id uint) map[string]string {
	var discipline model.Discipline
	GetByID(&discipline, id)
	if discipline.ID == 0 {
		log.Printf("Discipline with ID %d not found for deletion\n", id)
		return map[string]string{"status": "Disciplina não encontrada!"}
	}
	err := DeleteFromTable(&discipline)
	var status string
	if err == nil {
		status = "Disciplina deletada com sucesso!"
	} else {
		status = "Erro ao deletar disciplina: " + err.Error()
	}
	log.Printf("Deleted Discipline ID %d: %s\n", id, discipline.Name)
	return map[string]string{"status": status}
}
