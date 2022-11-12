import React from 'react'
import { Text, TextInput, TouchableNativeFeedback, TouchableOpacity, View } from 'react-native'
import FormStyle from "../Form.style";
import TopBar from "../component/TopBar"
import handleRemovePlayer from "../../event-handler/HandleRemovePlayer"

export default function RemovePlayer(){
    const [managerUsername, setManagerUsername] = React.useState("");
    const [playerUsername, setPlayerUsername] = React.useState("");
    const [accepted, setAccepted] = React.useState("");
    
    return (
        <> 
        <TopBar/>
        <View style={FormStyle.container}>
            <View style={FormStyle.inputView} >
             <TextInput
                style={FormStyle.inputText}
                placeholder="Manager Username..."
                placeholderTextColor="white"
                onChangeText={setManagerUsername}
                autoCapitalize={false}/>
            </View>
            <View style={FormStyle.inputView} >
             <TextInput
                style={FormStyle.inputText}
                placeholder="Player Username..."
                placeholderTextColor="white"
                onChangeText={setPlayerUsername}
                autoCapitalize={false}/>
            </View>
            <View style={FormStyle.inputView}>
            <TextInput
                style={FormStyle.inputText}
                placeholder="Remove Player? (Y/N)"
                placeholderTextColor="white"
                onChangeText={setAccepted}
                autoCapitalize={false}/>
            </View>
            <View>
                <TouchableOpacity style={FormStyle.button} onPress={() => handleRemovePlayer(managerUsername, playerUsername, accepted)}>
                    <Text style={FormStyle.loginText}>Submit</Text>
                </TouchableOpacity>
            </View> 
            </View>
        </>
        )
}