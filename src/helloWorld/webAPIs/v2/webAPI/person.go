package webAPI

import (
	"encoding/json"
	"github.com/DanielRenne/GoCore/core/extensions"
	"github.com/DanielRenne/GoCore/core/ginServer"
	"github.com/gin-gonic/gin"
	"helloWorld/models/v2/model"
	"io/ioutil"
	"strings"
)

func init() {

	ginServer.AddRouterGroup("/api/v2", "/singlePersons", "GET", getSinglePersons)
	ginServer.AddRouterGroup("/api/v2", "/searchPersons", "GET", getSearchPersons)
	ginServer.AddRouterGroup("/api/v2", "/sortPersons", "GET", getSortPersons)
	ginServer.AddRouterGroup("/api/v2", "/rangePersons", "GET", getRangePersons)
	ginServer.AddRouterGroup("/api/v2", "/persons", "GET", getPersons)
	ginServer.AddRouterGroup("/api/v2", "/person", "POST", postPerson)
	ginServer.AddRouterGroup("/api/v2", "/person", "PUT", putPerson)
	ginServer.AddRouterGroup("/api/v2", "/person", "DELETE", deletePerson)
}

func getSinglePersons(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	value := c.DefaultQuery("value", "")
	items := model.Persons{}
	itemsArray, _ := items.Single(field, value)
	ginServer.RespondJSON(itemsArray, c)
}

func getSearchPersons(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	value := c.DefaultQuery("value", "")
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	items := model.Persons{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.SearchAdvanced(field, value, limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.Search(field, value)
	ginServer.RespondJSON(itemsArray, c)
}

func getSortPersons(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	items := model.Persons{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.AllByIndexAdvanced(field, limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.AllByIndex(field)
	ginServer.RespondJSON(itemsArray, c)
}

func getRangePersons(c *gin.Context) {
	field := strings.Title(c.DefaultQuery("field", ""))
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	min := c.DefaultQuery("min", "")
	max := c.DefaultQuery("max", "")
	items := model.Persons{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.RangeAdvanced(min, max, field, limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.Range(min, max, field)
	ginServer.RespondJSON(itemsArray, c)
}

func getPersons(c *gin.Context) {
	limit := extensions.StringToInt(c.DefaultQuery("limit", ""))
	skip := extensions.StringToInt(c.DefaultQuery("skip", ""))
	items := model.Persons{}
	if limit != 0 || skip != 0 {
		itemsArray, _ := items.AllAdvanced(limit, skip)
		ginServer.RespondJSON(itemsArray, c)
		return
	}
	itemsArray, _ := items.All()
	ginServer.RespondJSON(itemsArray, c)
}

func postPerson(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	var obj model.Person
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

func putPerson(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	var obj model.Person
	errMarshal := json.Unmarshal(x, &obj)
	if errMarshal != nil {
		c.Data(406, gin.MIMEHTML, ginServer.RespondError(errMarshal.Error()))
		return
	}
	objPersons := model.Persons{}
	_, errSingle := objPersons.Single("Id", obj.Id)
	if errSingle != nil {
		c.Data(404, gin.MIMEHTML, ginServer.RespondError("Existing Person not found.  Please check your primary key to update."))
		return
	}
	errSave := obj.Save()
	if errSave != nil {
		c.Data(500, gin.MIMEHTML, ginServer.RespondError(errSave.Error()))
		return
	}
	ginServer.RespondJSON(obj, c)
}

func deletePerson(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	var obj model.Person
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
