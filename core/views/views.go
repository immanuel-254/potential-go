package views

import (
	"fmt"
	"net/http"

	"github.com/immanuel-254/potential-go/core/database"
	"github.com/immanuel-254/potential-go/core/models"
)

const UserRouteGroup = "/user"

type View struct {
	Route       string
	Middlewares []func(http.Handler) http.Handler
	Handler     http.Handler
}

var (
	UserCreateView = View{
		Route: fmt.Sprintf("%s/create", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data map[string]string
			err := GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.Email, user.Password = data["email"], data["password"]
			err = user.UserCreate(database.DB, data["confirm_password"], w, r)
			if err != nil {
				return
			}
		}),
	}

	UserReadView = View{
		Route: fmt.Sprintf("%s/read/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/read/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			err = user.UserRead(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserReadEmailView = View{
		Route: fmt.Sprintf("%s/read-email", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data map[string]string
			err := GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.Email = data["email"]
			err = user.UserReadByEmail(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserListView = View{
		Route: fmt.Sprintf("%s/list", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var user models.User
			err := user.UserList(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserUpdateEmailView = View{
		Route: fmt.Sprintf("%s/update-email/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/update-email/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var data map[string]string
			err = GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			user.Email = data["email"]
			err = user.UserUpdateEmail(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserUpdatePasswordView = View{
		Route: fmt.Sprintf("%s/update-password/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/update-password/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var data map[string]string
			err = GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			user.Password = data["password"]
			err = user.UserUpdatePassword(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserUpdateActiveView = View{
		Route: fmt.Sprintf("%s/update-active/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/update-active/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var data map[string]string
			err = GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			if data["active"] == "true" {
				user.Active = true
			}
			if data["active"] == "false" {
				user.Active = false
			}
			err = user.UserUpdateActive(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserUpdateAdminView = View{
		Route: fmt.Sprintf("%s/update-admin/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/update-admin/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var data map[string]string
			err = GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			if data["admin"] == "true" {
				user.Admin = true
			}
			if data["admin"] == "false" {
				user.Admin = false
			}
			err = user.UserUpdateAdmin(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserUpdateStaffView = View{
		Route: fmt.Sprintf("%s/update-staff/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/update-staff/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var data map[string]string
			err = GetData(&data, w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			if data["staff"] == "true" {
				user.Staff = true
			}
			if data["staff"] == "false" {
				user.Staff = false
			}
			err = user.UserUpdateStaff(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserDeleteView = View{
		Route: fmt.Sprintf("%s/delete/", UserRouteGroup),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, err := GetId(fmt.Sprintf("%s/delete/", UserRouteGroup), w, r)
			if err != nil {
				return
			}

			var user models.User
			user.ID = int64(id)
			err = user.UserDelete(database.DB, w, r)
			if err != nil {
				return
			}
		}),
	}

	UserViews = []View{
		UserCreateView,
		UserDeleteView,
		UserListView,
		UserReadEmailView,
		UserReadView,
		UserUpdateActiveView,
		UserUpdateAdminView,
		UserUpdateEmailView,
		UserUpdatePasswordView,
		UserUpdateStaffView,
	}
)

// Middleware chaining
func chainMiddlewares(handler http.Handler, middlewares []func(http.Handler) http.Handler) http.Handler {
	if len(middlewares) != 0 {
		for i := 0; i < len(middlewares); i++ { // Apply middlewares in normal order
			handler = middlewares[i](handler)
		}
		return handler
	}
	return handler
}

// Routes function
func Routes(mux *http.ServeMux, views []View) {
	for _, view := range views {
		handlerWithMiddlewares := chainMiddlewares(view.Handler, view.Middlewares)
		mux.HandleFunc(view.Route, func(w http.ResponseWriter, r *http.Request) {
			handlerWithMiddlewares.ServeHTTP(w, r)
		})

	}
}
