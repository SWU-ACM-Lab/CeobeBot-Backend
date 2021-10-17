package words

import (
	"CeobeBot-Backend/controller/word_controller"
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/response_module"
	"CeobeBot-Backend/module/response_module/v1/words"
	"github.com/gin-gonic/gin"
	"time"
)

type WordInterface struct {
}

func (c WordInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	route := engine.Group("/v1/words")

	c.getWord(route, db)
	//c.addClient(route, db)

	return nil
}

func (c WordInterface) getWord(route *gin.RouterGroup, db database.MysqlConnection) {
	route.GET("/get", func(context *gin.Context) {
		spell := context.Query("spell")
		// 验证参数是否足够
		if spell == "" {
			context.JSON(400, nil)
			return
		} else {
			controller := word_controller.WordController{}
			if controller.Init(db) != nil {
				context.JSON(500, nil)
				return
			} else {
				// 获取Client
				result, word, err := controller.GetWordBySpell(spell)
				if err != nil {
					context.JSON(500, nil)
					return
				} else if result != true {
					context.JSON(404, nil)
					return
				} else {
					// 返回Client
					context.JSON(200, words.GetClientResponse{
						BaseResponse: response_module.BaseResponse{
							Message: "Success",
							Time:    time.Now().String(),
						},
						Id:         word.Id,
						Spell:      word.Spell,
						Phonetic:   word.Phonetic,
						Pos:        word.Pos,
						WordForm:   word.WordForms,
						Paraphrase: word.Paraphrase,
					})
					return
				}
			}
		}
	})
	route.GET("/random", func(context *gin.Context) {
		controller := word_controller.WordController{}
		if controller.Init(db) != nil {
			context.JSON(500, nil)
			return
		} else {
			// 获取Client
			result, word, err := controller.GetRandomWord()
			if err != nil {
				context.JSON(500, nil)
				return
			} else if result != true {
				context.JSON(404, nil)
				return
			} else {
				// 返回Client
				context.JSON(200, words.GetClientResponse{
					BaseResponse: response_module.BaseResponse{
						Message: "Success",
						Time:    time.Now().String(),
					},
					Id:         word.Id,
					Spell:      word.Spell,
					Phonetic:   word.Phonetic,
					Pos:        word.Pos,
					WordForm:   word.WordForms,
					Paraphrase: word.Paraphrase,
				})
				return
			}
		}
	})
}
