01 - Obter a key de autorizacao da openAI em https://platform.openai.com/api-keys OBERVACAO: a conta que vai gerar a key da API deve ter menos de 3 meses desde a criacao, se nao nao funciona pois o periodo de teste da openAI eh 3 meses.

02 - Copiar a key e colocar no arquivo .env no campo OPENAI_API_KEY

03 - Colocar o contexto que o chat ira trabalhar necessario no campo INITIAL_CHAT_MESSAGE no arquivo .env

04 - iniciar servicos -> na pasta raiz digitar "docker compose up -d"

05 - acessar servico do golang -> " docker exec -it chatservice_app bash"

06 - Executar o web server -> "go run cmd/chatservice/main.go"

07 - Testar a API utilizando requisicoes REST no insomnia, postman ou na extensao REST Client do vscode
