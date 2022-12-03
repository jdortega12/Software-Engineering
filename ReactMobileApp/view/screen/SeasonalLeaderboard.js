import React, {useState} from 'react';
import { FlatList, Text,View } from 'react-native';
import FormStyle from "../Form.style";
import TopBar from "../component/TopBar";
import handleGetTeams from "../../event-handler/HandleGetTeams";

export default function SeasonalLeaderboard(){
    /*
    const [teams, setTeams] = useState([
        {id: 1, name: 'Cardinals', wins: 1, loses:0},
        {id: 2, name: 'Giants', wins: 0, loses: 1}
    ]);
    */

    const [teamsStruct, setTeamsStruct] = useState({
        responseReceived: false,
        teams: [],
    });

    function getTeams(){
        handleGetTeams().then((data) => {
            console.log("DATA", data)
            setTeamsStruct({
                responseReceived: true,
                teams: data,})
        })
    };

    function FlatList_Header(){
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
    
    
      if(!teamsStruct.responseReceived) {  
        getTeams();  
        console.log("TEAMS",teamsStruct.teams)
    }

    return(
        <>
        <TopBar/>
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', marginTop: '10%'}}>
            <FlatList
                ListHeaderComponent={FlatList_Header}
                data={teamsStruct.teams}
                renderItem={({item}) => (
                    <View style={{ flexDirection: 'row' }}>
                        <View style={FormStyle.leaderboardBackgroud}>
                            <Text style={FormStyle.leaderboardText}>{item.id}</Text>
                        </View>
                        <View style={{ width: 100, backgroundColor: '#e32636'}}>
                            <Text style={FormStyle.leaderboardText}>{item.team_name}</Text>
                        </View>
                        <View style={FormStyle.leaderboardBackgroud}>
                            <Text style={FormStyle.leaderboardText}>{item.wins}</Text>
                        </View>
                        <View style={FormStyle.leaderboardBackgroud}>
                            <Text style={FormStyle.leaderboardText}>{item.loses}</Text>
                        </View>
                    </View>
                )}/>
        </View>
        </>
    )


}
