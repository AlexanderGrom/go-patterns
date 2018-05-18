// Паттерн Состояние (State)
//
// В примере описываются варианты оповещения пользователя мобильного телефона.
// На самом деле, хорошо было бы реализовать смену состояния внутри системы
// оповещения (MobileAlert), по каким-либо критериям. Например, задействуя файл конфигруции.
// Но для упращения примера, используется внешняя смена состояния, посредствам метода SetState()

package state

// Тип MobileAlertStater, описывает общий интерфейс для различных состояний
type MobileAlertStater interface {
	Alert() string
}

// Тип MobileAlert, реализует оповещение в зависимости от своего состояния
type MobileAlert struct {
	state MobileAlertStater
}

// Оповещение в зависимости от внутреннего состояния
func (self *MobileAlert) Alert() string {
	return self.state.Alert()
}

// Меняет состояние
func (self *MobileAlert) SetState(state MobileAlertStater) {
	self.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
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
