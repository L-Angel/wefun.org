package controllers

import (
	"github.com/astaxie/beego"
        "github.com/slene/blackfriday"
        "io/ioutil"
        "os"
)

type MarkdownController struct {
	baseController
}

func (c *MarkdownController) Get() {
	c.Data["Website"] = "www.wefun.org"
	c.Data["Email"] = "Dean.Wefun@gmail.com"
        c.Data["IsMarkdown"] = true
	c.TplName = "Markdown.tpl"

        link := c.Ctx.Request.URL.Path
        beego.Info(link)
        // beego.Info(c.Ctx.Request.URL.Path)
        str, _ := findFile(link)
        body := Md2html(str)
        //      c.Ctx.Output.Status = 401                                    //设置状态码
        //      c.Ctx.Output.Header("WWW-Authenticate", "Basic realm=zenyu") //设置请求头部
        c.Data["MarkdownContent"] = body
}

func findFile(path string) ([]byte, error) {
        path = "./" + path
        beego.Info("=====", path)
        f, err := os.Open(path)
        if err != nil {
                panic(err)
        }
        defer f.Close()
        return ioutil.ReadAll(f)
}



func Md2html(raw []byte) string {
        htmlFlags := 0
        //htmlFlags |= blackfriday.HTML_USE_XHTML
        htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
        htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
        htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
        htmlFlags |= blackfriday.HTML_GITHUB_BLOCKCODE
        htmlFlags |= blackfriday.HTML_OMIT_CONTENTS
        //htmlFlags |= blackfriday.HTML_COMPLETE_PAGE
        renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

        // set up the parser
        extensions := 0
        extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
        extensions |= blackfriday.EXTENSION_TABLES
        extensions |= blackfriday.EXTENSION_FENCED_CODE
        extensions |= blackfriday.EXTENSION_AUTOLINK
        extensions |= blackfriday.EXTENSION_STRIKETHROUGH
        extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK
        extensions |= blackfriday.EXTENSION_SPACE_HEADERS
        extensions |= blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

        body := blackfriday.Markdown(raw, renderer, extensions)
        return string(body)
}

