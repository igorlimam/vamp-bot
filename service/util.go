package service

import (
	"log"
	"strconv"
	"strings"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func ModalToMap(interaction *discordgo.InteractionCreate) map[string]interface{} {
	dataModal := interaction.ModalSubmitData().Components
	values := map[string]interface{}{}

	for _, row := range dataModal {
		if row, ok := row.(*discordgo.ActionsRow); ok {
			for _, component := range row.Components {
				if input, ok := component.(*discordgo.TextInput); ok {
					values[input.CustomID] = input.Value
				}
			}
		}
	}

	return values
}

func ConvertStringToInt(levelStr string) int {
	level, err := strconv.Atoi(levelStr)
	if err != nil {
		log.Printf("Failed to convert string to int: %v\n", err)
		return -1
	}
	return level
}

func ConvertStringToUint(idStr string) uint {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Failed to convert string to uint: %v\n", err)
		return 0
	}
	return uint(id)
}

func SplitModalInput(input string, separator string, expected int) []string {
	arr := strings.Split(input, separator)
	if len(arr) != expected {
		return []string{}
	}
	return arr
}

func CheckDDLService() {
	repository.CheckDDL()
}
