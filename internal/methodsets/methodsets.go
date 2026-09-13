package methodsets

import (
	"fmt"
	"strings"
)

type Counter struct {
	Name  string
	Value int
}

// TODO: NewCounter должен создать Counter с переданными Name и Value.
func NewCounter(name string, value int) Counter {
	return Counter{
		Name:  name,
		Value: value,
	}
}

// TODO: Snapshot должен вернуть текущее значение счётчика. Value receiver не меняет объект.
func (c Counter) Snapshot() int {
	return c.Value
}

// TODO: Label должен вернуть "<Name>=<Value>".
func (c Counter) Label() string {
	return fmt.Sprintf("%s=%d", c.Name, c.Value)
}

// TODO: Add должен увеличить исходный Counter на delta. nil receiver нужно спокойно игнорировать.
func (c *Counter) Add(delta int) {
	if c == nil {
		return
	}
	c.Value = c.Value + delta

}

// TODO: Reset должен сбросить исходный Counter в 0. nil receiver нужно спокойно игнорировать.
func (c *Counter) Reset() {
	if c == nil {
		return
	}
	c.Value = 0
}

type Snapshoter interface{ Snapshot() int }
type Labeler interface{ Label() string }
type Adder interface{ Add(delta int) }
type Resetter interface{ Reset() }

// TODO: UseSnapshoter должен вернуть Snapshot или 0 для nil интерфейса.
func UseSnapshoter(s Snapshoter) int {
	if s == nil {
		return 0
	}
	return s.Snapshot()
}

// TODO: UseLabeler должен вернуть Label или пустую строку для nil интерфейса.
func UseLabeler(l Labeler) string {
	if l == nil {
		return ""
	}
	return l.Label()
}

// TODO: UseAdder должен вызвать Add и вернуть новое значение, если объект также умеет Snapshot. Иначе вернуть 0.
func UseAdder(a Adder, delta int) int {
	if a == nil {
		return 0
	}
	a.Add(delta)

	if t, ok := a.(Snapshoter); ok {
		return t.Snapshot()
	}

	return 0
}

// TODO: IsSnapshoter должен проверить, реализует ли значение интерфейс Snapshoter.
func IsSnapshoter(value any) bool {
	if _, ok := value.(Snapshoter); ok {
		return true
	}
	return false
}

// TODO: IsAdder должен проверить, реализует ли значение интерфейс Adder.
func IsAdder(value any) bool {
	if _, ok := value.(Adder); ok {
		return true
	}
	return false
}

// TODO: CloneAndAdd должен вернуть копию Counter с увеличенным Value, не меняя оригинал.
func CloneAndAdd(c Counter, delta int) Counter {
	a := c
	a.Value = a.Value + delta
	return a
}

// TODO: AddInPlace должен изменить исходный Counter по указателю и вернуть новое значение. nil -> 0.
func AddInPlace(c *Counter, delta int) int {
	if c == nil {
		return 0
	}
	c.Value = c.Value + delta
	return c.Value
}

type Profile struct {
	Name string
	Age  int
}

// TODO: Display должен вернуть "<Name>(<Age>)".
func (p Profile) Display() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

// TODO: Rename должен изменить Name у исходного Profile. nil receiver нужно игнорировать.
func (p *Profile) Rename(name string) {
	if p == nil {
		return
	}
	p.Name = name
}

type Renamer interface{ Rename(name string) }

// TODO: IsRenamer должен проверить, реализует ли значение Renamer.
func IsRenamer(value any) bool {
	if _, ok := value.(Renamer); ok {
		return true
	}
	return false
}

func Example() string {
	var out strings.Builder
	counter := NewCounter("orders", 10)
	fmt.Fprintf(&out, "value: %s\n", UseLabeler(counter))
	fmt.Fprintf(&out, "pointer before: %s\n", UseLabeler(&counter))
	UseAdder(&counter, 5)
	fmt.Fprintf(&out, "pointer after: %s\n", UseLabeler(&counter))
	profile := Profile{Name: "Maria", Age: 25}
	profile.Rename("Masha")
	fmt.Fprintf(&out, "profile: %s", profile.Display())
	return out.String()
}
