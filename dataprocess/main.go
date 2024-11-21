package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

const coordAdd float64 = 0.00002

func main() {
	// 打开一个 .xlsx 文件
	file, err := excelize.OpenFile("h.xlsx")
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer func() {
		// 确保文件关闭
		if err := file.Close(); err != nil {
			log.Fatalf("无法关闭文件: %v", err)
		}
	}()

	// 获取指定工作表的所有行
	rows, err := file.GetRows("distance")
	if err != nil {
		log.Fatalf("无法读取工作表: %v", err)
	}

	// 遍历行并打印内容
	var p time.Time
	var result [][]string
	var r [][]string
	var z float64
	for i := 1; i < len(rows); i += 3 {
		if len(rows[i]) == 0 {
			continue
		}
		timestr := rows[i][0]
		t, _ := time.Parse("2006-01-02_15:04:05", timestr)
		if p.IsZero() {
			p = t
		}
		if !p.Equal(t) {
			ProcessResult(result, p)
			result = make([][]string, 0)
			p = t
		}
		colx := rows[i+1]
		coly := rows[i+2]
		for ii := 2; ii < len(colx); ii++ {
			z += coordAdd
			tf := strconv.FormatFloat(z, 'f', 20, 64)
			result = append(result, []string{colx[ii], coly[ii], tf})
			r = append(r, []string{colx[ii], coly[ii], tf})
		}
	}
	ProcessResultT(r)
}

func ProcessResult(d [][]string, t time.Time) {
	dataLen := len(d)
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(`# .PCD v.7 - Point Cloud Data file format
VERSION .7
FIELDS x y z
SIZE 4 4 4
TYPE F F F
COUNT 1 1 1
WIDTH %d
HEIGHT 1
VIEWPOINT 0 0 0 1 0 0 0
POINTS %d
DATA ascii`, dataLen, dataLen))
	for _, v := range d {
		sb.WriteString("\n")
		sb.WriteString(v[0])
		sb.WriteString("        ")
		sb.WriteString(v[1])
		sb.WriteString("        ")
		sb.WriteString(v[2])
	}
	os.WriteFile("./data/"+t.Format("2006-01-02_15:04:05")+".pcd", []byte(sb.String()), 0664)
}

func ProcessResultT(d [][]string) {
	dataLen := len(d)
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(`# .PCD v.7 - Point Cloud Data file format
VERSION .7
FIELDS x y z
SIZE 4 4 4
TYPE F F F
COUNT 1 1 1
WIDTH %d
HEIGHT 1
VIEWPOINT 0 0 0 1 0 0 0
POINTS %d
DATA ascii`, dataLen, dataLen))
	for _, v := range d {
		sb.WriteString("\n")
		sb.WriteString(v[0])
		sb.WriteString("        ")
		sb.WriteString(v[1])
		sb.WriteString("        ")
		sb.WriteString(v[2])
	}
	os.WriteFile("./data/result.pcd", []byte(sb.String()), 0664)
}
