//Create Account Screen
import React from "react";
import { Button, Text, TextInput, TouchableOpacity, View } from "react-native";
import FormStyle from "../Form.style";
import TopBar from "../component/TopBar"
import handleCreateTeam from "../../event-handler/HandleCreateTeam"

export default function CreateTeam(){
    const [Name, setName] = React.useState("")
    const [TeamLocation, setTeamLocation] = React.useState("");

    return (<>
        <TopBar/>
            <View style={FormStyle.container}>
                <Text style={FormStyle.fllText}> FLL </Text>
                <Text style={FormStyle.logo}> Create Team </Text>
                <View style={FormStyle.inputView} >
                 <TextInput
                    style={FormStyle.inputText}
                    placeholder="Team Name..."
                    placeholderTextColor="white"
                    onChangeText={setName}
                    autoCapitalize={false}/>
                </View>
                <View style={FormStyle.inputView} >
                 <TextInput
                    style={FormStyle.inputText}
                    placeholder="Team Location..."
                    placeholderTextColor="white"
                    onChangeText={setTeamLocation}
                    autoCapitalize={false}/>
                </View>
                { /* change handleSubmit to handleChangeTeam */}
                <TouchableOpacity style={FormStyle.button}
                                 onPress={()=> handleCreateTeam(Name, TeamLocation)}>
                  <Text style={FormStyle.loginText}>REGISTER TEAM</Text>
                </TouchableOpacity>
            </View>
            </>
        );
    }
