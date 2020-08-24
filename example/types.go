package main

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/jonahgeorge/zwiftpower-go"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"allLeagues": &graphql.Field{
			Type: graphql.NewList(LeagueType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := zpp.zpClient.ListLeagues(
					context.TODO(),
					&zwiftpower.ListLeaguesRequest{},
				)
				if err != nil {
					return nil, err
				}
				return resp.Leagues, nil
			},
		},
		"allEvents": &graphql.Field{
			Type: graphql.NewList(EventType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := zpp.zpClient.ListEvents(
					context.TODO(),
					&zwiftpower.ListEventsRequest{},
				)
				if err != nil {
					return nil, err
				}
				return resp.Events, nil
			},
		},
		"allTeams": &graphql.Field{
			Type: graphql.NewList(TeamType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := zpp.zpClient.ListTeams(
					context.TODO(),
					&zwiftpower.ListTeamsRequest{},
				)
				if err != nil {
					return nil, err
				}
				return resp.Teams, nil
			},
		},
		"allActivities": &graphql.Field{
			Type: graphql.NewList(ActivityType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := zpp.zpClient.ListActivities(
					context.TODO(),
					&zwiftpower.ListActivitiesRequest{},
				)
				if err != nil {
					return nil, err
				}
				return resp.Activities, nil
			},
		}, "team": &graphql.Field{
			Type: TeamType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := zpp.zpClient.ListTeams(
					context.TODO(),
					&zwiftpower.ListTeamsRequest{},
				)
				if err != nil {
					return nil, err
				}
				for _, t := range resp.Teams {
					if t.TeamId == p.Args["id"].(string) {
						return t, nil
					}
				}
				return nil, nil
			},
		},
		"event": &graphql.Field{
			Type: EventType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				resp, err := zpp.zpClient.ListEvents(
					context.TODO(),
					&zwiftpower.ListEventsRequest{},
				)
				if err != nil {
					return nil, err
				}
				for _, e := range resp.Events {
					if e.Zid == p.Args["id"].(string) {
						return e, nil
					}
				}
				return nil, nil
			},
		},
	},
})

var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"zid": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if event, ok := p.Source.(*zwiftpower.Event); ok {
					return event.Zid, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if event, ok := p.Source.(*zwiftpower.Event); ok {
					return event.Name, nil
				}
				return nil, nil
			},
		},
	},
})

var TeamType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Team",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if team, ok := p.Source.(*zwiftpower.Team); ok {
					return team.TeamId, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if team, ok := p.Source.(*zwiftpower.Team); ok {
					return team.Tln, nil
				}
				return nil, nil
			},
		},
		"contact": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if team, ok := p.Source.(*zwiftpower.Team); ok {
					if team.Contact == "" {
						return nil, nil
					}
					return team.Contact, nil
				}
				return nil, nil
			},
		},
		"tag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if team, ok := p.Source.(*zwiftpower.Team); ok {
					return team.Tag, nil
				}
				return nil, nil
			},
		},
		"riders": &graphql.Field{
			Type: graphql.NewList(RiderType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if team, ok := p.Source.(*zwiftpower.Team); ok {
					resp, err := zpp.zpClient.GetTeamRiders(
						context.TODO(),
						&zwiftpower.GetTeamRidersRequest{
							Id: team.TeamId,
						},
					)
					if err != nil {
						return nil, err
					}
					return resp.Riders, nil
				}
				return nil, nil
			},
		},
	},
})

var RiderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Rider",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if event, ok := p.Source.(*zwiftpower.Rider); ok {
					return event.Name, nil
				}
				return nil, nil
			},
		},
	},
})

var LeagueType = graphql.NewObject(graphql.ObjectConfig{
	Name: "League",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if league, ok := p.Source.(*zwiftpower.League); ok {
					return league.Id, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if league, ok := p.Source.(*zwiftpower.League); ok {
					return league.Name, nil
				}
				return nil, nil
			},
		},
	},
})

var ActivityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Activity",
	Fields: graphql.Fields{
		// "id": &graphql.Field{
		// 	Type: graphql.String,
		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// 		if league, ok := p.Source.(*zwiftpower.Activity); ok {
		// 			return league.Id, nil
		// 		}
		// 		return nil, nil
		// 	},
		// },
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if league, ok := p.Source.(*zwiftpower.Activity); ok {
					return league.Title, nil
				}
				return nil, nil
			},
		},
	},
})
