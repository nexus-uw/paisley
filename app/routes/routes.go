// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tApplication struct {}
var Application tApplication


func (_ tApplication) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.AddUser", args).URL
}

func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).URL
}

func (_ tApplication) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Register", args).URL
}

func (_ tApplication) SaveUser(
		user interface{},
		verifyPassword string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "verifyPassword", verifyPassword)
	return revel.MainRouter.Reverse("Application.SaveUser", args).URL
}

func (_ tApplication) Login(
		username string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("Application.Login", args).URL
}

func (_ tApplication) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Logout", args).URL
}


type tFeed struct {}
var Feed tFeed


func (_ tFeed) GetFeed(
		SubscriptionID string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "SubscriptionID", SubscriptionID)
	return revel.MainRouter.Reverse("Feed.GetFeed", args).URL
}


type tSubscriptions struct {}
var Subscriptions tSubscriptions


func (_ tSubscriptions) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Subscriptions.Index", args).URL
}

func (_ tSubscriptions) Create(
		subscription interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "subscription", subscription)
	return revel.MainRouter.Reverse("Subscriptions.Create", args).URL
}

func (_ tSubscriptions) Delete(
		SubscriptionID string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "SubscriptionID", SubscriptionID)
	return revel.MainRouter.Reverse("Subscriptions.Delete", args).URL
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).URL
}


type tController struct {}
var Controller tController


func (_ tController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Controller.Begin", args).URL
}

func (_ tController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Controller.Rollback", args).URL
}

func (_ tController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Controller.Commit", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


