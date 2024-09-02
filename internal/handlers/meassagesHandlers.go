package handlers

import (
	"context"
	//"fmt"
	"go_go/internal/messagesService"
	"go_go/internal/web/messages"
)

type Handler struct {
	Service *messagesService.MessageService
}

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	//fmt.Println("\n==============********", *request.Body.Message, "**************===========")
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchMessages(_ context.Context, request messages.PatchMessagesRequestObject) (messages.PatchMessagesResponseObject, error) {
	messageRequest := request.Body
	idToUpdate := uint(*messageRequest.Id)
	messageToUpdate := messagesService.Message{Text: *messageRequest.Message}

	//fmt.Println("\n********============\nPATCH_ID=", idToUpdate, "\n***********=============")
	updateMessage, err := h.Service.UpdateMessageByID(idToUpdate, messageToUpdate)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PatchMessages200JSONResponse{
		Id:      &updateMessage.ID,
		Message: &updateMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) DeleteMessages(_ context.Context, request messages.DeleteMessagesRequestObject) (messages.DeleteMessagesResponseObject, error) {
	messageRequest := request.Body
	idToDelete := uint(*messageRequest.Id)
	err := h.Service.DeleteMessageByID(idToDelete)
	if err != nil {
		return nil, err
	}
	return nil, err
}
