package main

import (
	"math/rand"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var token = "" /* insert your bot token here         */
var tileamt = 3 /* amount of tiles u want to predict */

func predictor(grid []string) []string {
	for i := 0; i < tileamt; i++ {
		num := rand.Intn(25)
		grid[num] = "✅"
	}

	return grid
}

func predictindex() string {
	var x = []string{
		"❌", "❌", "❌", "❌", "❌",
		"❌", "❌", "❌", "❌", "❌",
		"❌", "❌", "❌", "❌", "❌",
		"❌", "❌", "❌", "❌", "❌",
		"❌", "❌", "❌", "❌", "❌",
	}

	grid := predictor(x)
	fixedgrid := fmt.Sprintf(`
%s%s%s%s%s
%s%s%s%s%s
%s%s%s%s%s
%s%s%s%s%s
%s%s%s%s%s`,
		grid[0], grid[1], grid[2], grid[3], grid[4],
		grid[5], grid[6], grid[7], grid[8], grid[9],
		grid[10], grid[11], grid[12], grid[13], grid[14],
		grid[15], grid[16], grid[17], grid[18], grid[19],
		grid[20], grid[21], grid[22], grid[23], grid[24],
	)

	return fixedgrid
}

func main() {
	sess, err := discordgo.New("Bot "+token); if err != nil { log.Fatal(err) }
	sess.Identify.Intents = discordgo.IntentsAll

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.Bot {
			return
		}

		if strings.HasPrefix(m.Content, "-") {
			if strings.HasPrefix(m.Content, "-predict ") {
				text := strings.ReplaceAll(m.Content, "-predict", "")
				prediction := predictindex()
				final := fmt.Sprintf("**Round ID**:%s\n**Tile Amount**: %d```\n%s```", text, tileamt, prediction)
				
				embed := &discordgo.MessageEmbed{
					Title: "github.com/latepatches",
					Description: final,
				}

				s.ChannelMessageSendEmbed(m.ChannelID, embed)
				return
			}

			s.ChannelMessageSend(m.ChannelID, "Command Not Found.\nTry -predict (round id)")
		}
	})

	sess.Open()
	select {}
}
