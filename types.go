package main

type ReqData struct {
	ConversationId            string          `json:"conversationId"`
	ChatbotCorpId             string          `json:"chatbotCorpId"`
	ChatbotUserId             string          `json:"chatbotUserId"`
	MsgId                     string          `json:"msgId"`
	SenderNick                string          `json:"senderNick"`
	IsAdmin                   bool            `json:"isAdmin"`
	SenderStaffId             string          `json:"senderStaffId"`
	SessionWebhookExpiredTime int64           `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64           `json:"createAt"`
	SenderCorpId              string          `json:"senderCorpId"`
	ConversationType          string          `json:"conversationType"`
	SenderId                  string          `json:"senderId"`
	SessionWebhook            string          `json:"sessionWebhook"`
	Text                      *ReqDataContent `json:"text"`
	Msgtype                   string          `json:"msgtype"`
}

type ReqDataContent struct {
	Content string `json:"content"`
}
