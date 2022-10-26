import React from "react"
import {Button, Text, View} from "react-native"
import HandleLogout from "../../event-handler/HandleLogout"
import TopBar from "../component/TopBar"
import {launchImageLibrary} from 'react-native-image-picker'
import ImgToBase64 from 'react-native-image-base64'


export default class UploadImage extends React.Component {
    state = {
        photo: null,
    }

    handleChoosePhoto = () => {
        const options = {
            noData: true, 
        }

        launchImageLibrary(options, response => {
            console.log(response)
            if (response.assets[0].uri) {
                console.log('recieved response')
                this.state.photo = response
            }
        })
    }

    handleUploadPhoto = () => {
        console.log(this.state.photo)
        if(this.state.photo != null) {
            console.log(this.state.photo.assets[0].uri)
            ImgToBase64.getBase64String(this.state.photo.assets[0].uri).then( 
                (base64String) => {
                    fetch('http://localhost:8080/createPhoto', {
                        method: "POST", 
                        headers: {
                            Accept: "application/json",
                            "Content-Type": "application/json" 
                        },
                        body: JSON.stringify({
                            photo: base64String, 
                            type: this.state.photo.assets[0].type
                        })
                    })
                }
            )
        }

    }

    render() {
        const {photo} = this.state
        return (
            <View>
                {photo && (
                    <Image 
                        source={{ uri: photo.photo.uri }}
                        style={{ width: 300, height: 300 }}
                    />
                )}

                <Button title="Choose Photo" onPress={this.handleChoosePhoto} />
                <Button title="Upload" onPress={this.handleUploadPhoto} />
            </View>
        )
    }
}

