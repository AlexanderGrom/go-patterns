// Паттерн Команда (Command)
//

package command

// Тип Command, описывает общий интерфейс для команд
type Command interface {
	Execute() string
}

// Тип ToggleOnCommand, реализует команду включения тумблера
type ToggleOnCommand struct {
	receiver *Receiver
}

// Выполнение команды получателем
func (self *ToggleOnCommand) Execute() string {
	return self.receiver.ToggleOn()
}

// Тип ToggleOffCommand, реализует команду выключения тумблера
type ToggleOffCommand struct {
	receiver *Receiver
}

// Выполнение команды получателем
func (self *ToggleOffCommand) Execute() string {
	return self.receiver.ToggleOff()
}

// Тип Receiver, реализует получателя команд
// Реализует набор действие которые получатель должен
// выполнять, взависимости от полученой команды
type Receiver struct {
}

// Действие на команду "поднять тумблер (ToggleOnCommand)"
func (self *Receiver) ToggleOn() string {
	return "Toggle On"
}

// Действие на команду "опустить тумблер (ToggleOffCommand)"
func (self *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Тип Invoker, реализует инициатора команды
type Invoker struct {
	commands []Command
}

func (self *Invoker) StoreCommand(command Command) {
	self.commands = append(self.commands, command)
}

func (self *Invoker) UnStoreCommand() {
	if len(self.commands) != 0 {
		self.commands = self.commands[:len(self.commands)-1]
	}
}

func (self *Invoker) Execute() string {
	var result string
	for _, command := range self.commands {
		result += command.Execute() + "\n"
	}
	return result
}
