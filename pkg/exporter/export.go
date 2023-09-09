package exporter

import (
	"fmt"
	"github.com/google/go-github/v55/github"
	"github.com/xuri/excelize/v2"
	"log"
)

const s string = "提交记录"
const ff string = "宋体"

func setTitleRow(f *excelize.File, repo *github.Repository) {
	log.Println("export: Formatting the title row...")

	err := f.MergeCell(s, "A1", "H1")
	if err != nil {
		panic(err)
	}

	err = f.SetRowHeight(s, 1, 60)
	if err != nil {
		panic(err)
	}

	style := excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Family: ff,
			Size:   24,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}

	id, err := f.NewStyle(&style)
	if err != nil {
		panic(err)
	}

	err = f.SetCellStyle(s, "A1", "A1", id)
	if err != nil {
		panic(err)
	}

	title := fmt.Sprintf("%s提交记录表", *repo.Name)
	err = f.SetCellValue(s, "A1", title)
	if err != nil {
		panic(err)
	}

	err = f.MergeCell(s, "E2", "H2")
	if err != nil {
		panic(err)
	}
}

func setTableHeader(f *excelize.File) {
	log.Println("export: Setting table header...")
	_ = f.SetCellValue(s, "A5", "序号")
	_ = f.SetCellValue(s, "B5", "日期")
	_ = f.SetCellValue(s, "C5", "SHA")
	_ = f.SetCellValue(s, "D5", "作者GitHub账户")
	_ = f.SetCellValue(s, "E5", "作者名称")
	_ = f.SetCellValue(s, "F5", "作者邮箱")
	_ = f.SetCellValue(s, "G5", "提交信息")
	_ = f.SetCellValue(s, "H5", "是否签名")

	_ = f.SetColWidth(s, "A", "A", 8.11)
	_ = f.SetColWidth(s, "B", "C", 12.78)
	_ = f.SetColWidth(s, "D", "D", 15.56)
	_ = f.SetColWidth(s, "E", "E", 13)
	_ = f.SetColWidth(s, "F", "F", 27)
	_ = f.SetColWidth(s, "G", "G", 43.3)
	_ = f.SetColWidth(s, "H", "H", 8.2)
}

func exportRepoInfo(f *excelize.File, repo *github.Repository, commitCount int) {
	log.Println("export: Exporting repository information...")
	_ = f.SetCellValue(s, "A2", "仓库名")
	_ = f.SetCellValue(s, "A3", "提交量")
	_ = f.SetCellValue(s, "E2", "所有者")

	_ = f.MergeCell(s, "B2", "D2")
	_ = f.MergeCell(s, "B3", "D3")
	//_ = f.MergeCell(s, "F2", "H2")

	_ = f.SetCellValue(s, "B2", *repo.Name)
	_ = f.SetCellValue(s, "B3", fmt.Sprintf("%d", commitCount))
	_ = f.SetCellValue(s, "F2", *repo.Owner.Login)
}

func exportRow(f *excelize.File, index int, commit *github.RepositoryCommit) {
	log.Printf("export: Exporting commit %s...", *commit.SHA)

	var commitGHUser string
	if commit.Committer == nil {
		commitGHUser = ""
	} else {
		commitGHUser = *commit.Committer.Login
	}

	data := []interface{}{
		index + 1,
		commit.Commit.Author.Date.String(),
		(*commit.SHA)[:7],
		commitGHUser,
		*commit.Commit.Author.Name,
		*commit.Commit.Author.Email,
		*commit.Commit.Message,
		*commit.Commit.Verification.Verified,
	}

	cell, err := excelize.CoordinatesToCellName(1, index+6)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.SetSheetRow(s, cell, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExportToXlsx(repository *github.Repository, commits []*github.RepositoryCommit, out string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	_ = f.SetSheetName("Sheet1", s)

	setTitleRow(f, repository)
	setTableHeader(f)
	exportRepoInfo(f, repository, len(commits))
	for i, commit := range commits {
		exportRow(f, i, commit)
	}

	if err := f.SaveAs(out); err != nil {
		fmt.Println(err)
	}
}
