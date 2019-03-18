package app

import (
	"fmt"
	"net/http"

	"github.com/nexus-uw/paisley/app/models"

	rgorp "github.com/revel/modules/orm/gorp/app"
	"github.com/revel/revel"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v2"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	revel.AddInitEventHandler(func(event revel.Event, i interface{}) revel.EventResponse {
		switch event {
		case revel.ENGINE_BEFORE_INITIALIZED:

			if revel.RunMode == "dev-fast" {
				revel.AddHTTPMux("/this/is/a/test", fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
					fmt.Fprintln(ctx, "Hi there, it worked", string(ctx.Path()))
					ctx.SetStatusCode(200)
				}))
				revel.AddHTTPMux("/this/is/", fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
					fmt.Fprintln(ctx, "Hi there, shorter prefix", string(ctx.Path()))
					ctx.SetStatusCode(200)
				}))
			} else {
				revel.AddHTTPMux("/this/is/a/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "Hi there, it worked", r.URL.Path)
					w.WriteHeader(200)
				}))
				revel.AddHTTPMux("/this/is/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "Hi there, shorter prefix", r.URL.Path)
					w.WriteHeader(200)
				}))

			}
		}
		return 0
	})

	revel.OnAppStart(func() {
		revel.AppLog.Debug("SIMON TEST MSG TABLES  BE HERE")

		Dbm := rgorp.Db.Map
		setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
			for col, size := range colSizes {
				t.ColMap(col).MaxSize = size
			}
		}
		t := Dbm.AddTableWithName(models.User{}, "users").SetKeys(false, "UserId")
		t.ColMap("Password").Transient = true
		setColumnSizes(t, map[string]int{
			"Username": 20,
			"Name":     100,
		})

		t = Dbm.AddTableWithName(models.Subscription{}, "subscriptions").SetKeys(false, "SubscriptionID")
		setColumnSizes(t, map[string]int{
			"Subredit": 50,
		})

		rgorp.Db.TraceOn(revel.AppLog)
		Dbm.CreateTablesIfNotExists()

		count, _ := rgorp.Db.SelectInt(rgorp.Db.SqlStatementBuilder.Select("count(*)").From("User"))
		if count < 1 {
			bcryptPassword, _ := bcrypt.GenerateFromPassword(
				[]byte("demo"), bcrypt.DefaultCost)
			demoUser := &models.User{0, "Demo User", "demo", "demo", bcryptPassword}
			if err := Dbm.Insert(demoUser); err != nil {
				panic(err)
			}
		}

		revel.AppLog.Debug("SIMON TEST MSG TABLES SHOUL BE HERE")

	}, 5)
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
