import React from 'react'
import { FlatList, Text,View } from 'react-native'
import FormStyle from "../Form.style";
import TopBar from "../component/TopBar"

export default function SeasonalLeaderboard(){
    const data = [
        {id: 1, name: 'Cardinals', wins: 1, loses:0},
        {id: 2, name: 'Giants', wins: 0, loses: 1}
    ]

    const item = ({ item }) => (
        <View style={{ flexDirection: 'row' }}>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>{item.id}</Text>
            </View>
            <View style={{ width: 100, backgroundColor: '#e32636'}}>
                <Text style={FormStyle.leaderboardText}>{item.name}</Text>
            </View>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>{item.wins}</Text>
            </View>
            <View style={FormStyle.leaderboardBackgroud}>
                <Text style={FormStyle.leaderboardText}>{item.loses}</Text>
            </View>
        </View>
    )

    const FlatList_Header = () => {
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

    return(
        <>
        <TopBar/>
        <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', marginTop: '10%'}}>
            <FlatList ListHeaderComponent={FlatList_Header} data={data} renderItem={item} keyExtractor={item => item.id.toString()} />
        </View>
        </>
    )
}

