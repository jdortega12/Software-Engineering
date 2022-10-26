//Login Screen
import React from "react";
import { Button, Text, TextInput, TouchableOpacity, View } from "react-native";
import FormStyle from "../Form.style";

export default function LoginScreen(){
    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");

    function handleSubmit(pEmail, pPwd){
        const user = {login: pEmail, password: pPwd};
        // I would call a method from a different Layer

        console.log(JSON.stringify(user));
        alert( JSON.stringify(user) );
    }

    return (<>
        <View style={FormStyle.container}>
            <Text style={FormStyle.fllText}> FLL </Text>
            <Text style={FormStyle.logo}> Login </Text>
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
            <TouchableOpacity style={FormStyle.button}
                             onPress={()=> handleSubmit(email,password)}>
              <Text style={FormStyle.loginText}>LOGIN</Text>
            </TouchableOpacity>
        </View>
        </>
    );
}
