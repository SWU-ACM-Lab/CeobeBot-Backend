package words

import "CeobeBot-Backend/module/response_module"

type GetClientResponse struct {
	response_module.BaseResponse
	Id           string `json:"word_id"`       // 单词ID
	Spell        string `json:"spell"`         // 单词拼写
	Phonetic     string `json:"phonetic"`      // 单词音标
	Pos          string `json:"pos"`           // 单词释义
	WordForm     string `json:"word_form"`     // 单词形式
	Paraphrase   string `json:"paraphrase"`    // 单词英英释义
	AudioSources string `json:"audio_sources"` // 单词读音URL
}
