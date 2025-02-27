/*
 Copyright 2025 The Nephio Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 You may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package nmath

import "testing"

func TestAdd(t *testing.T) {
	t.Log("TestAdd")

	result := Add(10, 10)
	expected := 20

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	t.Log("TestSubtract")
	result := Subtract(10, 10)
	expected := 0

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	t.Log("TestMultiply")
	result := Multiply(10, 10)
	expected := 100

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	t.Log("TestMultiply")
	result := Divide(10, 10)
	expected := 1

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
