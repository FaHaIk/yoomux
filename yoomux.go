package yoomux

import (
	"net/http"
)

type Yoomux struct {
	mux        *http.ServeMux
	middleware []func(http.Handler) http.Handler
}

func (yoomux *Yoomux) Use(middleware func(http.Handler) http.Handler) *Yoomux {
	return &Yoomux{
		mux:        yoomux.mux,
		middleware: append(yoomux.middleware, middleware),
	}
}

func (yoomux *Yoomux) UseGlobal(middleware func(http.Handler) http.Handler) *Yoomux {
	yoomux.middleware = append(yoomux.middleware, middleware)
	return yoomux
}

func (yoomux *Yoomux) Get(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodGet+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Head(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodHead+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Post(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodPost+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Put(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodPut+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Patch(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodPatch+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Delete(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodDelete+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Connect(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodConnect+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Options(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodOptions+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) Trace(path string, handler func(http.ResponseWriter, *http.Request)) *Yoomux {
	yoomux.mux.HandleFunc(http.MethodTrace+" "+path, yoomux.applyMiddleware(handler))
	return yoomux
}

func (yoomux *Yoomux) applyMiddleware(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var h http.Handler = http.HandlerFunc(handler)
		for i := len(yoomux.middleware) - 1; i >= 0; i-- {
			h = yoomux.middleware[i](h)
		}
		h.ServeHTTP(w, r)
	}
}

func (yoomux *Yoomux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	yoomux.mux.ServeHTTP(w, r)
}

func NewYoomux() *Yoomux {
	return &Yoomux{
		mux:        http.NewServeMux(),
		middleware: make([]func(http.Handler) http.Handler, 0),
	}
}

func (yoomux *Yoomux) Subrouter(path string) *Yoomux {
	submux := http.NewServeMux()
	yoomux.mux.Handle(path+"/", http.StripPrefix(path, submux))
	return &Yoomux{
		mux:        submux,
		middleware: yoomux.middleware,
	}
}

// func MiddlewareOne(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Print("Executing middlewareOne")
// 		ctx := context.WithValue(r.Context(), "user", "fideloper")
// 		newReq := r.WithContext(ctx)
// 		next.ServeHTTP(w, newReq)
// 		log.Print("Executing middlewareOne again")
// 	})
// }

// func MiddlewareTwo(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Print("Executing middlewareTwo")
// 		if r.URL.Path == "/foo" {
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 		log.Print("Executing middlewareTwo again")
// 	})
// }

// func main() {
// 	yoomux := NewYoomux()

// 	// yoomux.Use(MiddlewareOne).Use(MiddlewareTwo).Delete("/dd", func(w http.ResponseWriter, r *http.Request) {
// 	// 	// Handler logic for GET /dd
// 	// 	fmt.Println("dd")
// 	// 	w.Write([]byte("dd"))
// 	// }).Post("/kk", func(w http.ResponseWriter, r *http.Request) {
// 	// 	// Handler logic for POST /dd
// 	// 	fmt.Println("post kk")
// 	// 	w.Write([]byte("post kk"))
// 	// })

// 	// global middleware and routes
// 	yoomux.UseGlobal(MiddlewareOne)

// 	yoomux.Get("/kiki", func(w http.ResponseWriter, r *http.Request) {
// 		// Handler logic for GET /dd
// 		fmt.Println("kiki")
// 		fmt.Printf("HERE IS YOUR USER: %s\n", r.Context().Value("user"))

// 		w.Write([]byte("kiki"))
// 	})

// 	apiRouter := yoomux.Subrouter("/api").Use(MiddlewareOne)

// 	apiRouter.Get("/users", func(w http.ResponseWriter, r *http.Request) {
// 		// Handler logic for GET /api/users
// 		fmt.Println("Get users")
// 		w.Write([]byte("Get users"))
// 	})

// 	apiRouter.Use(MiddlewareTwo).Post("/users", func(w http.ResponseWriter, r *http.Request) {
// 		// Handler logic for POST /api/users
// 		fmt.Println("Create user")
// 		w.Write([]byte("Create user"))
// 	})

// 	http.ListenAndServe(":3000", yoomux)
// }
