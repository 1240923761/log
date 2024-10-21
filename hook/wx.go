package hook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/1240923761/log/entity"
	"net/http"
	"sync"
)

type wxWorkHook struct {
	sync.Mutex
	client  *http.Client
	address string
}

func NewWxWorkHook(address string) *wxWorkHook {
	return &wxWorkHook{
		client:  http.DefaultClient,
		address: address,
	}
}

// todo: 完善具体内容, 支持不同格式等
func (w *wxWorkHook) Process(entity *entity.Entity) error {
	text := fmt.Sprintf(entity.Msg, entity.Args...)

	requestBody := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content": text,
		},
	}
	// 将请求体编码成JSON
	bs, _ := json.Marshal(requestBody)

	req, err := http.NewRequestWithContext(entity.Ctx, "POST", w.address, bytes.NewBuffer(bs))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	_, err = w.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}

	return nil
}
