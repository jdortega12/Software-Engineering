import React from "react"
import {Button, Text, View} from "react-native"
import NavBarStyle from "../component/NavBar.style"
import NavBar from "../component/NavBar"
import HandleLogout from "../../event-handler/HandleLogout"

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

export default function HomeScreen({navigation}) {
    return (<>
        <View style={NavBarStyle.titleBarView}>
            <Button
            color="e32636"
            title="Create Account"
            onPress={() => navigation.navigate('CreateAccount')}
            />
            <Button color="e32636" title="Login" onPress={() => navigation.navigate('Login')}/>
            <Button color="e32636" title="Logout (temp)" onPress={HandleLogout}/>
        </View>
        <View style={NavBarStyle.centeredText}>
            <Text style ={NavBarStyle.homeText}> Home Screen </Text>
        </View>

        </>
    )
}
