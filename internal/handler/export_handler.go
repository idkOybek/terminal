// internal/handler/export_handler.go

package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/csv"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
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
// @Summary Export data to CSV
// @Description Export given data to CSV format
// @Tags export
// @Accept json
// @Produce text/csv
// @Param data body []interface{} true "Data to export"
// @Param filename query string false "Name of the CSV file (without extension)"
// @Success 200 {file} string "exported_data.csv"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /export [post]
func (h *ExportHandler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	// Получаем имя файла из query параметров
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		// Если имя файла не указано, используем текущую дату и время
		filename = fmt.Sprintf("export_%s", time.Now().Format("2006-01-02_15-04-05"))
	}

	// Добавляем расширение .csv
	filename = filename + ".csv"

	// Декодируем JSON данные из тела запроса
	var data []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Заменяем user_id на user_login и получаем логины пользователей
	for _, item := range data {
		if userID, ok := item["user_id"]; ok {
			user, err := h.userService.GetByID(r.Context(), int(userID.(float64)))
			if err == nil {
				item["user_login"] = user.Username
				delete(item, "user_id")
			}
		}
	}

	// Устанавливаем заголовки для CSV файла
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))

	// Записываем данные в CSV
	err = csv.WriteCSV(data, w)
	if err != nil {
		h.logger.Error("Failed to write CSV", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate CSV")
		return
	}
}
