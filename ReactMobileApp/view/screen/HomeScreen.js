import React from "react"
import {TouchableOpacity, Text, View, ScrollView} from "react-native"
import NavBarStyle from "../component/NavBar.style"
import HandleLogout from "../../event-handler/HandleLogout"
import FormStyle from "../Form.style";

//Imports for navigation
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';


export default function HomeScreen({navigation}) {
    return (
    <>
    <ScrollView>
        <View style={FormStyle.container}>
            <Text style={FormStyle.fllText}> FLL </Text>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('CreateAccount')}>
                <Text style={FormStyle.linkText}>Create Account</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('Login')}>
                <Text style={FormStyle.linkText}>Login</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('CreateTeam')}>
                <Text style={FormStyle.linkText}>Create Team</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('UpdateUserPersonalInfo')}>
                <Text style={FormStyle.linkText}>Update Personal Info</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('SeasonalLeaderboard')}>
                <Text style={FormStyle.linkText}>Seasonal Leaderboard</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('ChangeRoster')}>
                <Text style={FormStyle.linkText}>Change Roster</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('StartMatchForm')}>
                <Text style={FormStyle.linkText}>Start Match</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('ViewMatch')}>
                <Text style={FormStyle.linkText}>View Match</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('AskManagerRequestForm')}>
                <Text style={FormStyle.linkText}>Ask Manager</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('InvitePlayerRequestForm')}>
                <Text style={FormStyle.linkText}>Invite Player</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('TeamRequestForm')}>
                <Text style={FormStyle.linkText}>Team Request</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('AcceptOrDeny')}>
                <Text style={FormStyle.linkText}>Accept or Deny</Text>
            </TouchableOpacity>
            <TouchableOpacity style={FormStyle.link}
                onPress={()=> navigation.navigate('Logout')}>
                <Text style={FormStyle.linkText}>Logout</Text>
            </TouchableOpacity>  
        </View>
        </ScrollView>
        </>
    )
}
