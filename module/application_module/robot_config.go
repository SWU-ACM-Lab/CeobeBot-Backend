package application_module

import "CeobeBot-Backend/middleware/config"

type RobotConfig struct {
	SolveProblemSuccessIntimacy int //答对题目亲密度加成
	SolveProblemFailIntimacy    int // 答错题目亲密度加成
	RememberWordIntimacy        int // 记住单词亲密度加成
	ActiveIntimacy              int // 连续登录亲密度加成
	AddProblemIntimacy          int // 上传题目亲密度加成
	AddDialogIntimacy           int // 上传对话亲密度加成
}

func (c *RobotConfig) Init(config config.Config) error {
	var err error
	if c.SolveProblemSuccessIntimacy, err = config.Int("robot.solve_problem_success_intimacy"); err != nil {
		return err
	}
	if c.SolveProblemFailIntimacy, err = config.Int("robot.solve_problem_fail_intimacy"); err != nil {
		return err
	}
	if c.RememberWordIntimacy, err = config.Int("robot.remember_word_intimacy"); err != nil {
		return err
	}
	if c.ActiveIntimacy, err = config.Int("robot.active_intimacy"); err != nil {
		return err
	}
	if c.AddProblemIntimacy, err = config.Int("robot.add_problem_intimacy"); err != nil {
		return err
	}
	if c.AddDialogIntimacy, err = config.Int("robot.add_dialog_intimacy"); err != nil {
		return err
	}

	return nil
}
