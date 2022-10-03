package handler

import "github.com/gin-gonic/gin"

//создаем задачу
func (h *Handler) createItem(c *gin.Context){}

//Получаем элементы задачи
func (h *Handler) getAllItems(c *gin.Context){}

//Получаем элемент задачи по Id
func (h *Handler) getItemtById(c *gin.Context){}

//Обновляем задачу
func (h *Handler) updateItem(c *gin.Context){}

//Удаляем элемент задачи
func (h *Handler) deleteItem(c *gin.Context){}