package main

import (
	"strings"

	"go.mau.fi/whatsmeow/types"
)

func cmdGetGroup(args []string) {
	if len(args) < 1 {
		log.Errorf("Usage: getgroup <jid>")
		return
	}
	group, ok := parseJID(args[0])
	if !ok {
		return
	} else if group.Server != types.GroupServer {
		log.Errorf("Input must be a group JID (@%s)", types.GroupServer)
		return
	}
	resp, err := cli.GetGroupInfo(group)
	if err != nil {
		log.Errorf("Failed to get group info: %v", err)
	} else {
		log.Infof("Group info: %+v", resp)
	}
}

func cmdListGroups(args []string) {
	groups, err := cli.GetJoinedGroups()
	if err != nil {
		log.Errorf("Failed to get group list: %v", err)
	} else {
		for _, group := range groups {
			log.Infof("%+v: %+v", group.GroupName.Name, group.JID)
		}
	}
}

func cmdSendSpoofedReply(args []string) {
	if len(args) < 4 {
		log.Errorf("Usage: send-spoofed-reply <chat_jid> <msgID:!|#ID> <spoofed_jid> <spoofed_text>|<text>")
		return
	}

	chat_jid, ok := parseJID(args[0])
	if !ok {
		return
	}

	msgID := args[1]
	if msgID[0] == '!' {
		msgID = cli.GenerateMessageID()
	}

	spoofed_jid, ok2 := parseJID(args[2])
	if !ok2 {
		return
	}

	parameters := strings.SplitN(strings.Join(args[3:], " "), "|", 2)
	spoofed_text := parameters[0]
	text := parameters[1]

	_, resp, err := sendSpoofedReplyMessage(chat_jid, spoofed_jid, msgID, spoofed_text, text)
	if err != nil {
		log.Errorf("Error on sending spoofed msg: %v", err)
	} else {
		log.Infof("spoofed msg sended: %+v", resp)
	}
}

func cmdSendSpoofedImgReply(args []string) {
	if len(args) < 5 {
		log.Errorf("Usage: send-spoofed-img-reply <chat_jid> <msgID:!|#ID> <spoofed_jid> <spoofed_file> <spoofed_text>|<text>")
		return
	}
	chat_jid, ok := parseJID(args[0])
	if !ok {
		return
	}

	msgID := args[1]
	if msgID[0] == '!' {
		msgID = cli.GenerateMessageID()
	}

	spoofed_jid, ok2 := parseJID(args[2])
	if !ok2 {
		return
	}

	spoofed_file := args[3]

	parameters := strings.SplitN(strings.Join(args[4:], " "), "|", 2)
	spoofed_text := parameters[0]
	text := parameters[1]

	_, resp, err := sendSpoofedReplyImg(chat_jid, spoofed_jid, msgID, spoofed_file, spoofed_text, text)
	if err != nil {
		log.Errorf("Error on sending spoofed msg: %v", err)
	} else {
		log.Infof("spoofed msg sended: %+v", resp)
	}
}

func cmdSendSpoofedDemo(args []string) {
	if len(args) < 4 {
		log.Errorf("Usage: send-spoofed-demo <toGender:boy|girl> <language:br|en> <chat_jid> <spoofed_jid>")
		return
	}

	var toGender string
	if args[0] != "boy" && args[0] != "girl" {
		log.Errorf("Error: <boy|girl>")
		return
	} else {
		toGender = args[0]
	}

	var language string
	if args[1] != "br" && args[1] != "en" {
		log.Errorf("Error: <br|en>")
		return
	} else {
		language = args[1]
	}

	chat_jid, ok := parseJID(args[2])
	if !ok {
		log.Errorf("Error: chat_jid")
		return
	}
	spoofed_jid, ok2 := parseJID(args[3])
	if !ok2 {
		log.Errorf("Error: spoofed_jid")
		return
	}
	sendSpoofedTalkDemo(chat_jid, spoofed_jid, toGender, language, "")

}

func cmdSendSpoofedDemoImg(args []string) {
	if len(args) < 5 {
		log.Errorf("Usage: send-spoofed-demo-img <toGender:boy|girl> <language:br|en> <chat_jid> <spoofed_jid> <spoofed_img>")
		return
	}

	var toGender string
	if args[0] != "boy" && args[0] != "girl" {
		log.Errorf("Error: <boy|girl>")
		return
	} else {
		toGender = args[0]
	}

	var language string
	if args[1] != "br" && args[1] != "en" {
		log.Errorf("Error: <br|en>")
		return
	} else {
		language = args[1]
	}

	chat_jid, ok := parseJID(args[2])
	if !ok {
		log.Errorf("Error: chat_jid")
		return
	}
	spoofed_jid, ok2 := parseJID(args[3])
	if !ok2 {
		log.Errorf("Error: spoofed_jid")
		return
	}

	spoofed_img := args[4]

	sendSpoofedTalkDemo(chat_jid, spoofed_jid, toGender, language, spoofed_img)
}
