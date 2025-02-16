// SPDX-License-Identifier: 0BSD
// SPDX-FileCopyrightText: 2025 Hajime Hoshi

package foo

import "fmt"

var message = "Original"

func Foo() {
	fmt.Println(message)
}
