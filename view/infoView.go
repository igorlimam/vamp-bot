package view

import (
	"fmt"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func AddInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, infoData *model.Info) {
	customID := "add-info-modal|0"
	title := "Adicionar Nova Informação Mecânica"

	nameValue := ""
	kindValue := ""
	infoValue := ""

	if infoData != nil {
		customID = "update-info-modal|" + fmt.Sprintf("%d", infoData.ID)
		title = "Atualizar Informação Mecânica"
		nameValue = infoData.Name
		kindValue = infoData.Kind
		infoValue = infoData.Info
	}

	fields := []map[string]string{
		{
			"customID": "info-name",
			"label":    "Nome da Informação",
			"style":    "short",
			"value":    nameValue,
		},
		{
			"customID": "info-kind",
			"label":    "Tipo da Informação",
			"style":    "short",
			"value":    kindValue,
		},
		{
			"customID": "info-info",
			"label":    "Descrição da Informação",
			"style":    "paragraph",
			"value":    infoValue,
		},
	}

	Modal(s, interaction, customID, title, fields)
}

func InfoAutocompleteView(s *discordgo.Session, interaction *discordgo.InteractionCreate, query string, infos []model.Info) {
	options := []map[string]string{}
	for _, info := range infos {
		if (query == "" || strings.Contains(strings.ToLower(info.Name), strings.ToLower(query))) && len(options) < 25 {
			options = append(options, map[string]string{
				"label": info.Name,
				"value": fmt.Sprintf("%d", info.ID),
			})
		}
	}

	AutoComplete(s, interaction, options)
}

func ShowInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, info model.Info) {
	title := info.Name
	description := info.Kind
	embedFields := []map[string]string{}

	embedFields = append(embedFields, map[string]string{
		"name":   "Info",
		"value":  info.Info,
		"inline": "false",
	})

	EmbedMessage(s, interaction, embedFields, title, description)
}

func ConfirmDeleteInfo(s *discordgo.Session, interaction *discordgo.InteractionCreate, info model.Info) {
	contentMessage := fmt.Sprintf("Você tem certeza que deseja deletar a informação mecânica **%s**?", info.Name)
	customIDConfirmation := "confirm-delete-info|" + fmt.Sprintf("%d", info.ID)
	customIDCancel := "cancel-delete-info"
	ConfirmationButton(s, interaction, customIDConfirmation, customIDCancel, contentMessage)
}
