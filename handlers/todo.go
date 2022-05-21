package handlers

import (
	"models"

	"github.com/gin-gonic/gin"
)

// GetTodos ...
// @Summary Todo一覧を配列で返す                                          -> ①どのようなエンドポイントなのかを端的に示すコメント
// @Tags Todo                                                          -> ②tagの指定、tagベースでエンドポイントをグループ化する時に役立つ
// @Produce  json                                                      -> ③どのような形のデータを返すかを指定
// @Success 200 {object} responses.SuccessResponse{data=[]models.Todo} -> ④successレスポンス
// @Failure 400 {object} responses.ErrorResponse                       -> ⑤errorレスポンス
// @Router /todos [get]                                                -> ⑥ルーティング
func (t *Todo) GetTodos(c *gin.Context) {
	todos, err := models.PostTodo() // ←何かしらのTodo配列を返す処理
	if err != nil {
		// エラーハンドリング
	}
	c.JSON(200, todos)
}
