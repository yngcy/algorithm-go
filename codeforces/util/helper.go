package util

import (
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func removeExtraSpace(s string) string {
	s = strings.TrimSpace(s)
	sp := strings.Split(s, "\n")
	for i := range sp {
		sp[i] = strings.TrimSpace(sp[i])
	}
	return strings.Join(sp, "\n")
}

const minLineToFile = 750

func handleLongOutput(s string) {
	name := time.Now().Format("150405.000")
	if err := os.WriteFile(name+".txt", []byte(s), 0644); err != nil {
		panic(err)
	}
}

func handleOutput(s string) {
	if len(s) >= minLineToFile {
		handleOutput(s)
	}
}

// IsDebugging 检查当前进程是否正在被 Delve 调试器调试。
// 该函数通过追溯进程的父进程，判断是否为 Delve 调试器。
func IsDebugging() bool {
	// 获取当前进程的父进程ID。
	pid := int32(os.Getppid())

	// 循环检查直到找到根进程或确定没有被 Delve 调试。
	for pid != 0 {
		// 尝试创建一个新进程对象来表示父进程。
		p, err := process.NewProcess(pid)
		if err != nil {
			// 如果创建进程对象失败，认为当前进程不是被 Delve 调试的。
			return false
		}
		// 获取父进程的名称。
		name, err := p.Name()
		if err != nil {
			// 如果获取进程名称失败，认为当前进程不是被 Delve 调试的。
			return false
		}
		// 检查进程名称是否为 Delve。
		if strings.HasPrefix(name, "dlv") {
			// 如果是，返回 true，表示当前进程正在被 Delve 调试。
			return true
		}
		// 获取父进程的父进程ID，继续向上追溯。
		pid, _ = p.Ppid()
	}
	// 如果没有找到 Delve 调试器，返回 false。
	return false
}

func TransEdge(edges [][2]int) [][]int {
	es := make([][]int, len(edges))
	for i, e := range edges {
		es[i] = []int{e[0], e[1]}
	}
	return es
}
