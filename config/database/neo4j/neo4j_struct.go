package neo4j

import "github.com/mindstand/gogm/v2"

type Family struct {
	gogm.BaseNode

	Name     string `gogm:"name=name" json:"name"`
	Username string `gogm:"name=username" json:"username"`
	Uuid     string `gogm:"name=uuid"`

	Group    *Group    `gogm:"direction=outgoing;relationship=OWNER"`
	Children []*Person `gogm:"direction=incoming;relationship=CHILDREN"`
	Wife     []*Person `gogm:"direction=incoming;relationship=WIFE"`
	Husband  []*Person `gogm:"direction=incoming;relationship=HUSBAND"`
}

type Group struct {
	gogm.BaseNode

	Name string `gogm:"name=name"`
	Uuid string `gogm:"name=uuid"`

	Families []*Family `gogm:"direction=incoming;relationship=OWNER"`
}

type Person struct {
	gogm.BaseNode

	Name string `gogm:"name=name"`
	Uuid string `gogm:"name=uuid"`

	Husband  *Family `gogm:"direction=outgoing;relationship=HUSBAND"`
	Wife     *Family `gogm:"direction=outgoing;relationship=WIFE"`
	Children *Family `gogm:"direction=outgoing;relationship=CHILDREN"`
}

type ContingentUponEdge struct {
	gogm.BaseNode

	Start       *Group  `json:"group"`
	End         *Person `json:"person"`
	Criticality string  `gogm:"name=criticality"`
}
