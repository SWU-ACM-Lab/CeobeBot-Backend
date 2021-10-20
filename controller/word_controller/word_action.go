package word_controller

import (
	"CeobeBot-Backend/module/orm_module/orm_word"
	"CeobeBot-Backend/module/runtime_module/runtime_word"
	"encoding/json"
	"fmt"
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
