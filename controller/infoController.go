package controller

import (
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddInfo(s *discordgo.Session, interaction *discordgo.InteractionCreate) string {
	status := service.AddInfoService(interaction, "")["status"]
	return status
}

func GetInfoByID(id string) model.Info {
	info := service.GetInfoByID(id)
	return info
}

func GetAllInfos() []model.Info {
	infos := service.GetAllInfos()
	return infos
}

func UpdateInfo(s *discordgo.Session, interaction *discordgo.InteractionCreate, id string) string {
	status := service.AddInfoService(interaction, id)["status"]
	return status
}

func DeleteInfo(s *discordgo.Session, interaction *discordgo.InteractionCreate, idStr string) string {
	status := service.DeleteInfo(idStr)
	return status
}
