package main

import "testing"

func TestArea(t *testing.T) {

    checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    }

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{1,2}
        checkArea(t, rectangle, 2)
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{2}
        checkArea(t, circle, 12.56)
    })

}

func TestR(t *testing.T){
    // rectangle := Rectangle{1,2}
    want:=2.00
    var s Shape
    s = Rectangle{1,2}
    got := s.Area()
    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }

}

func TestC(t *testing.T){
    // rectangle := Rectangle{1,2}
    want:=12.56
    var s Shape
    s = Circle{2}
    got := s.Area()
    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }

}

func TestTableArea(t *testing.T) {

    areaTests := []struct {
        name    string
        shape   Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{radius: 2}, hasArea: 12.56},
        
    }

    for _, tt := range areaTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
            }
        })

    }

}