package main

import (
	"NetLikePlate/properties"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func BasicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//makeLogEntry(c,"").Info("incoming request")
		return next(c)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	var reportCode = 200
	reportMessage := ""
	if ok {
		reportCode = report.Code
		reportMessage = fmt.Sprintf("%v", report.Message)
	} else {
		reportCode = http.StatusInternalServerError
		reportMessage = fmt.Sprintf("%v", err.Error())
	}
	//u := uuid.New().String()
	//makeLogEntry(c, u).Error(report.Message)
	err = c.JSON(report.Code, properties.NResponse(fmt.Sprintf(reportMessage), strconv.Itoa(reportCode), nil))
	if err != nil {
		return
	}
}

//func makeLogEntry(c echo.Context, supportId string) *log.Entry {
//	if c == nil {
//		return log.WithFields(log.Fields{
//			"at": time.Now().Format("2006-01-02 15:04:05"),
//		})
//	}
//
//	return log.WithFields(log.Fields{
//		"supportId": supportId,
//		"at":        time.Now().Format("2006-01-02 15:04:05"),
//		"method":    c.Request().Method,
//		"uri":       c.Request().URL.String(),
//		"ip":        c.Request().RemoteAddr,
//	})
//}

//func InitializeLogging(logFile string) {
//	var file, err = OpenFile(logFile, O_APPEND|O_CREATE|O_WRONLY, 0666)
//	if err != nil {
//		fmt.Println("Could Not Open Log File : " + err.Error())
//	}
//	log.SetOutput(file)
//	log.SetFormatter(&log.TextFormatter{})
//	//defer file.Close()
//
//}

func startup() {
	e := echo.New()
	e.Use(middlewareLogging)
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST},
	}))
	Route(e)
	e.Logger.Debug(e.Start(":1234"))
}
