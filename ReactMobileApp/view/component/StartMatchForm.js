import React from "react"
import { Text, View, TextInput, TouchableOpacity } from "react-native"
import TopBar from "./TopBar"
import FormStyle from "../Form.style"
import handleStartMatch from "../../event-handler/HandleStartMatch"

export default function StartMatchForm() {
    const [homeTeam, setHomeTeam] = React.useState("")
    const [awayTeam, setAwayTeam] = React.useState("")
    const [location, setLocation] = React.useState("")

    return (
    <>
        <TopBar/>

        <View style={FormStyle.container}>
            <Text style={FormStyle.logo}> Start Match </Text>

            <View style={FormStyle.inputView} >
                <TextInput
                    style={FormStyle.inputText}
                    placeholder="Home Team"
                    placeholderTextColor="white"
                    autoCapitalize={false}
                    onChangeText={setHomeTeam}
                />
            </View>

            <View style={FormStyle.inputView} >
                <TextInput
                    style={FormStyle.inputText}
                    placeholder="Away Team"
                    placeholderTextColor="white"
                    autoCapitalize={false}
                    onChangeText={setAwayTeam}
                />
            </View>
    
            <View style={FormStyle.inputView} >
                <TextInput
                    style={FormStyle.inputText}
                    placeholder="Location"
                    placeholderTextColor="white"
                    autoCapitalize={false}
                    onChangeText={setLocation}
                />
            </View>

            <TouchableOpacity style={FormStyle.button} onPress={()=> handleStartMatch(homeTeam, awayTeam, location)}>
                <Text style={FormStyle.loginText}> Start Match </Text>
            </TouchableOpacity>
        </View>
    </>
    )
}