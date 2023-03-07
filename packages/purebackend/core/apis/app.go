package apis

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuremlHQ/PureML/packages/purebackend/core"
	config "github.com/PuremlHQ/PureML/packages/purebackend/core/config"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

var (
	allowedOrigins = []string{"*"}
)

func Serve(app core.App, hideStartBanner bool) error {
	dataDir := app.DataDir()
	httpAddr := config.GetHttpAddr()
	httpsAddr := config.GetHttpsAddr()

	router, err := InitApi(app)
	if err != nil {
		panic(err)
	}

	// configure cors
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// start http server
	// ---
	mainAddr := httpAddr
	if httpsAddr != "" {
		mainAddr = httpsAddr
	}

	mainHost, _, _ := net.SplitHostPort(mainAddr)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(filepath.Join(dataDir, ".autocert_cache")),
		HostPolicy: autocert.HostWhitelist(mainHost, "www."+mainHost),
	}

	serverConfig := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			NextProtos:     []string{acme.ALPNProto},
		},
		ReadTimeout:       5 * time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
		// WriteTimeout: 60 * time.Second, // breaks sse!
		Handler: router,
		Addr:    mainAddr,
	}

	// if showStartBanner {
	schema := "http"
	if httpsAddr != "" {
		schema = "https"
	}
	address := serverConfig.Addr
	if strings.HasPrefix(address, "0.0.0.0") {
		address = strings.Replace(address, "0.0.0.0", "localhost", 1)
	}
	if !hideStartBanner {
		regular := color.New()
		bold := color.New(color.Bold).Add(color.FgGreen)
		bold.Printf("> Server started at: %s\n", color.CyanString("%s://%s", schema, address))
		regular.Printf("  - REST API: %s\n", color.CyanString("%s://%s/api/", schema, address))
		regular.Printf("  - API Docs: %s\n", color.HiGreenString("%s://%s/api/swagger/index.html", schema, address))
		// regular.Printf("  - Admin UI: %s\n", color.CyanString("%s://%s/_/", schema, address))
	}

	var serveErr error
	if httpsAddr != "" {
		// if httpAddr is set, start an HTTP server to redirect the traffic to the HTTPS version
		if httpAddr != "" {
			// nolint:errcheck
			go http.ListenAndServe(httpAddr, certManager.HTTPHandler(nil))
		}

		// start HTTPS server
		serveErr = serverConfig.ListenAndServeTLS("", "")
	} else {
		// start HTTP server
		serveErr = serverConfig.ListenAndServe()
	}

	if serveErr != http.ErrServerClosed {
		log.Fatalln(serveErr)
	}

	return nil
}
