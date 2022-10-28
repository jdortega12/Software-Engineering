import React from 'react'
import { Text, TextInput, TouchableNativeFeedback, TouchableOpacity, View } from 'react-native'
import FormStyle from "../Form.style";


export default class TeamRequestForm extends React.Component {
    state = {
        ReceiverID: null,
        Message: null
    }

    async handleSubmit() {
        console.log(this.state);
        try {
            const response = await fetch('http://localhost:8080/api/v1/createTeamRequest', {
                method: 'POST',
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                },
                body: {
                    'ReceiverID': this.state.ReceiverID,
                    'Message': this.state.Message
                }
            })

            console.log(response);
        } catch (e) {
            console.error(e);
        }        
    }

    render() {
        console.log(this.props);
        this.state.ReceiverID = this.props.ReceiverID;
        const titleString = this.props.type == '0' ? 'Request to join ' + this.state.ReceiverID + '\'s team' : 'Invite ' + this.state.ReceiverID + ' to Your Team';
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