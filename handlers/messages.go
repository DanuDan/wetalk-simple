package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	messagesdto "wetalk/dto/messages"
	dto "wetalk/dto/result"
	"wetalk/models"
	"wetalk/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerMessage struct {
	MessageRepository repositories.MessageRepository
}

func HandlerMessage(MessageRepository repositories.MessageRepository) *handlerMessage {
	return &handlerMessage{MessageRepository}
}

func (h *handlerMessage) FindMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	messages, err := h.MessageRepository.FindMessages(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: messages}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerMessage) GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	message, err := h.MessageRepository.GetMessage(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponseMessage(message)}
	json.NewEncoder(w).Encode(response)
	return
}

func (h *handlerMessage) CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	request := messagesdto.CreateMessageRequest{
		Message: r.FormValue("message"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	message := models.Message{
		Message: request.Message,
		UserID:  userId,
	}

	message, err = h.MessageRepository.CreateMessage(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	message, _ = h.MessageRepository.GetMessage(message.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: message}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerMessage) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	message, err := h.MessageRepository.GetMessage(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteMessage, err := h.MessageRepository.DeleteMessage(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: deleteMessage}
	json.NewEncoder(w).Encode(response)
}

func convertResponseMessage(u models.Message) messagesdto.MessageResponse {
	return messagesdto.MessageResponse{
		ID:      u.ID,
		Message: u.Message,
		UserID:  u.User.ID,
	}
}
