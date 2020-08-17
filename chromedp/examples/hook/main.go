package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/chromedp/cdproto"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"net/http"
	"time"
)

var (
	msgChann = make(chan cdproto.Message)
	ctxt     context.Context
)

var flagPort = flag.Int("port", 8544, "port")

func devToolHandler(s string, is ...interface{}) {
	/*
	 Uncomment the following line to have a log of the events
	 log.Printf(s, is...)
	*/
	/*
	 We need this to be on a separate gorutine
	 otherwise we block the browser and we don't receive messages
	*/
	go func() {

		for _, elem := range is {
			var msg cdproto.Message
			// The CDP messages are sent as strings so we need to convert them back
			json.Unmarshal([]byte(fmt.Sprintf("%s", elem)), &msg)
			msgChann <- msg
		}
	}()
}

func main() {
	flag.Parse()
	// start cookie server
	go cookieServer(fmt.Sprintf(":%d", *flagPort))
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	chromedp.WithLogf(devToolHandler)
	err := chromedp.Run(ctx, hooks(
		fmt.Sprintf("http://localhost:%d", *flagPort),
	))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1000)
}

func hooks(host string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// This is in an action function because the Executor is needed in some method
			go func() {
				for {
					/*
						Right now I have no guaranties (did not run enough tests) that the messages are evaluated
						in the same order they are received
					*/
					msg := <-msgChann

					switch msg.Method.String() {
					case "Network.requestWillBeSent":
						var reqWillSend network.EventRequestWillBeSent
						json.Unmarshal(msg.Params, &reqWillSend)
						fmt.Println(reqWillSend.Request.URL)

					case "Network.responseReceived":
						var respevent network.EventResponseReceived
						// json.Unmarshal(msg.Params, &respevent)
						// This contains a bunch of data, like the response headers
						fmt.Println(respevent.Response)
						/*
							Uncomment the following lines if you want to print the response body
							rbp := network.GetResponseBody(respevent.RequestID)
							b, e := rbp.Do(ctxt, h)
							if e != nil {
								fmt.Println(e)
								continue
							}
							fmt.Printf("%s\n", b)
						*/

					default:
						continue
					}
				}
			}()
			return nil
		}),
		// navigate to site
		chromedp.Navigate(host),
		// read the returned values
		//chromedp.Text(`#result`, res, chromedp.ByID, chromedp.NodeVisible),
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}

// cookieServer creates a simple HTTP server that logs any passed cookies.
func cookieServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookies := req.Cookies()
		for i, cookie := range cookies {
			log.Printf("from %s, server received cookie %d: %v", req.RemoteAddr, i, cookie)
		}
		buf, err := json.MarshalIndent(req.Cookies(), "", "  ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(res, indexHTML, string(buf))
	})
	return http.ListenAndServe(addr, mux)
}

const (
	indexHTML = `<!doctype html>
<html>
<body>
  <div id="result">%s</div>
</body>
</html>`
)
