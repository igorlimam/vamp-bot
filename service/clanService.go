package service

import (
	"log"
	"strings"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddClanService(interaction *discordgo.InteractionCreate, disciplineIDsSuffix string, isUpdate bool) map[string]string {
	dataModal := ModalToMap(interaction)
	disciplinesSplited := strings.Split(disciplineIDsSuffix, "-")
	var clanId uint
	var disciplinesVector []model.Discipline

	for i, idStr := range disciplinesSplited {
		id := ConvertStringToUint(idStr)
		if i == 0 && isUpdate {
			clanId = id
		} else {
			disciplinesVector = append(disciplinesVector, repository.GetDisciplineById(id))
		}
	}

	var status map[string]string
	if isUpdate {
		status = repository.UpdateClan(
			uint(clanId),
			dataModal["clan-name"].(string),
			dataModal["clan-description"].(string),
			dataModal["clan-bane"].(string),
			dataModal["clan-compulsion"].(string),
			disciplinesVector,
		)

		log.Printf("Updated Clan ID %d: %s\n", clanId, dataModal["clan-name"].(string))
	} else {
		status = repository.AddClan(
			dataModal["clan-name"].(string),
			dataModal["clan-description"].(string),
			dataModal["clan-bane"].(string),
			dataModal["clan-compulsion"].(string),
			disciplinesVector,
		)
	}

	return status
}

func GetAllClansService() []model.Clan {
	clans := repository.GetAllClans()
	return clans
}

func GetClanByIDService(idStr string) model.Clan {
	id := ConvertStringToUint(idStr)
	clan := repository.GetClanByID(id)
	return clan
}

func DeleteClanService(interaction *discordgo.InteractionCreate, clanID string) map[string]string {
	id := ConvertStringToUint(clanID)
	status := repository.DeleteClan(id)
	log.Printf("Deleted Clan ID %d\n", id)
	return status
}
