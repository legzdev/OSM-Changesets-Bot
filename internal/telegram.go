package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/legzdev/OSM-Changesets-Bot/env"
	"github.com/legzdev/OSM-Changesets-Bot/types"
)

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type Message struct {
	ChatID                int64                `json:"chat_id"`
	Text                  string               `json:"text"`
	ParseMode             string               `json:"parse_mode"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`
	DisableWebPagePreview bool                 `json:"disable_web_page_preview"`
}

func SendToTelegram(changeset types.Changeset) error {
	var builder strings.Builder
	builder.WriteString("Changeset ")

	changesetURL := fmt.Sprintf("https://openstreetmap.org/changeset/%d", changeset.ID)
	changesetFmt := fmt.Sprintf("<a href=\"%s\">%d</a>", changesetURL, changeset.ID)
	builder.WriteString(changesetFmt)

	builder.WriteString(" by ")

	userURL := fmt.Sprintf("https://openstreetmap.org/user/%s", url.QueryEscape(changeset.Username))
	userFmt := fmt.Sprintf("<a href=\"%s\">%s</a>", userURL, changeset.Username)
	builder.WriteString(userFmt)

	builder.WriteString("\n\n")
	builder.WriteString(changeset.Description)
	builder.WriteString("\n\n")

	date := changeset.Date.Format(time.DateTime)
	builder.WriteString(date)
	builder.WriteString("\n")

	builder.WriteString("🟢")
	builder.WriteString(changeset.Create)

	builder.WriteString(" | 🟠")
	builder.WriteString(changeset.Modify)

	builder.WriteString(" | 🔴")
	builder.WriteString(changeset.Delete)

	osmChaBtn := InlineKeyboardButton{}
	osmChaBtn.Text = "🌐 OSMCha"
	osmChaBtn.URL = fmt.Sprintf("https://osmcha.org/changesets/%d", changeset.ID)

	overPassBtn := InlineKeyboardButton{}
	overPassBtn.Text = "🌐 Overpass"
	overPassBtn.URL = fmt.Sprintf("https://overpass-api.de/achavi/?changeset=%d", changeset.ID)

	firstRow := []InlineKeyboardButton{
		osmChaBtn, overPassBtn,
	}

	inline_keyboard := [][]InlineKeyboardButton{
		firstRow,
	}

	markup := InlineKeyboardMarkup{
		InlineKeyboard: inline_keyboard,
	}

	message := Message{
		ChatID:                env.ChannelID,
		Text:                  builder.String(),
		ParseMode:             "HTML",
		ReplyMarkup:           markup,
		DisableWebPagePreview: true,
	}

	encodedMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", env.BotToken)

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(encodedMessage))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(data))
	}

	return nil
}
