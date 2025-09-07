package slices

import "fmt"

type Part string

func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func Exercise() {
	assemblyLine := make([]Part, 3)
	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Sheet"
	fmt.Println("3 Parts:")
	showLine(assemblyLine)
	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyLine)
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced")
	showLine(assemblyLine)

}
