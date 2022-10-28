import React from 'react'
import { Text, TextInput, TouchableNativeFeedback, TouchableOpacity, View } from 'react-native'
import TeamRequestForm from './TeamRequestForm'

export default class AskManagerRequestForm extends React.Component {
    render() {
        return (
            <TeamRequestForm ReceiverID={this.props.ReceiverID} type="0"/>
        )
    }
}