package game

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"math/rand"
	"strconv"
)

type RmGame struct {
	MaxRow      int     `json:"MaxRow" from:"MaxRow"`
	MaxCol      int     `json:"MaxCol" from:"MaxCol"`
	Table       [][]Tag `json:"Table" from:"Table"`
	Integration int     `json:"Integration" from:"Integration"`
}

var tagBase = []string{"A", "B", "C", "D", "E", "F", "G"}
var tags = [][]string{
	{"A", "B", "C", "D", "E", "F", "G", "H", "I", "A"},
	{"B", "A", "E", "F", "A", "I", "H", "G", "E", "G"},
	{"C", "D", "A", "C", "H", "G", "I", "E", "H", "I"},
	{"D", "A", "D", "A", "I", "E", "H", "E", "C", "C"},
	{"E", "C", "E", "D", "A", "C", "H", "C", "I", "H"},
	{"F", "B", "C", "C", "I", "A", "C", "B", "H", "A"},
	{"A", "D", "A", "H", "A", "G", "H", "G", "C", "G"},
	{"B", "E", "C", "E", "I", "H", "E", "A", "E", "C"},
	{"C", "F", "B", "I", "C", "G", "G", "C", "A", "C"},
	{"D", "F", "G", "H", "A", "C", "B", "C", "C", "A"},
}

// PrintTable
//
//	@Description: 打印游戏表格
//	@receiver rg
func (rg *RmGame) PrintTable() {
	fmt.Print("-----------------------------\r\n")
	for _, row := range rg.Table {

		for _, col := range row {
			if col.Name != "" {
				fmt.Print(col.Name + " ")
			} else {
				fmt.Print("X" + " ")
			}

		}
		fmt.Print("\r\n")
	}
	fmt.Print("-----------------------------\r\n")
}
func NewGame() (newGame *RmGame) {
	newGame = &RmGame{
		MaxRow: 10,
		MaxCol: 10,
		Table:  make([][]Tag, 0),
	}
	for i, row := range tags {
		tags := make([]Tag, 0)
		for c, _ := range row {
			tags = append(tags, Tag{
				RowIndex: i,
				ColIndex: c,
				Name:     GetRandomTag(),
				Uid:      strconv.Itoa(i) + strconv.Itoa(c),
				Status:   1,
				Style: Style{
					Top:     0,
					Left:    0,
					Opacity: 1,
				},
			})
		}
		newGame.Table = append(newGame.Table, tags)
	}
	return newGame
}

// ExchangePosition
//
//	@Description: 交换两个位置的tag
//	@receiver rg
//	@param tag1  第一个tag
//	@param tag2 第二个tag
//	@return changes 交换后的位置
//	@return err
func (rg *RmGame) ExchangePosition(tag1 Tag, tag2 Tag) (err error) {
	if tag1.RowIndex == tag2.RowIndex && tag1.ColIndex == tag2.ColIndex {
		return errors.New("不能交换相同的位置")
	}
	if ((tag1.RowIndex-tag2.RowIndex == 1 || tag1.RowIndex-tag2.RowIndex == -1) ||
		(tag1.ColIndex-tag2.ColIndex == 1 || tag1.ColIndex-tag2.ColIndex == -1)) &&
		(tag1.RowIndex == tag2.RowIndex || tag1.ColIndex == tag2.ColIndex) {

		// 交换位置
		rg.Table[tag1.RowIndex][tag1.ColIndex].Name, rg.Table[tag2.RowIndex][tag2.ColIndex].Name = rg.Table[tag2.RowIndex][tag2.ColIndex].Name, rg.Table[tag1.RowIndex][tag1.ColIndex].Name
		return nil
	} else {
		return errors.New("只能交换相邻的位置")
	}
	return nil
}
func (rg *RmGame) SameExtraction(tag Tag, isPlay bool) (tags []Tag) {
	if tag.Name == "" {
		return
	}
	LRTags := []Tag{rg.Table[tag.RowIndex][tag.ColIndex]}
	UDTags := []Tag{rg.Table[tag.RowIndex][tag.ColIndex]}
	//向左
	for i := tag.ColIndex - 1; i >= 0; i-- {
		if rg.Table[tag.RowIndex][i].Name == rg.Table[tag.RowIndex][tag.ColIndex].Name {
			LRTags = append(LRTags, rg.Table[tag.RowIndex][i])
		} else {
			break
		}
	}
	//向右
	for i := tag.ColIndex + 1; i < rg.MaxCol; i++ {
		if rg.Table[tag.RowIndex][i].Name == rg.Table[tag.RowIndex][tag.ColIndex].Name {
			LRTags = append(LRTags, rg.Table[tag.RowIndex][i])
		} else {
			break
		}
	}
	//向上
	for i := tag.RowIndex - 1; i >= 0; i-- {
		if rg.Table[i][tag.ColIndex].Name == rg.Table[tag.RowIndex][tag.ColIndex].Name {
			UDTags = append(UDTags, rg.Table[i][tag.ColIndex])
		} else {
			break
		}
	}
	//向下
	for i := tag.RowIndex + 1; i < rg.MaxRow; i++ {
		if rg.Table[i][tag.ColIndex].Name == rg.Table[tag.RowIndex][tag.ColIndex].Name {
			UDTags = append(UDTags, rg.Table[i][tag.ColIndex])
		} else {
			break
		}
	}
	if len(LRTags) >= 3 {
		for _, v := range LRTags {
			rg.Table[v.RowIndex][v.ColIndex].Style.Opacity = 0.0
			rg.Table[v.RowIndex][v.ColIndex].Name = ""
		}
		tags = lo.Union(tags, LRTags)
		if isPlay {
			rg.Integration += (len(LRTags) - 2) * 100
		}
	}
	if len(UDTags) >= 3 {
		for _, v := range UDTags {
			rg.Table[v.RowIndex][v.ColIndex].Style.Opacity = 0.0
			rg.Table[v.RowIndex][v.ColIndex].Name = ""
		}
		tags = lo.Union(tags, UDTags)
		if isPlay {
			rg.Integration += (len(UDTags) - 2) * 100
		}
	}
	if len(tags) == 1 {
		return
	}
	return tags
}

func (rg *RmGame) Drop(tags []Tag) {
	cols := lo.GroupBy(tags, func(tag Tag) int {
		return tag.ColIndex
	})
	for col, _ := range cols {
		count := 0
		for row := rg.MaxRow - 1; row >= 0; row-- {
			if rg.Table[row][col].Name == "" {
				count++
			} else {
				rg.Table[row][col].Name, rg.Table[row+count][col].Name = rg.Table[row+count][col].Name, rg.Table[row][col].Name
				rg.Table[row][col].Style.Opacity, rg.Table[row+count][col].Style.Opacity = rg.Table[row+count][col].Style.Opacity, rg.Table[row][col].Style.Opacity
			}
		}
	}
}

func (rg *RmGame) AddTag(tags []Tag) {
	cols := lo.GroupBy(tags, func(tag Tag) int {
		return tag.ColIndex
	})
	for col, _ := range cols {
		for row := rg.MaxRow - 1; row >= 0; row-- {
			if rg.Table[row][col].Name == "" {
				rg.Table[row][col].Name = GetRandomTag()
				rg.Table[row][col].Style.Opacity = 1.0
			}
		}
	}
}
func GetRandomTag() string {
	return tagBase[rand.Intn(len(tagBase))]
}
