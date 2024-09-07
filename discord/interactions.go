package discord

import (
	"fmt"

	"github.com/QuartzWarrior/discordgo-self/types"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func GetSlashCommands(gateway *Gateway, guildID string) (types.ApplicationCommandIndex, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Brave\";v=\"128\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("referrer", fmt.Sprintf("https://discord.com/channels/%s", guildID))
	req.Header.Set("referrerPolicy", "strict-origin-when-cross-origin")
	req.SetRequestURI("https://discord.com/api/v9/guilds/" + guildID + "/application-command-index")
	resp := fasthttp.AcquireResponse()
	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.ApplicationCommandIndex{}, err
	}
	var applicationCommandIndex types.ApplicationCommandIndex
	err = json.Unmarshal(resp.Body(), &applicationCommandIndex)
	if err != nil {
		fmt.Println("Error:", err)
		return types.ApplicationCommandIndex{}, err
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return applicationCommandIndex, nil
}

func GetUserSlashCommands(gateway *Gateway) (types.ApplicationCommandIndex, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Brave\";v=\"128\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("referrer", "https://discord.com/channels/@me")
	req.Header.Set("referrerPolicy", "strict-origin-when-cross-origin")
	req.SetRequestURI("https://discord.com/api/v9/users/@me/application-command-index")
	resp := fasthttp.AcquireResponse()
	err := requestClient.Do(req, resp)
	if err != nil {
		fmt.Println("Error:", err)
		return types.ApplicationCommandIndex{}, err
	}
	var applicationCommandIndex types.ApplicationCommandIndex
	err = json.Unmarshal(resp.Body(), &applicationCommandIndex)
	if err != nil {
		fmt.Println("Error:", err)
		return types.ApplicationCommandIndex{}, err
	}
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return applicationCommandIndex, nil
}

func ClickButton(gateway *Gateway, e *types.MessageEventData, interactionID string) bool {
	sessionID := gateway.SessionID
	if sessionID == "" {
		gateway.SessionID = GenerateSessionID()
		sessionID = gateway.SessionID
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.Header.Set("authorization", gateway.Selfbot.Token)
	req.Header.Set("x-super-properties", GenerateSuperProperties(gateway))
	req.Header.Set("x-discord-locale", gateway.Selfbot.User.Locale)
	req.Header.Set("x-discord-timezone", "America/Denver")
	req.Header.Set("x-debug-options", "bugReporterEnabled")
	req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Google Chrome\";v=\"122\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Chrome OS\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("referrer", fmt.Sprintf("https://discord.com/channels/%s/%s", e.GuildID, e.ChannelID))
	req.Header.Set("referrerPolicy", "strict-origin-when-cross-origin")
	req.Header.SetUserAgent(gateway.Config.UserAgent)
	req.SetRequestURI("https://discord.com/api/v9/interactions")
	req.SetBodyString(fmt.Sprintf("{\"type\":3,\"nonce\":\"%s\",\"guild_id\":\"%s\",\"channel_id\":\"%s\",\"message_flags\":0,\"message_id\":\"%s\",\"application_id\":\"%s\",\"session_id\":\"%s\",\"data\":{\"component_type\":2,\"custom_id\":\"%s\"}}", GenerateNonce(), e.GuildID, e.ChannelID, e.ID, e.Author.ID, sessionID, interactionID))
	resp := fasthttp.AcquireResponse()
	err := requestClient.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode() == 204 {
		fasthttp.ReleaseResponse(resp)
		return true
	} else {
		fasthttp.ReleaseResponse(resp)
		return false
	}
}
