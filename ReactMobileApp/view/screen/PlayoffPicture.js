import React from "react";
import {Text, TouchableNativeFeedback, View, SafeAreaView} from "react-native";
import TopBar from "../component/TopBar";
import FormStyle from "../Form.style";
import PlayoffStyle from "./Playoff.style";

export default class PlayoffPicture extends React.Component {
    state = {
        teams: [], 
        loaded: false
    }

    getTeams = async() => {
        const teams = await fetch("http://10.0.2.2:8080/api/v1/getPlayoffs", {
            method: "GET", 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        });

        if(teams.status != 202) {
            return; 
        }

        const response = await teams.json()

        matchups = []

        for(let i = 0; i < response.length / 2; i++) {
            matchups.push({
                name1: response[i],
                name2: response[response.length - i - 1],
                seed1: i+1, 
                seed2: response.length - i
            });
        }

        this.setState({
            teams: matchups,
            loaded: true
        });
    }

    boxWithTeams(name1, name2, seed1, seed2) {
        return (
            <View style={PlayoffStyle.bracketBox} key={seed1}>
                <Text>{seed1}. {name1}</Text>
                <Text>{seed2}. {name2}</Text>
            </View>
        )
    }

    render() {
        if(!this.state.loaded) {
            this.getTeams()
            return (
                <TopBar/>
            )
        }
        return (
            <SafeAreaView>
                <TopBar/>
                <View style={PlayoffStyle.container}>
                    <Text style={FormStyle.fllText}>Playoff Picture</Text>
                </View>
                <View style={PlayoffStyle.playoffContainer}>
                        {this.state.teams.map((item) => (
                            <View style={PlayoffStyle.row} key={item.name1}>
                                {this.boxWithTeams(item.name1, item.name2, item.seed1, item.seed2)}
                            </View>
                        ))}
                </View>
            </SafeAreaView> 
        )
    }
}