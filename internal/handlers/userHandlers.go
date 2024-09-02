package handlers

import (
	"context"

	//"fmt"
	"go_go/internal/userService"
	"go_go/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUser(_ context.Context, _ users.GetUserRequestObject) (users.GetUserResponseObject, error) {
	// Получение всех сообщений из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUser200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, u := range allUsers {
		formattedCreatedAt := u.CreatedAt.Format("2006-01-02 15:04:05")
		formattedUpdatedAt := u.UpdatedAt.Format("2006-01-02 15:04:05")
		user := users.User{

			Id:        &u.ID,
			Email:     &u.Email,
			Password:  &u.Password,
			CreatedAt: &formattedCreatedAt,
			UpdatedAt: &formattedUpdatedAt,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUser(_ context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	//fmt.Println("\n==============********", *request.Body.Message, "**************===========")
	userRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	userToCreate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	formattedCreatedAt := createdUser.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := createdUser.UpdatedAt.Format("2006-01-02 15:04:05")
	response := users.PostUser201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &formattedCreatedAt,
		UpdatedAt: &formattedUpdatedAt,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) PatchUser(_ context.Context, request users.PatchUserRequestObject) (users.PatchUserResponseObject, error) {
	userRequest := request.Body
	idToUpdate := uint(*userRequest.Id)
	// for future userToUpdate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password, Model: gorm.Model{ID: idToUpdate}}
	userToUpdate := userService.User{Email: *userRequest.Email, Password: *userRequest.Password}

	//fmt.Println("\n********============\nPATCH_ID=", idToUpdate, "\n***********=============")
	updateUser, err := h.Service.UpdateUserByID(idToUpdate, userToUpdate)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	formattedCreatedAt := updateUser.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := updateUser.UpdatedAt.Format("2006-01-02 15:04:05")
	response := users.PatchUser200JSONResponse{
		Id:        &updateUser.ID,
		Email:     &updateUser.Email,
		Password:  &updateUser.Password,
		CreatedAt: &formattedCreatedAt,
		UpdatedAt: &formattedUpdatedAt,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) DeleteUser(_ context.Context, request users.DeleteUserRequestObject) (users.DeleteUserResponseObject, error) {
	messageRequest := request.Body
	idToDelete := uint(*messageRequest.Id)
	err := h.Service.DeleteUserByID(idToDelete)
	if err != nil {
		return nil, err
	}
	return nil, err
}
