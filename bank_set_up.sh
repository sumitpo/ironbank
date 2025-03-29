function apiSetup() {
  mkdir -p shorturl/api
  (cd shorturl; go mod init shorturl)
  (cd shorturl/api; goctl api -o shorturl.api)
  echo "change to shorturl/api/shorturl.api file to define your services and [goctl api go -api shorturl.api -dir .] to generate code"
}

apiSetup
