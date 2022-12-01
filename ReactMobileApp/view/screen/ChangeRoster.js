import React, {useState} from "react";
import {Text, View, ScrollView, Button, TextInput} from "react-native";
import TopBar from "../component/TopBar";
import FormStyle from "../Form.style";
import handleChangeRoster from "../../event-handler/HandleChangeRoster"
import handleGetUserTeamData from "../../event-handler/HandleGetUserTeamData"

export default function ChangeRoster(){
    /*
    const [players, setPlayers] = useState([
        {ID: 1, Teamname: "Jaymins", Firstname: "Jaymin", Lastname: "Ortega"},
        {ID: 2, Teamname: "Evans", Firstname: "Evan", Lastname: "Sisitsky"},
        {ID: 3, Teamname: "Wills", Firstname: "Will", Lastname: "Farrell"},
        {ID: 4, Teamname: "Dougs", Firstname: "Doug", Lastname: "Ortega"},
        {ID: 5, Teamname: "Bobbys", Firstname: "Bobby", Lastname: "Shmurda"},
        {ID: 6, Teamname: "Jacks", Firstname: "Jack", Lastname: "McCormick"},
        {ID: 7, Teamname: "Jackies", Firstname: "Jackie", Lastname: "Pineda"},
    ]);/*/

    const [playersStruct, setPlayersStruct] = useState({
        responseReceived: false,
        players: [],
    });

    const [newTeam, setNewTeam] = React.useState("");
    
    function getData(){
        handleGetUserTeamData().then((data) => {
            console.log("DATA", data)
            setPlayersStruct({
                responseReceived: true,
                players: data,
            })
        })
    };

    if(!playersStruct.responseReceived) {  
        getData()
        console.log("PLAYERS", playersStruct.players)
    }
        return (
        <>
        <Text style={FormStyle.rosterTitle}>Admin Roster Change Page</Text>
        <ScrollView>
        <View>
        {playersStruct.players.map((item) => {
            return(
                <View key={item.ID} style={FormStyle.player}>
                    <Text style={FormStyle.playerText}>Name: {item.firstname} {item.lastname}</Text>
                    <Text style={FormStyle.playerText}>Team: {item.teamname}</Text>
                    <View style={{padding: 10}}>
                    <TextInput
                        style={FormStyle.changeRosterInput}
                        placeholder="New Team Name..."
                        placeholderTextColor="black"
                        onChangeText={setNewTeam}
                        autoCapitalize={false}/>          
                    </View>
                    <View style={{padding: 10}}>
                    <Button title="Change Roster" onPress={()=> handleChangeRoster(item.ID, newTeam)}/>
                    </View>
                </View>   
                )
            })}
        </View>
        </ScrollView>
        </>
        );
        

}