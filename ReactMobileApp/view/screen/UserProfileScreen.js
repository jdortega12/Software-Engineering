import React from "react"
import {Text, View, TouchableOpacity, Image} from "react-native"
import TopBar from "../component/TopBar"
import ProfileStyle from "../Profile.style"
import handleGetUserProfile from "../../event-handler/HandleGetUserProfile"

// Screen for displaying a user's profile. Accepts bool isSelf for whether
// the profile to display belongs to the logged in user or not.
export default class UserProfileScreen extends React.Component {
    state = {
        responseReceived: false,
        userData: null,
    }

    componentDidMount(){
        handleGetUserProfile(this.props.username)
        .then((data) => {
            this.setState({
                responseReceived: true,
                userData: data,
            })
        })
    }

    capitalizeFirstLetter(str){
        return str[0].toUpperCase() + str.substring(1)
    }

    render () {
        // if api call has not received a response yet
        if(!this.state.responseReceived) {
            return (
                <TopBar/>
            )
        }

        // if no data was returned (bad request or some other error)
        if(this.state.userData == null) {
            return (
                <Text>
                    User not found.
                </Text>
            )
        }

        let user = this.state.userData.user
        let persInfo = this.state.userData.personal_info
        let teamName = this.state.userData.team_name

        let name = persInfo.firstname + " " + persInfo.lastname
        let email = user.email
        let role = user.role
        let position = user.position
        let height = persInfo.height
        let weight = persInfo.weight

        role = this.capitalizeFirstLetter(role)
        position = this.capitalizeFirstLetter(position)

        let subtitle 

        if (teamName === "") {
            teamName = "Not currently on a team"
        }

        b64ProfilePhoto = 'data:image/png;base64,' + user.photo

        return (
            <>
                <TopBar/>
                
                <View style={ProfileStyle.headerSectionView}>
                    <Text style={ProfileStyle.headerText}> 
                        {user.username} 
                    </Text>
                </View>

                <View style={ProfileStyle.photoSectionView}>
                    <Image style={ProfileStyle.proilePicImage} source={{uri: b64ProfilePhoto}}/>
                    <Text style={ProfileStyle.likesText}> likes and stuff </Text>
                </View>

                <View style={ProfileStyle.titleSectionView}>
                    <View>
                        <Text style={ProfileStyle.titleText}>
                                {name} {"\n" + email}
                            </Text>
                        <Text style={ProfileStyle.subtitleText}>
                            {role + "\n"}
                            {position + "\n"}
                            {teamName + "\n"}
                        </Text>
                    </View>
                    <View>
                        <Text style={ProfileStyle.titleText}>
                            Stats
                        </Text>
                        <Text style={ProfileStyle.subtitleText}>
                            {height + " inches"} 
                            {"\n" + weight + " lbs"}
                        </Text>
                    </View>
                </View>

                <View style={ProfileStyle.interactionsSectionView}> 
                    <TouchableOpacity style={ProfileStyle.interactionArea}>
                        <Text style={ProfileStyle.interactionText}>
                            Contextual button
                        </Text>
                    </TouchableOpacity>
                </View>

                <View style={ProfileStyle.emptySpaceView}/>
            </>
        )
    }   
}