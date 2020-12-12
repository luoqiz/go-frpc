package frp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-frpc/src/utils"
	"go-frpc/src/utils/cmd"
	"go-frpc/src/utils/cmd/linux"
	"go-frpc/src/utils/cmd/windows"

	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

const githubDomain = "https://github.com"

//检查frp状态，未启动返回0，运行中返回其pid
func CheckStatus() int {
	pid, _ := cmd.CMDFactory{}.Generate().GetPID("frpc")
	fmt.Println("****************************", pid)
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

	client := cmd.CMDFactory{}.Generate()
	if (client == linux.Linux{}) {
		client.RunCommandBg("nohup " + frpc + "/frpc -c " + workspace + "/frpc.ini &")
	} else if (client == windows.Windows{}) {
		client.RunCommandBg(frpc + "/frpc.exe -c " + workspace + "/frpc.ini")
	}
	utils.Log.Info("frpc start success...")
	return CheckStatus()
}

// 关闭frpc
func Stop() {
	out := ""
	client := cmd.CMDFactory{}.Generate()
	if (client == linux.Linux{}) {
		out, _ = client.RunCommand("kill -9 $(pidof frpc)")
	} else {
		out, _ = client.RunCommand("taskkill /f /im frpc.exe")
	}

	utils.Log.Info("frpc close success...")
	if out == "" {
		println("frp 未启动")
	}
	utils.Log.Info(out)
}

// 下载并解压frp
func Download(fb func(length, downLen int64)) {

	workDir, _ := filepath.Abs("")
	utils.Log.Infof("frp download start into %s\n", workDir)

	frpName, frpUrl := ValidUrlOnOs()
	frpDownloadPath := filepath.Join(workDir, frpName)
	utils.Log.Infof("frp download path : %s", frpDownloadPath)
	utils.DownloadProcess(frpUrl, frpDownloadPath, fb)
	utils.Log.Info("frp download end...")

	// 解压之前先删除已解压文件
	utils.Log.Infof("frp unzip start into %s", workDir)
	utils.DeleteDir(workDir + "/frpc")
	utils.UnCompress(frpDownloadPath, workDir+"/frpc", true)
	utils.DeleteFile(frpDownloadPath)
	utils.Log.Info("frp.zip end unzip...")
}

// 获取最新tag
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

	return githubDomain + href

}

// 根据系统判定最新版本
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
	return fileName, githubDomain + resurl
}

func ValidUrlOnOs() (string, string) {
	baseUrl := "https://github.com/fatedier/frp/tags"
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

	resFileName := ""
	resFileUrlPath := ""

	items := doc.Find("h4.commit-title")
	// 遍历有效url
	items.EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, isValid := s.Find("a").Attr("href")
		if isValid {
			fileName, fileUrlPath := TagUrlValid(githubDomain + href)
			if fileName != "" {
				resFileName = fileName
				resFileUrlPath = fileUrlPath
				return false
			}
		}
		return true
	})
	return resFileName, resFileUrlPath
}

// 根据系统判定最新版本
func TagUrlValid(tagUrl string) (string, string) {
	// Request the HTML page.
	res, err := http.Get(tagUrl)
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
	return fileName, githubDomain + resurl
}
