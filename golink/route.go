package golink

import (
    "github.com/QLeelulu/goku"
)

// routes
var Routes []*goku.Route = []*goku.Route{
    &goku.Route{
        Name:     "static",
        IsStatic: true,
        Pattern:  "/assets/(.*)",
    },
    &goku.Route{
        Name:    "topicInfo",
        Pattern: "/t/{name}",
        Default: map[string]string{"controller": "topic", "action": "show"},
    },
    &goku.Route{
        Name:       "edit",
        Pattern:    "/{controller}/{id}/{action}",
        Default:    map[string]string{"action": "show"},
        Constraint: map[string]string{"id": "\\d+"},
    },
    &goku.Route{
        Name:       "vote",
        Pattern:    "/{controller}/{action}/{id}/{votetype}/{topid}", //1 == vote up, 2 == votet down
        Default:    map[string]string{"action": "link"},
        Constraint: map[string]string{"id": "\\d+", "topid": "\\d+", "votetype": "\\d+"},
    },
    &goku.Route{
        Name:    "default",
        Pattern: "/{controller}/{action}",
        Default: map[string]string{"controller": "home", "action": "index"},
    },
}