import React from 'react'
import { Text, TextInput, TouchableNativeFeedback, TouchableOpacity, View } from 'react-native'
import FormStyle from "../Form.style";


export default class TeamRequestForm extends React.Component {
    state = {
        ReceiverUsername: null,
        Message: null
    }

    async handleSubmit() {
        console.log(this.state);
        try {
            const response = await fetch('http://10.0.2.2:8080/api/v1/createTeamRequest', {
                method: 'POST',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                },
                body: {
                    'receiver_username': this.state.ReceiverUsername,
                    'message': this.state.Message
                }
            })

            console.log(response);
        } catch (e) {
            console.error(e);
        }        
    }

    render() {
        console.log(this.props);
        this.state.ReceiverUsername = this.props.ReceiverUsername;
        const titleString = this.props.type == '0' ? 'Request to join ' + this.state.ReceiverUsername + '\'s team' : 'Invite ' + this.state.ReceiverUsername + ' to Your Team';
        return (
            <View style={FormStyle.container}>
                <Text style={FormStyle.teamRequest}>{titleString}</Text>
                <Text>Message:</Text>
                <View style={FormStyle.inputView}>
                    <TextInput style={FormStyle.inputText} onChangeText={(message) => this.state.Message = message}/>
            </View>
            <View>
                <TouchableOpacity style={FormStyle.button} onPress={() => this.handleSubmit()}>
                    <Text style={FormStyle.loginText}>Submit</Text>
                </TouchableOpacity>
                </View>
            </View>
        )
    }   
}