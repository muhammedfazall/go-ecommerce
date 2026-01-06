package controllers


// // GET admin/products
// func (h *productHandler) AllProducts(ctx *gin.Context) {
// 	var req ProductPagination
// 	var err error

// 	req.Page, err = strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 	if err != nil {
// 		req.Page = 1
// 	}
// 	req.Limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "5"))

// 	if err != nil {
// 		req.Limit = 10
// 	}
// 	req.Query = ctx.Query("q")

// 	products, err := h.svc.AllProducts(&req)
// 	if err != nil {
// 		utils.RenderError(ctx, http.StatusInternalServerError, "Failed to load products", err)
// 		return
// 	}

// 	role, _ := ctx.Get("role")
// 	if role == "admin" {
// 		req.TotalPages = int(math.Ceil(float64(req.Total) / float64(req.Limit)))
// 		ctx.HTML(http.StatusOK, "pages/product/products.html", gin.H{
// 			"Products":    products,
// 			"Query":       req.Query,
// 			"Page":        req.Page,
// 			"Limit":       req.Limit,
// 			"TotalPages":  req.TotalPages,
// 			"CurrentYear": time.Now().Year(),
// 		})
// 		return
// 	}
// 	utils.RenderSuccess(ctx, http.StatusOK, "products list", products)

// }