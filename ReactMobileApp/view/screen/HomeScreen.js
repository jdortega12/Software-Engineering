import React from "react"
import {Button, Text, View} from "react-native"
import NavBarStyle from "../component/NavBar.style"
import NavBar from "../component/NavBar"
import HandleLogout from "../../event-handler/HandleLogout"
import TopBar from "../component/TopBar"

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

export default function HomeScreen({navigation}) {
    return (<>
        <View backgroundColor="#e32636">
            <Button color="e32636" title="Create Account" onPress={() => navigation.navigate('CreateAccount')}/>
            <Button color="e32636" title="Login" onPress={() => navigation.navigate('Login')}/>
            <Button color="e32636" title="Create Team" onPress={() => navigation.navigate('CreateTeam')}/>
            <Button color="e32636" title="Update User Personal Info" onPress={() => navigation.navigate('UpdateUserPersonalInfo')}/>
            <Button color="e32636" title="Ask Manager Request Form" onPress={() => navigation.navigate('AskManagerRequestForm')}/>
            <Button color="e32636" title="Invite Player Request Form" onPress={() => navigation.navigate('InvitePlayerRequestForm')}/>
            <Button color="e32636" title="Team Request Form" onPress={() => navigation.navigate('TeamRequestForm')}/>
            <Button color="e32636" title="Logout (temp)" onPress={HandleLogout}/>
        </View>
        
        <View style={NavBarStyle.centeredText}>
            <Text style ={NavBarStyle.homeText}> Fake Football League Home </Text>
        </View>
        </>
    )
}
