package webAPI

import (
	"encoding/json"
	"github.com/DanielRenne/GoCore/core/extensions"
	"github.com/DanielRenne/GoCore/core/ginServer"
	"github.com/gin-gonic/gin"
	"helloWorld/models/v1/model"
	"io/ioutil"
	"strings"
)

func init() {

	ginServer.AddRouterGroup("/api/v1", "/singleCars", "GET", getSingleCars)
	ginServer.AddRouterGroup("/api/v1", "/searchCars", "GET", getSearchCars)
	ginServer.AddRouterGroup("/api/v1", "/sortCars", "GET", getSortCars)
	ginServer.AddRouterGroup("/api/v1", "/rangeCars", "GET", getRangeCars)
	ginServer.AddRouterGroup("/api/v1", "/cars", "GET", getCars)
	ginServer.AddRouterGroup("/api/v1", "/car", "POST", postCar)
	ginServer.AddRouterGroup("/api/v1", "/car", "PUT", putCar)
	ginServer.AddRouterGroup("/api/v1", "/car", "DELETE", deleteCar)
}

func getSingleCars(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	value := c.DefaultQuery("value", "")
	items := model.Cars{}
	itemsArray, _ := items.Single(field, value)
	ginServer.RespondJSON(itemsArray, c)
}

func getSearchCars(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	value := c.DefaultQuery("value", "")
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	items := model.Cars{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.SearchAdvanced(field, value, limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.Search(field, value)
	ginServer.RespondJSON(itemsArray, c)
}

func getSortCars(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	items := model.Cars{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.AllByIndexAdvanced(field, limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.AllByIndex(field)
	ginServer.RespondJSON(itemsArray, c)
}

func getRangeCars(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	min := c.DefaultQuery("min", "")
	max := c.DefaultQuery("max", "")
	items := model.Cars{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.RangeAdvanced(min, max, field, limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.Range(min, max, field)
	ginServer.RespondJSON(itemsArray, c)
}

func getCars(c *gin.Context) {
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	items := model.Cars{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.AllAdvanced(limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.All()
	ginServer.RespondJSON(itemsArray, c)
}

func postCar(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	var obj model.Car
	errMarshal := json.Unmarshal(x, &obj)
	if errMarshal != nil {
		c.Data(406, gin.MIMEHTML, ginServer.RespondError(errMarshal.Error()))
		return
	}
	errSave := obj.Save()
	if errSave != nil {
		c.Data(500, gin.MIMEHTML, ginServer.RespondError(errSave.Error()))
		return
	}
	ginServer.RespondJSON(obj, c)
}

func putCar(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	var obj model.Car
	errMarshal := json.Unmarshal(x, &obj)
	if errMarshal != nil {
		c.Data(406, gin.MIMEHTML, ginServer.RespondError(errMarshal.Error()))
		return
	}
	objCars := model.Cars{}
	_, errSingle := objCars.Single("Id", obj.Id)
	if errSingle != nil {
		c.Data(404, gin.MIMEHTML, ginServer.RespondError("Existing Car not found.  Please check your primary key to update."))
		return
	}
	errSave := obj.Save()
	if errSave != nil {
		c.Data(500, gin.MIMEHTML, ginServer.RespondError(errSave.Error()))
		return
	}
	ginServer.RespondJSON(obj, c)
}

func deleteCar(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	var obj model.Car
	errMarshal := json.Unmarshal(x, &obj)
	if errMarshal != nil {
		c.Data(406, gin.MIMEHTML, ginServer.RespondError(errMarshal.Error()))
		return
	}
	errDelete := obj.Delete()
	if errDelete != nil {
		c.Data(500, gin.MIMEHTML, ginServer.RespondError(errDelete.Error()))
		return
	}
	ginServer.RespondJSON(obj, c)
}
