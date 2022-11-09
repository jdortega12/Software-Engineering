import React from "react"
import {Button, Text, View} from "react-native"
import NavBarStyle from "../component/NavBar.style"
import NavBar from "../component/NavBar"
import HandleLogout from "../../event-handler/HandleLogout"
import TopBar from "../component/TopBar"
import Icon from 'react-native-vector-icons/FontAwesome';

//Imports for navigation
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';


export default function HomeScreen({navigation}) {
    return (<>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color="white" name="user-plus" size={30} onPress={() => navigation.navigate('CreateAccount')}/>
            <Text style={NavBarStyle.iconText}> Create Account </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color="white" name="sign-in" size={30} onPress={() => navigation.navigate('Login')}/>
            <Text style={NavBarStyle.iconText}> Login </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color="white" name="male" size={30} onPress={() => navigation.navigate('CreateTeam')}/>
            <Text style={NavBarStyle.iconText}> Create Team </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color="white" name="user" size={30} onPress={() => navigation.navigate('UpdateUserPersonalInfo')}/>
            <Text style={NavBarStyle.iconText}> Update User Info </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color ="white" name="envelope" size={30} onPress={() => navigation.navigate('AskManagerRequestForm')}/>
            <Text style={NavBarStyle.iconText}> Ask Manager </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>    
            <Icon  color ="white" name="send" size={30} onPress={() => navigation.navigate('InvitePlayerRequestForm')}/>
            <Text style={NavBarStyle.iconText}> Invite Player </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color ="white" name="file" size={30} onPress={() => navigation.navigate('TeamRequestForm')}/>
            <Text style={NavBarStyle.iconText}> Team Request </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color ="white" name="file" size={30} onPress={() => navigation.navigate('AcceptOrDeny')}/>
            <Text style={NavBarStyle.iconText}> Accept Or Deny </Text>
        </View>
        <View style={NavBarStyle.titleBarView}>
            <Icon  color ="white" name="power-off" size={30} onPress={HandleLogout}/>
            <Text style={NavBarStyle.iconText}> Logout </Text>
        </View>
        
        <View style={NavBarStyle.centeredText}>
            <Text style ={NavBarStyle.homeText}> Fake Football League </Text>
        </View>
        </>
    )
}
