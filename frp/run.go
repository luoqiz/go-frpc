package frp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-frpc/utils"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

//检查frp状态，未启动返回0，运行中返回其pid
func CheckStatus() int {
	out, _ := utils.ExecCmd("cmd.exe", "/c tasklist | findstr frp")

	if out == "" {
		println("frp 未启动")
		return 0
	}
	utils.Log.Info(out)
	re, _ := regexp.Compile(`\d+`)
	//查找符合正则的第一个
	all := re.FindAll([]byte(out), -1)
	pid := 0
	for index, item := range all {
		if index == 0 {
			pid, _ = strconv.Atoi(string(item))
			break
		}

	}
	return pid
}

//开启frpc
func Start() int {
	workspace, _ := filepath.Abs("")
	dir := utils.GetDirectory(workspace + "/frpc")
	if len(dir) < 1 {
		return 0
	}
	frpc := dir[0]
	frpStatus := exec.Command(frpc+"/frpc.exe", "-c", frpc+"/frpc.ini")
	frpStatus.Start()

	utils.Log.Info("frpc start success...")
	return CheckStatus()
}

// 关闭frpc
func Stop() {
	//pid := CheckStatus()
	out, _ := utils.ExecCmd("cmd.exe", "/c taskkill /f /im frpc.exe")
	if out == "" {
		println("frp 未启动")
	}
	utils.Log.Info(out)
	println("frpc closed...")
}

// 下载frp
func Download(fb func(length, downLen int64)) {
	utils.Log.Info("frp download start...")
	workDir, _ := filepath.Abs("")
	utils.Log.Debug("frp download path: " + workDir)
	utils.DownloadProcess(LatestUrlOnOs(), workDir+"/frp.zip", fb)
	utils.Log.Info("frp download end...")

	utils.Log.Info("frp.zip start unzip...")
	utils.Log.Debug("frp unzip path:" + workDir)
	utils.UnZip(workDir+"/frp.zip", workDir+"/frpc", true)
	utils.Log.Info("frp.zip end unzip...")
}

func LatestTag(baseUrl string) string {
	// Request the HTML page.
	res, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	items := doc.Find("h4.commit-title")

	href, _ := items.First().Find("a").Attr("href")

	return "https://github.com" + href

}

func LatestUrlOnOs() string {
	baseUrl := "https://github.com/fatedier/frp/tags"
	latestUrl := LatestTag(baseUrl)
	println("------------")
	// Request the HTML page.
	res, err := http.Get(latestUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	osname := runtime.GOOS
	ostype := runtime.GOARCH
	resurl := ""
	// Find the review items
	items := doc.Find("div[class='d-flex flex-justify-between flex-items-center py-1 py-md-2 Box-body px-2']")
	println(len(items.Nodes))
	items.Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		href, _ := s.Find("a").Attr("href")
		fmt.Printf("Review %d: %s - %s\n", i, band, "https://github.com"+href)
		if strings.Contains(href, osname) && strings.Contains(href, ostype) {
			resurl = href
		}
	})
	//href, _ := items.First().Find("a").Attr("href")

	println(runtime.GOOS)
	println(runtime.GOARCH)
	println(strconv.IntSize)
	return "https://github.com" + resurl
}
