package frp

import (
	"github.com/PuerkitoBio/goquery"
	"go-frpc/utils"
	"go-frpc/utils/cmd"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//检查frp状态，未启动返回0，运行中返回其pid
func CheckStatus() int {
	pid, _ := cmd.CMDFactory{}.Generate().GetPID("frp")
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
	out, _ := cmd.CMDFactory{}.Generate().RunCommand("taskkill /f /im frpc.exe")
	if out == "" {
		println("frp 未启动")
	}
	utils.Log.Info(out)
	println("frpc closed...")
}

// 下载并解压frp
func Download(fb func(length, downLen int64)) {
	workDir, _ := filepath.Abs("")
	utils.Log.Infof("frp download start into %s\n", workDir)

	frpName, frpUrl := LatestUrlOnOs()
	frpDownloadPath := workDir + "/" + frpName
	utils.Log.Infof("frp download path : %s", frpDownloadPath)
	utils.DownloadProcess(frpUrl, frpDownloadPath, fb)
	utils.Log.Info("frp download end...")

	//utils.Log.Infof("frp unzip start into %s", workDir)
	utils.UnCompress(frpDownloadPath, workDir+"/frpc", true)
	//utils.Log.Info("frp.zip end unzip...")
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

func LatestUrlOnOs() (string, string) {
	baseUrl := "https://github.com/fatedier/frp/tags"
	latestUrl := LatestTag(baseUrl)
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
	fileName := ""
	// Find the review items
	items := doc.Find("div[class='d-flex flex-justify-between flex-items-center py-1 py-md-2 Box-body px-2']")
	items.Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := strings.TrimSpace(s.Find("a").Text())
		href, _ := s.Find("a").Attr("href")
		//fmt.Printf("Review %d: %s - %s\n", i, band, "https://github.com"+href)
		if strings.Contains(href, osname) && strings.Contains(href, ostype) {
			fileName = band
			resurl = href
		}
	})
	//println(runtime.GOOS)
	//println(runtime.GOARCH)
	//println(strconv.IntSize)
	return fileName, "https://github.com" + resurl
}
