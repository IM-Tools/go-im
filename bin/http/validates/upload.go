/**
  @author:panliang
  @data:2021/7/9
  @note
**/
package validates

type Upload struct {
	Repo     string `json:"repo" validate:"required"`
	Path string `json:"path" validate:"required"`
	Content string `json:"content" validate:"required"`
	Message string `json:"message" validate:"required"`
}

