dapr run --app-id hello-go-server \
         --dapr-http-port 3501 \
         --dapr-grpc-port 3500 \
         --app-port 8200 \
         go run main.go