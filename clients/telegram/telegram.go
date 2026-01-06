package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"tgbot/lib/e"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	SendMessageMethod = "sendMessage"
)

func New(host string, token string) Client { // возвращаем тип Client
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string { // реализация для типа Client, вернет путь токена
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) { // Метод получения сообщений
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("offset", strconv.Itoa(limit))

	data, err := c.doRequst(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil

}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequst(SendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil

}

func (c *Client) doRequst(method string, query url.Values) (data []byte, err error) { // Метод отправки сообщений
	defer func() { err = e.WrapIfErr("can't do request", err) }() // обработчик ошибок, функциф дефер вызвается после завершения основной функции

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method), // path.Join с аргументами уберет лишние слеши из пути и склеит как надо
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}
