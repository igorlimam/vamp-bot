package service

import (
	"log"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineService(interaction *discordgo.InteractionCreate, idStr string) map[string]string {
	dataModal := ModalToMap(interaction)

	var status map[string]string

	if idStr != "" {
		id := ConvertStringToUint(idStr)
		status = repository.UpdateDiscipline(
			id,
			dataModal["discipline-name"].(string),
			dataModal["discipline-type"].(string),
			dataModal["discipline-resonance"].(string),
			dataModal["discipline-threat"].(string),
			dataModal["discipline-description"].(string),
		)

		log.Printf("Updated Discipline ID %d: %s\n", id, dataModal["discipline-name"].(string))
	} else {
		status = repository.AddDiscipline(
			dataModal["discipline-name"].(string),
			dataModal["discipline-type"].(string),
			dataModal["discipline-resonance"].(string),
			dataModal["discipline-threat"].(string),
			dataModal["discipline-description"].(string),
		)

		log.Printf("Inserted Discipline: %s\n", dataModal["discipline-name"].(string))
	}

	return status
}

func GetDisciplineByID(idStr string) model.Discipline {
	id := ConvertStringToUint(idStr)
	discipline := repository.GetDisciplineById(id)
	return discipline
}

func GetAllDisciplines() []model.Discipline {
	disciplines := repository.GetAllDisciplines()
	return disciplines
}

func DeleteDiscipline(idStr string) string {
	id := ConvertStringToUint(idStr)
	status := repository.DeleteDiscipline(id)
	return status["status"]
}
