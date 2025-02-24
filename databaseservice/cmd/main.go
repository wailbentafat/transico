package main

import (
	//"context"


	"databaseservice/internal/di"
	
	"databaseservice/pkj/utils"
	"fmt"
	"log"
	"net/http"
	// "net/http"
)

func main() {
	logger:=utils.NewLogger()
	// config:=config.Config{}
	
	
	// rdb,err:=redis.NewRedisconfig(config.REDIS_ADDR,config.Channelname,&logger)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	container,err:=di.Newcontainer(&logger)
	if err != nil {
		logger.LogError().Msg(err.Error())
	}
	fmt.Println("heyy container is here " ,container)



	
	
 
	 

	
    
	 port := ":8080"
	log.Printf("Go server running on http://localhost%s", port)
	 log.Fatal(http.ListenAndServe(port, nil))

}
