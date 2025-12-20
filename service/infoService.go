package service

import (
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddInfoService(interaction *discordgo.InteractionCreate, id string) map[string]string {
	modalData := ModalToMap(interaction)

	name := modalData["info-name"].(string)
	kind := modalData["info-kind"].(string)
	info := modalData["info-info"].(string)

	if id == "" {
		return repository.AddInfo(name, kind, info)
	} else {
		idUint := ConvertStringToUint(id)
		return repository.UpdateInfo(idUint, name, kind, info)
	}
}

func GetInfoByID(id string) model.Info {
	idUint := ConvertStringToUint(id)
	info := repository.GetInfoByID(idUint)
	return info
}

func GetAllInfos() []model.Info {
	infos := repository.GetAllInfos()
	return infos
}

func DeleteInfo(idStr string) string {
	id := ConvertStringToUint(idStr)
	status := repository.DeleteInfo(id)
	return status["status"]
}
