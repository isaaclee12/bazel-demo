package main

import (
	"fmt"

	"github.com/isaaclee12/bazel-demo/mathutils"
	"github.com/isaaclee12/bazel-demo/stringutils"
)

func main() {
	result := mathutils.Add(5, 3)
	reversed := stringutils.Reverse("Bazel", true)
	fmt.Printf("5 + 3 = %d\n", result)
	fmt.Printf("Reversed: %s\n", reversed)
}
