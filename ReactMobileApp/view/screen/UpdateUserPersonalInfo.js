import React from 'react-native'
import { Button, Text, TextInput, TouchableOpacity, View } from "react-native";
import TopBar from "../component/TopBar"
import FormStyle from "../Form.style"


export default function UpdateUserPersonalInfoScreen() {
    return (
        <>
            <TopBar/>
            <View style={FormStyle.container}>
            <Text style={FormStyle.logo}> Update Info </Text>
            <View style={FormStyle.inputView} >
             <TextInput
                style={FormStyle.inputText}
                placeholder="Firstname"
                placeholderTextColor="white"
                //onChangeText={setEmail}
                autoCapitalize={false}
                />
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                style={FormStyle.inputText}
                placeholder="Lastname"
                placeholderTextColor="white"
                //onChangeText={setPassword}
                autoCapitalize={false}
                />
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                style={FormStyle.inputText}
                placeholder="Height (inches)"
                placeholderTextColor="white"
                //onChangeText={setPassword}
                autoCapitalize={false}
                />
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                style={FormStyle.inputText}
                placeholder="Weight (lbs)"
                placeholderTextColor="white"
                //onChangeText={setPassword}
                autoCapitalize={false}
                />
            </View>
            <TouchableOpacity style={FormStyle.button}
                             //</View>onPress={()=> handleSubmit(email,password)}
            >
              <Text style={FormStyle.loginText}> UPDATE </Text>
            </TouchableOpacity>
        </View>
        </>
    )
}