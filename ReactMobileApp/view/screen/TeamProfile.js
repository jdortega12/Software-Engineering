import React, {useState} from "react"; 
import { Text, TouchableHighlightBase, View, RefreshControl, Modal, TouchableOpacity } from "react-native";
import TopBar from "../component/TopBar"
import TeamStyle from "../Team.style";
import FormStyle from "../Form.style";
import AskManagerRequestForm from "../component/AskManagerRequestForm";

export default class TeamProfile extends React.Component {

    state = {
        name: null, 
        location: null,
        manager_id: null,
        manager_name: null,
        loaded: false,
        modal_visible: false,
        players: []
    };

    setModalVisible = (visible) => {
        this.setState({
            name: this.state.name,
            location: this.state.location,
            manager_id: this.state.manager_id,
            manager_name: this.state.manager_name,
            loaded: this.state.loaded,
            players: this.state.players,
            modal_visible: visible
        });
    }
  
    getTeam = async () => {
        const team = await fetch("http://10.0.2.2:8080/api/v1/getTeam/" + this.props.id, {
            method: "GET",
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        }); 
        
        if(team.status != 202) {
            return;
        }

        const response = await team.json();

        await this.getPlayers()

        this.setState({
            name: response.name,
            location: response.location,
            manager_id: response.manager_id,
            manager_name: response.manager_name,
            modal_visible: false,
            loaded: true,
            players: this.state.players
        });
    }

    getPlayers = async() => {
        const players = await fetch("http://10.0.2.2:8080/api/v1/getTeamPlayers/" + this.props.id, {
            method: "GET", 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        });

        if(players.status != 202) {
            return;
        }

        const response = await players.json();
 
        this.state = {
            name: this.state.name,
            location: this.state.location,
            manager_id: this.state.manager_id,
            manager_name: this.state.manager_name,
            modal_visible: false, 
            loaded: false, 
            players: response
        };

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
                <View style={TeamStyle.container}>
                    <Modal 
                        visible={this.state.modal_visible}
                        animationType="slide"
                        onRequestClose={() => {
                            this.setModalVisible(!this.state.modal_visible)
                        }}
                    >
                        <AskManagerRequestForm ReceiverUsername={this.state.manager_name}/>
                    </Modal>
                    <Text style={TeamStyle.nameText}>{this.state.name}</Text>
                    <Text style={TeamStyle.locationText}>{this.state.location}</Text>
                </View>
                <View style={TeamStyle.row}>
                    <View style={TeamStyle.col}>
                        <Text>Manager: {this.state.manager_name}</Text>
                    </View>
                    <View style={TeamStyle.col}>
                        <Text>Record: </Text>
                    </View>
                </View>
                <View style={TeamStyle.container}>
                    <TouchableOpacity 
                            onPress={() => this.setModalVisible(true)}
                            style={FormStyle.button}
                    >
                        <Text style={FormStyle.loginText}>Request to Join</Text>
                    </TouchableOpacity>
                </View>
                <View>
                    {this.state.players.map((item, key) =>(
                        <View style={TeamStyle.playerRow}>
                            <Text style={TeamStyle.playerText}>{key + 1}. </Text><Text style={TeamStyle.playerText}>{item.username}</Text><Text style={TeamStyle.playerText}> {item.position}</Text>
                        </View>
                    ))}
                </View>
            </>   
        )
    }
}