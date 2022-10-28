import React from "react"
import {Button, Text, View, TouchableOpacity} from "react-native"
import {launchImageLibrary} from 'react-native-image-picker'
import ImgToBase64 from 'react-native-image-base64'
import FormStyle from "../Form.style";


export default class UploadImage extends React.Component {
    state = {
        photo: null,
    }

    handleChoosePhoto = () => {
        const options = {
            noData: true, 
        };

        launchImageLibrary(options, response => {
            console.log(response);
            if (response.assets[0].uri) {
                console.log('recieved response');
                this.state.photo = response;
            }
        })
    }

    handleUploadPhoto = () => {
        console.log(this.state.photo);
        if(this.state.photo != null) {
            console.log(this.state.photo.assets[0].uri);
            ImgToBase64.getBase64String(this.state.photo.assets[0].uri).then( 
                (base64String) => {
                    fetch('http://localhost:8080/api/v1/createPhoto', {
                        method: "POST", 
                        headers: {
                            Accept: "application/json",
                            "Content-Type": "application/json" 
                        },
                        body: JSON.stringify({
                            photo: base64String, 
                            type: this.state.photo.assets[0].type
                        })
                    });
                }
            )
        }

    }

    render() {
        const {photo} = this.state;
        return (
            <View style={FormStyle.container}>
                {photo && (
                    <Image 
                        source={{ uri: photo.photo.uri }}
                        style={{ width: 300, height: 300 }}
                    />
                )}

                <TouchableOpacity style={FormStyle.button} onPress={this.handleChoosePhoto}>
                    <Text style={FormStyle.loginText}>Select Image</Text>
                </TouchableOpacity>
                <TouchableOpacity style={FormStyle.button} onPress={this.handleUploadPhoto} >
                    <Text style={FormStyle.loginText}>Upload Image</Text>
                </TouchableOpacity>
            </View>
        )
    }
}

