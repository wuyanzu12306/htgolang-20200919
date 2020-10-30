package funcs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
)

// ShowUser show a user based on Id
func ShowUser(id int64) {
	t := tablewriter.NewWriter(os.Stdout)
	for _, user := range define.UserList {
		if user.Id == id {
			s := strconv.FormatInt(id, 10)
			p := fmt.Sprintf("%x", user.Passwd)
			t.Append([]string{s, user.Name, user.Cell, user.Address, user.Born.Format("2006.01.02"), p})
			//md5.Sum([]byte(user.Passwd))})
		}
	}
	t.Render()
}

// ShowUserList display the UserList's contents
func ShowUserList() {
	ul := &define.UserList
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)
	t.SetHeader([]string{"ID", "Name", "Cell", "Address", "Born", "Passwd"})
	for _, user := range *ul {
		id := strconv.FormatUint(uint64(user.Id), 10)
		p := fmt.Sprintf("%x", user.Passwd)
		t.Append([]string{id, user.Name, user.Cell, user.Address,
			user.Born.Format("2006.01.02"), p})
	}
	t.Render()
}
