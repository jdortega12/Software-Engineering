import React from 'react'
import { Button, Text, TextInput, TouchableOpacity, View } from "react-native";
import TopBar from "../component/TopBar"
import FormStyle from "../Form.style"
import handleUpdateUserPersonalInfo from '../../event-handler/HandleUpdateUserPersonalInfo'

export default function UpdateUserPersonalInfoScreen() {
    const [firstname, setFirstname] = React.useState("");
    const [lastname, setLastname] = React.useState("");
    const [height, setHeight] = React.useState("");
    const [weight, setWeight] = React.useState("");

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
                onChangeText={setFirstname}
                autoCapitalize={false}
                />
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                style={FormStyle.inputText}
                placeholder="Lastname"
                placeholderTextColor="white"
                onChangeText={setLastname}
                autoCapitalize={false}
                />
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                style={FormStyle.inputText}
                placeholder="Height (inches)"
                placeholderTextColor="white"
                onChangeText={setHeight}
                autoCapitalize={false}
                keyboardType="numeric"
                />
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                style={FormStyle.inputText}
                placeholder="Weight (lbs)"
                placeholderTextColor="white"
                onChangeText={setWeight}
                autoCapitalize={false}
                keyboardType="numeric"
                />
            </View>
            <TouchableOpacity style={FormStyle.button}
                onPress={()=> handleUpdateUserPersonalInfo(firstname,lastname,height,weight)}
            >
              <Text style={FormStyle.loginText}> UPDATE </Text>
            </TouchableOpacity>
        </View>
        </>
    )
}