import React from "react";
import { Text, View, Button } from "react-native";
import TopBar from "./view/component/TopBar"

//Screens
import HomeScreen from "./view/screen/HomeScreen";
import UpdateUserPersonalInfoScreen from "./view/screen/UpdateUserPersonalInfo";
import CreateAccount from "./view/screen/CreateAccountScreen"
import Login from "./view/screen/LoginScreen"
import CreateTeam from "./view/screen/CreateTeam"
import AcceptOrDeny from "./view/screen/AcceptOrDeny"
import SeasonalLeaderboard from "./view/screen/SeasonalLeaderboard"
import ChangeRoster from "./view/screen/ChangeRoster"
import ProfileScreen from "./view/screen/user_profile/UserProfileScreen"
import ProfileScreenPersonal from "./view/screen/user_profile/UserProfileScreenPersonal"
import ProfileScreenNotPersonal from "./view/screen/user_profile/UserProfileScreenNotPersonal"
import ViewMatch from "./view/screen/ViewMatch"
import PlayoffPicture from "./view/screen/PlayoffPicture"

//Forms/Notifications
import AskManagerRequestForm from "./view/component/AskManagerRequestForm"
import InvitePlayerRequestForm from "./view/component/InvitePlayerRequestForm"
import TeamRequestForm from "./view/component/TeamRequestForm"
import AdminNotificationScreen from "./view/screen/admin_notifications/AdminNotificationsScreen"
import StartMatchForm from "./view/component/StartMatchForm"


//Import navigation files
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

//Create stack
const Stack = createNativeStackNavigator()
export default function App(){
    return (
        <NavigationContainer>{
            <Stack.Navigator>
                <Stack.Screen name="Home" component={HomeScreen} />
                <Stack.Screen name="CreateAccount" component={CreateAccount} />
                <Stack.Screen name="Login" component={Login} />
                <Stack.Screen name="CreateTeam" component={CreateTeam} />
                <Stack.Screen name="UpdateUserPersonalInfo" component={UpdateUserPersonalInfoScreen} />
                <Stack.Screen name="AskManagerRequestForm" component={AskManagerRequestForm} />
                <Stack.Screen name="InvitePlayerRequestForm" component={InvitePlayerRequestForm} />
                <Stack.Screen name="TeamRequestForm" component={TeamRequestForm} />
                <Stack.Screen name="AcceptOrDeny" component={AcceptOrDeny} />
                <Stack.Screen name="SeasonalLeaderboard" component={SeasonalLeaderboard} />
                <Stack.Screen name="ChangeRoster" component={ChangeRoster} />
                <Stack.Screen name="ViewMatch" component={ViewMatch} />
                <Stack.Screen name="StartMatchForm" component={StartMatchForm} />
                <Stack.Screen name="PlayoffPicture" component={PlayoffPicture} />
                <Stack.Screen name="ProfileScreenPersonal" component={ProfileScreenPersonal} />
            </Stack.Navigator>
        }</NavigationContainer>     
    );
}
