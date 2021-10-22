package words

import (
	"CeobeBot-Backend/controller/statistic_controller"
	"CeobeBot-Backend/controller/word_controller"
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/response_module"
	"CeobeBot-Backend/module/response_module/v1/words"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type WordInterface struct {
}

func (c WordInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	route := engine.Group("/v1/words")

	c.getWord(route, db)
	c.getWordProblem(route, db)
	c.getWordProblemWithAuth(route, db)
	c.judgeWordProblemWithAuth(route, db)
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
					context.JSON(200, words.GetWordResponse{
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
				context.JSON(200, words.GetWordResponse{
					BaseResponse: response_module.BaseResponse{
						Message: "Success",
						Time:    time.Now().String(),
					},
					Id:           word.Id,
					Spell:        word.Spell,
					Phonetic:     word.Phonetic,
					Pos:          word.Pos,
					WordForm:     word.WordForms,
					Paraphrase:   word.Paraphrase,
					AudioSources: word.AudioSources,
				})
				return
			}
		}
	})
}

func (c WordInterface) getWordProblem(route *gin.RouterGroup, db database.MysqlConnection) {
	route.GET("/get-problem", func(context *gin.Context) {
		controller := word_controller.WordController{}
		if controller.Init(db) != nil {
			context.JSON(500, nil)
			return
		} else {
			result, problem, err := controller.CreateWordProblem()
			if result != true || err != nil {
				context.JSON(500, nil)
				return
			} else {
				context.JSON(200, words.GetWordProblemResponse{
					BaseResponse: response_module.BaseResponse{
						Message: "Success",
						Time:    time.Now().String(),
					},
					Spell:    problem.Spell,
					OptionsA: problem.OptionsA,
					OptionsB: problem.OptionsB,
					OptionsC: problem.OptionsC,
					OptionsD: problem.OptionsD,
					Answer:   problem.Answer,
				})
			}
		}
	})
}

func (c WordInterface) getWordProblemWithAuth(route *gin.RouterGroup, db database.MysqlConnection) {
	route.GET("/problem", func(context *gin.Context) {
		cid := context.Query("cid")
		token := context.Query("token")
		uid := context.Query("uid")
		uType := context.Query("type")

		if cid == "" || token == "" || uid == "" || uType == "" {
			context.JSON(400, nil)
			return
		} else {
			intCid, err1 := strconv.ParseUint(cid, 10, 64)
			intUid, err2 := strconv.ParseUint(uid, 10, 64)
			intUtype, err3 := strconv.Atoi(uType)
			if err1 != nil || err2 != nil || err3 != nil {
				context.JSON(500, nil)
				return
			} else {
				controller := word_controller.WordController{}
				if err := controller.Init(db); err != nil {
					context.JSON(500, nil)
					return
				} else {
					result, problem, err := controller.CreateWordProblemWithAuth(intCid, token, intUid, uint(intUtype))
					if result != true || err != nil {
						context.JSON(500, nil)
						return
					} else {
						context.JSON(200, words.GetWordProblemResponse{
							BaseResponse: response_module.BaseResponse{
								Message: "Success",
								Time:    time.Now().String(),
							},
							Spell:    problem.Spell,
							OptionsA: problem.OptionsA,
							OptionsB: problem.OptionsB,
							OptionsC: problem.OptionsC,
							OptionsD: problem.OptionsD,
							Answer:   problem.Answer,
						})
					}
				}
			}
		}
	})
}

func (c WordInterface) judgeWordProblemWithAuth(route *gin.RouterGroup, db database.MysqlConnection) {
	route.POST("/problem", func(context *gin.Context) {
		cid := context.Query("cid")
		token := context.Query("token")
		uid := context.Query("uid")
		uType := context.Query("type")
		answer := context.Query("answer")
		uname := context.Query("uname")

		if cid == "" || token == "" || uid == "" || uType == "" || answer == "" || uname == "" {
			context.JSON(400, nil)
			return
		} else {
			intCid, err1 := strconv.ParseUint(cid, 10, 64)
			intUid, err2 := strconv.ParseUint(uid, 10, 64)
			intUtype, err3 := strconv.Atoi(uType)
			intAnswer, err4 := strconv.Atoi(answer)

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				context.JSON(500, nil)
				return
			} else {
				controller := word_controller.WordController{}
				if err0 := controller.Init(db); err0 != nil {
					context.JSON(500, nil)
					return
				} else {
					result, success, correct, err := controller.JudgeWordProblemWithAuth(intCid, token, intUid, uint(intUtype), intAnswer)
					if err != nil {
						context.JSON(500, nil)
						return
					} else {
						if result == true && success == true {
							statistic := statistic_controller.StatisticController{}
							statistic.Init(db)
							statistic.SolveWordProblem(intCid, token, intUid, uname)
						}
						context.JSON(200, words.JudgeWordProblemResponse{
							BaseResponse: response_module.BaseResponse{
								Message: "Success",
								Time:    time.Now().String(),
							},
							Correct: result,
							Success: success,
							Answer:  correct,
						})
						return
					}
				}
			}
		}

	})
}
