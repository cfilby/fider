package handler

import (
	"net/http"

	"github.com/WeCanHearYou/wchy/context"
	"github.com/WeCanHearYou/wchy/model"
	"github.com/labstack/echo"
)

type indexHandlder struct {
	ctx *context.WchyContext
}

// Index handles initial page
func Index(ctx *context.WchyContext) echo.HandlerFunc {
	return indexHandlder{ctx: ctx}.get()
}

func (h indexHandlder) get() echo.HandlerFunc {
	return func(c echo.Context) error {
		tenant := c.Get("Tenant").(*model.Tenant)
		ideas, _ := h.ctx.Idea.GetAll(tenant.ID)
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Tenant": tenant,
			"Ideas":  ideas,
		})
	}
}
