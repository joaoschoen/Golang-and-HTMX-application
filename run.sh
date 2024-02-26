(
  export PORT="8080" &&
  nodemon --watch './**/*.go' --ext 'go,html,css' --signal SIGTERM --exec 'go run main.go'

)&
(
  npx tailwind \
  -i 'styles.css' \
  -o 'public/styles.css' \
  --watch
)
