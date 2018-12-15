// Package command is an example of the Command Pattern.
package command

// Command provides a command interface.
type Command interface {
	Execute() string
}

// ToggleOnCommand implements the Command interface.
type ToggleOnCommand struct {
	receiver *Receiver
}

// Execute command.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand implements the Command interface.
type ToggleOffCommand struct {
	receiver *Receiver
}

// Execute command.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Receiver implementation.
type Receiver struct {
}

// ToggleOn implementation.
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

// ToggleOff implementation.
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Invoker implementation.
type Invoker struct {
	commands []Command
}

// StoreCommand adds command.
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand removes command.
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Execute all commands.
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}
