package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
	"github.com/idkOybek/newNewTerminal/pkg/xlsx"
)

type ExportHandler struct {
	logger      *logger.Logger
	userService *service.UserService
}

func NewExportHandler(logger *logger.Logger, userService *service.UserService) *ExportHandler {
	return &ExportHandler{
		logger:      logger,
		userService: userService,
	}
}

// @Security Bearer
// @Summary Export data to XLSX
// @Description Export given data to XLSX format
// @Tags export
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param request body models.ExportRequest true "Export request"
// @Success 200 {file} string "exported_data.xlsx"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /export [post]
func (h *ExportHandler) ExportXLSX(w http.ResponseWriter, r *http.Request) {
	var req models.ExportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	filename := req.Filename
	if filename == "" {
		filename = fmt.Sprintf("export_%s", time.Now().Format("2006-01-02_15-04-05"))
	}
	filename = filename + ".xlsx"

	// Заменяем user_id на user_login и получаем логины пользователей
	for _, item := range req.Objects {
		if userID, ok := item["user_id"]; ok {
			user, err := h.userService.GetByID(r.Context(), int(userID.(float64)))
			if err == nil {
				item["user_login"] = user.Username
				delete(item, "user_id")
			}
		}
	}

	// Создаем XLSX файл
	xlsxFile, err := xlsx.WriteXLSX(req.Objects)
	if err != nil {
		h.logger.Error("Failed to create XLSX", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate XLSX")
		return
	}

	// Устанавливаем заголовки для XLSX файла
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Записываем XLSX файл в ответ
	err = xlsxFile.Write(w)
	if err != nil {
		h.logger.Error("Failed to write XLSX to response", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to send XLSX")
		return
	}
}
