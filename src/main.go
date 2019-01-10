package main

import (
	"flag"
	"fmt"
	"os"

	API "./transapi"
)

const __version__ = "0.0.1"

var (
	filename   string
	configfile string
	logON      bool
	outputfile string
)

func init() {
	flag.StringVar(&filename, "f", "none", "file name")
	flag.StringVar(&configfile, "c", "./conf.json", "config file name")
	flag.BoolVar(&logON, "l", false, "show result on stdout")
	flag.StringVar(&outputfile, "o", "./translate.txt", "output file name")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "trans-cli version: trans-cli/%v\nUsage: trans [-f filename] [-c configfile] [-l logres] [-o outputfile]\n", __version__)
		flag.PrintDefaults()
	}
}

func trans(text string, api API.TransAPI) (content string) {
	segs := segment(text)
	content = ""

	for _, val := range segs {
		res := api.TransTo("en", val)

		if logON {
			fmt.Println(val)
			fmt.Println(res)
		}

		if val != "" {
			content += val + "\n"
			content += res + "\n\n"
		}

	}
	return
}

func main() {
	flag.Parse()

	if filename == "none" {
		fmt.Println("[error] pls select src file.")
		return
	}

	text := readf(filename)
	conf := loadConf(configfile)

	bt := API.NewBaiduTrans(conf.Baidu.URL, conf.Baidu.AppID, conf.Baidu.Key)

	outText := trans(text, bt)

	savef(outText, outputfile)
}
