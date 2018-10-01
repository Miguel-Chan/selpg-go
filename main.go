// Copyright © 2018 Miguel Chan <vvchan@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/Miguel-Chan/selpg-go/cmd"
)

func main() {
	//startPage := pflag.IntP("start_page", "s", -1, "The first page to be selected.")
	//endPage := pflag.IntP("end_page", "e", -1, "The last page to be selected.")
	//if *startPage < 0 {
	//	fmt.Printf("%v: No valid start_page specified", os.Args[0])
	//}
	//if *endPage < 0 {
	//	fmt.Printf("%v: No valid end_page specified", os.Args[0])
	//}
	cmd.Execute()
}
