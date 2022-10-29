//Login Screen
import React from "react";
import { Button, Text, TextInput, TouchableOpacity, View } from "react-native";
import FormStyle from "../Form.style";
import TopBar from "../component/TopBar"
import handleLogin from "../../event-handler/HandleLogin"

export default function LoginScreen(){
    const [username, setUsername] = React.useState("");
    const [password, setPassword] = React.useState("");

    /*
    function handleSubmit(pEmail, pPwd){
        const userInfo = {login: pEmail, password: pPwd};
        handleLogin(userInfo)
        //console.log(JSON.stringify(user));
        //alert( JSON.stringify(user) );
    }
    */
    return (<>
    <TopBar/>
        <View style={FormStyle.container}>
            <Text style={FormStyle.fllText}> FLL </Text>
            <Text style={FormStyle.logo}> Login </Text>
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
                secureTextEntry={true}
                style={FormStyle.inputText}
                placeholder="Password..."
                placeholderTextColor="white"
                onChangeText={setPassword}/>
            </View>
            <TouchableOpacity style={FormStyle.button}
                             onPress={()=> handleLogin(username,password)}>
              <Text style={FormStyle.loginText}>LOGIN</Text>
            </TouchableOpacity>
        </View>
        </>
    );
}
