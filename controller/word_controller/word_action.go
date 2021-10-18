package word_controller

import (
	"CeobeBot-Backend/module/orm_module/orm_word"
	"math/rand"
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
