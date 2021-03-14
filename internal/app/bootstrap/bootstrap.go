package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/raspiantoro/transaction-wrapper/internal/app/appcontext"
)

func Run() {

	stop := make(chan bool)

	appContext := appcontext.New()

	db, err := appContext.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := appContext.GetRepository(db)
	svc := appContext.GetService(repo)
	handler := appContext.GetHandler(svc)
	r := appContext.GetRouter(handler)

	r.Init()

	go func() {
		interupt := make(chan os.Signal)
		signal.Notify(interupt, os.Interrupt)

		<-interupt
		stop <- true
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("run http server on port 0.0.0.0:8080")
	<-stop

	fmt.Println("bye")
}
