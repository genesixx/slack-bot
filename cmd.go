package main

type CMD struct {
	Command         string
	Option          string
	Channel         string
	User            string
	Timestamp       string
	ThreadTimestamp string
}

type cmdFunc func(cmd *CMD) (*Response, error)

type Cog struct {
	cmd         string
	helper      string
	description string
	function    cmdFunc
}

var (
	commands = make(map[string]*Cog)
)

//RegisterCommand Register command
func RegisterCommand(cmd, helper, description string, function cmdFunc) {
	commands[cmd] = &Cog{
		cmd:         cmd,
		helper:      helper,
		description: description,
		function:    function,
	}
}

func (b *Bot) handleCMD(cmd *CMD) {
	c := commands[cmd.Command]
	if c == nil {
		logger.Printf("Command %s not found\n", cmd.Command)
	}

	logger.Printf("%#v\n", cmd)

	resp, err := c.function(cmd)

	if err != nil {
		logger.Printf("Command %s error: %s\n", cmd.Command, err)
	}

	b.sendResponse(resp)
}