package main

import (
	"database/sql"
	"fmt"

	"github.com/pedrosandrini/fclx/chatservice/configs"
	"github.com/pedrosandrini/fclx/chatservice/internal/infra/grpc/server"
	"github.com/pedrosandrini/fclx/chatservice/internal/infra/repository"
	"github.com/pedrosandrini/fclx/chatservice/internal/infra/web"
	"github.com/pedrosandrini/fclx/chatservice/internal/infra/web/webserver"
	"github.com/pedrosandrini/fclx/chatservice/internal/usecase/chatcompletion"
	"github.com/pedrosandrini/fclx/chatservice/internal/usecase/chatcompletionstream"
	"github.com/sashabaranov/go-openai"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	repo := repository.NewChatRepositoryMySql(conn)
	client := openai.NewClient(configs.OpenAIApiKey)

	chatConfig := chatcompletion.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	chatConfigStream := chatcompletionstream.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	usecase := chatcompletion.NewChatCompletionUseCase(repo, client)

	streamChannel := make(chan chatcompletionstream.ChatCompletionOutputDTO)
	useCaseStream := chatcompletionstream.NewChatCompletionUseCase(repo, client, streamChannel)

	fmt.Println("Starting gRPC server on port " + configs.GRPCServerPort)
	grpcServer := server.NewGRPCServer(
		*useCaseStream,
		chatConfigStream,
		configs.GRPCServerPort,
		configs.AuthToken,
		streamChannel,
	)
	go grpcServer.Start()

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverChatHandler := web.NewWebChatGPTHandler(*usecase, chatConfig, configs.AuthToken)
	webserver.AddHandler("/chat", webserverChatHandler.Handle)
	webserver.Start()

}
