package word_controller

import (
	"CeobeBot-Backend/module/orm_module/orm_client"
	"CeobeBot-Backend/module/orm_module/orm_word"
	"CeobeBot-Backend/module/runtime_module/runtime_word"
	"encoding/json"
	"fmt"
	"time"

	//"encoding/json"
	"math/rand"
	"sort"
)

func (c WordController) GetWordById(wid uint64) (result bool, word orm_word.Words, err error) {
	queryWord := orm_word.Words{}
	queryResult, queryError := c.db.Connection.ID(wid).Get(&queryWord)
	if queryResult != true || queryError != nil {
		return false, orm_word.Words{}, queryError
	} else {
		return true, queryWord, nil
	}
}

func (c WordController) GetWordBySpell(spell string) (result bool, word orm_word.Words, err error) {
	queryWord := orm_word.Words{Spell: spell}
	queryResult, queryError := c.db.Connection.Get(&queryWord)
	if queryResult != true || queryError != nil {
		return false, orm_word.Words{}, queryError
	} else {
		return true, queryWord, nil
	}
}

func (c WordController) GetRandomWord() (result bool, word orm_word.Words, err error) {
	idx := rand.Intn(13281)
	queryWord := orm_word.Words{}
	queryResult, queryError := c.db.Connection.Table(word).Limit(1, idx).Get(&queryWord)
	if queryResult != true || queryError != nil {
		return false, orm_word.Words{}, queryError
	} else {
		return true, queryWord, nil
	}
}

func (c WordController) CreateWordProblem() (result bool, problem runtime_word.WordProblem, err error) {
	options := make([]int, 0)

	// 生成随机数
	for i := 0; i < 4; i++ {

		num := rand.Intn(13281)
		exist := false
		for _, v := range options {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			options = append(options, num)
		}
	}

	fmt.Println(options)

	sort.Ints(options)

	problems := make([]orm_word.Words, 4)
	for i := 0; i < 4; i++ {
		problemResult, problemError := c.db.Connection.Table(orm_word.Words{}).Limit(1, options[i]).Get(&problems[i])
		if problemResult != true || problemError != nil {
			return false, runtime_word.WordProblem{}, problemError
		}
		fmt.Println(problems[i].Pos)
	}

	var problemPos []string
	for i := 0; i < 4; i++ {
		var Poses []runtime_word.WordPos
		if errs := json.Unmarshal([]byte(problems[i].Pos), &Poses); errs != nil {
			return false, runtime_word.WordProblem{}, errs
		}

		temp := ""
		for j := 0; j < len(Poses); j++ {
			temp += fmt.Sprintf("%v. %v。", Poses[j].Type, Poses[j].Definition)
		}

		problemPos = append(problemPos, temp)
	}

	fmt.Println(problemPos)

	// 确定正确单词并组装
	index := rand.Intn(4)
	wordProblem := runtime_word.WordProblem{
		Spell:    problems[index].Spell,
		OptionsA: problemPos[0],
		OptionsB: problemPos[1],
		OptionsC: problemPos[2],
		OptionsD: problemPos[3],
		Answer:   index,
	}
	return true, wordProblem, nil
}

func (c WordController) CreateWordProblemWithAuth(cid uint64, token string, uid uint64, uType uint) (result bool, problem runtime_word.WordProblem, err error) {
	// 获取cid对应的Client
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	if queryResult != true || queryError != nil {
		return false, runtime_word.WordProblem{}, queryError
	} else if client.Token != token {
		return false, runtime_word.WordProblem{}, nil
	} else {
		// 创建题目
		results, problems, errs := c.CreateWordProblem()
		if results != true || errs != nil {
			return false, runtime_word.WordProblem{}, errs
		} else {
			// 判断题目记录是否存在
			record := orm_word.WordProblemRecord{}
			recordResult, recordError := c.db.Connection.Table(orm_word.WordProblemRecord{}).Where("UserId = ?", uid).Where("UserType = ?", uType).Get(&record)
			if recordError != nil {
				return false, runtime_word.WordProblem{}, recordError
			} else {
				// 如果题目记录存在，则覆盖
				if recordResult == true {
					// 更新题目记录
					record.Answer = problems.Answer
					record.UpdateTime = time.Now()
					record.Status = 1
					// 更新数据库
					if _, updateError := c.db.Connection.Table(orm_word.WordProblemRecord{}).ID(record.Id).Update(record); updateError != nil {
						return false, runtime_word.WordProblem{}, updateError
					} else {
						return true, problems, nil
					}
					// 否则则创建
				} else {
					record = orm_word.WordProblemRecord{
						UserId:     uid,
						UserType:   uType,
						UpdateTime: time.Now(),
						Answer:     problems.Answer,
						Status:     1,
					}
					// 写入数据库
					if _, insertError := c.db.Connection.Table(orm_word.WordProblemRecord{}).InsertOne(record); insertError != nil {
						return false, runtime_word.WordProblem{}, insertError
					} else {
						return true, problems, nil
					}
				}
			}
		}
	}
}

func (c WordController) JudgeWordProblemWithAuth(cid uint64, token string, uid uint64, uType uint, answer int) (result, success bool, err error) {
	// 获取cid对应的Client
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	if queryResult != true || queryError != nil {
		return false, false, queryError
	} else if client.Token != token {
		return false, false, nil
	} else {
		// 获取uid, uType对应的题目
		record := orm_word.WordProblemRecord{}
		recordResult, recordError := c.db.Connection.Table(orm_word.WordProblemRecord{}).Where("UserId = ?", uid).Where("UserType = ?", uType).Get(&record)
		if recordResult != true || recordError != nil {
			return false, false, recordError
		} else {
			// 判题
			if record.Status == 2 {
				return record.Answer == answer, false, nil
			} else {
				if record.Answer == answer {
					record.Status = 2
					if _, updateError := c.db.Connection.ID(record.Id).Update(record); updateError != nil {
						return false, false, updateError
					} else {
						return true, true, nil
					}
				} else {
					return false, false, nil
				}
			}
		}
	}
}
