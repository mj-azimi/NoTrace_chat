package lib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendRequest(url string, jsonData map[string]interface{}) (string, error) {
	// تبدیل داده‌ها به فرمت JSON
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	// ایجاد درخواست POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}

	// تنظیم هدر Content-Type
	req.Header.Set("Content-Type", "application/json")

	// ارسال درخواست
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// خواندن بدنه پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// بازگشت بدنه پاسخ به صورت رشته
	return string(body), nil
}
