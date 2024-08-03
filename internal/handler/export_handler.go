package handler

import (
	"encoding/json"
	"net/http"

	"github.com/idkOybek/newNewTerminal/pkg/csv"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type ExportHandler struct {
	logger *logger.Logger
}

func NewExportHandler(logger *logger.Logger) *ExportHandler {
	return &ExportHandler{
		logger: logger,
	}
}

// @Summary Export data to CSV
// @Description Export given data to CSV format
// @Tags export
// @Accept json
// @Produce text/csv
// @Param data body []interface{} true "Data to export"
// @Param filename query string true "Name of the CSV file"
// @Success 200 {file} string "exported_data.csv"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /export [post]
func (h *ExportHandler) ExportCSV(w http.ResponseWriter, r *http.Request) {
	// Получаем имя файла из query параметров
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		filename = "exported_data.csv"
	}

	// Декодируем JSON данные из тела запроса
	var data []interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Устанавливаем заголовки для CSV файла
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	// Записываем данные в CSV
	err = csv.WriteCSV(data, w)
	if err != nil {
		h.logger.Error("Failed to write CSV", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to generate CSV")
		return
	}
}
