/**
  @author:panliang
  @data:2021/8/26
  @note
**/
package wordsfilter

import (
	"bufio"
	"github.com/syyongx/go-wordsfilter"
	"log"
	"os"
)

var Wf *wordsfilter.WordsFilter
var samples []string
var root map[string]*wordsfilter.Node

//敏感词过滤

func SetTexts()  {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		samples = append(samples, s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	Wf = wordsfilter.New()
	root = Wf.Generate(samples)
	return
}

func MsgFilter(val string) bool  {
    return 	Wf.Contains(val,root)
}