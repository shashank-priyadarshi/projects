package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/knadh/koanf/v2"

	chrome "github.com/chromedp/chromedp"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	logger "go.uber.org/zap"
)

type config struct {
	url       string
	target    string
	reqMethod string
	username  string
	password  string
	userAgent string
	selectors []string
}

type content struct{}

var k = koanf.New(".")

// Navigate to target page: https://github.com/aldesantis/the-codeless-code
// Go to: the-codeless-code folder
// Go to: en-qi folder
// Go to each file: click on raw, copy all content
// Merge text content of each file into a master file
// Convert file to pdf

func main() {
	os.Setenv("CONFIG", "./config.json")
	c := &config{}
	if err := c.loadConfig(); err != nil {
		fmt.Println(err)
		logger.Error(fmt.Errorf("main(): %v", err))
		os.Exit(1)
	}
	if err := c.downloadHTMLUsingChromeHeadless(); err != nil {
		fmt.Printf("%v: error: %v", time.Now(), err)
	}
	// if resp, err := c.getReq(client); err != nil {
	// 	fmt.Println(err)
	// 	logger.Error(fmt.Errorf("main(): %v", err))
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println(resp)
	// }
}

func (c *config) loadConfig() (err error) {
	if err = k.Load(file.Provider(os.Getenv("CONFIG")), json.Parser()); err != nil {
		err = fmt.Errorf("loadConfig(): %v", err)
		return
	}
	c.url = k.String("url")
	c.reqMethod = k.String("reqMethod")
	c.username = k.String("username")
	c.password = k.String("password")
	c.userAgent = k.String("userAgent")
	c.target = k.String("target")
	c.selectors = k.Strings("selectors")
	return
}

//	func (c *config) getReq(client *http.Client) (*http.Response, error) {
//		if req, err := http.NewRequest(c.reqMethod, c.url, nil); err != nil {
//			err = fmt.Errorf("getReq(): %v", err)
//			return &http.Response{}, err
//		} else {
//			req.Header.Set("User-Agent", c.userAgent)
//			if resp, err := client.Do(req); err != nil {
//				err = fmt.Errorf("getReq(): %v", err)
//				return &http.Response{}, err
//			} else {
//				return resp, nil
//			}
//		}
//	}
//
//	func createHTTPClient() *http.Client {
//		return &http.Client{
//			Timeout: time.Second * 5,
//		}
//	}

func (c *config) downloadHTMLUsingChromeHeadless() (err error) {
	ctx, _ := chrome.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	folderSel := c.selectors[0]
	enqiFolderSel := c.selectors[1]

	if err = chrome.Run(ctx, chrome.Navigate(c.url+c.target)); err != nil {
		return fmt.Errorf("%v: error navigating to %v: %v", time.Now(), c.url+c.target, err)
	}

	if err = chrome.Run(ctx, chrome.WaitVisible(folderSel)); err != nil {
		return fmt.Errorf("%v: error XPath to %v: %v", time.Now(), folderSel, err)
	}

	if err = chrome.Run(ctx, chrome.Click(".js-navigation-open .Link--primary")); err != nil {
		return fmt.Errorf("%v: error clicking on button", time.Now(), folderSel, err)
	}

	if err = chrome.Run(ctx, chrome.WaitVisible(enqiFolderSel), chrome.Click(enqiFolderSel)); err != nil {
		return fmt.Errorf("%v: error navigating to %v: %v", time.Now(), enqiFolderSel, err)
	}

	//copyRawTxtSel := c.selectors[3]
	for i := 2; i < 5; i++ {
		txtFolderSel := fmt.Sprintf("%s%d%s", c.selectors[2], i, c.selectors[3])
		if err = chrome.Run(ctx, chrome.Click(txtFolderSel)); err != nil {
			//if err := chrome.Run(ctx, chrome.Click(txtFolderSel), chrome.WaitVisible(copyRawTxtSel)); err != nil {
			fmt.Printf("%v: second html: %v", time.Now(), err)
			continue
		}
		var res string
		chrome.Text(".Box-sc-g0xbh4-0 TCenl", &res, chrome.NodeVisible)
		fmt.Printf("%v: now in: %s", time.Now(), txtFolderSel)
		time.Sleep(time.Second * 2)
		//fmt.Println(chrome.Click(copyRawTxtSel))
	}
	return
}

func (c *content) parseHTML() {}

func downloadFiles() {}

func mergeFiles() {}

func convertToPDF() {}
