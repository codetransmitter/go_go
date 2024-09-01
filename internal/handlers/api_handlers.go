package handlers

import (
	"context"
	"encoding/json"

	//"fmt"
	"go_go/internal/messagesService"
	"go_go/internal/web/messages"
	"net/http"
)

type Handler struct {
	Service *messagesService.MessageService
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

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) UpdateMessageByID(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateMessage, err := h.Service.UpdateMessageByID(message.ID, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateMessage)
}

func (h *Handler) DeleteMessageByID(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteMessageByID(message.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
