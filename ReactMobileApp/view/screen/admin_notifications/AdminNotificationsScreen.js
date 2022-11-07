import React from "react"
import { View, Text } from "react-native"
import NotificationsStyle from "../../Notifications.style"
import TopBar from "../../component/TopBar"
import handleGetPromotionToManagerRequests from "../../../event-handler/request_manager_promotion/HandleGetPromotionToManagerRequests"

// Displays all promotion to manager requests currently open
export default class AdminNotificationsScreen extends React.Component {
    state = {
        responseReceived: false,
        promotionToManagerRequests: {},
    }

    componentDidMount() {
        handleGetPromotionToManagerRequests()
        .then((requests) => {
            this.setState({
                responseReceived: true,
                promotionToManagerRequests: requests
            })
        })
    }
    
    render() {
        if (!this.state.responseReceived) {
            return (
                <>
                    <TopBar/>
                    <Text>
                        Loading...
                    </Text>
                </>
            )
        }

        console.log(this.state.promotionToManagerRequests) 
        console.log(typeof(this.state.promotionToManagerRequests))

        elements = []        

        for (var i = 0; i < this.state.promotionToManagerRequests.length; i++) {
            console.log(this.state.promotionToManagerRequests[i])
    
            requestElement = 
                <View key={this.state.promotionToManagerRequests[i].sender_username}
                    style={NotificationsStyle.container}>
                    <Text style={NotificationsStyle.titleText}>
                        Promotion to Manager Request
                    </Text>
                    <Text>
                        {this.state.promotionToManagerRequests[i].sender_username}
                    </Text>
                    <Text>
                        {this.state.promotionToManagerRequests[i].message}
                    </Text>
                </View>

            elements.push(requestElement)
        }
        
        
        return (
            <>
                <TopBar/>

                <View style={NotificationsStyle.container}>
                    <Text style={NotificationsStyle.pageTitleText}>
                        Admin Notifications
                    </Text>
                </View>

                <View>
                    {elements}
                </View>
            </>
        )
    }
}