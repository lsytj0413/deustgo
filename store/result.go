// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

const (
	// Get is const value of Get Action
	Get = "get"
	// Set is const value of Set Action
	Set = "set"
	// Update is const value of Update Action
	Update = "update"
	// Create is const value of Create Action
	Create = "create"
	// Delete is const value of Delete Action
	Delete = "delete"
)

// Result is basic Action Result
type Result struct {
	Action   string
	CurrNode *Node
	PrevNode *Node
}

func newResult(action string, key string) *Result {
	return &Result{
		Action: action,
		CurrNode: &Node{
			Key: key,
		},
	}
}

func (r *Result) Clone() *Result {
	return &Result{
		Action:   r.Action,
		CurrNode: r.CurrNode.Clone(),
		PrevNode: r.PrevNode.Clone(),
	}
}
