package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	const html = `
<html>
    <head>
        <script src="https://unpkg.com/vue@next"></script>
    </head>
    <body>
        <div id="app">
        </div>
        <script src="js/app" ></script>
    </body>
</html>
`
	rw.Write([]byte(html))
}

func JavaScript(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/javascript")
	vars := mux.Vars(r)

	switch vars["file"] {
	case "app":
		const js = `
const App = {
    data() {
        return {
        counter: 0
        }
    },
    template: "<log-portal/>"
}
const app = Vue.createApp(App)
app.component('log-portal', {
    template: "<div>ok</div>"
})
 
app.mount('#app')
`
		rw.Write([]byte(js))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/js/{file}", JavaScript)
	address := "0.0.0.0:8000"
	fmt.Printf("Starting http server on %s\n", address)
	http.ListenAndServe(address, r)
}
