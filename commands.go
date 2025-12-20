package main

import (
	"log"
	"vtm-go-bot/controller"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(session *discordgo.Session) {
	commands := map[string]string{
		"ping":           "AM I ALIVE?",
		"add-disciplina": "Adiciona uma nova disciplina",
		"add-poder":      "Adiciona um novo poder a uma disciplina existente",
		"add-clan":       "Adiciona um novo clã",
		"add-merito":     "Adiciona um novo mérito",
		"add-info":       "Adiciona uma nova informação mecânica",
	}

	selectionCommands := map[string][]map[string]string{
		"disciplina": {
			{"name": "disciplina", "description": "Fornece informações sobre uma disciplina específica"},
		},
		"clan": {
			{"name": "clan", "description": "Fornece informações sobre um clã específico"},
		},
		"poder": {
			{"name": "disciplina", "description": "Disciplina do poder"},
			{"name": "poder", "description": "Fornece informações sobre um poder específico"},
		},
		"merito": {
			{"name": "tipo-merito", "description": "Tipo de mérito - Vantagem, Desvantagem, Antecedente"},
			{"name": "merito", "description": "Fornece informações sobre um mérito específico"},
		},
		"info": {
			{"name": "info", "description": "Fornece informações sobre uma mecânica específica"},
		},
		"update-merito": {
			{"name": "tipo-merito", "description": "Tipo de mérito - Vantagem, Desvantagem, Antecedente"},
			{"name": "merito", "description": "Merito a ser atualizado"},
		},
		"update-disciplina": {
			{"name": "disciplina", "description": "Disciplina a ser atualizada"},
		},
		"update-clan": {
			{"name": "clan", "description": "Clã a ser atualizado"},
		},
		"update-poder": {
			{"name": "disciplina", "description": "Disciplina do poder"},
			{"name": "poder", "description": "Poder a ser atualizado"},
		},
		"update-info": {
			{"name": "info", "description": "Informação a ser atualizada"},
		},
		"delete-disciplina": {
			{"name": "disciplina", "description": "Disciplina a ser deletada"},
		},
		"delete-clan": {
			{"name": "clan", "description": "Clã a ser deletado"},
		},
		"delete-poder": {
			{"name": "disciplina", "description": "Disciplina do poder"},
			{"name": "poder", "description": "Poder a ser deletado"},
		},
		"delete-merito": {
			{"name": "tipo-merito", "description": "Tipo de mérito - Vantagem, Desvantagem, Antecedente"},
			{"name": "merito", "description": "Mérito a ser deletado"},
		},
		"delete-info": {
			{"name": "info", "description": "Informação a ser deletada"},
		},
	}

	var all []*discordgo.ApplicationCommand

	for cmd, desc := range commands {
		all = append(all, &discordgo.ApplicationCommand{
			Name:        cmd,
			Description: desc,
		})
	}

	for cmd, cmdMap := range selectionCommands {
		var options []*discordgo.ApplicationCommandOption
		for _, params := range cmdMap {
			options = append(options, &discordgo.ApplicationCommandOption{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         params["name"],
				Description:  params["description"],
				Required:     true,
				Autocomplete: true,
			})
		}
		all = append(all, &discordgo.ApplicationCommand{
			Name:        cmd,
			Description: "Mostrar informações sobre " + cmd,
			Options:     options,
		})
	}

	_, err := session.ApplicationCommandBulkOverwrite(
		session.State.User.ID,
		session.State.Guilds[0].ID,
		all,
	)
	if err != nil {
		log.Fatalf("Cannot bulk overwrite commands: %v", err)
	}

	log.Println("Commands registered successfully.")
	controller.CheckDDLController()
}
