// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Выходные параметры для проверки авторизации
type AuthcheckInput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

// Выходные параметры для проверки авторизации
type AuthcheckOutput struct {
	//  Информация о юзере
	UserInfo *User `json:"userInfo"`
}

// Мутации связанные с авторизацией
type AuthorizationMutations struct {
	//  Регистрация
	Signup *SignupOutput `json:"signup"`
	//  Авторизация
	Signin *SigninOutput `json:"signin"`
	//  Выход
	Logout *LogoutOutput `json:"logout,omitempty"`
}

type CoordsInput struct {
	//  Координата X
	X string `json:"x"`
	//  Координата Y
	Y string `json:"y"`
	//  Координата Z
	Z string `json:"z"`
}

type CoordsResults struct {
	//  Координата X
	X string `json:"x"`
	//  Координата Y
	Y string `json:"y"`
	//  Координата Z
	Z string `json:"z"`
}

// Входные параметры для create device
type CreateDeviceInput struct {
	//  Название девайса
	Name string `json:"Name"`
	//  Токен
	Token string `json:"Token"`
	//  Описание
	Description *string `json:"Description,omitempty"`
	//  Координаты
	Coords *CoordsInput `json:"Coords"`
}

// Выходные параметры для create device
type CreateDeviceOutput struct {
	Device *Device `json:"device"`
}

// Входные параметры для создания спутника
type CreateSatelliteInput struct {
	//  Внешний индетификатор спутника
	ExternalSatelliteID string `json:"ExternalSatelliteId"`
	//  Название спутника
	SatelliteName string `json:"SatelliteName"`
}

// Выходные параметры для создания спутника
type CreateSatelliteOutput struct {
	Satellite *SatelliteInfo `json:"satellite"`
}

// Входные параметры для create task
type CreateTaskInput struct {
	//  Индентификатор спутника
	SatelliteID string `json:"satelliteId"`
	//  Тип сигнала
	SignalType SignalType `json:"signalType"`
	//  Тип группировки
	GroupingType GroupingType `json:"groupingType"`
	//  Время начала
	StartAt time.Time `json:"startAt"`
	//  Время конца
	EndAt time.Time `json:"endAt"`
}

// Выходные параметры для create task
type CreateTaskOutput struct {
	Task *Task `json:"task"`
}

// Входные параметры для delete task
type DeleteTaskInput struct {
	//  Индетификатор
	ID string `json:"id"`
}

// Входные параметры для delete task
type DeleteTaskOutput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

type DeviceFilter struct {
	//  Индетификатор
	Ids []string `json:"Ids,omitempty"`
	//  Название девайса
	Names []string `json:"Names,omitempty"`
	//  Токен
	Tokens []string `json:"Tokens,omitempty"`
}

type DevicePagination struct {
	//  Загруженные элементы
	Items []*Device `json:"items,omitempty"`
}

type Gnss struct {
	//  Индентификатор
	ID string `json:"Id"`
	//  id спутника
	SatelliteID string `json:"SatelliteId"`
	//  Координаты спутника
	Coordinates *CoordsResults `json:"Coordinates"`
	//  Время создания
	CreatedAt time.Time `json:"CreatedAt"`
}

type GNSSFilter struct {
	//  Фильтр по индетификаторам
	Coordinates *CoordsInput `json:"Coordinates"`
}

type GNSSPagination struct {
	//  Загруженные элементы
	Items []*Gnss `json:"items,omitempty"`
}

// Мутации связанные с gnss
type GnssMutations struct {
	//  Обновить device
	UpdateDevice *UpdateDeviceOutput `json:"updateDevice"`
	//  Создать device
	CreateDevice *CreateDeviceOutput `json:"createDevice"`
	//  Создать task
	CreateTask *CreateTaskOutput `json:"createTask"`
	//  Обновить task
	UpdateTask *UpdateTaskOutput `json:"updateTask"`
	//  Удалить task
	DeleteTask *DeleteTaskOutput `json:"deleteTask"`
	//  Создать спутник
	CreateSatellite *CreateSatelliteOutput `json:"createSatellite"`
}

type Header struct {
	RinexVersion       string    `json:"rinex_version"`
	FileType           string    `json:"file_type"`
	PgmRunByDate       string    `json:"pgm_run_by_date"`
	Comments           []string  `json:"comments"`
	MarkerName         string    `json:"marker_name"`
	MarkerNumber       string    `json:"marker_number"`
	ObserverAgency     string    `json:"observer_agency"`
	RecInfo            string    `json:"rec_info"`
	AntInfo            string    `json:"ant_info"`
	ApproxPositionXyz  []float64 `json:"approx_position_xyz"`
	AntennaDeltaHen    []float64 `json:"antenna_delta_hen"`
	WavelengthFactL1L2 []int     `json:"wavelength_fact_l1_l2"`
	TypesOfObs         []string  `json:"types_of_obs"`
	Interval           float64   `json:"interval"`
	TimeOfFirstObs     string    `json:"time_of_first_obs"`
	TimeOfLastObs      string    `json:"time_of_last_obs"`
	EndOfHeader        bool      `json:"end_of_header"`
}

// Выходные параметры для выхода
type LogoutInput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

// Выходные параметры для выхода
type LogoutOutput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

type Mutation struct {
}

type Observation struct {
	Time       string       `json:"time"`
	EpochFlag  int          `json:"epoch_flag"`
	Satellites []*Satellite `json:"satellites"`
}

type Query struct {
}

type RinexInput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

type RinexPagination struct {
	Items []*RinexResults `json:"items,omitempty"`
}

type RinexResults struct {
	Header       *Header        `json:"header"`
	Observations []*Observation `json:"observations"`
}

type Satellite struct {
	SatelliteID  string   `json:"satellite_id"`
	Observations []string `json:"observations"`
}

type SatellitesFilter struct {
	//  Индетификатор
	IDS []string `json:"IdS,omitempty"`
	//  Внешний индетификатор спутника
	ExternalSatelliteIds []string `json:"ExternalSatelliteIds,omitempty"`
	//  Название спутника
	SatelliteNames []string `json:"SatelliteNames,omitempty"`
}

type SatellitesPagination struct {
	//  Загруженные элементы
	Items []*SatelliteInfo `json:"items,omitempty"`
}

// Входные параметры для авторизации
type SigninInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Выходные параметры для авторизации
type SigninOutput struct {
	//  Информация о юзере
	UserInfo *User `json:"userInfo"`
}

// Входные параметры для регистрации
type SignupInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Выходные параметры для регистрации
type SignupOutput struct {
	//  Информация о юзере
	UserInfo *User `json:"userInfo"`
}

type TaskFilter struct {
	//  Фильтр по индетификаторам
	Ids           []string       `json:"ids,omitempty"`
	SatelliteIds  []string       `json:"satelliteIds,omitempty"`
	SatelliteName []string       `json:"satelliteName,omitempty"`
	SignalType    []SignalType   `json:"signalType,omitempty"`
	GroupingType  []GroupingType `json:"groupingType,omitempty"`
	StartAt       *time.Time     `json:"startAt,omitempty"`
	EndAt         *time.Time     `json:"endAt,omitempty"`
}

type TaskPagination struct {
	//  Загруженные элементы
	Items []*Task `json:"items,omitempty"`
}

// Входные параметры для update device
type UpdateDeviceInput struct {
	//  Индетификатор
	ID string `json:"Id"`
	//  Название девайса
	Name string `json:"Name"`
	//  Токен
	Token string `json:"Token"`
	//  Описание
	Description *string `json:"Description,omitempty"`
	//  Координаты
	Coords *CoordsInput `json:"Coords"`
}

// Выходные параметры для update device
type UpdateDeviceOutput struct {
	Device *Device `json:"device"`
}

// Входные параметры для update task
type UpdateTaskInput struct {
	//  Индетификатор
	ID string `json:"id"`
	//  Индетификатор спутника
	SatelliteID string `json:"satelliteId"`
	//  Тип сигнала
	SignalType SignalType `json:"signalType"`
	//  Тип группировки
	GroupingType GroupingType `json:"groupingType"`
	//  Время начала
	StartAt time.Time `json:"startAt"`
	//  Время конца
	EndAt time.Time `json:"endAt"`
}

// Выходные параметры для update task
type UpdateTaskOutput struct {
	Task *Task `json:"task"`
}

// Бизнес ошибки
type Error string

const (
	//  Уже cуществует
	ErrorAlreadyExists Error = "ALREADY_EXISTS"
	//  Не авторизован
	ErrorNotAuthorized Error = "NOT_AUTHORIZED"
	//  Не найден
	ErrorNotFound Error = "NOT_FOUND"
	//  Нет прав
	ErrorPermissionDenied Error = "PERMISSION_DENIED"
	//  Ошибка сервера
	ErrorInternalError Error = "INTERNAL_ERROR"
	//  Ошибка запроса
	ErrorBadRequest Error = "BAD_REQUEST"
)

var AllError = []Error{
	ErrorAlreadyExists,
	ErrorNotAuthorized,
	ErrorNotFound,
	ErrorPermissionDenied,
	ErrorInternalError,
	ErrorBadRequest,
}

func (e Error) IsValid() bool {
	switch e {
	case ErrorAlreadyExists, ErrorNotAuthorized, ErrorNotFound, ErrorPermissionDenied, ErrorInternalError, ErrorBadRequest:
		return true
	}
	return false
}

func (e Error) String() string {
	return string(e)
}

func (e *Error) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Error(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Error", str)
	}
	return nil
}

func (e Error) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GroupingType string

const (
	//  Неизвестно
	GroupingTypeGroupingTypeUnknown GroupingType = "GROUPING_TYPE_UNKNOWN"
	//  GPS
	GroupingTypeGroupingTypeGps GroupingType = "GROUPING_TYPE_GPS"
	//  Glonass
	GroupingTypeGroupingTypeGlonass GroupingType = "GROUPING_TYPE_GLONASS"
	//  Galileo
	GroupingTypeGroupingTypeGalileo GroupingType = "GROUPING_TYPE_GALILEO"
	//  Beidou
	GroupingTypeGroupingTypeBeidou GroupingType = "GROUPING_TYPE_BEIDOU"
)

var AllGroupingType = []GroupingType{
	GroupingTypeGroupingTypeUnknown,
	GroupingTypeGroupingTypeGps,
	GroupingTypeGroupingTypeGlonass,
	GroupingTypeGroupingTypeGalileo,
	GroupingTypeGroupingTypeBeidou,
}

func (e GroupingType) IsValid() bool {
	switch e {
	case GroupingTypeGroupingTypeUnknown, GroupingTypeGroupingTypeGps, GroupingTypeGroupingTypeGlonass, GroupingTypeGroupingTypeGalileo, GroupingTypeGroupingTypeBeidou:
		return true
	}
	return false
}

func (e GroupingType) String() string {
	return string(e)
}

func (e *GroupingType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GroupingType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GroupingType", str)
	}
	return nil
}

func (e GroupingType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SignalType string

const (
	//  Неизвестно
	SignalTypeSignalTypeUnknown SignalType = "SIGNAL_TYPE_UNKNOWN"
	//  L1
	SignalTypeSignalTypeL1 SignalType = "SIGNAL_TYPE_L1"
	//  L2
	SignalTypeSignalTypeL2 SignalType = "SIGNAL_TYPE_L2"
	//  L3
	SignalTypeSignalTypeL3 SignalType = "SIGNAL_TYPE_L3"
)

var AllSignalType = []SignalType{
	SignalTypeSignalTypeUnknown,
	SignalTypeSignalTypeL1,
	SignalTypeSignalTypeL2,
	SignalTypeSignalTypeL3,
}

func (e SignalType) IsValid() bool {
	switch e {
	case SignalTypeSignalTypeUnknown, SignalTypeSignalTypeL1, SignalTypeSignalTypeL2, SignalTypeSignalTypeL3:
		return true
	}
	return false
}

func (e SignalType) String() string {
	return string(e)
}

func (e *SignalType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SignalType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SignalType", str)
	}
	return nil
}

func (e SignalType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
