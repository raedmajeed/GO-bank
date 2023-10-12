package controllers

type HttpControllers interface {
	Handlers()
	Start()
}