package main

import "fmt"

// 新しい構造体型の定義。各フィールドの識別子とその型名を指定。
type Person struct {
    name string // 名前。
    age int // 年齢。
}

// メソッドの定義。
func (m *Person) PrintName() {
    fmt.Printf("Name = %s\n", m.name)
    //fmt.Printf("Name (%T) = %#v\n", m.name, m.name)
}

func (m *Person) PrintAge() {
    fmt.Printf("Age = %d\n", m.age)
    //fmt.Printf("Age (%T) = %#v\n", m.age, m.age)
}

func main() {
    var person Person
    person.name = "John Smith"
    person.age = 24
    // 以下のように書くこともできる。
    //person := Person{ "John Smith", 24 }
    person.PrintName()
    person.PrintAge()
    // 以下のように書式指定すると、型名とリテラル表現を出力できる。
    fmt.Printf("person (%T) = %#v\n", person, person)
}
