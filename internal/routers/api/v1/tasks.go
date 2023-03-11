package v1

import "github.com/gin-gonic/gin"

type Tasks struct{}

func NewTasks() Tasks {
	return Tasks{}
}

// GetTPOPData
// @Summary 获取《今日各项目时间占比图》数据
// @Produce json
// @Param today path uint32 true "今日日期"
// @Success 200 {object} model.TPOPData "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /menu/statistics/tpop/:today [get]
func (t Tasks) GetTPOPData(c *gin.Context) {

}

// GetTPOCData
// @Summary 获取《今日各类项目时间占比图》数据
// @Produce json
// @Param today path uint32 true "今日日期"
// @Success 200 {object} model.TPOCData "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /menu/statistics/tpoc/:today [get]
func (t Tasks) GetTPOCData(c *gin.Context) {

}

// GetTDOWData
// @Summary 获取《本周每天工作时长》数据
// @Produce json
// @Param today path uint32 true "今日日期"
// @Success 200 {object} model.TDOWData "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /menu/statistics/tdow/:today [get]
func (t Tasks) GetTDOWData(c *gin.Context) {

}

// GetMoneyData
// @Summary 获取《本周每天工作时长》数据
// @Produce json
// @Param today path uint32 true "今日日期"
// @Success 200 {object} model.MoneyData "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /menu/statistics/money/:today [get]
func (t Tasks) GetMoneyData(c *gin.Context) {

}

// UploadTimerData
// @Summary 上传定时器数据
// @Param start body uint32 true "当前时间"
// @Param duration body uint32 true "持续时间"
// @Param end body uint32 true "结束时间"
// @Param task_name body string true "任务名称"
// @Param tag body string true "任务标签"
// @Param cid path string true "用户唯一表示符"
// @Success 200  "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router  /timer/:cid [post]
func (t Tasks) UploadTimerData(c *gin.Context) {

}
