package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	HookRequest("http://binghuan.github.io/debug/testm3u8withkey.html", func(url string) bool {
		return true
	})
}

type Hooker struct {
	ReqHooks    map[string]interface{}
	RespHooks   map[string]interface{}
	MatchString func(url string) bool
}

func HookRequest(url string, matchString func(url string) bool) {
	hooker := &Hooker{
		MatchString: matchString,
		ReqHooks:    make(map[string]interface{}),
		RespHooks:   make(map[string]interface{}),
	}
	hooker.Hooking(url)
}

func (h *Hooker) Hooking(url string) {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "50,400"),
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout
	taskCtx, cancel = context.WithTimeout(taskCtx, 10*time.Second)
	defer cancel()

	// ensure that the browser process is started
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}

	// listen network event
	listenForNetworkEvent(taskCtx, h)

	chromedp.Run(taskCtx,
		network.Enable(),
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.BySearch),
	)
	log.Print(h.RespHooks)
	log.Print(h.RespHooks)
}

func listenForNetworkEvent(ctx context.Context, hooker *Hooker) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			req := ev.Request
			match := hooker.MatchString(req.URL)
			if match {
				hooker.ReqHooks[req.URL] = req.Headers
			}
			//if len(req.Headers) != 0 {
			//	log.Printf("received req headers: %s", req.Headers)
			//	log.Printf("received req url: %s", req.URL)
			//}
		case *network.EventResponseReceived:
			resp := ev.Response
			match := hooker.MatchString(resp.URL)
			if match {
				hooker.RespHooks[resp.URL] = resp.Headers
			}
			//if len(resp.Headers) != 0 {
			//	log.Printf("received resp headers: %s", resp.Headers)
			//	log.Printf("received resp url: %s", resp.URL)
			//	log.Printf("received resp protocol: %s", resp.Protocol)
			//}
		}
	})
}
