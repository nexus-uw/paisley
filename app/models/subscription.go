package models

import (
	"github.com/revel/revel"
)

type Subscription struct {
	SubscriptionID      int
	Frequency           int
	OwnerID             int
	Subredit            string
	MaxDailyPosts       int
	MinVoteRatio        int
	MinVoteCount        int
	MinNumberOfComments int
	IncludeComments     bool
	IncludeUserNames    bool
}

func (subscription *Subscription) Validate(v *revel.Validation) {
	v.Check(subscription.Subredit,
		revel.Required{},
		revel.MaxSize{50},
	)
}
