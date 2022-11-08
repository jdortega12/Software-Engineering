import React, {useState} from "react"; 
import { Text, TouchableHighlightBase, View } from "react-native";
import TopBar from "../component/TopBar"

export default class TeamProfile extends React.Component {

    state = {
        name: null, 
        location: null,
        loaded: false,
    };
  
    getTeam = async () => {
        const team = await fetch("http://10.0.2.2:8080/api/v1/getTeam/" + this.props.id, {
            method: "GET",
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        }); 
        
        if(team.status != 302) {
            return 
        }

        const response = await team.json();

        this.setState({
            name: response.name,
            location: response.location,
            loaded: true,
        });
    }

    render() {
        if(!this.state.loaded) {
            this.getTeam();
            return (
                <>
                    <TopBar />
                </>
            )
        } 
        return (
            <>
                <TopBar />
                <View>
                    <Text>{this.state.name}</Text>
                    <Text>{this.state.location}</Text>
                </View>
            </>   
        )
    }
}