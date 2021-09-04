/**
  @author:panliang
  @data:2021/8/26
  @note
**/
package wordsfilter

import (
	"bufio"
	"github.com/syyongx/go-wordsfilter"
	"go.uber.org/zap"
	"go_im/pkg/log"
	"os"
)

var Wf *wordsfilter.WordsFilter
var samples []string
var root map[string]*wordsfilter.Node

//敏感词过滤

func SetTexts()  {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Logger.Error("没有找到敏感词文件",zap.Error(err))
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		samples = append(samples, s)
	}
	if err := scanner.Err(); err != nil {
		log.Logger.Error("加载scanner失败",zap.Error(err))
	}
	Wf = wordsfilter.New()
	root = Wf.Generate(samples)
	return
}

func MsgFilter(val string) bool  {
    return 	Wf.Contains(val,root)
}