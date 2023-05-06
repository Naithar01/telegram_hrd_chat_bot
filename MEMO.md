switch {
case u.Message != nil:
	return u.Message.Chat
case u.EditedMessage != nil:
	return u.EditedMessage.Chat
case u.ChannelPost != nil:
	return u.ChannelPost.Chat
case u.EditedChannelPost != nil:
	return u.EditedChannelPost.Chat
case u.CallbackQuery != nil:
	return u.CallbackQuery.Message.Chat
default:
	return nil
}