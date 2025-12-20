package service

import (
	"slices"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func GetMeritKindName(kindID string) string {
	switch kindID {
	case "1":
		return "Vantagem"
	case "2":
		return "Desvantagem"
	case "3":
		return "Antecedente"
	default:
		return "Desconhecido"
	}
}

func AddMeritService(interaction *discordgo.InteractionCreate, meritID string) map[string]string {
	dataModal := ModalToMap(interaction)

	meritKind := dataModal["merit-kind"].(string)
	kinds := []string{"Vantagem", "Desvantagem", "Antecedente"}
	if !slices.Contains(kinds, meritKind) {
		return map[string]string{"status": "MÉRITO NÃO CADASTRADO! Tipo de mérito inválido!"}
	}

	var status map[string]string
	if meritID == "" {
		status = repository.AddMerit(
			dataModal["merit-name"].(string),
			dataModal["merit-description"].(string),
			meritKind,
			dataModal["merit-levels-info"].(string),
		)
	} else {
		id := ConvertStringToUint(meritID)
		status = repository.UpdateMerit(
			id,
			dataModal["merit-name"].(string),
			dataModal["merit-description"].(string),
			meritKind,
			dataModal["merit-levels-info"].(string),
		)
	}

	return status
}

func GetMeritByID(meritID string) model.Merit {
	id := ConvertStringToUint(meritID)
	return repository.GetMeritByID(id)
}

func GetMeritsByKind(meritKind string) []model.Merit {
	return repository.GetMeritsByKind(meritKind)
}

func DeleteMeritService(interaction *discordgo.InteractionCreate, meritID string) map[string]string {
	id := ConvertStringToUint(meritID)
	status := repository.DeleteMerit(id)
	return status
}
