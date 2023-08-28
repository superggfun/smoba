package doTask

import (
	"encoding/json"
	"fmt"
)

const RunGameURL = "https://ssl.kohsocialapp.qq.com:10001/play/gettaskconditiondata"

type RunGameResponse struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

// RunGame executes the game and returns any encountered error
func (m *Account) RunGame() error {
	return m.doTaskWithDataType(2)
}

func (m *Account) ViewRecords() error {
	return m.doTaskWithDataType(6)
}

func (m *Account) ShareGame() error {
	return m.doTaskWithDataType(1)
}

// doTaskWithDataType is a helper function to perform tasks with different data types.
func (m *Account) doTaskWithDataType(dtype int) error {
	data := fmt.Sprintf("type=%v&token=%v&userId=%v", dtype, m.Token, m.UserId)

	bodyText, err := m.DoTask(RunGameURL, data)
	if err != nil {
		return fmt.Errorf("error executing DoTask: %w", err)
	}

	var response RunGameResponse
	if err = json.Unmarshal(bodyText, &response); err != nil {
		return fmt.Errorf("error unmarshalling response: %w", err)
	}

	if response.ReturnCode != 0 {
		return fmt.Errorf("server returned non-OK status: %s", response.ReturnMsg)
	}

	return nil
}
