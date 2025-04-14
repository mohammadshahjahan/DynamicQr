package routes

import (
	Middleware "backend/middlewares"
	DashboardService "backend/routehandler/DashboardService"
	proxyurl "backend/routehandler/ProxyUrl"
	QrService "backend/routehandler/QrService"
	UserService "backend/routehandler/UserService"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users/signup", UserService.Signup).Methods("POST")
	router.HandleFunc("/users/login", UserService.Login).Methods("POST")
	router.HandleFunc("/proxy/qr/{qrID}", proxyurl.ProxyUrl).Methods("GET")
	router.HandleFunc("/user/profile", Middleware.AuthMiddleWare(DashboardService.Profile)).Methods("GET")
	router.HandleFunc("/user/qr", Middleware.AuthMiddleWare(QrService.GetQrsByOffest)).Methods("GET")
	router.HandleFunc("/qr", Middleware.AuthMiddleWare(QrService.GetQrById)).Methods("GET")
	router.HandleFunc("/create-qr", Middleware.AuthMiddleWare(QrService.CreateQR)).Methods("POST")
	router.HandleFunc("/update-qr", Middleware.AuthMiddleWare(QrService.UpdateQrByID)).Methods("PUT")

	return router
}
