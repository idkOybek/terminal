// pkg/csv/csv.go

package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

var headerTranslations = map[string]string{
	"id":                    "ID",
	"user_login":            "Логин пользователя",
	"username":              "Имя пользователя",
	"email":                 "Электронная почта",
	"created_at":            "Дата создания",
	"updated_at":            "Дата обновления",
	"is_active":             "Активен",
	"is_admin":              "Администратор",
	"fiscal_number":         "Фискальный номер",
	"factory_number":        "Заводской номер",
	"inn":                   "ИНН",
	"company_name":          "Название компании",
	"address":               "Адрес",
	"cash_register_number":  "Номер кассового аппарата",
	"module_number":         "Номер модуля",
	"assembly_number":       "Номер сборки",
	"last_request_date":     "Дата последнего запроса",
	"database_update_date":  "Дата обновления базы данных",
	"status":                "Статус",
	"free_record_balance":   "Баланс свободных записей",
	"password":              "Пароль",
	"phone":                 "Телефон",
	"role":                  "Роль",
	"last_login":            "Последний вход",
	"registration_date":     "Дата регистрации",
	"balance":               "Баланс",
	"activation_date":       "Дата активации",
	"expiration_date":       "Дата истечения срока",
	"notes":                 "Примечания",
	"department":            "Отдел",
	"position":              "Должность",
	"salary":                "Зарплата",
	"manager":               "Менеджер",
	"region":                "Регион",
	"city":                  "Город",
	"postal_code":           "Почтовый индекс",
	"country":               "Страна",
	"website":               "Веб-сайт",
	"tax_number":            "Налоговый номер",
	"legal_entity":          "Юридическое лицо",
	"contract_number":       "Номер договора",
	"contract_date":         "Дата договора",
	"service_plan":          "Тарифный план",
	"last_payment_date":     "Дата последнего платежа",
	"next_payment_date":     "Дата следующего платежа",
	"total_transactions":    "Общее количество транзакций",
	"total_revenue":         "Общая выручка",
	"average_check":         "Средний чек",
	"loyalty_points":        "Баллы лояльности",
	"referral_code":         "Реферальный код",
	"last_maintenance_date": "Дата последнего обслуживания",
	"software_version":      "Версия ПО",
	"hardware_model":        "Модель оборудования",
	"connection_type":       "Тип подключения",
	"ip_address":            "IP-адрес",
	"mac_address":           "MAC-адрес",
	"last_sync_date":        "Дата последней синхронизации",
	"timezone":              "Часовой пояс",
	"language":              "Язык",
	"currency":              "Валюта",
}

func WriteCSV(data []map[string]interface{}, writer io.Writer) error {
	if len(data) == 0 {
		return nil
	}

	csvWriter := csv.NewWriter(writer)
	defer csvWriter.Flush()

	// Собираем все возможные заголовки
	headers := make(map[string]bool)
	for _, item := range data {
		for key := range item {
			headers[key] = true
		}
	}

	// Сортируем заголовки, но "id" должен быть первым
	var sortedHeaders []string
	if headers["id"] {
		sortedHeaders = append(sortedHeaders, "id")
		delete(headers, "id")
	}
	for header := range headers {
		sortedHeaders = append(sortedHeaders, header)
	}
	sort.Strings(sortedHeaders[1:]) // Сортируем все, кроме первого элемента (id)

	// Переводим заголовки на русский
	var translatedHeaders []string
	for _, header := range sortedHeaders {
		if translation, ok := headerTranslations[header]; ok {
			translatedHeaders = append(translatedHeaders, translation)
		} else {
			translatedHeaders = append(translatedHeaders, header)
		}
	}

	// Записываем заголовки
	if err := csvWriter.Write(translatedHeaders); err != nil {
		return err
	}

	// Записываем данные
	for _, item := range data {
		var row []string
		for _, header := range sortedHeaders {
			value := fmt.Sprintf("%v", item[header])
			row = append(row, value)
		}

		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	return nil
}
