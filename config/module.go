package config

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

type Module interface {
	Name() string
	Complete() bool
	Run()
}

var modules []Module

func SetupModules(cmd *cobra.Command, args []string) {
	ns := make([]string, len(modules)+1)
	for i, m := range modules {
		ns[i] = m.Name()
	}
	ns[len(modules)] = "Quit"

	for {
		cs := make([]bool, len(modules))
		for i, m := range modules {
			cs[i] = m.Complete()
		}

		var mi int
		p := &survey.Select{
			Options: ns,
			Help:    "\u2714 - Complete, \u2718 - Incomplete",
			Description: func(_ string, i int) string {
				if i == len(modules) {
					return ""
				}
				if cs[i] {
					return "\u2714"
				}
				return "\u2718"
			},
		}
		survey.AskOne(p, &mi)

		if mi == len(modules) {
			return
		}

		modules[mi].Run()
	}
}

func Register(m Module) {
	modules = append(modules, m)
	slices.SortFunc(modules, func(a, b Module) bool {
		return strings.ToLower(a.Name()) < strings.ToLower(b.Name())
	})
}
