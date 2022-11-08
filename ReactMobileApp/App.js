import React from "react";
import { Text, View, Button } from "react-native";
import TopBar from "./view/component/TopBar"
import HomeScreen from "./view/screen/HomeScreen";
import UpdateUserPersonalInfoScreen from "./view/screen/UpdateUserPersonalInfo";
import CreateAccount from "./view/screen/CreateAccountScreen"
import Login from "./view/screen/LoginScreen"
import CreateTeam from "./view/screen/CreateTeam"

//import UploadImage from "./view/screen/UploadImage"
import AskManagerRequestForm from "./view/component/AskManagerRequestForm"
import InvitePlayerRequestForm from "./view/component/InvitePlayerRequestForm"
import TeamRequestForm from "./view/component/TeamRequestForm"

import ProfileScreen from "./view/screen/user_profile/UserProfileScreen"
import ProfileScreenPersonal from "./view/screen/user_profile/UserProfileScreenPersonal"
import ProfileScreenNotPersonal from "./view/screen/user_profile/UserProfileScreenNotPersonal"

import AdminNotificationScreen from "./view/screen/admin_notifications/AdminNotificationsScreen"

//Import navigation files
import { NavigationContainer } from '@react-navigation/native'
import { createNativeStackNavigator } from '@react-navigation/native-stack'

//Create stack
const Stack = createNativeStackNavigator()

// The plan for the acceptance test presentation is to 
// manually switch screens becuase we do not yet have the 
// navigation bar working.
export default function App(){
    return (
        
        <NavigationContainer>{
            <Stack.Navigator>
                <Stack.Screen name="Home" component={HomeScreen} options={{ title: 'Fake Football League'}}/>
                <Stack.Screen name="CreateAccount" component={CreateAccount} />
                <Stack.Screen name="Login" component={Login} />
                <Stack.Screen name="CreateTeam" component={CreateTeam} />
            </Stack.Navigator>
        }</NavigationContainer>

        //<HomeScreen />
        //<CreateAccount />
        //<Login />
        //<UpdateUserPersonalInfoScreen />
        //<CreateTeam />
        //<AskManagerRequestForm />
        //<InvitePlayerRequestForm />
        //<TeamRequestForm />
        //<UploadImage />
        
    );
}