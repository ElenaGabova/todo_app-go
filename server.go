package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct{
	httpServer *http.Server;
}

//Метод запуска
func(s * Server) Run(port string) error{
	//Конфигурация параметров сервера
	s.httpServer = &http.Server{
		Addr: ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,

	}

	//Метод слушает все входящие запросы для обработки
	return s.httpServer.ListenAndServe()

}

//Метод запуска
func(s * Server) Shoutdown(ctx context.Context) error{	
	return s.httpServer.Shutdown(ctx)
}
