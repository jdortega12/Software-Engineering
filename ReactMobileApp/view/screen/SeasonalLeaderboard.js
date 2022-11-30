import React, {useState} from 'react';
import { FlatList, Text,View } from 'react-native';
import FormStyle from "../Form.style";
import TopBar from "../component/TopBar";
import handleGetTeams from "../../event-handler/HandleGetTeams";

//export default function SeasonalLeaderboard(){
export default class SeasonalLeaderboard extends React.Component {  
    /*const data = [
        {id: 1, name: 'Cardinals', wins: 1, loses:0},
        {id: 2, name: 'Giants', wins: 0, loses: 1}
    ]*/
    
    state = {
        responseReceived: false,
        teams: [],
    };
    
    
    getTeams(){
        handleGetTeams().then((data) => {
            this.setState({
                responseReceived: true,
                teams: data,
            })
            console.log("GET TEAMS", this.teams.data)
        })
    };
    
    /*
    getTeams = async() => {
        const receivedTeams = await fetch("http://10.0.2.2:8080/api/v1/getTeams/", {
            method: "GET", 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        });

        if(receivedTeams.status != 202) {
            return;
        }

        const response = await receivedTeams.json();
        /*
        this.state = {
            responseReceived: true,
            teams: response,
        };
        this.setState({
            responseReceived: true,
            teams: response,
        })
        console.log("TEAMS FROM GET TEAMS", this.state.teams)
    }*/
    
    team = ({ team }) => (
        <View style={{ flexDirection: 'row' }}>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>{team.id}</Text>
            </View>
            <View style={{ width: 100, backgroundColor: '#e32636'}}>
                <Text style={FormStyle.leaderboardText}>{team.name}</Text>
            </View>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>{team.wins}</Text>
            </View>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>{team.loses}</Text>
            </View>
        </View>
    )

    FlatList_Header = () => {
        return (
        <>
        <View style={{justifyContent: 'center', alignItems: 'center', marginTop: '10%', backgroundColor: '#e32636'}}>
            <Text style={{fontSize: 24, textAlign: 'center', color: 'white'}}>Seasonal Leaderboard</Text>
        </View>
          <View style={{flexDirection: 'row'}}>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>Rank</Text>
            </View>
            <View style={{ width: 100, backgroundColor: '#e32636'}}>
                <Text style={FormStyle.leaderboardText}>Name</Text>
            </View>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>Wins</Text>
            </View>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>Loses</Text>
            </View>
          </View>
        </>
        );
      }

      render () {
        // if api call has not received a response yet
        if(!this.state.responseReceived) {  
            this.getTeams();  
        }

        console.log("TEAMS FROM RENDER", this.state.teams)

        return(
            <>
            <TopBar/>
            <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', marginTop: '10%'}}>
                <FlatList ListHeaderComponent={FlatList_Header} data={this.state.teams} renderItem={team} keyExtractor={item => team.id.toString()} />
            </View>
            </>
        )
    }
}
