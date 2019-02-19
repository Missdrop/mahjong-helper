package main

import (
	"strings"
	"strconv"
	"fmt"
)

type tenhouMessage struct {
	Tag string `json:"tag"`

	//Name string `json:"name"` // id
	//Sex  string `json:"sx"`

	//UserName    string `json:"uname"`
	//RatingScale string `json:"ratingscale"`

	//N string `json:"n"`
	//J string `json:"j"`
	//G string `json:"g"`

	// round 开始 tag=INIT
	Seed string `json:"seed"` // 本局信息：场数，连庄棒数，立直棒数，骰子A减一，骰子B减一，宝牌指示牌 1,0,0,3,2,92
	Ten  string `json:"ten"`  // 各家点数 280,230,240,250
	Oya  string `json:"oya"`  // 庄家 2
	Hai  string `json:"hai"`  // 初始手牌 30,114,108,31,78,107,25,23,2,14,122,44,49

	// 摸牌 tag=T编号，如 T68

	// 副露 tag=N
	Who string `json:"who"` // 副露者 0
	M   string `json:"m"`   // 副露编号 35914

	// 立直声明 tag=REACH, step=1
	// `json:"who"` // 立直者 2
	Step string `json:"step"` // 1

	// TODO: 立直成功后会收到宣言牌的消息，提示宣言牌是否为摸切！

	// 立直成功，扣1000点 tag=REACH, step=2
	// `json:"who"` // 立直者 2
	// `json:"ten"` // 立直成功后的各家点数 250,250,240,250
	// `json:"step"` // 2

	// 自摸/有人放铳 tag=牌, t>=8
	T string `json:"t"` // 选项

	// 和牌 tag=AGARI
	// ba, hai, m, machi, ten, yaku, doraHai, who, fromWho, sc
	//Ba string `json:"ba"` // 0,0
	// `json:"hai"` // 和牌型 8,9,11,14,19,125,126,127
	// `json:"m"` // 副露编号 13527,50794
	//Machi string `json:"machi"` // (待ち) 自摸/荣和的牌 126
	// `json:"ten"` // 符数和点数 30,7700,0
	//Yaku        string `json:"yaku"`       // 役（编号，翻数） 18,1,20,1,34,2
	//DoraTile    string `json:"doraHai"`    // 宝牌 123
	//UraDoraTile string `json:"doraHaiUra"` // 里宝牌 77
	// `json:"who"` // 和牌者 1
	//FromWho string `json:"fromWho"` // 点炮者（自摸时为和牌者） 0
	//Score   string `json:"sc"`      // 各家增减分 260,-77,310,77,220,0,210,0

	// 游戏结束 tag=PROF

	// 重连 tag=GO
	// type, lobby, gpid
	//Type  string `json:"type"`
	//Lobby string `json:"lobby"`
	//GPID  string `json:"gpid"`

	// 重连 tag=REINIT
	// `json:"seed"` // 1,0,0,3,2,92
	// `json:"ten"` // 各家点数 230,211,230,329
	// `json:"oya"` // 2
	// `json:"hai"` // 自家手牌 4,11,17,57,58,59,61,77,81,85,96,97,99
	//M1    string `json:"m1"` // 副露编号 17450
	//M2    string `json:"m2"`
	//M3    string `json:"m3"`
	//Kawa0 string `json:"kawa0"` // 牌河 112,73,3,131,43,98,78,116
	//Kawa1 string `json:"kawa1"`
	//Kawa2 string `json:"kawa2"`
	//Kawa3 string `json:"kawa3"`
}

// 0-35 m
// 36-71 p
// 72-107 s
// 108- z
func (*tenhouMessage) _parseTenhouTile(tile string) int {
	t, err := strconv.Atoi(tile)
	if err != nil {
		panic(err)
	}
	return t / 4
}

func (m *tenhouMessage) analysis(counts []int) error {
	fmt.Println("收到", m.Tag)

	switch m.Tag {
	case "INIT", "REINIT":
		// round 开始/重连
		for i := range counts {
			counts[i] = 0
		}
		for _, pai := range strings.Split(m.Hai, ",") {
			index := m._parseTenhouTile(pai)
			counts[index]++
		}
		return _analysis(13, counts)
	case "N":
		// 某人副露
		if m.Who == "0" {

		}
	case "DORA":
		// 杠宝牌
		// 1. 能摸的牌减少
		// 2. 打点提高
	case "REACH":
		// 如果是他家立直，进入攻守判断模式
		if m.Step != "1" {
			break
		}
	case "AGARI", "RYUUKYOKU":
		// round 结束
	case "PROF":
		// 游戏结束
	case "BYE":
		// 某人退出
	case "REJOIN", "GO":
		// 重连
	case "FURITEN":
		// 振听
	case "TAIKYOKU", "UN", "LN":
		// 其他
	default:
		switch m.Tag[0] {
		case 'T':
			// 摸牌
			index := m._parseTenhouTile(m.Tag[1:])
			counts[index]++
			return _analysis(14, counts)
		case 'D':
			// 自家舍牌
			index := m._parseTenhouTile(m.Tag[1:])
			counts[index]--
			return _analysis(13, counts)
		case 'e', 'f', 'g', 'E', 'F', 'G':
			// 他家舍牌
			//index := m._parseTenhouTile(m.Tag[1:])
			//isTsumogiri := m.Tag[0] >= 'a' // 自模切
			// TODO: 添加到舍牌列表中
			if m.T != "" {
				// 是否副露，何切
			}
		default:
		}
	}

	return nil
}