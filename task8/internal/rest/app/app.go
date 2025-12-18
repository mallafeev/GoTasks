package app

import (
	"GoTasks/task8/internal/rest/app/service"
	"GoTasks/task8/internal/rest/bootstrap"
	"GoTasks/task8/internal/rest/config"
	"GoTasks/task8/internal/rest/repository/database"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) error {

	db, err := bootstrap.InitDB(cfg)
	if err != nil {
		return err
	}

	employeeService := service.NewService(database.NewDatabase(db))

	router := http.NewServeMux()
	router.HandleFunc("POST /employees", employeeService.Create)
	router.HandleFunc("GET /employees/{id}", employeeService.Get)
	router.HandleFunc("GET /employees", employeeService.GetAll)
	router.HandleFunc("PUT /employees/{id}", employeeService.Update)
	router.HandleFunc("DELETE /employees/{id}", employeeService.Delete)

	srv := http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("run server: http://localhost%s", cfg.Port)
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("error when listen and serve", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)
	sig := <-ch
	log.Printf("received signal: %s", sig)
	return srv.Shutdown(context.Background())
}
