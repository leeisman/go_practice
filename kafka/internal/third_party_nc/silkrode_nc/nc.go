package silkrode_nc

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.silkrode.com.tw/golang/errors"
	db "gitlab.silkrode.com.tw/golang/gopher/db/v2/db"
	"gorm.io/gorm"
	"nc/pkg/model"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

// 串接內部推播系統
// https://gitlab.silkrode.com.tw/team_golang/nc
type SilkrodeNC struct {
	Config Config
	Write  *gorm.DB
}

func NewSilkrodeNC(config *Config, conn *db.Connections) *SilkrodeNC {
	return &SilkrodeNC{
		Config: Config{
			AppID:     config.AppID,
			SecretKey: config.SecretKey,
			Hostname:  config.Hostname,
		},
		Write: conn.WriteDB,
	}
}

// 串接推播需要帶的headers
func (nc *SilkrodeNC) getHeaders(requestBody string) map[string]string {
	reqTime := strconv.FormatInt(time.Now().Unix(), 10)
	m := md5.New()
	m.Write([]byte(string([]byte{}) + reqTime + nc.Config.SecretKey))
	headers := make(map[string]string)
	headers["app_id"] = nc.Config.AppID
	headers["time"] = reqTime

	md5Str := requestBody + reqTime + nc.Config.SecretKey
	data := []byte(md5Str)
	has := md5.Sum(data)
	sign := fmt.Sprintf("%x", has)
	headers["sign"] = sign
	return headers
}

// 主要推播 func
func (nc *SilkrodeNC) Push(ctx context.Context, users []*model.User, notification *model.Notification) {
	// 組資料
	deviceIDs := make([]string, 0)
	data := ""
	ncMsg, err := json.Marshal(struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		Title:   notification.Title,
		Content: notification.Content,
	})
	if err != nil {
		nc.recordError(nil, nil, err)
	}
	for _, users := range users {
		for _, userDeviceID := range users.UserDeviceIDs {
			deviceIDs = append(deviceIDs, userDeviceID)
		}
	}

	client := http.Client{}
	body := struct {
		DeviceIDs string `json:"device_ids"`
		Msg       string `json:"msg"`
		Data      string `json:"data"`
	}{
		DeviceIDs: strings.Join(deviceIDs, ","),
		Msg:       string(ncMsg),
		Data:      data,
	}
	input, err := json.Marshal(body)
	if err != nil {
		nc.recordError(nil, nil, err)
	}
	u, _ := url.Parse(nc.Config.Hostname)
	u.Path = path.Join(u.Path, "notify_messages")
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bytes.NewReader(input))
	if err != nil {
		nc.recordError(nil, nil, err)
	}
	for key, value := range nc.getHeaders(string(input)) {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		nc.recordError(req, resp, err)
	}
	if resp.StatusCode != 200 {
		nc.recordError(req, resp, errors.New("status code :", resp.Status))
	}
}

// 錯誤紀錄
func (nc *SilkrodeNC) recordError(req *http.Request, resp *http.Response, err error) {
	ncNcError := model.ThirdPartyNCError{
		Type:     model.SilkrodeNC,
		Req:      fmt.Sprintf("%s", req),
		Resp:     fmt.Sprintf("%s", resp),
		ErrorMsg: err.Error(),
	}
	err = nc.Write.Create(&ncNcError).Error
	if err != nil {
		log.Error().Msgf("recordError err: %s", err.Error())
	}
}
