// Паттерн Состояние (State)
//
// В примере описываются варианты оповещения пользователя мобильного телефона.
// На самом деле, хорошо было бы реализовать смену состояния внутри системы
// оповещения (ModileAlert), покаким либо критериям. Например, задействуя файл конфигруции.
// Но для упращения примера, используется внешняя смена состояния, посредствам метода SetState()

package state

// Тип ModileAlertStater, описывает общий интерфейс для различных состояний
type ModileAlertStater interface {
	Alert() string
}

// Тип ModileAlert, реализует оповещение в зависимости от своего состояния
type ModileAlert struct {
	state ModileAlertStater
}

// Оповещение в зависимости от внутреннего состояния
func (self *ModileAlert) Alert() string {
	return self.state.Alert()
}

// Меняет состояние
func (self *ModileAlert) SetState(state ModileAlertStater) {
	self.state = state
}

func NewModileAlert() *ModileAlert {
	return &ModileAlert{state: &MobileAlertVibration{}}
}

// Тип MobileAlertVibration, реализует оповещение только вибрацией
type MobileAlertVibration struct {
}

// Оповещение вибрацией
func (self *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// Тип MobileAlertSong, реализует оповещение только звуковым сигналом
type MobileAlertSong struct {
}

// Оповещение звуком
func (self *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}
