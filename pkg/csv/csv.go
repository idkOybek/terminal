package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

func WriteCSV(data []interface{}, writer io.Writer) error {
	if len(data) == 0 {
		return nil
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Собираем все возможные заголовки
	headers := make(map[string]bool)
	for _, item := range data {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return err
		}

		var itemMap map[string]interface{}
		if err := json.Unmarshal(jsonData, &itemMap); err != nil {
			return err
		}

		for key := range itemMap {
			headers[key] = true
		}
	}

	// Сортируем заголовки для консистентности
	var sortedHeaders []string
	for header := range headers {
		sortedHeaders = append(sortedHeaders, header)
	}
	sort.Strings(sortedHeaders)

	// Записываем заголовки
	if err := csvWriter.Write(sortedHeaders); err != nil {
		return err
	}

	// Записываем данные
	for _, item := range data {
		jsonData, err := json.Marshal(item)
		if err != nil {
			return err
		}

		var itemMap map[string]interface{}
		if err := json.Unmarshal(jsonData, &itemMap); err != nil {
			return err
		}

		var row []string
		for _, header := range sortedHeaders {
			value := fmt.Sprintf("%v", itemMap[header])
			row = append(row, value)
		}

		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	return nil
}
