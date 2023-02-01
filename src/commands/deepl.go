package commands

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var supportLangs = [...]string{
	"BG",
	"CS",
	"DA",
	"DE",
	"EL",
	"EN",
	"ES",
	"ET",
	"FI",
	"FR",
	"HU",
	"IT",
	"JA",
	"LT",
	"LV",
	"NL",
	"PL",
	"PT",
	"RO",
	"RU",
	"SK",
	"SL",
	"SV",
	"ZH",
}

type DeeplResult struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func Deepl(message tgbotapi.Message) string {
	lang := strings.TrimSpace(message.CommandArguments())
	lang = strings.ToUpper(lang)
	if message.ReplyToMessage == nil || message.ReplyToMessage.Text == "" {
		return "No text to translate"
	} else {
		for _, eachItem := range supportLangs {
			if eachItem == lang {
				return trans(message.ReplyToMessage.Text, lang)
			}
		}
		return "Language not supported"
	}
}

func trans(msg string, lang string) string {
	client := &http.Client{}
	form := url.Values{}
	form.Set("text", msg)
	form.Set("target_lang", lang)
	req, err := http.NewRequest("POST", os.Getenv("DEEPL_API_ULR"), strings.NewReader(form.Encode()))
	if err != nil {
		return "Error Build Request Error"
	}
	req.Header.Set("Authorization", "DeepL-Auth-Key "+os.Getenv("DEEPL_APIKEY"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "Error - Request error"
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "Error - Deepl API error code: " + resp.Status
	}
	var result DeeplResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "Error - Decoder error"
	}

	return result.Translations[0].Text
}
