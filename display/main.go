package display

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mgutz/ansi"
)

var Opts = survey.WithIcons(func(is *survey.IconSet) {
	is.Error.Text = "\u22A5"
	is.Error.Format = ansi.LightRed
	is.Help.Text = "\u2318"
	is.Help.Format = ansi.Cyan
	is.Question.Text = "\u21D2"
	is.Question.Format = ansi.Green

	is.SelectFocus.Text = "\u21E2"
	is.SelectFocus.Format = ansi.LightBlue

	is.MarkedOption.Text = "\u2299"
	is.MarkedOption.Format = ansi.Blue
	is.UnmarkedOption.Text = "\u2218"
	is.UnmarkedOption.Format = ansi.Blue
})
