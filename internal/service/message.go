package service

var messages = []string{"Hello", "World"}

func GetMessages() []string {
    return messages
}

func AddMessage(msg string) {
    messages = append(messages, msg)
}
