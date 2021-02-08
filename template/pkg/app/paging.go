package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pjchender/go-snippets/template/global"
)

// Paging 是用來表示頁籤資訊
type Paging struct {
	Page       int   `form:"page" json:"page"`         // 要顯示第幾頁的資料
	PageSize   int   `form:"pageSize" json:"pageSize"` // 每頁要顯示幾筆資料
	TotalPages int64 `json:"totalPages"`
}

func NewPaging(ctx *gin.Context) (*Paging, error) {
	paging := Paging{}
	err := ctx.ShouldBind(&paging)
	if err != nil {
		return nil, err
	}

	return &Paging{
		Page:       GetPage(paging.Page),
		PageSize:   GetPageSize(paging.PageSize),
		TotalPages: 0,
	}, nil
}

func ToListResponse(data interface{}, paging *Paging) gin.H {
	return gin.H{
		"data":   data,
		"paging": paging,
	}
}

// GetPage 可以用來取得 API request 想要請求哪一頁的資料
func GetPage(page int) int {
	if page <= 0 {
		return 1
	}

	return page
}

// GetPageSize 可以用來取得 API 請求時每頁顯示的資料筆數
func GetPageSize(pageSize int) int {
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// WithTotalPage 可以根據 pageSize 和 totalCount（總資料筆數）換算共有幾頁資料
func (p *Paging) WithTotalPage(totalCount int64) *Paging {
	totalPage := totalCount / int64(p.PageSize)
	remainderPage := totalCount % int64(p.PageSize)

	if remainderPage > 0 {
		totalPage++
	}

	return &Paging{
		Page:       p.Page,
		PageSize:   p.PageSize,
		TotalPages: totalPage,
	}
}

// GetPageOffset 可以根據 page 和 pageSize 取得開始呈現的是第幾筆的資料
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
