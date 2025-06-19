// @title           Simple Withdraw API
// @version         1.0
// @description     This is a simple withdrawal and inquiry service
// @host            localhost:6005
// @BasePath        /api
package main

import "simple-withdraw-api/internal/infrastructure"

func main() {
	infrastructure.Run()
}
