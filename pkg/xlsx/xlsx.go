package xlsx

import (
	"fmt"
	"sort"

	"github.com/xuri/excelize/v2"
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

func WriteXLSX(data []map[string]interface{}) (*excelize.File, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("no data to export")
	}

	f := excelize.NewFile()
	sheet := "Sheet1"

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

	// Записываем заголовки
	for col, header := range sortedHeaders {
		cell := fmt.Sprintf("%s1", string(rune('A'+col)))
		translation, ok := headerTranslations[header]
		if !ok {
			translation = header
		}
		f.SetCellValue(sheet, cell, translation)
	}

	// Записываем данные
	for row, item := range data {
		for col, header := range sortedHeaders {
			cell := fmt.Sprintf("%s%d", string(rune('A'+col)), row+2)
			f.SetCellValue(sheet, cell, item[header])
		}
	}

	return f, nil
}
