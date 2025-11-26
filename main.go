package main

import (
	"flag" // флаги нужны для определения параметров запуска приложения ./myapp -port=8080 -debug=true
	"log"
)

func main() {

	t := mustToken()

	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	//fetcher = fetcher.New()

	//processor = processor.New()

	//consumer.Start(fetcher, processor)

}

func mustToken() string {

	token := flag.String("token-bot-token", "", "token for access to telegram bot") // в переменной токен будет лежать не значение, а лишь ссылка на значение

	flag.Parse() // передает значение в token

	if *token == "" {
		log.Fatal("token is not specified")
	}

}
