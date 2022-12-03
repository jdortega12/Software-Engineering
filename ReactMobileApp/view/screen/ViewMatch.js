import { ThemeProvider } from "@react-navigation/native";
import React from "react";
import {Text, TouchableNativeFeedback, View, SafeAreaView, TouchableOpacity, TextInput} from "react-native";
import TopBar from "../component/TopBar";
import MatchStyle from "./Match.style";
import FormStyle from "../Form.style";

export default class ViewMatch extends React.Component {
    state = {
        location: null, 
        quarter: null,
        time: null,
        hometeam: null,
        awayteam: null, 
        matchtype: null,
        comments: null,
        hometeamscore: null,
        awayteamscore: null,
        likes: 0,
        dislikes: 0, 
        userComment: null
    };

    getMatch = async() => {
        const match = await fetch("http://10.0.2.2:8080/api/v1/getMatch/" + this.props.id, {
            method: "GET", 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        });

        if(match.status != 202) {
            return; 
        }

        const response = await match.json(); 

        const comments = await this.getComments(this.props.id)
        console.log("likes:");
        console.log(this.state.likes);
        const comment_display = 4;
        const comment_limit = comments.length >= comment_display ? comment_display : comments.length;
        const hometeam = await this.getTeam(response.home_id);
        const awayteam = await this.getTeam(response.away_id);

        date = new Date(response.quarter_time);
        time_string = date.getMinutes() + ":" + date.getSeconds();
        
        this.setState({
            location: response.location,
            quarter: response.quarter,
            time: time_string,
            hometeam: hometeam,
            awayteam: awayteam,
            matchtype: response.match_type,
            comments: comments.slice(comments.length - comment_limit - 1, comments.length),
            comments_cache: comments.slice(0, comments.length - comments.limit - 1),
            hometeamscore: response.home_team_score,
            awayteamscore: response.away_team_score, 
            likes: response.likes, 
            dislikes: response.dislikes
        });
    }

    postComment = async(message) => {
        comment = await fetch("http://10.0.2.2:8080/api/v1/createComment/", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                'message': message,
                'match_id': this.props.id,
            })
        });
    };

    postLikes = async(id) => {
        liked = await fetch("http://10.0.2.2:8080/api/v1/postLikes/" + id, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                'id': id,
            })
        });
    };

    postDislikes = async(id) => {
        liked = await fetch("http://10.0.2.2:8080/api/v1/postDislikes/" + id, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                'id': id,
            })
        });
    };

    getComments = async(id) => {
        comments = await fetch("http://10.0.2.2:8080/api/v1/getComments/" + id, {
            method: "GET", 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        });
        if(comments.status != 202) {
            return; 
        }

        response = await comments.json()

        return response; 
    }

    getTeam = async(id) => {
        team = await fetch("http://10.0.2.2:8080/api/v1/getTeam/" + id, {
            method: "GET", 
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        });

        if(team.status != 202) {
            return; 
        }

        response = await team.json();
        return response.name; 
    }

    render() {
        if(this.state.hometeamscore == null) {
            this.getMatch(); 
            return (
                <TopBar />
            )
        }

        return (
            <>
                <TopBar />
                <View style={MatchStyle.col}>
                    <View style={MatchStyle.row}>
                        <View style={MatchStyle.teamcol}>
                            <View style={MatchStyle.row2}>
                                <Text style={MatchStyle.teamtext}>{this.state.hometeam}</Text>
                            </View>
                            <View style={MatchStyle.row}>
                                <Text style={MatchStyle.teamtext}>{this.state.hometeamscore}</Text>
                            </View>
                        </View>
                        <View style={MatchStyle.teamcol}>
                            <View style={MatchStyle.row2}>
                                <Text style={MatchStyle.teamtext}>{this.state.awayteam}</Text>
                            </View>
                            <View style={MatchStyle.row}>
                                <Text style={MatchStyle.teamtext}>{this.state.awayteamscore}</Text>
                            </View>
                        </View>
                    </View>
                    <View style={MatchStyle.row3}>
                        <View style={MatchStyle.col2}>
                            <Text style={MatchStyle.quartertext}>Quarter: {this.state.quarter}</Text>
                            <Text style={MatchStyle.timetext}>{this.state.time}</Text>
                        </View>
                    </View>
                </View>
                <View style={MatchStyle.row3}>
                        <View style={MatchStyle.col2}>
                        <Text style={MatchStyle.quartertext}>Likes: {this.state.likes}</Text>
                            <TouchableOpacity style={MatchStyle.buttonLike}
                                onPress={()=> this.postLikes(this.props.id)}>
                                <Text style={MatchStyle.quartertext}>    Like    </Text>
                            </TouchableOpacity>
                        </View>
                        <View style={MatchStyle.col2}>
                        <Text style={MatchStyle.quartertext}>Dislikes: {this.state.dislikes}</Text>
                            <TouchableOpacity style={MatchStyle.buttonLike}
                                onPress={()=> this.postDislikes(this.props.id)}>
                                <Text style={MatchStyle.quartertext}>Dislike</Text>
                            </TouchableOpacity>
                        </View>
                    </View>
                <View style={MatchStyle.textcol}>
                        {this.state.comments.map((item) => (
                            <View style={MatchStyle.commentrow}>
                                <Text>{item.username}: {item.message}</Text>
                            </View>
                        ))}
                </View>
                <View style={MatchStyle.textcol}>
            <Text style={MatchStyle.timetext}> Comment </Text>
            <View style={MatchStyle.inputView} >
                <View style={MatchStyle.col3}>
                    <TextInput
                        style={MatchStyle.inputText}
                        placeholder="Add a comment..."
                        placeholderTextColor="#D6CFC7"
                        onChangeText={(userComment) => this.state.userComment = userComment}
                        autoCapitalize={false}/>
                </View>
                <View style={MatchStyle.col4}>
                    <TouchableOpacity style={MatchStyle.button}
                             onPress={()=> this.postComment(this.state.userComment)}>
                    <Text style={MatchStyle.quartertext}>Post</Text>
                </TouchableOpacity>
                </View>
            </View>
        </View>
            </>
        )
}
}