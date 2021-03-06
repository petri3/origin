package codelocation

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/onsi/ginkgo/types"
)

func New(skip int) types.CodeLocation {
	_, file, line, _ := runtime.Caller(skip + 1)
	return types.CodeLocation{FileName: file, LineNumber: line}
}

func PruneStack(fullStackTrace string, skip int) string {
	stack := strings.Split(fullStackTrace, "\n")
	if len(stack) > 2*(skip+1) {
		stack = stack[2*(skip+1):]
	}
	prunedStack := []string{}
	re := regexp.MustCompile(`\/ginkgo\/|\/pkg\/testing\/|\/pkg\/runtime\/`)
	for i := 0; i < len(stack)/2; i++ {
		if !re.Match([]byte(stack[i*2])) {
			prunedStack = append(prunedStack, stack[i*2])
			prunedStack = append(prunedStack, stack[i*2+1])
		}
	}
	return strings.Join(prunedStack, "\n")
}
