//Create Account Screen
import React from "react";
import { Button, Text, TextInput, TouchableOpacity, View } from "react-native";
import FormStyle from "../Form.style";

export default function CreateAccount(){
    const [username, setUsername] = React.useState("")
    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");


    function handleSubmit(pUsername, pEmail, pPwd){
        const user = {username:pUsername, login: pEmail, password: pPwd};
        event-handler.handleCreateAccount(user)

        console.log(JSON.stringify(user));
        alert( JSON.stringify(user) );
    }


    return (<>
        <View style={FormStyle.container}>
            <Text style={FormStyle.fllText}> FLL </Text>
            <Text style={FormStyle.logo}> Create Account </Text>
            <View style={FormStyle.inputView} >
             <TextInput
                style={FormStyle.inputText}
                placeholder="Username..."
                placeholderTextColor="white"
                onChangeText={setUsername}
                autoCapitalize={false}/>
            </View>
            <View style={FormStyle.inputView} >
             <TextInput
                style={FormStyle.inputText}
                placeholder="Email..."
                placeholderTextColor="white"
                onChangeText={setEmail}
                autoCapitalize={false}/>
            </View>
            <View style={FormStyle.inputView} >
            <TextInput
                secureTextEntry={true}
                style={FormStyle.inputText}
                placeholder="Password..."
                placeholderTextColor="white"
                onChangeText={setPassword}/>
            </View>
            {/*change handleSubmit to handleChangeAccount */}
            <TouchableOpacity style={FormStyle.button}
                             onPress={()=> handleSubmit(username,email,password)}>
              <Text style={FormStyle.loginText}>REGISTER</Text>
            </TouchableOpacity>
        </View>
        </>
    );
}