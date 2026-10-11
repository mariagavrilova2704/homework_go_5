package interfaces

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	ID   int
	Name string
}

// TODO: String должен вернуть пользователя в формате "user#<ID>:<Name>".
func (u User) String() string {
	str := fmt.Sprintf("user#%d:%s", u.ID, u.Name)
	return str
}

// TODO: FormatOne должен вызвать String у переданного значения. Если значение nil, верни пустую строку.
func FormatOne(value fmt.Stringer) string {
	if value == nil {
		return ""
	}
	return value.String()
}

// TODO: FormatMany должен вернуть строки для всех элементов в исходном порядке. nil элементы пропускаются.
func FormatMany(values []fmt.Stringer) []string {

	if values == nil {
		return []string{}
	}
	stringValues := make([]string, 0, len(values))
	for _, value := range values {
		if value == nil {
			continue
		}
		stringValues = append(stringValues, value.String())
	}
	return stringValues
}

type EmailNotifier struct {
	Address string
}

// TODO: Notify должен вернуть "email:<Address>:<message>".
func (n EmailNotifier) Notify(message string) string {
	str := fmt.Sprintf("email:%s:%s", n.Address, message)
	return str
}

type PushNotifier struct {
	DeviceID string
}

// TODO: Notify должен вернуть "push:<DeviceID>:<message>".
func (n PushNotifier) Notify(message string) string {
	str := fmt.Sprintf("push:%s:%s", n.DeviceID, message)
	return str
}

type Notifier interface {
	Notify(message string) string
}

// TODO: SendNotification должен отправить одно сообщение через Notifier. Если Notifier nil, верни пустую строку.
func SendNotification(n Notifier, message string) string {
	if n == nil {
		return ""
	}
	return n.Notify(message)
}

// TODO: SendBatch должен отправить все сообщения и вернуть результаты в том же порядке. nil Notifier даёт пустой результат.
func SendBatch(n Notifier, messages []string) []string {
	if n == nil {
		return []string{}
	}
	newBatch := make([]string, 0, len(messages))
	for _, message := range messages {
		newBatch = append(newBatch, n.Notify(message))
	}
	return newBatch
}

type StaticLoader struct {
	Value string
	Err   error
}

// TODO: Load должен вернуть поля Value и Err без дополнительной обработки.
func (l StaticLoader) Load() (string, error) {
	return l.Value, l.Err
}

type Loader interface {
	Load() (string, error)
}

// TODO: LoadUpper должен загрузить строку, привести её к верхнему регистру и вернуть ошибку без потери, если загрузка не удалась.
func LoadUpper(l Loader) (string, error) {
	if l == nil {
		return "", errors.New("Nil loader")
	}

	str, err := l.Load()
	if err != nil {
		return "", err
	}
	str = strings.ToUpper(str)
	return str, nil
}

// TODO: Describe должен вернуть описание динамического типа: nil, string:<value>, user:<String>, notifier:<Notify("ping")>, loader:<loaded>, error:<text> или unknown:<T>.
func Describe(value any) string {

	switch value := value.(type) {
	case nil:
		return "nil"
	case string:
		return fmt.Sprintf("string:%s", value)
	case User:
		return fmt.Sprintf("user:%s", value.String())
	case Notifier:
		return fmt.Sprintf("notifier:%s", value.Notify("ping"))
	case Loader:
		str, err := value.Load()
		if err != nil {
			return fmt.Sprintf("loader-error:%s", err.Error())
		}
		return fmt.Sprintf("loader:%s", str)
	case error:
		return fmt.Sprintf("error:%s", value.Error())
	default:
		return fmt.Sprintf("unknown:%T", value)
	}
}

// TODO: OnlyErrors должен вернуть тексты только ненулевых ошибок в исходном порядке.
func OnlyErrors(values []error) []string {
	errorsText := make([]string, 0, len(values))

	for _, value := range values {
		if value != nil {
			errorsText = append(errorsText, value.Error())
		}
	}

	return errorsText
}

// TODO: JoinStringers должен отформатировать ненулевые Stringer и объединить их через sep.
func JoinStringers(values []fmt.Stringer, sep string) string {

	finalResult := make([]string, 0, len(values))

	for _, value := range values {
		if value != nil {
			finalResult = append(finalResult, value.String())
		}
	}
	return strings.Join(finalResult, sep)
}

// TODO: FirstNonEmptyStringer должен вернуть первую непустую строку среди Stringer. nil и пустые строки пропускаются.
func FirstNonEmptyStringer(values []fmt.Stringer) string {
	if values == nil {
		return ""
	}
	var str string
	for _, value := range values {
		if value != nil && value.String() != "" {
			str = value.String()
			break
		}
	}
	return str
}

// TODO: CountStringers должен посчитать, сколько элементов реально реализуют fmt.Stringer.
func CountStringers(values []any) int {
	count := 0
	for _, value := range values {
		if _, ok := value.(fmt.Stringer); ok {
			count++
		}
	}
	return count
}

// TODO: NormalizeAndFormat должен обрезать пробелы вокруг String() и схлопнуть пустой результат в "empty".
func NormalizeAndFormat(value fmt.Stringer) string {
	if value == nil {
		return "empty"
	}

	str := strings.TrimSpace(value.String())

	if str == "" {
		return "empty"
	}
	return str
}

func Example() string {
	var out strings.Builder
	fmt.Fprintf(&out, "user: %s\n", FormatOne(User{ID: 7, Name: "Maria"}))
	fmt.Fprintf(&out, "email: %s\n", SendNotification(EmailNotifier{Address: "team@example.com"}, "hello"))
	fmt.Fprintf(&out, "push: %s\n", SendNotification(PushNotifier{DeviceID: "ios-1"}, "hello"))
	fmt.Fprintf(&out, "describe: %s", Describe(User{ID: 8, Name: "Alex"}))
	return out.String()
}

//check ci/cd
