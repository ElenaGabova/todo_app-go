package handler

import "github.com/gin-gonic/gin"

//создаем список
func (h *Handler) createList(c *gin.Context){}

//Получаем элементы списка
func (h *Handler) getAllLists(c *gin.Context){}

//Получаем элемент списка по Id
func (h *Handler) getListById(c *gin.Context){}

//Обновляем список
func (h *Handler) updateList(c *gin.Context){}

//Удаляем элемент списка
func (h *Handler) deleteList(c *gin.Context){}