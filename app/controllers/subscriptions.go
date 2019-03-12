package controllers

import (
	"github.com/nexus-uw/paisley/app/routes"

	"github.com/revel/revel"

	"github.com/nexus-uw/paisley/app/models"
)

type Subscriptions struct {
	Application
}

func (c Subscriptions) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.Application.Index())
	}
	return nil
}

func (c Subscriptions) Index() revel.Result {
	c.Log.Info("Fetching index")
	var subscriptions []*models.Subscription
	_, err := c.Txn.Select(&subscriptions,
		c.Db.SqlStatementBuilder.Select("*").
			From("Subscription")) /*.Where("UserId = ?", c.connected().UserId))*/

	if err != nil {
		panic(err)
	}

	return c.Render(subscriptions)
}

func (c Subscriptions) Create(subscription models.Subscription) revel.Result {
	subscription.OwnerID = c.connected().UserId

	subscription.Validate(c.Validation)

	err := c.Txn.Insert(&subscription)
	if err != nil {
		panic(err)
	}
	c.Flash.Success("Thank you, %s, created sub",
		c.connected().Name)
	return c.Redirect(routes.Subscriptions.Index())

}
