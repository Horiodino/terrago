package Alerts

// import (
// 	"fmt"

// 	"github.com/slack-go/slack"
// )

// func Notify(podname string, namespace string, message string, channel string, logs string) error {

// 	return nil
// }

// func SendFailureAlert() error {

// 	return nil
// }

// // use this function to send notification to  alternative channel channel or discord or email if slack is not available or any error occurs
// func NotificationFailuarBackup() {

// }

// func main() {
// 	api := slack.New("YOUR_TOKEN_HERE")
// 	user, err := api.GetUserInfo("U023BECGF")
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 		return
// 	}
// 	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
// }

// func slack() {
// 	api := slack.New("YOUR_TOKEN_HERE")
// 	// If you set debugging, it will log all requests to the console
// 	// Useful when encountering issues
// 	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
// 	groups, err := api.GetUserGroups(slack.GetUserGroupsOptionIncludeUsers(false))
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 		return
// 	}
// 	for _, group := range groups {
// 		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
// 	}
// }
