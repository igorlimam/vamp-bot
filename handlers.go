package main

import (
	"fmt"
	"log"
	"strings"
	"vtm-go-bot/controller"
	"vtm-go-bot/view"

	"github.com/bwmarrin/discordgo"
)

func checkGuildOwner(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var guild, _ = s.State.Guild(interaction.GuildID)
	log.Printf("Comparing Guild Owner ID: %s, with Interaction User ID: %s", guild.OwnerID, interaction.Member.User.ID)
	if interaction.Member.User.ID != guild.OwnerID {
		view.ResolveResponse(s, interaction, "Apenas o dono do servidor pode usar este comando.")
	}
}

func getFocusedOption(interaction *discordgo.InteractionCreate) string {
	for _, option := range interaction.ApplicationCommandData().Options {
		if option.Focused {
			return option.StringValue()
		}
	}
	return ""
}

func getQueryOption(interaction *discordgo.InteractionCreate, optionName string) string {
	for _, option := range interaction.ApplicationCommandData().Options {
		if option.Name == optionName {
			return option.StringValue()
		}
	}
	return ""
}

func RegisterHandlers(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

	if interaction.Type == discordgo.InteractionModalSubmit {
		modalHandlers(interaction, s)
		return
	}

	if interaction.Type == discordgo.InteractionMessageComponent {
		messageComponentHandlers(interaction, s)
		return
	}

	if interaction.Type == discordgo.InteractionApplicationCommandAutocomplete {
		autoCompleteHandlers(interaction, s)
		return
	}

	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	appData := interaction.ApplicationCommandData()
	switch appData.Name {
	case "ping":
		s.InteractionRespond(
			interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong",
				},
			},
		)
	case "add-disciplina":
		checkGuildOwner(s, interaction)
		view.AddDisciplineView(s, interaction, nil)
	case "add-poder":
		checkGuildOwner(s, interaction)
		view.PowerSelectDisciplineView(s, interaction, controller.GetAllDisciplines())
	case "add-clan":
		checkGuildOwner(s, interaction)
		view.StringSelectClanDisciplines(s, interaction, controller.GetAllDisciplines(), nil, "")
	case "add-merito":
		checkGuildOwner(s, interaction)
		view.StringSelectMeritKindView(s, interaction)
	case "add-info":
		checkGuildOwner(s, interaction)
		view.AddInfoView(s, interaction, nil)
	case "merito":
		meritoID := appData.Options[1].StringValue()
		meritoID = controller.GetMeritKindName(meritoID)
		view.ShowMeritInfoView(s, interaction, controller.GetMeritByID(meritoID))
	case "disciplina":
		disciplinaID := appData.Options[0].StringValue()
		view.ShowDisciplineInfoView(s, interaction, controller.GetDisciplineByID(disciplinaID))
	case "clan":
		clanID := appData.Options[0].StringValue()
		disciplines := controller.GetClanDisciplinesById(clanID)
		view.ShowClanInfoView(s, interaction, controller.GetClanByID(clanID), disciplines)
	case "poder":
		power := controller.GetPowerById(appData.Options[1].StringValue())
		view.ShowPowerInfoView(s, interaction, power)
	case "info":
		infoID := appData.Options[0].StringValue()
		view.ShowInfoView(s, interaction, controller.GetInfoByID(infoID))
	case "update-merito":
		checkGuildOwner(s, interaction)
		meritID := appData.Options[1].StringValue()
		merit := controller.GetMeritByID(meritID)
		view.AddMeritView(s, interaction, "", &merit)
	case "update-disciplina":
		checkGuildOwner(s, interaction)
		discipline := controller.GetDisciplineByID(appData.Options[0].StringValue())
		view.AddDisciplineView(s, interaction, &discipline)
	case "update-clan":
		checkGuildOwner(s, interaction)
		clanID := appData.Options[0].StringValue()
		selectedDisciplines := controller.GetClanDisciplinesById(clanID)
		disciplines := controller.GetAllDisciplines()
		view.StringSelectClanDisciplines(s, interaction, disciplines, selectedDisciplines, clanID)
	case "update-poder":
		checkGuildOwner(s, interaction)
		disciplineID := appData.Options[0].StringValue()
		powerID := appData.Options[1].StringValue()
		power := controller.GetPowerById(powerID)
		view.AddPowerView(s, interaction, disciplineID, &power)
	case "update-info":
		checkGuildOwner(s, interaction)
		infoID := appData.Options[0].StringValue()
		info := controller.GetInfoByID(infoID)
		view.AddInfoView(s, interaction, &info)
	case "delete-disciplina":
		checkGuildOwner(s, interaction)
		disciplineID := appData.Options[0].StringValue()
		view.ConfirmDeleteDiscipline(s, interaction, controller.GetDisciplineByID(disciplineID))
	case "delete-clan":
		checkGuildOwner(s, interaction)
		clanID := appData.Options[0].StringValue()
		view.ConfirmDeleteClan(s, interaction, controller.GetClanByID(clanID))
	case "delete-poder":
		checkGuildOwner(s, interaction)
		disciplineID := appData.Options[0].StringValue()
		discipline := controller.GetDisciplineByID(disciplineID)
		powerID := appData.Options[1].StringValue()
		power := controller.GetPowerById(powerID)
		view.ConfirmDeletePower(s, interaction, power, discipline.Name)
	case "delete-merito":
		checkGuildOwner(s, interaction)
		meritKindID := appData.Options[0].StringValue()
		meritKind := controller.GetMeritKindName(meritKindID)
		meritID := appData.Options[1].StringValue()
		merit := controller.GetMeritByID(meritID)
		view.ConfirmDeleteMerit(s, interaction, meritKind, merit)
	case "delete-info":
		checkGuildOwner(s, interaction)
		infoID := appData.Options[0].StringValue()
		info := controller.GetInfoByID(infoID)
		view.ConfirmDeleteInfo(s, interaction, info)
	default:
		status := fmt.Sprintf("Comando %s não reconhecido.", appData.Name)
		view.ResolveResponse(s, interaction, status)
	}

}

func modalHandlers(interaction *discordgo.InteractionCreate, s *discordgo.Session) {
	modalData := interaction.ModalSubmitData()
	customID := strings.Split(modalData.CustomID, "|")[0]
	customData := strings.Split(modalData.CustomID, "|")[1]
	log.Printf("CUSTOM ID: %s", customID)
	switch customID {
	case "add-discipline-modal":
		status := controller.AddDiscipline(s, interaction)
		view.ResolveResponse(s, interaction, status)
	case "update-discipline-modal":
		status := controller.UpdateDiscipline(s, interaction, customData)
		view.ResolveResponse(s, interaction, status)
	case "add-power-modal":
		status := controller.AddPower(s, interaction, customData)
		view.ResolveResponse(s, interaction, status)
	case "add-clan-modal":
		status := controller.AddClan(s, interaction, customData)
		view.ResolveResponse(s, interaction, status)
	case "add-merit-modal":
		status := controller.AddMerit(s, interaction)
		view.ResolveResponse(s, interaction, status)
	case "add-info-modal":
		status := controller.AddInfo(s, interaction)
		view.ResolveResponse(s, interaction, status)
	case "update-clan-modal":
		status := controller.UpdateClan(s, interaction, customData)
		view.ResolveResponse(s, interaction, status)
	case "update-power-modal":
		powerID := strings.Split(modalData.CustomID, "|")[1]
		disciplineID := strings.Split(modalData.CustomID, "|")[2]
		status := controller.UpdatePower(s, interaction, powerID, disciplineID)
		view.ResolveResponse(s, interaction, status)
	case "update-merit-modal":
		status := controller.UpdateMerit(s, interaction, customData)
		view.ResolveResponse(s, interaction, status)
	case "update-info-modal":
		status := controller.UpdateInfo(s, interaction, customData)
		view.ResolveResponse(s, interaction, status)
	default:
		status := "Interação EM MODAL Cancelada"
		view.ResolveResponse(s, interaction, status)
	}
}

func messageComponentHandlers(interaction *discordgo.InteractionCreate, s *discordgo.Session) {
	data := interaction.MessageComponentData()
	customID := strings.Split(data.CustomID, "|")[0]
	switch customID {
	case "select-discipline-for-power":
		view.AddPowerView(s, interaction, data.Values[0], nil)
	case "select-disciplines-for-clan":
		clanID := strings.Split(data.CustomID, "|")[1]

		if clanID != "0" {
			clan := controller.GetClanByID(clanID)
			view.AddClanView(s, interaction, data.Values, &clan)
		} else {
			view.AddClanView(s, interaction, data.Values, nil)
		}

		log.Printf("Selected disciplines for clan: %v", data.Values)
	case "select-merit-kind":
		meritKind := data.Values[0]
		meritKind = controller.GetMeritKindName(meritKind)
		view.AddMeritView(s, interaction, meritKind, nil)
	case "confirm-delete-discipline":
		disciplineID := strings.Split(data.CustomID, "|")[1]
		status := controller.DeleteDiscipline(s, interaction, disciplineID)
		view.ResolveResponse(s, interaction, status)
	case "confirm-delete-clan":
		clanID := strings.Split(data.CustomID, "|")[1]
		status := controller.DeleteClan(s, interaction, clanID)
		view.ResolveResponse(s, interaction, status)
	case "confirm-delete-power":
		powerID := strings.Split(data.CustomID, "|")[1]
		status := controller.DeletePower(s, interaction, powerID)
		view.ResolveResponse(s, interaction, status)
	case "confirm-delete-merit":
		meritID := strings.Split(data.CustomID, "|")[1]
		status := controller.DeleteMerit(s, interaction, meritID)
		view.ResolveResponse(s, interaction, status)
	case "confirm-delete-info":
		infoID := strings.Split(data.CustomID, "|")[1]
		status := controller.DeleteInfo(s, interaction, infoID)
		view.ResolveResponse(s, interaction, status)
	default:
		status := "Interação cancelada"
		view.ResolveResponse(s, interaction, status)
	}
}

func autoCompleteHandlers(interaction *discordgo.InteractionCreate, s *discordgo.Session) {
	for _, opt := range interaction.ApplicationCommandData().Options {
		if opt.Focused && opt.Name == "disciplina" {
			log.Printf("Focused option for disciplina: %s", opt.StringValue())
			view.DisciplinaInfoView(s, interaction, controller.GetAllDisciplines())
			return
		}
		if opt.Focused && opt.Name == "clan" {
			view.ClanInfoView(s, interaction, controller.GetAllClans())
			return
		}
		if opt.Focused && opt.Name == "poder" {
			disciplineID := getQueryOption(interaction, "disciplina")
			query := getFocusedOption(interaction)
			log.Printf("Autocomplete query for power: %s", disciplineID)
			disciplinePowers := controller.GetDisciplinePowersByID(disciplineID)
			view.PowerInfoView(s, interaction, query, disciplinePowers)
			return
		}
		if opt.Focused && opt.Name == "tipo-merito" {
			view.MeritKindInfoView(s, interaction)
			return
		}
		if opt.Focused && opt.Name == "merito" {
			chosenKind := getQueryOption(interaction, "tipo-merito")
			query := getFocusedOption(interaction)
			chosenKind = controller.GetMeritKindName(chosenKind)
			view.MeritInfoView(s, interaction, query, controller.GetMeritsByKind(chosenKind))
			return
		}
		if opt.Focused && opt.Name == "info" {
			query := getFocusedOption(interaction)
			view.InfoAutocompleteView(s, interaction, query, controller.GetAllInfos())
			return
		}
	}
}
