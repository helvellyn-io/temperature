package main

import (
	"demo/src/api"
	"demo/src/tests"
	"flag"
	"fmt"
	"net/http"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start()
	defer tracer.Stop()
	runOptions := flag.String("mode", "none", "--mode app runs the main app, --mode test runs the tests")
	flag.Parse()

	switch *runOptions {
	case "app":
		App()
	case "test":
		Tests()
	}

}

func Tests() {
	tracer.Start()

	defer tracer.Stop()
	t := tests.ReturnResults()
	r := fmt.Sprintf("test returned HTTP STATUS %v", t)
	print(r)
}

func App() {

	tracer.Start()
	defer tracer.Stop()
	mux := httptrace.NewServeMux()

	data := TempBoulder()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(data))

	})

	http.ListenAndServe(":8080", mux)

}

func TempBoulder() string {

	weather := new(api.Weather)
	api.GetJson("https://api.openweathermap.org/data/2.5/weather?q=Boulder&appid=050224087da57dabdc13099b40e747e0&units=imperial", weather)
	data := fmt.Sprintf("Location: %v Temp: %v F \n", weather.Name, weather.Main.Temp)

	return data
}
