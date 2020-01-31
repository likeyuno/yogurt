/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package message

import (
	"fmt"
	"time"
)

func DefaultMessage() string {
	return fmt.Sprintf(
		"我還在接受主人開發，暫時還是一個呆瓜沒有智慧，請通過 /start 來獲取幫助列表",
	)
}

func StartMessage() string {
	return fmt.Sprintf(
		"你好，我是人工智障 - 酸奶姬\n\n我目前已支持以下服務\n\n- /start 獲取幫助\n- /bind key 綁定訂閱及賬號\n- /subscription 查看賬號的訂閱信息\n- /link key 獲取訂閱鏈接\n\n如果我幫不上忙，請在群內留言等待管理",
	)
}

func WelcomeMessage(nickname, groupName string, userID int, chatID int64) string {
	return fmt.Sprintf(
		"歡迎 [%s](tg://user?id=%d) 加入 [%s](tg://group?id=%d)\n請及時閱讀群規，需要帮助可以找我或在群内留言，续费及购买请 At 管理",
		nickname, userID, groupName, chatID,
	)
}

func ArabicMessage(userID int, nickname string) string {
	return fmt.Sprintf("Fa 現一隻清真 [@%s](tg://user?id=%d) 已扔出群聊", nickname, userID)
}

func AccountAlreadyBindMessage() string {
	return fmt.Sprintf("豁呀，該賬號已經被綁定過了")
}

func AccountBindSuccessMessage(username string) string {
	return fmt.Sprintf("恭喜，賬號 %s 已與 Telegram 綁定", username)
}

func BindKeyEmptyMessage() string {
	return fmt.Sprintf("欸，似乎沒有提供 Key 呢，請使用 /bind key")
}

func LinkKeyEmptyMessage() string {
	return fmt.Sprintf("欸，似乎沒有提供 Key 呢，請使用 /link key")
}

func AccountNotFoundMessage() string {
	return fmt.Sprintf("欸，賬號似乎不存在")
}

func AccountNotBindMessage() string {
	return fmt.Sprintf("需要先通過 /bind key 綁定賬號噢")
}

func GetSubscriptionLinkMessage(text string) string {
	return fmt.Sprintf("以下是你的訂閱鏈接\n\n%s\n\n截圖請打碼，防止洩漏", text)
}

func SubscriptionListMessage(text string) string {
	return fmt.Sprintf("以下是你的訂閱列表\n\n%s\n\n使用 /link key 获取订阅链接", text)
}

func SubscriptionEmptyMessage() string {
	return fmt.Sprintf("你的賬號中似乎沒有訂閱噢")
}

func SubscriptionNotFoundMessage() string {
	return fmt.Sprintf("欸，訂閱似乎不存在")
}

func ErrorMessage() string {
	return fmt.Sprintf("哎呀，情緒崩壞無法提供服務，問題已報告給管理員")
}

func ChatNotPrivateMessage(nickname string, userID int) string {
	return fmt.Sprintf("[@%s](tg://user?id=%d) 這種事情悄悄地私信我噢", nickname, userID)
}

func InfoMessage(username string, expiredAt time.Time) string {
	return fmt.Sprintf(
		"该用户已绑定账号 %s\n到期时间为 %s",
		username, expiredAt.Format("2006-01-02"),
	)
}

func CreateParamErrorMessage() string {
	return fmt.Sprintf(
		"哎呀，缺少參數或格式錯誤",
	)
}

func PackageNotFoundMessage() string {
	return fmt.Sprintf("哎呀，沒有找到這種套餐")
}

func CreateAccountSuccessMessage() string {
	return fmt.Sprintf("好啦，創建賬號成功")
}

func CreateSubscriptionSuccessMessage(key, expiredAt string) string {
	return fmt.Sprintf("好啦，創建訂閱成功\n\n訂閱 Key %s\n到期時間 %s", key, expiredAt)
}

func NotIsMasterMessage() string {
	return fmt.Sprintf("哼，我才不聽你的呢")
}
