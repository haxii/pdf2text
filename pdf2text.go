package main

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func convertPDF(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != "POST" {
		ctx.Error("post only", 405)
		return
	}
	if ctx.PostBody() == nil {
		ctx.Error("Body is missing", 400)
		return
	}
	tmpFileName := fmt.Sprintf("/tmp/pdf2text-%s.tmp", xid.New().String())
	query := ctx.QueryArgs()
	x := string(query.Peek("x"))
	y := string(query.Peek("y"))
	w := string(query.Peek("w"))
	h := string(query.Peek("h"))

	bodyBytes := ctx.PostBody()
	err := ioutil.WriteFile(tmpFileName, bodyBytes, 0600)
	if err != nil {
		ctx.Error("Failed to open the file for writing", 500)
		return
	}
	defer func() {
		_ = os.Remove(tmpFileName)
	}()

	args := make([]string, 0)
	if len(x) > 0 {
		args = append(args, "-x", x)
	}
	if len(y) > 0 {
		args = append(args, "-y", y)
	}
	if len(w) > 0 {
		args = append(args, "-W", w)
	}
	if len(h) > 0 {
		args = append(args, "-H", h)
	}
	args = append(args, "-nopgbrk", "-enc", "UTF-8", tmpFileName, "-")

	body, err := exec.Command("pdftotext", args...).Output()
	if err != nil {
		log.Printf("pdftotext error: %s", err)
		ctx.Error(err.Error(), 500)
	}
	_, _ = fmt.Fprintf(ctx, string(body))
}

func main() {
	h := convertPDF
	if err := fasthttp.ListenAndServe(":5000", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
