(
  npx nodemon --watch './**/*.go' \
  --ext 'go,html,css' \
  --signal SIGTERM \
  --exec 'go run main.go'
)&
(
  npx nodemon --watch './**/*' \
  --ext 'html,css,go' \
  --signal SIGTERM \
  --exec "tailwind -i 'styles.css' -o 'public/styles.css'"
)& 
(
  npx browser-sync start --proxy 'localhost:1323' -e * --ignore=node_modules -w --reload-delay="1000" --files '**/*.html, **/*.css, **/*.go'
)
