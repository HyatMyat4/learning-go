package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/controllers"
)

func Invoice_router(req *gin.Engine) {
	req.GET("/invoices", controllers.Getinvoices())
	req.GET("/invoice/:invoice_id", controllers.Getsingle_invoice())
	req.POST("/invoice/create", controllers.Createinvoice())
	req.PATCH("/invoice/:invoice_id", controllers.Updateinvoice())

}
