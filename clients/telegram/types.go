package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"update_id"` //теги структур, апдейты с сервера будут приходить ввиде джейсона и стандартный парсер будет искать в этом ответе поле id, но его там не будет, а будет update_id, и с помощью этого тега мы подсказываем парсеру, куда надо смотреть
	Message string `json:"message"`
}
