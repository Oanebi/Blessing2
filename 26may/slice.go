package main
import "fmt"

func main() {
    // 3D Slice: [Floor][Row][Column]
    building := [][][]string{
        // Floor 0
        {
            {"Room 101", "Room 102"},
            {"Room 103", "Room 104"},
        },
        // Floor 1
        {
            {"Room 201", "Room 202"},
            {"Room 203", "Room 204"},
        },
    }

    // To print "Room 203" (Floor 1, Row 1, Column 0)
    fmt.Println(building[0][0][0]) 
}