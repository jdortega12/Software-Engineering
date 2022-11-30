import React, {useState} from "react";
import {Text, View, ScrollView} from "react-native";
import TopBar from "../component/TopBar";
import FormStyle from "../Form.style";

export default function ChangeRoster(){

    const [players, setPlayers] = useState([
        {ID: 1, Teamname: "Jaymins", Firstname: "Jaymin", Lastname: "Ortega"},
        {ID: 2, Teamname: "Evans", Firstname: "Evan", Lastname: "Sisitsky"},
        {ID: 3, Teamname: "Wills", Firstname: "Will", Lastname: "Farrell"},
        {ID: 4, Teamname: "Dougs", Firstname: "Doug", Lastname: "Ortega"},
        {ID: 5, Teamname: "Bobbys", Firstname: "Bobby", Lastname: "Shmurda"},
        {ID: 6, Teamname: "Jacks", Firstname: "Jack", Lastname: "McCormick"},
        {ID: 7, Teamname: "Jackies", Firstname: "Jackie", Lastname: "Pineda"},
    ]);

        return (
        <>
        <TopBar/>
        <Text style={FormStyle.rosterTitle}>Admin Change Roster</Text>
        <ScrollView>
        <View>
        {players.map((item) => {
            return(
                <View key={item.ID} style={FormStyle.player}>
                    <Text style={FormStyle.playerText}>Name: {item.Firstname} {item.Lastname}</Text>
                    <Text style={FormStyle.playerText}>Team: {item.Teamname}</Text>
                </View>   
                )
            })}
        </View>
        </ScrollView>
        </>
        );
        

}