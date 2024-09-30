package DP

import (
	"math/rand/v2"
)

type Pricer interface {
	Price()
}

type Microphone struct {
	Product_Price int
}

func (m *Microphone) Price() int {
	m.Product_Price = rand.IntN(8)
	return m.Product_Price
}

type Speacker struct {
	Product_Price int
}

func (m *Speacker) Price() int {
	m.Product_Price = rand.IntN(12)
	return m.Product_Price
}

type Key struct {
	Product_Price int
}

func (m *Key) Price() int {
	m.Product_Price = rand.IntN(2)
	return m.Product_Price
}

type Keyboard struct {
	Product_Price int
	initialized   bool
	keys          []Key
}

func (m *Keyboard) Initializer() {
	var app Key
	for i := 0; i < rand.IntN(12); i++ {
		m.keys = append(m.keys, app)
	}
	m.initialized = true
}

func (m *Keyboard) Price() int {
	if !m.initialized {
		m.Initializer()
	}
	for i, _ := range m.keys {
		m.Product_Price = m.Product_Price + m.keys[i].Price()
	}
	return m.Product_Price
}

type Smartphone struct {
	Product_Price int
	K_keyboard    Keyboard
	M_microphope  Microphone
	S_speaker     Speacker
}

func (m *Smartphone) Price() int {
	m.Product_Price = m.K_keyboard.Price() + m.M_microphope.Price() + m.S_speaker.Price()
	return m.Product_Price
}
