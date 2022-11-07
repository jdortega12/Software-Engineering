import React from "react"
import {Button, Text, View} from "react-native"
import HandleLogout from "../../event-handler/HandleLogout"
import TopBar from "../component/TopBar"
import TopBarStyle from "../component/TopBar.style"

//Import navigation files
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

export default function HomeScreen({ navigation }) {
    return (
        <View style={TopBarStyle.titleBarView}>
            <Button
            color="e32636"
            title="Create Account"
            onPress={() => navigation.navigate('CreateAccount')}
            />
            <Button color="e32636" title="Login" onPress={() => navigation.navigate('Login')}/>
            <Button color="e32636" title="Logout (temp)" onPress={HandleLogout}/>
        </View>
    )
}
