package main

import (
	"context"
	"fmt"
	"github.com/tigranftp/grpc_random_py_script/script_server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	mainContext := context.Background()

	dial, err := grpc.DialContext(mainContext,
		"localhost:50012",
		grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()

	srv := script_server.NewScriptServerClient(dial)
	scriptResult, err := srv.RunScript(
		mainContext,
		&script_server.ScriptRequest{Req: "выфвфыв"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(scriptResult.Rsp)
}
