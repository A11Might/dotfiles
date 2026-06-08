package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	authToken string
	apiURL    string
)

func init() {
	authToken = "Bearer " + os.Getenv("READECK_API_TOKEN")
	apiURL = os.Getenv("READECK_API_URL")
}

type Bookmark struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	SiteName    string   `json:"site_name"`
	Description string   `json:"description"`
	ReadingTime int      `json:"reading_time"`
	IsArchived  bool     `json:"is_archived"`
	Labels      []string `json:"labels"`
}

type BookmarkDetail struct {
	Resources struct {
		Article struct {
			Src string `json:"src"`
		} `json:"article"`
	} `json:"resources"`
}

func main() {
	if authToken == "Bearer " {
		log.Fatal("请设置环境变量 READECK_API_TOKEN")
	}
	if apiURL == "" {
		log.Fatal("请设置环境变量 READECK_API_URL")
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "archive":
			if len(os.Args) < 3 {
				log.Fatal("用法: main.go archive <id>")
			}
			archiveBookmark(os.Args[2])
			fmt.Println("已归档")
		case "skip":
			if len(os.Args) < 3 {
				log.Fatal("用法: main.go skip <id>")
			}
			skipAndArchive(os.Args[2])
			fmt.Println("已标记不感兴趣并归档")
		default:
			log.Fatalf("未知命令: %s (可用: archive, skip)", os.Args[1])
		}
		return
	}

	// 默认: 获取随机未归档书签并输出
	total := fetchTotalCount()
	b := fetchRandomBookmark(total)
	detail := fetchBookmarkDetail(b.ID)

	fmt.Printf("ID: %s\n", b.ID)
	fmt.Printf("标题: %s\n", b.Title)
	fmt.Printf("URL: %s\n", b.URL)
	if b.SiteName != "" {
		fmt.Printf("来源: %s\n", b.SiteName)
	}
	fmt.Printf("预计阅读: %d 分钟\n", b.ReadingTime)
	if len(b.Labels) > 0 {
		fmt.Printf("标签: %s\n", strings.Join(b.Labels, ", "))
	}
	if b.Description != "" {
		fmt.Printf("描述: %s\n", b.Description)
	}
	fmt.Println("---")

	// 通过 detail 中的 resources.article.src 获取文章内容
	if detail.Resources.Article.Src != "" {
		article := fetchArticleBySrc(detail.Resources.Article.Src)
		fmt.Println(article)
	} else {
		fmt.Println("（该文章没有可用的文本内容）")
	}
}

func fetchTotalCount() int {
	req, _ := http.NewRequest(http.MethodGet, apiURL+"/bookmarks?is_archived=false&limit=1", nil)
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("API 返回错误: HTTP %d, %s", resp.StatusCode, string(body))
	}

	totalStr := resp.Header.Get("Total-Count")
	total, err := strconv.Atoi(totalStr)
	if err != nil || total == 0 {
		log.Fatalf("没有找到未归档的书签 (Total-Count: %q)", totalStr)
	}
	return total
}

func fetchRandomBookmark(total int) *Bookmark {
	offset := rand.Intn(total)
	req, _ := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/bookmarks?is_archived=false&limit=1&offset=%d", apiURL, offset), nil)
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	var bookmarks []Bookmark
	if err := json.NewDecoder(resp.Body).Decode(&bookmarks); err != nil {
		log.Fatalf("解析失败: %v", err)
	}
	if len(bookmarks) == 0 {
		log.Fatal("未获取到书签")
	}
	return &bookmarks[0]
}

func fetchBookmarkDetail(id string) *BookmarkDetail {
	req, _ := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/bookmarks/%s", apiURL, id), nil)
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return &BookmarkDetail{}
	}
	defer resp.Body.Close()

	var detail BookmarkDetail
	json.NewDecoder(resp.Body).Decode(&detail)
	return &detail
}

func fetchArticleBySrc(src string) string {
	req, _ := http.NewRequest(http.MethodGet, src, nil)
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return "（无法获取文章内容）"
	}
	defer resp.Body.Close()

	htmlBytes, _ := io.ReadAll(resp.Body)
	text := stripHTMLTags(string(htmlBytes))

	const maxChars = 8000
	if utf8.RuneCountInString(text) > maxChars {
		text = string([]rune(text)[:maxChars]) + "..."
	}
	return text
}

func skipAndArchive(id string) {
	body, _ := json.Marshal(map[string]any{
		"add_labels":  []string{"不感兴趣"},
		"is_archived": true,
	})
	req, _ := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("%s/bookmarks/%s", apiURL, id), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("操作失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		log.Fatalf("操作失败: HTTP %d, %s", resp.StatusCode, string(respBody))
	}
}

func archiveBookmark(id string) {
	body, _ := json.Marshal(map[string]bool{"is_archived": true})
	req, _ := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("%s/bookmarks/%s", apiURL, id), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("归档失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		log.Fatalf("归档失败: HTTP %d, %s", resp.StatusCode, string(respBody))
	}
}

func stripHTMLTags(html string) string {
	text := regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`).ReplaceAllString(html, " ")
	text = regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	text = strings.ReplaceAll(text, "&nbsp;", " ")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	return strings.TrimSpace(text)
}

