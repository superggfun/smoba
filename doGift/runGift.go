package doGift

import (
	"encoding/json"
	"errors"
	"fmt"
)

type gift struct {
	ReturnCode int    `json:"returnCode"`
	ReturnMsg  string `json:"returnMsg"`
}

func (m *Account) RunGift(taskIds []string) error {
	a, err := json.Marshal(taskIds)
	if err != nil {
		return err
	}
	if len(a) == 2 {
		return nil
	}
	data := fmt.Sprintf(`{"cSystem":"android","h5Get":1,"taskIds":%v,"roleId":%v}`, string(a), m.RoleId)
	bodyText, err := m.DoGift("https://kohcamp.qq.com/operation/action/rewardtask", data, m.UserId, m.OriginalRoleId)
	if err != nil {
		return err
	}
	var gift gift
	err = json.Unmarshal(bodyText, &gift)
	if err != nil {
		return err
	} else if gift.ReturnCode != 0 {
		return errors.New(gift.ReturnMsg)
	} else {
		return nil
	}

}
