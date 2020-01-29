/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	WelcomeAnimationID = "CgACAgUAAxkBAANyXjEx1OAytW9umKzEsXpYtjprZw8AAvUAA9JHiFXAsP2LeDy6AxgE"
	HelpAnimationID    = "CgACAgUAAxkBAANqXjEsshi0jsxpVCM9vT8XHCCi11QAAmMAA436CVXl-e9zFqNojxgE"
)

var JoinKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("查看群規", "https://t.me/yogurtCloud/5"),
	),
)

var ArabicKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Ta 並不清真", "https://t.me/yogurtCloud"),
	),
)

func StartMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "你好，我是人工智障 - 酸奶姬\n\n我目前已支持以下服務\n\n- /start 獲取幫助\n- /bind key 綁定訂閱及賬號\n- /subscription 查看賬號的訂閱信息\n- /link key 獲取訂閱鏈接\n\n如果我幫不上忙，請在群內留言等待管理")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func ErrorMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "哎呀，情緒崩壞無法提供服務，問題已報告給管理員")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func ChatNotPrivateMessage(chatID int64, replyID int, nickname string, userID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("[@%s](tg://user?id=%d) 這種事情悄悄地私信我噢", nickname, userID)
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func SubscriptionNotFoundMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("欸，訂閱似乎不存在")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func SubscriptionNotHaveMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("你的賬號中似乎沒有訂閱噢")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func SubscriptionListMessage(chatID int64, replyID int, data string) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("以下是你的訂閱列表\n\n%s\n\n使用 /link key 获取订阅链接", data)
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func GetSubscriptionLinkMessage(chatID int64, replyID int, data string) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("以下是你的訂閱鏈接\n\n%s\n\n截圖請打碼，防止洩漏", data)
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func AccountNotBindMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("需要先通過 /bind key 綁定賬號噢")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func AccountNotFoundMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("欸，賬號似乎不存在")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func GetSubscriptionLinEmptyMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("欸，似乎沒有提供 Key 呢，請使用 /link key")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func AccountBindEmptyMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("欸，似乎沒有提供 Key 呢，請使用 /bind key")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func AccountBindSuccessMessage(chatID int64, replyID int, username string) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf("恭喜，賬號 %s 已與 Telegram 綁定", username)
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func AccountAlreadyBindMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "豁呀，該賬號已經被綁定過了")
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func ArabicMessage(chatID int64, replyID int, userID int, nickname string) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Fa 現一隻清真 [@%s](tg://user?id=%d) 已扔出群聊", nickname, userID))
	msg.ReplyMarkup = ArabicKeyboard
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	return msg
}

func WelcomeMessage(chatID int64, nickname, groupName string, userID int) tgbotapi.Chattable {
	ac := tgbotapi.NewAnimationShare(chatID, WelcomeAnimationID)
	ac.Caption = fmt.Sprintf(
		"歡迎 [%s](tg://user?id=%d) 加入 [%s](tg://group?id=%d)\n請及時閱讀群規，需要帮助可以找我或在群内留言，续费及购买请 At 管理", nickname, userID, groupName, chatID,
	)
	ac.ParseMode = "Markdown"
	ac.ReplyMarkup = JoinKeyboard
	return ac
}

func DefaultMessage(chatID int64, replyID int) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "")
	msg.Text = fmt.Sprintf(
		"我還在接受主人開發，暫時還是一個呆瓜沒有智慧，請通過 /start 來獲取幫助列表",
	)
	msg.ParseMode = "Markdown"
	msg.ReplyToMessageID = replyID
	msg.ReplyMarkup = JoinKeyboard
	return msg
}
