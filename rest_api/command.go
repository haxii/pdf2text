package rest_api

import (
	"fmt"
	"github.com/haxii/pdf2text/rest_api/service"
	"os/exec"
	"strconv"
)

func pdf2Text(param service.PDF2TextParams) ([]byte, error) {
	if param.Pdf == nil {
		return nil, nil
	}
	args := buildArgs(param)
	cmd := exec.Command("pdftotext", args...)
	cmd.Stdin = param.Pdf
	return cmd.Output()
}

func buildArgs(param service.PDF2TextParams) (args []string) {
	args = make([]string, 0)
	for _, arg := range [][]string{
		buildIntArg(param.F, "-f"),
		buildIntArg(param.L, "-l"),
		buildIntArg(param.R, "-r"),
		buildIntArg(param.X, "-x"),
		buildIntArg(param.Y, "-y"),
		buildIntArg(param.W, "-W"),
		buildBoolArg(param.Layout, "-layout"),
		buildIntArg(param.Fixed, "-fixed"),
		buildBoolArg(param.Raw, "-raw"),
		buildBoolArg(param.Nodiag, "-nodiag"),
		buildBoolArg(param.Htmlmeta, "-htmlmeta"),
		buildBoolArg(param.Tsv, "-tsv"),
		buildStrArg(param.Enc, "-enc"),
		buildStrArg(param.Eol, "-eol"),
		buildBoolArg(param.Bbox, "-bbox"),
		buildBoolArg(param.BboxLayout, "-bbox-layout"),
		buildBoolArg(param.Cropbox, "-cropbox"),
		buildFloatArg(param.Colspacing, "-colspacing"),
		buildStrArg(param.Opw, "-opw"),
		buildStrArg(param.Upw, "-upw"),
	} {
		for _, s := range arg {
			args = append(args, s)
		}
	}
	if param.Pgbrk != nil && *param.Pgbrk {
		// this indicates a page break needed
	} else {
		args = append(args, "-nopgbrk")
	}
	return append(args, "-", "-")
}

func buildIntArg(arg *int64, name string) []string {
	if arg == nil {
		return nil
	}
	return []string{name, strconv.FormatInt(*arg, 10)}
}

func buildStrArg(arg *string, name string) []string {
	if arg == nil {
		return nil
	}
	return []string{name, *arg}
}

func buildFloatArg(arg *float64, name string) []string {
	if arg == nil {
		return nil
	}
	return []string{name, fmt.Sprintf("%.1f", *arg)}
}

func buildBoolArg(arg *bool, name string) []string {
	if arg == nil {
		return nil
	}
	if !*arg {
		return nil
	}
	return []string{name}
}
