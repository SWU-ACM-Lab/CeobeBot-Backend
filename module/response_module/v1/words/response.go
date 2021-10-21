package words

import "CeobeBot-Backend/module/response_module"

type GetWordResponse struct {
	response_module.BaseResponse
	Id           string `json:"word_id"`       // 单词ID
	Spell        string `json:"spell"`         // 单词拼写
	Phonetic     string `json:"phonetic"`      // 单词音标
	Pos          string `json:"pos"`           // 单词释义
	WordForm     string `json:"word_form"`     // 单词形式
	Paraphrase   string `json:"paraphrase"`    // 单词英英释义
	AudioSources string `json:"audio_sources"` // 单词读音URL
}

type GetWordProblemResponse struct {
	response_module.BaseResponse
	Spell    string `json:"spell"`     // 正确单词拼写
	OptionsA string `json:"options_a"` // 选项A
	OptionsB string `json:"options_b"` // 选项B
	OptionsC string `json:"options_c"` // 选项C
	OptionsD string `json:"options_d"` // 选项D
	Answer   int    `json:"answer"`    // 正确答案
}

type JudgeWordProblemResponse struct {
	response_module.BaseResponse
	Correct bool `json:"correct"`
	Success bool `json:"success"`
	Answer  int  `json:"answer"`
}
